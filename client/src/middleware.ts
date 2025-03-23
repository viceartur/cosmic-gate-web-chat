import { NextResponse } from "next/server";
import type { NextRequest } from "next/server";
import { getToken } from "next-auth/jwt";

// Middleware to protect routes
export async function middleware(request: NextRequest) {
  const token = await getToken({
    req: request,
    secret: process.env.AUTH_SECRET,
  });

  // If the token exists, the user is authenticated
  if (token) {
    return NextResponse.next();
  }
  // If not authenticated, redirect to the sign-in page
  return NextResponse.redirect(new URL("/sign-in", request.url));
}

// Configuration for the middleware
export const config = {
  // Define the paths where the middleware should be applied
  matcher: [
    "/((?!api/auth|api|_next/static|_next/image|sign-in|sign-up|.*\\.png$).*)",
  ],
};
