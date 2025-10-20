"use client";

import { useAuth } from "@/context/AuthContext";
import { getCourts } from "@/services/courtService";
import { formatRupiah } from "@/utils/format";
import Image from "next/image";
import { useEffect, useState } from "react";

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
        <button className="mt-8 text-xl bg-[var(--accent)] px-3 py-1 rounded-lg font-bold w-fit">Book a Court Now</button>
        </div>
        <div className="z-10 h-full w-full absolute bottom-0 bg-linear-to-b from-[rgba(0,0,0,0.2)] to-[rgba(0,0,0,0.5)]"></div>
        </div>
      </section>
      <section className="px-8 py-4 md:px-20 md:py-16" id="courts">
        <div className="grid lg:grid-cols-3 md:grid-cols-2 md:gap-6 gap-4">
          {courts?.map(({id, name, imageUrl, type, description, price}) =>
            <div key={id} className="md:p-4 p-2 rounded-lg bg-white grid md:gap-4 gap-2 shadow-[0_2px_10px_2px_rgba(0,0,0,0.1)]">
              <Image src={imageUrl} alt={name} width={500} height={250} className="w-full h-48 object-cover"/>
              <div>
              <h3 className="text-lg font-bold">{name}</h3>
              <p className="text-gray-500 text-sm">{type}</p>
              </div>
              <p className="">{description}</p>
              <div className="grid gap-2 md:flex md:justify-between md:items-end">
              <p className="font-bold text-xl">{formatRupiah(price)}<span className="text-sm font-medium">/hour</span></p>
              <button className="px-4 py-1 bg-[var(--primary)] font-bold text-white rounded">Book Now</button>
              </div>
            </div>
          )}
        </div>
      </section>
    </div>
  );
}
