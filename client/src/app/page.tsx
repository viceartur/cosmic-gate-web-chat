"use client";

import { useSession } from "next-auth/react";

export default function Home() {
  const { data: session } = useSession();
  return (
    <section>
      <div className="section-headers">
        <h1>Cosmic Gate Chat</h1>
        <h3>Welcome, {session?.user.username}</h3>
      </div>
    </section>
  );
}
