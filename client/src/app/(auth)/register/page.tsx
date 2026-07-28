"use client";

import { useState } from "react";
import { useRouter, useSearchParams } from "next/navigation";
import { Visibility } from '@mui/icons-material';
import { useAuthStore } from "@/store/useAuthStore";
import { register } from "@/services/authService";

export default function RegisterPage() {
  const loginStore = useAuthStore((state) => state.login)
  const router = useRouter();
  const searchParams = useSearchParams()

  const [form, setForm] = useState({ name: "", email: "", password: "", confirmPassword: "" });
  const [error, setError] = useState("");
  const [hidePassword, setHidePassword] = useState(true);
  const [hideConfirmPassword, setHideConfirmPassword] = useState(true);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const prevPath = searchParams.get('redirect') || '/'

  const handleRegister = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");

    try {
      const res = await register(form.name, form.email, form.password);
      loginStore(res.user, res.token)

      router.push(prevPath)
    } catch (err: unknown) {
      setError("Email is already used");
      console.error("Failed to register", err)
    }
  };

  return (
    <>
      <h2 className="text-2xl font-bold mb-6 text-center">Register</h2>

      <form onSubmit={handleRegister} className="bg-white p-6 rounded-lg max-w-lg mx-auto space-y-4">
        <div>
          <label className="font-bold">Name</label>
          <input
            type="text"
            name="name"
            required
            className="w-full border px-3 py-2 rounded-md border-gray-500"
            onChange={handleChange}
          />
        </div>
        
        <div>
          <label className="font-bold">Email</label>
          <input
            type="email"
            name="email"
            required
            className="w-full border px-3 py-2 rounded-md border-gray-500"
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
            className="w-full px-3 py-2 rounded-l-md border-y border-l border-gray-500"
            onChange={handleChange}
          />
          <div 
          onClick={() => setHidePassword((prev) => !prev)}
          className={`${hidePassword ? "text-gray-300" : "text-gray-700"} border-gray-500 rounded-r-md border-y border-r flex items-center px-2`}>
            <Visibility/>
          </div>
          </div>
        </div>

        <div>
          <label className="font-bold">Confirm Password</label>
          <div className="flex">
          <input
            type={hideConfirmPassword ? "password" : "text"}
            name="confirmPassword"
            required
            className="w-full px-3 py-2 rounded-l-md border-y border-l border-gray-500"
            onChange={handleChange}
          />
          <div 
          onClick={() => setHideConfirmPassword((prev) => !prev)}
          className={`${hideConfirmPassword ? "text-gray-300" : "text-gray-700"} border-gray-500 rounded-r-md border-y border-r flex items-center px-2`}>
            <Visibility/>
          </div>
          </div>
        </div>

        {error && <p className="text-red-500">{error}</p>}

        <button
          type="submit"
          className="w-full bg-green-500 text-white py-2 rounded-md hover:bg-green-600 transition"
        >
          Register
        </button>

        <p className="text-center text-sm">
        Already have an account?{" "}
          <a href="/login" className="text-green-600 hover:underline">
            Login here
          </a>
        </p>
      </form>
    </>
  );
}
