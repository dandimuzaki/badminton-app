"use client";

import { Court } from "@/types/court";
import { formatRupiah } from "@/utils/format";
import Image from "next/image";

export default function CourtCard({ court }: {court: Court}) {
  return (
    <div key={court.ID} className="space-y-1 md:space-y-3">
      <Image src={court.image_url} alt={court.name} width={500} height={250} className="w-full h-48 object-cover rounded-lg mb-2"/>
      <div className="space-y-2">
        <p className="text-sm md:text-sm text-secondary px-3 py-1 rounded bg-secondary/20 w-fit">{court.type}</p>
        <h3 className="text-lg md:text-xl font-bold">{court.name}</h3>
      </div>
      <p className="text-sm md:text-base">{court.description}</p>
      <div className="grid gap-2 md:flex md:justify-between md:items-end">
        <p className="text-lg md:text-xl">Start from <span className="font-bold">{formatRupiah(court.price)}</span><span className="text-sm font-medium">/hour</span></p>
      </div>
    </div>
  )
}