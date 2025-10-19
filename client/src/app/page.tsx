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
    <div className="md:pt-36 md:p-20 pt-24 p-8">
      <section>

      </section>
      <section>
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
