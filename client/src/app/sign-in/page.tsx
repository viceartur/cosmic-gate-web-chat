"use client";

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
    const password = formData.get("password") as string;

    try {
      // Make a request to the Next API
      const response: Response = await fetch("/api/auth/signin", {
        method: "POST",
        body: JSON.stringify({ email, password }),
        headers: {
          "Content-Type": "application/json",
        },
      });

      const data = await response.json();

      if (!response.ok) {
        setErrorMessage(data.message);
      } else {
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
        {loading ? "Signing In..." : "Sign In"}
      </button>
      <span className="error-message">{errorMessage}</span>
      <div className="sign-up-link">
        First time user? <a href="/sign-up">Sign Up Here</a>
      </div>
    </form>
  );
}
