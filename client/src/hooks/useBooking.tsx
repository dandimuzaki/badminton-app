"use client"

import { getAvailableCourts, getAvailableTimeslots } from "@/services/availabilityService"
import { createReservation } from "@/services/reservationService"
import { useAuthStore } from "@/store/useAuthStore"
import { ReservationRequest } from "@/types/reservation"
import { Timeslot } from "@/types/timeslot"
import { formatDate } from "@/utils/format"
import { useRouter } from "next/navigation"
import { useEffect, useState } from "react"
import usePayment from "./usePayment"

export default function useBooking() {
  const [date, setDate] = useState<Date | undefined>(undefined)
  const [timeSlots, setTimeSlots] = useState<Timeslot[]>([])
  const [courts, setCourts] = useState([])
  const [selectedCourt, setSelectedCourt] = useState<number>()
  const [selectedSlot, setSelectedSlot] = useState<number>()
  const [selectedPrice, setSelectedPrice] = useState(0)
  const [error, setError] = useState<string>("")
  const { user } = useAuthStore()
  const router = useRouter()
  const {handlePayment} = usePayment()

  useEffect(() => {
    if (!date) return
    const fetchSlots = async () => {
      const res = await getAvailableTimeslots({date: formatDate(date)})
      setTimeSlots(res)
    }
    fetchSlots()
  }, [date])

  useEffect(() => {
    if (!date || !selectedSlot) {
      setCourts([])
      setError("Please select a date and slot")
      return
    }

    const fetchCourts = async () => {
      const res = await getAvailableCourts({date: formatDate(date), time_slot_id: selectedSlot})
      setCourts(res)
      if (res.length === 0) {
        setError("No courts available for this time slot")
      } else {
        setError("")
      }
    }

    fetchCourts()
  }, [date, selectedSlot])

  const handleSelectCourt = (id: number, price: number) => {
    setSelectedCourt(id)
    setSelectedPrice(price)
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    if (!user) {
      // Store booking data temporarily
      const bookingData = { date, selectedSlot, selectedCourt }
      sessionStorage.setItem("pendingBooking", JSON.stringify(bookingData))
      router.push(`/login?redirect=${encodeURIComponent('/book/checkout')}`)
      return
    }

    try {
      const dataReservation: ReservationRequest = {
        date: formatDate(date),
        court_id: Number(selectedCourt),
        time_slot_id: Number(selectedSlot),
      };
      const reservation = await createReservation(dataReservation);

      const dataPayment = {
        reservation_id: reservation.ID,
        amount: selectedPrice,
      };
      
      handlePayment(dataPayment)
    } catch (err) {
      console.error("Failed to create reservation or payment:", err);
      alert("Something went wrong while processing your payment.");
    }
  };

  return {
    date, setDate, timeSlots, courts, selectedCourt, setSelectedCourt, 
    selectedSlot, setSelectedSlot, selectedPrice, setSelectedPrice, 
    error, handleSelectCourt, handleSubmit
  }
}