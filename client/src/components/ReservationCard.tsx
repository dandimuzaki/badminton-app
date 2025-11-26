"use client";

import { capitalize, formatDate, formatRupiah } from "@/utils/format";
import Image from "next/image";
import Link from "next/link";

export default function ReservationCard({reservation}: {reservation: any}) {
  return (
    <div key={reservation.id} className="md:p-4 p-2 rounded-lg bg-white flex md:gap-4 gap-2 shadow-[0_2px_10px_2px_rgba(0,0,0,0.1)]">
      <Image src={reservation.Court.imageUrl} alt={reservation.Court.name} width={500} height={250} className="rounded w-48 h-full object-cover"/>
        <div className="flex-1 grid gap-2">
        <div
        className={`${reservation.status === "cancelled" ? "bg-red-200 text-red-800" : reservation.status === "paid" ? "bg-green-200 text-green-800" : "bg-yellow-200 text-yellow-800" } w-fit px-2 py-1 rounded`}
        >{reservation.status === "pending" ? "Waiting for payment" : capitalize(reservation.status)}</div>
        <div>
        <h3 className="text font-bold">{reservation.Court.name}</h3>
        <p className="text-gray-500 text-sm">{reservation.Court.type}</p>
        </div>
        <div className="flex gap-4 text-xl">
        <p className="">{formatDate(reservation.date)}</p>
        <p className="font-bold">{reservation.Timeslot.startTime} - {reservation.Timeslot.endTime}</p>
        </div>
        <p className="font-bold">{formatRupiah(reservation.Court.price)}<span className="text-sm font-medium">/hour</span></p>
        <button className="w-fit justify-self-end bg-[var(--primary)] text-white font-bold text-lg px-3 py-1 rounded">{reservation.status === "pending" ? "Pay Now" : "Book Again"}</button>
        </div>
    </div>
  )
}