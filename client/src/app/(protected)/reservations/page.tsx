"use client";

import { useAuth } from "@/context/AuthContext";
import { getReservations } from "@/services/reservationService";
import { formatDate, formatRupiah } from "@/utils/format";
import Image from "next/image";
import { useEffect, useState } from "react";

export interface Reservation {
  id: number;
  userId: number;
  User: {
    id: number;
    name: string;
    email: string;
    password: string;
    createdAt: string;
  };
  courtId: number;
  Court: {
    id: number;
    name: string;
    imageUrl: string;
    type: string;
    description: string;
    location: string;
    price: number;
    createdAt: string;
  };
  date: string; // ISO date string
  timeSlotId: number;
  Timeslot: {
    id: number;
    startTime: string;
    endTime: string;
    createdAt: string;
  };
  status: "pending" | "paid" | "cancelled" | "confirmed"; // you can extend this union
  Payment: {
    id: number;
    userId: number;
    reservationId: number;
    reservation: Reservation | null; // recursive type allowed, or use `any` if not needed
    amount: number;
    status: "pending" | "paid" | "failed";
    transactionId: string;
    createdAt: string;
    updatedAt: string;
  };
  createdAt: string;
  updatedAt: string;
}

export default function MyReservationPage() {
  const { token } = useAuth()
  const [reservations, setReservations] = useState<Reservation[]>([]);
  const [loading, setLoading] = useState(true)

  const fetchReservations = async () => {
    setLoading(true)
    try {
      const result = await getReservations()
      setReservations(result.data)
    } catch (err) {
      console.error("Failed to load reservations", err)
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    if (token) {
      fetchReservations()
    }
  }, [token])

  return (
    <div className="md:pt-28 pt-20 px-8 py-4 md:px-20 md:py-16 min-h-screen">
      <h1 className="text-center text-2xl font-bold mb-2">My Reservations</h1>
      <p className="mb-10 text-center">See all your bookings and don’t forget your play time!</p>
      {loading ? (<p>Loading...</p>) : reservations.length > 0 ? (<section>
        <div className="grid lg:grid-cols-2 md:grid-cols-2 md:gap-6 gap-4">
          {reservations?.map((reservation) =>
            <div key={reservation.id} className="md:p-4 p-2 rounded-lg bg-white flex md:gap-4 gap-2 shadow-[0_2px_10px_2px_rgba(0,0,0,0.1)]">
              <Image src={reservation.Court.imageUrl} alt={reservation.Court.name} width={500} height={250} className="rounded w-48 h-full object-cover"/>
                <div className="flex-1 grid gap-2">
                <div
                className={`${reservation.status === "cancelled" ? "bg-red-200 text-red-800" : reservation.status === "confirmed" ? "bg-green-200 text-green-800" : "bg-yellow-200 text-yellow-800" } w-fit px-2 py-1 rounded`}
                >{reservation.status === "pending" ? "Waiting for payment" : reservation.status}</div>
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
          )}
        </div>
      </section>)
    : (<p className="text-gray-400">You haven’t made any reservations yet.</p>)  
    }
    </div>
  );
}
