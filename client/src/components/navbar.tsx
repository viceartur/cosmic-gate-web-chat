"use client";

import Link from "next/link";
import { signOut, useSession } from "next-auth/react";

export function NavBar() {
  const { data: session } = useSession();
  if (!session?.user) return;

  return (
    <nav>
      <Link href="/">Main</Link>
      <Link href="/profile">Profile</Link>
      <Link href="/users">People</Link>
      <Link href="/friend-requests">Friend Requests</Link>
      <Link href="/friends">Friends</Link>
      {session?.user && <button onClick={() => signOut()}>Sign Out</button>}
    </nav>
  );
}
