"use client";

import Link from "next/link";
import { signOut, useSession } from "next-auth/react";

export function NavBar() {
  const { data } = useSession();

  return (
    <nav>
      <Link href="/">Main</Link>
      <Link href="/profile">Profile</Link>
      <Link href="/people">People</Link>
      <Link href="/friends">Friend Requests</Link>
      <Link href="/friends">Friends</Link>
      <Link href="/messages">Messages</Link>
      {data?.user && <button onClick={() => signOut()}>Sign Out</button>}
    </nav>
  );
}
