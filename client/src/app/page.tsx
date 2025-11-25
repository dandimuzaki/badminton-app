"use client";

import CourtCard from "@/components/CourtCard";
import { useAuth } from "@/context/AuthContext";
import { getCourts } from "@/services/courtService";
import { AlarmOn, Dashboard, LocalActivity, NearMe } from "@mui/icons-material";
import Image from "next/image";
import Link from "next/link";
import { useEffect, useState } from "react";
import court from "./../../src/assets/court-racket.jpg"

interface Court {
  id: string;
  name: string;
  imageUrl: string;
  type: string;
  description: string;
  price: number;
}

export default function Home() {
  const { token } = useAuth()
  const [courts, setCourts] = useState([])

  const fetchCourts = async () => {
    try {
      const result = await getCourts()
      setCourts(result.data)
    } catch (err: unknown) {
      console.error("Failed to fetch courts", err)
    }
  }
  useEffect(() => {
    if (token) {
      fetchCourts()
    }
  }, [token])

  return (
    <div className="min-h-screen">
      <section>
        <div className="relative min-h-screen">
          <Image src={'/images/family-court.jpg'} alt="ShuttleTime" width={900} height={1600} className="w-full h-full object-cover absolute top-0" />
        <div className="z-15 flex flex-col gap-4 items-center justify-center text-center absolute w-full h-full">
        <h1 className="text-white text-5xl/15 font-bold">Smash Your<span className="ml-4 font-bold text-white px-3 py-1 bg-[var(--primary)] rounded-lg">Schedule</span><br/>Not Your Time</h1>
        <p className="text-white text-lg">Book badminton courts anytime, anywhere — no waiting, no hassle.<br/>Just pick a court, grab your racket, and play your best game.</p>
        <Link href="/reservations/new" className="mt-8 text-xl bg-[var(--accent)] px-3 py-1 rounded-lg font-bold w-fit">Book a Court Now</Link>
        </div>
        <div className="z-10 h-full w-full absolute bottom-0 bg-linear-to-b from-[rgba(0,0,0,0.2)] to-[rgba(0,0,0,0.5)]"></div>
        </div>
      </section>
      <section className="px-8 pt-4 pb-2 md:px-20 md:pt-24 md:pb-12">
        <h1 className="text-center text-4xl font-bold mb-2">Courts Ready For Game</h1>
        <p className="mb-10 text-center">Indoor or outdoor, wooden or synthetic — choose your ideal match and start playing now.</p>
        <div className="grid lg:grid-cols-3 md:grid-cols-2 md:gap-6 gap-4">
          {courts?.map((c: Court) => <CourtCard key={c.id} c={c} />)}
        </div>
      </section>
      <section className="px-8 py-2 md:px-20 md:py-12">
        <h1 className="text-center text-4xl font-bold mb-2">Play Smarter, Not Harder</h1>
        <p className="mb-10 text-center">From instant booking to smart scheduling — we make every game effortless</p>
        <div className="grid grid-cols-2 gap-6">
            <div>
              <Image src={court} alt="Values ShuttleTime" width={100} height={100} className="w-full h-full object-cover rounded-lg" />
            </div>
            <div className="grid grid-cols-2 gap-6">
              <div className="flex flex-col justify-center gap-2 p-6 rounded-lg bg-[var(--primary-light)]">
                <LocalActivity fontSize="large"/>
                <p className="mt-4 font-bold text-xl">Seamless Booking</p>
                <p>Find, book, and confirm courts with just a few taps</p>
              </div>
              <div className="flex flex-col justify-center gap-2 p-6 rounded-lg bg-[var(--primary-light)]">
                <NearMe fontSize="large"/>
                <p className="mt-4 font-bold text-xl">Trusted Venues</p>
                <p>Verified courts with clear pricing and facility info</p>
              </div>
              <div className="flex flex-col justify-center gap-2 p-6 rounded-lg bg-[var(--primary-light)]">
                <Dashboard fontSize="large"/>
                <p className="mt-4 font-bold text-xl">Stay Organized</p>
                <p>Manage all your bookings in one simple dashboard</p>
              </div>
              <div className="flex flex-col justify-center gap-2 p-6 rounded-lg bg-[var(--primary-light)]">
                <AlarmOn fontSize="large"/>
                <p className="mt-4 font-bold text-xl">Game Ready</p>
                <p>Notifications and reminders keep your playtime on track</p>
              </div>
            </div>
          </div>
      </section>
    </div>
  );
}
