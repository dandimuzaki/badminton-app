"use client";

import Image from "next/image";
import Link from "next/link";
import logo from "./../assets/badminton-logo.png"
import { usePathname } from "next/navigation";

export default function Navbar() {
  const pathname = usePathname()
  const minimalNavRoutes = ["/login", "/register"]

  if (minimalNavRoutes.includes(pathname)) {
    return (
    <nav className="h-16 flex justify-between items-center px-6 bg-[var(--primary)] text-white font-bold fixed top-0 w-full">
      <Link href="/">
        <Image src={logo} alt="badminton-logo" width={180}/>
      </Link>
    </nav>
    )
  }

  return (
    <nav className="h-16 flex justify-between items-center px-6 bg-[var(--primary)] text-white font-bold fixed top-0 w-full">
      <Link href="/">
        <Image src={logo} alt="badminton-logo" width={180}/>
      </Link>
      <div className="flex items-center gap-2 py-4 text-lg">
        <Link href="/" className="rounded px-4 py-1 hover:bg-[var(--primary-dark)]">Home</Link>
        <Link href="/reservations/new" className="rounded px-4 py-1 hover:bg-[var(--primary-dark)]">Make a Booking</Link>
        <Link href="/reservations" className="rounded px-4 py-1 hover:bg-[var(--primary-dark)]">My Reservations</Link>
      </div>
      <div className="flex items-center gap-4 text-lg">
        <Link href="/login" className="border-2 border-white px-4 py-1 rounded hover:bg-white hover:text-[var(--primary)]">Login</Link>
        <Link href="/register" className="rounded px-4 py-1 hover:bg-[var(--primary-dark)] hover:border-[var(--primary-dark)] border-2 border-transparent">Register</Link>
      </div>
    </nav>
  );
}
