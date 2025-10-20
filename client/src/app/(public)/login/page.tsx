"use client";

import { Suspense, useState } from "react";
import { useRouter, useSearchParams } from "next/navigation";
import { Visibility } from '@mui/icons-material';
import { useAuth } from "@/context/AuthContext";

export default function LoginPage() {
  const { login } = useAuth()
  const router = useRouter();
  const searchParams = useSearchParams();

  const prevPath = searchParams.get("redirect") || "/"

  const [form, setForm] = useState({ email: "", password: "" });
  const [error, setError] = useState("");
  const [hidePassword, setHidePassword] = useState(true);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");

    try {
      await login(form.email, form.password);
      router.push(prevPath);
    } catch (err: unknown) {
      setError("Invalid email or password");
      console.error("Failed to login", err)
    }
  };

  return (
    <div className="bg-[var(--primary)] pt-36 min-h-screen">
      <Suspense fallback={<div>Loading...</div>}>
      <h2 className="text-2xl font-bold mb-6 text-center">Login</h2>

      <form onSubmit={handleSubmit} className="bg-white p-6 rounded-lg max-w-md mx-auto space-y-4">
        <div>
          <label className="font-bold">Email</label>
          <input
            type="email"
            name="email"
            required
            className="w-full border px-3 py-2 rounded-md"
            onChange={handleChange}
          />
        </div>

        <div>
          <label className="font-bold">Password</label>
          <div className="flex">
          <input
            type={hidePassword ? "password" : "text"}
            name="password"
            required
            className="w-full px-3 py-2 rounded-l-md border-y border-l border-black"
            onChange={handleChange}
          />
          <div 
          onClick={() => setHidePassword((prev) => !prev)}
          className={`${hidePassword ? "text-gray-300" : "text-gray-700"} border-black rounded-r-md border-y border-r flex items-center px-2`}>
            <Visibility/>
          </div>
          </div>
        </div>

        {error && <p className="text-red-500">{error}</p>}

        <button
          type="submit"
          className="w-full bg-green-500 text-white py-2 rounded-md hover:bg-green-600 transition"
        >
          Login
        </button>

        <p className="text-center">
        Don’t have an account?{" "}
          <a href="/register" className="text-green-600 hover:underline">
            Register here
          </a>
        </p>
      </form>
      </Suspense>
    </div>
  );
}
