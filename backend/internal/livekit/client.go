package livekit

import (
	"context"
	"errors"
	"fmt"

	livekitproto "github.com/livekit/protocol/livekit"
	lksdk "github.com/livekit/server-sdk-go/v2"
)

var ErrCreateRoomFailed = errors.New("create livekit room failed")

type RoomClient interface {
	CreateRoom(context.Context, string) error
	DeleteRoom(context.Context, string) error
}

type NopClient struct{}

type Client struct {
	roomService *lksdk.RoomServiceClient
}

func NewClient(host string, apiKey string, apiSecret string) *Client {
	return &Client{
		roomService: lksdk.NewRoomServiceClient(host, apiKey, apiSecret),
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

func (NopClient) CreateRoom(context.Context, string) error {
	return nil
}

func (NopClient) DeleteRoom(context.Context, string) error {
	return nil
}
