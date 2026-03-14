package livekit

import (
	"context"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	livekitauth "github.com/livekit/protocol/auth"
	livekitproto "github.com/livekit/protocol/livekit"
	lksdk "github.com/livekit/server-sdk-go/v2"
	"google.golang.org/protobuf/encoding/protojson"
)

var ErrCreateRoomFailed = errors.New("create livekit room failed")
var ErrCreateTokenFailed = errors.New("create livekit token failed")
var ErrInvalidWebhook = errors.New("invalid livekit webhook")
var ErrParseWebhook = errors.New("parse livekit webhook failed")

const defaultTokenValidity = 6 * time.Hour

type RoomClient interface {
	CreateRoom(context.Context, string) error
	DeleteRoom(context.Context, string) error
}

type TokenClient interface {
	CreatePublisherToken(roomName string, identity string, name string) (string, error)
	CreateViewerToken(roomName string, identity string, name string) (string, error)
}

type WebhookVerifier interface {
	VerifyWebhookEvent(authHeader string, body []byte) (*livekitproto.WebhookEvent, error)
}

type NopClient struct{}

type Client struct {
	roomService *lksdk.RoomServiceClient
	apiKey      string
	apiSecret   string
}

func NewClient(host string, apiKey string, apiSecret string) *Client {
	return &Client{
		roomService: lksdk.NewRoomServiceClient(host, apiKey, apiSecret),
		apiKey:      apiKey,
		apiSecret:   apiSecret,
	}
}

func (c *Client) CreateRoom(ctx context.Context, roomName string) error {
	if _, err := c.roomService.CreateRoom(ctx, &livekitproto.CreateRoomRequest{Name: roomName}); err != nil {
		return fmt.Errorf("%w: %v", ErrCreateRoomFailed, err)
	}

	return nil
}

func (c *Client) DeleteRoom(ctx context.Context, roomName string) error {
	if _, err := c.roomService.DeleteRoom(ctx, &livekitproto.DeleteRoomRequest{Room: roomName}); err != nil {
		return fmt.Errorf("delete livekit room: %w", err)
	}

	return nil
}

func (c *Client) CreatePublisherToken(roomName string, identity string, name string) (string, error) {
	return c.createToken(roomName, identity, name, true)
}

func (c *Client) CreateViewerToken(roomName string, identity string, name string) (string, error) {
	return c.createToken(roomName, identity, name, false)
}

func (c *Client) createToken(roomName string, identity string, name string, canPublish bool) (string, error) {
	grant := &livekitauth.VideoGrant{
		RoomJoin: true,
		Room:     roomName,
	}
	grant.SetCanPublish(canPublish)
	grant.SetCanSubscribe(true)

	token, err := livekitauth.NewAccessToken(c.apiKey, c.apiSecret).
		SetIdentity(identity).
		SetName(name).
		SetValidFor(defaultTokenValidity).
		SetVideoGrant(grant).
		ToJWT()
	if err != nil {
		return "", fmt.Errorf("%w: %v", ErrCreateTokenFailed, err)
	}

	return token, nil
}

func (c *Client) VerifyWebhookEvent(authHeader string, body []byte) (*livekitproto.WebhookEvent, error) {
	if authHeader == "" {
		return nil, ErrInvalidWebhook
	}

	verifier, err := livekitauth.ParseAPIToken(authHeader)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidWebhook, err)
	}

	secret := livekitauth.NewSimpleKeyProvider(c.apiKey, c.apiSecret).GetSecret(verifier.APIKey())
	if secret == "" {
		return nil, ErrInvalidWebhook
	}

	_, claims, err := verifier.Verify(secret)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidWebhook, err)
	}

	sum := sha256.Sum256(body)
	hash := base64.StdEncoding.EncodeToString(sum[:])
	if subtle.ConstantTimeCompare([]byte(claims.Sha256), []byte(hash)) != 1 {
		return nil, ErrInvalidWebhook
	}

	event := &livekitproto.WebhookEvent{}
	if err := (protojson.UnmarshalOptions{
		DiscardUnknown: true,
		AllowPartial:   true,
	}).Unmarshal(body, event); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrParseWebhook, err)
	}

	return event, nil
}

func (NopClient) CreateRoom(context.Context, string) error {
	return nil
}

func (NopClient) DeleteRoom(context.Context, string) error {
	return nil
}

func (NopClient) CreatePublisherToken(string, string, string) (string, error) {
	return "", nil
}

func (NopClient) CreateViewerToken(string, string, string) (string, error) {
	return "", nil
}

func (NopClient) VerifyWebhookEvent(string, []byte) (*livekitproto.WebhookEvent, error) {
	return nil, ErrInvalidWebhook
}
