"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { Visibility } from '@mui/icons-material';
import { useAuth } from "@/context/AuthContext";

export default function RegisterPage() {
  const { register } = useAuth()
  const router = useRouter();

  const [form, setForm] = useState({ name: "", email: "", password: "", confirmPassword: "" });
  const [error, setError] = useState("");
  const [hidePassword, setHidePassword] = useState(true);
  const [hideConfirmPassword, setHideConfirmPassword] = useState(true);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");

    try {
      await register(form.name, form.email, form.password);
      router.push('/login');
    } catch (err: unknown) {
      setError("Email is already used");
      console.error("Failed to login", err)
    }
  };

  return (
    <div className="bg-[var(--primary)] p-20 min-h-screen">
      <h2 className="text-2xl font-bold mb-6 text-center">Register</h2>

      <form onSubmit={handleSubmit} className="bg-white p-6 rounded-lg max-w-md mx-auto space-y-4">
        <div>
          <label className="font-bold">Name</label>
          <input
            type="text"
            name="name"
            required
            className="w-full border px-3 py-2 rounded-md"
            onChange={handleChange}
          />
        </div>
        
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

        <div>
          <label className="font-bold">Confirm Password</label>
          <div className="flex">
          <input
            type={hideConfirmPassword ? "password" : "text"}
            name="confirmPassword"
            required
            className="w-full px-3 py-2 rounded-l-md border-y border-l border-black"
            onChange={handleChange}
          />
          <div 
          onClick={() => setHideConfirmPassword((prev) => !prev)}
          className={`${hideConfirmPassword ? "text-gray-300" : "text-gray-700"} border-black rounded-r-md border-y border-r flex items-center px-2`}>
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
        Already have an account?{" "}
          <a href="/login" className="text-green-600 hover:underline">
            Login here
          </a>
        </p>
      </form>
    </div>
  );
}
