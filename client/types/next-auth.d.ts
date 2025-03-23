import NextAuth from "next-auth";

declare module "next-auth" {
  interface User {
    id: string;
    email: string;
    username: string;
  }

  interface Session {
    user: User & DefaultSession["user"];
  }
}
