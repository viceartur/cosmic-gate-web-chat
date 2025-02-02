import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  title: "Cosmic Gate Chat",
  description: "Cosmic Gate Chat",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  );
}
