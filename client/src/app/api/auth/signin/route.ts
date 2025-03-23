import { NextResponse } from "next/server";
import { signIn } from "@/auth";

export async function POST(req: Request) {
  try {
    const { email, password } = await req.json();

    const user = await signIn("credentials", {
      email,
      password,
      redirect: false,
    });

    if (user) {
      return NextResponse.json(
        { message: "User Authenticated" },
        { status: 200 }
      );
    } else {
      return NextResponse.json(
        { message: "Invalid credentials" },
        { status: 401 }
      );
    }
  } catch (error: any) {
    return NextResponse.json(
      { message: error.cause?.err?.message || "Internal Server Error" },
      { status: 500 }
    );
  }
}
