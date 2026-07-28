"use client";

import Image from "next/image";
import Link from "next/link";
import { useState } from "react";
import { useAuthStore } from "@/store/useAuthStore";
import { Button } from "./ui/button";
import { Menu, X } from "lucide-react";
import { motion, AnimatePresence } from "framer-motion";

export default function Navbar() {
  const { user } = useAuthStore();
  const [isOpen, setIsOpen] = useState(false);

  const menuVariants = {
    hidden: { opacity: 0, y: -20 },
    visible: {
      opacity: 1,
      y: 0,
      transition: {
        staggerChildren: 0.1,
        duration: 0.3,
      },
    },
    exit: { opacity: 0, y: -20, transition: { duration: 0.2 } },
  };

  const itemVariants = {
    hidden: { opacity: 0, y: -10 },
    visible: { opacity: 1, y: 0 },
  };

  const auth = useAuthStore()

  return (
    <nav className="z-50 h-16 flex justify-between items-center px-6 bg-background/80 backdrop-blur-md text-white fixed top-0 w-full">
      {/* Logo */}
      <Link href="/">
        <Image
          src="/images/shuttle-time-logo-trp.png"
          alt="badminton-logo"
          width={130}
          height={80}
          className="cursor-pointer"
        />
      </Link>

      {/* Desktop */}
      <ul className="hidden md:flex items-center gap-2 text-secondary">
        {["home", "about", "court", "testimonial", "contact"].map((item) => (
          <li
            key={item}
            className="hover:font-bold hover:text-primary hover:bg-secondary px-4 py-2 rounded-full transition"
          >
            <Link href={`#${item}`}>
              {item.charAt(0).toUpperCase() + item.slice(1)}
            </Link>
          </li>
        ))}
      </ul>

      {/* Right Section */}
      <div className="hidden md:flex items-center gap-2">
        {user ? (<>
          <Link href="/reservations">
            <Button className="rounded-full bg-secondary text-white cursor-pointer">
              My Reservations
            </Button>
          </Link>
            <Button onClick={() => auth.logout()} className="cursor-pointer rounded-full border border-secondary text-secondary">
              Logout
            </Button>
          </>
        ) : (
          <>
            <Link href="/login">
              <Button className="cursor-pointer rounded-full">
                Login
              </Button>
            </Link>
            <Link href="/register">
              <Button className="cursor-pointer rounded-full bg-secondary text-white">
                Register
              </Button>
            </Link>
          </>
        )}
      </div>

      {/* Mobile Button */}
      <button
        className="md:hidden"
        onClick={() => setIsOpen(!isOpen)}
      >
        {isOpen ? <X /> : <Menu />}
      </button>

      {/* Animated Mobile Menu */}
      <AnimatePresence>
        {isOpen && (
          <motion.div
            variants={menuVariants}
            initial="hidden"
            animate="visible"
            exit="exit"
            className="absolute top-16 left-0 w-full bg-background/95 backdrop-blur-md flex flex-col items-center gap-4 py-6 md:hidden shadow-lg"
          >
            {["home", "about", "court", "testimonial", "contact"].map((item) => (
              <motion.div key={item} variants={itemVariants}>
                <Link
                  href={`#${item}`}
                  onClick={() => setIsOpen(false)}
                  className="text-secondary text-lg"
                >
                  {item.charAt(0).toUpperCase() + item.slice(1)}
                </Link>
              </motion.div>
            ))}

            <motion.div variants={itemVariants} className="flex flex-col gap-2">
              {user ? (
                <Link href="/reservations">
                  <Button className="rounded-full bg-secondary text-white">
                    My Reservations
                  </Button>
                </Link>
              ) : (
                <>
                  <Link href="/login">
                    <Button className="rounded-full">
                      Login
                    </Button>
                  </Link>
                  <Link href="/register">
                    <Button className="rounded-full bg-secondary text-white">
                      Register
                    </Button>
                  </Link>
                </>
              )}
            </motion.div>
          </motion.div>
        )}
      </AnimatePresence>
    </nav>
  );
}