"use client";

import { useEffect } from "react";
import { useRouter } from "next/navigation";
import { useAuthStore } from "@/store/useAuthStore";

export default function ProtectedLayout({ children }: { children: React.ReactNode }) {
  const token = useAuthStore((state) => state.token);
  const router = useRouter();

  useEffect(() => {
    if (!token) {
      const currentPath = window.location.pathname;
      router.push(`/login?redirect=${encodeURIComponent(currentPath)}`);
    }
  }, [token, router]);

  if (!token) return null;

  return <>{children}</>;
}
