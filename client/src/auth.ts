import { authUser } from "@/actions/auth";
import NextAuth from "next-auth";
import Credentials from "next-auth/providers/credentials";

export const { handlers, signIn, signOut, auth } = NextAuth({
  // Pages for authentication
  pages: {
    signIn: "/sign-in",
  },

  // Providers for authentication
  providers: [
    Credentials({
      credentials: {
        email: {},
        password: {},
      },
      // Authorize function to validate user credentials
      authorize: async (credentials) => {
        if (!credentials?.email || !credentials?.password) {
          throw new Error("Email and password are required.");
        }

        const user = await authUser(
          credentials.email as string,
          credentials.password as string
        );
        if (!user) {
          throw new Error("Invalid credentials.");
        }

        return user;
      },
    }),
  ],

  // Callbacks to handle JWT and session
  callbacks: {
    // JWT is created or updated
    async jwt({ token, user }) {
      if (user) {
        token.userId = user.id;
        token.email = user.email;
        token.username = user.username;
      }
      return token;
    },
    // Session is checked
    async session({ session, token }) {
      if (token) {
        session.user.id = token.userId as string;
        session.user.email = token.email as string;
        session.user.username = token.username as string;
      }
      return session;
    },
  },
});
