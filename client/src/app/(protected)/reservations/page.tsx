"use client";

import { useState } from "react";
import { useRouter, useSearchParams } from "next/navigation";
import { Visibility } from '@mui/icons-material';
import { useAuth } from "@/context/AuthContext";

export default function LoginPage() {
  const { login } = useAuth()
  const router = useRouter();
  const searchParams = useSearchParams();

  const prevPath = searchParams.get("redirect")

  const [form, setForm] = useState({ email: "", password: "" });
  const [error, setError] = useState("");

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");
    console.log(form)
    router.push(`/${prevPath}`);

    /*try {
      await login(form.email, form.password);
      router.push("/");
    } catch (err: unknown) {
      setError("Invalid email or password");
      console.error("Failed to login", err)
    }*/
  };

  return (
    <div className="bg-[var(--primary)] pt-36 min-h-screen">

    </div>
  );
}
