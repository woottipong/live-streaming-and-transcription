import "./globals.css";

export const metadata = {
  title: "Realtime Streaming",
  description: "Live streaming with realtime transcription",
};

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  );
}
