"use client"

import { getAvailableCourts, getAvailableTimeslots } from "@/services/availabilityService"
import { createPayment } from "@/services/paymentService"
import { createReservation } from "@/services/reservationService"
import { useAuthStore } from "@/store/useAuthStore"
import { Timeslot } from "@/types/timeslot"
import { useRouter } from "next/navigation"
import { useEffect, useState } from "react"

export default function useBooking() {
  const [date, setDate] = useState<Date | undefined>(undefined)
  const [timeSlots, setTimeSlots] = useState<Timeslot[]>([])
  const [courts, setCourts] = useState([])
  const [selectedCourt, setSelectedCourt] = useState<string>("")
  const [selectedSlot, setSelectedSlot] = useState<string>("")
  const [selectedPrice, setSelectedPrice] = useState(0)
  const [error, setError] = useState<string>("")
  const { user } = useAuthStore()
  const router = useRouter()

  const dateString = date?.toLocaleDateString().replaceAll("/", "-") ?? ""

  useEffect(() => {
    if (!date) return
    const fetchSlots = async () => {
      const res = await getAvailableTimeslots({date: dateString})
      setTimeSlots(res)
    }
    fetchSlots()
  }, [date, dateString])

  useEffect(() => {
    if (!date || !selectedSlot) {
      setCourts([])
      setError("Please select a date and slot")
      return
    }

    const fetchCourts = async () => {
      const res = await getAvailableCourts({date: dateString, time_slot_id: selectedSlot})
      setCourts(res)
      if (res.length === 0) {
        setError("No courts available for this time slot")
      } else {
        setError("")
      }
    }

    fetchCourts()
  }, [date, selectedSlot, dateString])

  const handleSelectCourt = (id: string, price: number) => {
    setSelectedCourt(id)
    setSelectedPrice(price)
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    console.log("user", user)

    if (!user) {
      // Store booking data temporarily
      const bookingData = { date, selectedSlot, selectedCourt }
      sessionStorage.setItem("pendingBooking", JSON.stringify(bookingData))
      router.push(`/login?redirect=${encodeURIComponent('/book/checkout')}`)
      return
    }

    try {
      const dataReservation = {
        date,
        courtId: Number(selectedCourt),
        timeSlotId: Number(selectedSlot),
      };
      const reservation = await createReservation(dataReservation);

      const dataPayment = {
        reservationId: reservation.data.id,
        amount: selectedPrice,
      };
      const payment = await createPayment(dataPayment);

      if (typeof window.snap === "undefined") {
        alert("Midtrans Snap is not loaded yet. Please refresh.");
        return;
      }

      window.snap.pay(payment.snapToken, {
        onSuccess: async function (result) {
          console.log("✅ Payment success:", result);
          alert("Payment successful!");
        },
        onPending: function (result) {
          console.log("⏳ Payment pending:", result);
          alert("Payment pending. Please wait for confirmation.");
        },
        onError: function (result) {
          console.error("❌ Payment error:", result);
          alert("Payment failed. Please try again.");
        },
        onClose: function () {
          console.log("⚠️ Payment popup closed before finishing");
        },
      });
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