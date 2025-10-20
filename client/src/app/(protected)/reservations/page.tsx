"use client";

import { useAuth } from "@/context/AuthContext";
import { getReservations } from "@/services/reservationService";
import { formatRupiah } from "@/utils/format";
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
  status: "pending" | "paid" | "cancelled"; // you can extend this union
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

  const fetchReservations = async () => {
    const result = await getReservations()
    setReservations(result.data)
  }

  useEffect(() => {
    if (token) {
      fetchReservations()
    }
  }, [token])

  return (
    <div className="pt-36 p-20 min-h-screen">
      <section>
        <div className="grid lg:grid-cols-2 md:grid-cols-2 md:gap-6 gap-4">
          {reservations?.map((reservation) =>
            <div key={reservation.id} className="md:p-4 p-2 rounded-lg bg-white grid md:gap-4 gap-2 shadow-[0_2px_10px_2px_rgba(0,0,0,0.1)]">
              <Image src={reservation.Court.imageUrl} alt={reservation.Court.name} width={500} height={250} className="w-full h-48 object-cover"/>
                <div>
                <h3 className="text-lg font-bold">{reservation.Court.name}</h3>
                <p className="text-gray-500 text-sm">{reservation.Court.type}</p>
                </div>
                <p className="font-bold">{reservation.date}</p>
                <p className="">{reservation.Timeslot.startTime} - {reservation.Timeslot.endTime}</p>
                <div className="grid gap-2 md:flex md:justify-between md:items-end">
                <p className="font-bold text-xl">{formatRupiah(reservation.Court.price)}<span className="text-sm font-medium">/hour</span></p>
                
                </div>
            </div>
          )}
        </div>
      </section>
    </div>
  );
}
