"use client";

import Link from "next/link";

export default function Navbar() {
  return (
    <nav className="flex justify-between items-center px-6 py-4 bg-green-600 text-white">
      <Link href="/" className="text-lg font-bold">
        🏸 BadmintonApp
      </Link>
      <div className="flex gap-4">
        <Link href="/reservations">My Reservations</Link>
        <Link href="/login">Login</Link>
      </div>
    </nav>
  );
}
