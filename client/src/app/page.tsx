"use client";

import { useSession } from "next-auth/react";

export default function Home() {
  const { data } = useSession();
  return <h1>Welcome, {data?.user.username}</h1>;
}
