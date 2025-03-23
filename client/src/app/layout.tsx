import type { Metadata } from "next";
import { SessionProvider } from "next-auth/react";

import "./globals.css";

import { NavBar } from "@/components/navbar";

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
      <body>
        <SessionProvider>
          <NavBar />
          {children}
        </SessionProvider>
      </body>
    </html>
  );
}
