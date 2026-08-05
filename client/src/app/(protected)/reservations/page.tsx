"use client";

import { Button } from "@/components/ui/button";
import usePayment from "@/hooks/usePayment";
import { getReservations } from "@/services/reservationService";
import { useAuthStore } from "@/store/useAuthStore";
import { Reservation } from "@/types/reservation";
import { capitalize, formatDateStr, formatRupiah } from "@/utils/format";
import Image from "next/image";
import { useEffect, useState } from "react";

export default function MyReservationPage() {
  const token = useAuthStore((state) => state.token)
  const [reservations, setReservations] = useState<Reservation[]>([]);
  const [loading, setLoading] = useState(true)
  const {handlePayment} = usePayment()

  const fetchReservations = async () => {
    setLoading(true)
    try {
      const result = await getReservations()
      setReservations(result)
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
        <div className="grid lg:grid-cols-2 md:gap-6 gap-4">
          {reservations?.map((reservation) =>
            (<div key={reservation.ID} className="md:p-4 p-2 rounded-lg bg-white flex md:gap-4 gap-2 shadow-[0_2px_10px_2px_rgba(0,0,0,0.1)]">
              <Image src={reservation.court.image_url} alt={reservation.court.name} width={500} height={250} className="rounded w-48 h-full object-cover"/>
                <div className="flex-1 grid gap-2">
                <div
                className={`${reservation.status === "cancelled" ? "bg-red-200 text-red-800" : reservation.status === "paid" ? "bg-green-200 text-green-800" : "bg-yellow-200 text-yellow-800" } w-fit px-2 py-1 rounded text-sm`}
                >{reservation.status === "pending" ? "Waiting for payment" : capitalize(reservation.status)}</div>
                <div>
                <h3 className="text font-bold">{reservation.court.name}</h3>
                <p className="text-gray-500 text-sm">{reservation.court.type}</p>
                </div>
                <div className="">
                <p className="">{formatDateStr(reservation.date)}</p>
                <p className="text-lg font-bold">{reservation.timeslot.start_time} - {reservation.timeslot.end_time}</p>
                </div>
                <p className="font-bold">{formatRupiah(reservation.court.price)}<span className="text-sm font-medium">/hour</span></p>
                {reservation.status == "pending" && <Button onClick={() => handlePayment({
                  reservation_id: reservation.ID,
                  amount: reservation.court.price
                })} className="rounded bg-secondary text-white cursor-pointer">
                  Pay Now
                </Button>}
                </div>
            </div>)
          )}
        </div>
      </section>)
    : (<p className="text-gray-400">You haven’t made any reservations yet.</p>)  
    }
    </div>
  );
}
