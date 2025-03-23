"use client";

import { signUp } from "@/actions/auth";
import { useState } from "react";

export default function SignIn() {
  const [errorMessage, setErrorMessage] = useState("");
  const [loading, setLoading] = useState(false);

  async function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
    event.preventDefault();
    setErrorMessage("");
    setLoading(true);

    const formData = new FormData(event.currentTarget);
    const email = formData.get("email") as string;
    const username = formData.get("username") as string;
    const password = formData.get("password") as string;

    try {
      // Create an Account for the User
      const { userId } = await signUp({ email, username, password });

      // Authenticate the created User
      const authUser: Response = await fetch("/api/auth/signin", {
        method: "POST",
        body: JSON.stringify({ email, password }),
        headers: {
          "Content-Type": "application/json",
        },
      });

      const data = await authUser.json();

      if (!authUser.ok) {
        setErrorMessage(data.message);
      } else {
        console.log(`Account created. Your user ID: ${userId}`);
        window.location.href = "/";
      }
    } catch (error: unknown) {
      if (error instanceof Error) {
        setErrorMessage(error.message);
      } else {
        setErrorMessage("An unknown error occurred.");
      }
    } finally {
      setLoading(false);
    }
  }

  return (
    <form onSubmit={handleSubmit} className="sign-in-form">
      <label htmlFor="email">
        Email
        <input name="email" id="email" required className="input-field" />
      </label>
      <label htmlFor="email">
        Username
        <input name="username" id="username" required className="input-field" />
      </label>
      <label htmlFor="password">
        Password
        <input
          name="password"
          id="password"
          type="password"
          required
          className="input-field"
        />
      </label>
      <button type="submit" disabled={loading} className="submit-button">
        {loading ? "Signing Up..." : "Sign Up"}
      </button>
      <span className="error-message">{errorMessage}</span>
      <div className="sign-up-link">
        Already have an Account? <a href="/sign-in">Sign In Here</a>
      </div>
    </form>
  );
}
