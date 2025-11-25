"use client";

import { formatRupiah } from "@/utils/format";
import Image from "next/image";
import Link from "next/link";

interface Court {
  id: string;
  name: string;
  imageUrl: string;
  type: string;
  description: string;
  price: number;
}

export default function CourtCard({ c }: {c: Court}) {
  return (
    <div key={c.id} className="md:p-4 p-2 rounded-lg bg-white grid md:gap-4 gap-2 shadow-[0_2px_10px_2px_rgba(0,0,0,0.1)]">
      <Image src={c.imageUrl} alt={c.name} width={500} height={250} className="w-full h-48 object-cover"/>
      <div>
      <h3 className="text-lg font-bold">{c.name}</h3>
      <p className="text-gray-500 text-sm">{c.type}</p>
      </div>
      <p className="">{c.description}</p>
      <div className="grid gap-2 md:flex md:justify-between md:items-end">
      <p className="font-bold text-xl">{formatRupiah(c.price)}<span className="text-sm font-medium">/hour</span></p>
      <Link href="/reservations/new" className="px-4 py-1 bg-[var(--primary)] font-bold text-white rounded">Book Now</Link>
      </div>
    </div>
  )
}