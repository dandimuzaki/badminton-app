"use client"

import { getAvailableCourts, getAvailableTimeslots } from "@/services/availabilityService"
import { createPayment } from "@/services/paymentService"
import { createReservation } from "@/services/reservationService"
import { formatRupiah } from "@/utils/format"
import Image from "next/image"
import { useEffect, useState } from "react"

interface TimeSlot {
  id: number
  startTime: string
  endTime: string
}

export default function BookPage() {
  const [date, setDate] = useState("")
  const [timeSlots, setTimeSlots] = useState<TimeSlot[]>([])
  const [courts, setCourts] = useState([])
  const [selectedCourt, setSelectedCourt] = useState<string>("")
  const [selectedSlot, setSelectedSlot] = useState<string>("")
  const [message, setMessage] = useState("")
  const [selectedPrice, setSelectedPrice] = useState(0)

  useEffect(() => {
    if (!date) return
    const fetchSlots = async () => {
      const res = await getAvailableTimeslots(new URLSearchParams({ date }))
      setTimeSlots(res.data)
    }
    fetchSlots()
  }, [date])

  useEffect(() => {
    if (!date || !selectedSlot) {
      setCourts([])
      setMessage("Please select a date and slot")
      return
    }

    const fetchCourts = async () => {
      const res = await getAvailableCourts(new URLSearchParams({ date, timeSlotId: selectedSlot }))
      setCourts(res.data)
      if (res.data.length === 0) {
        setMessage("No courts available for this time slot")
      } else {
        setMessage("")
      }
    }

    fetchCourts()
  }, [date, selectedSlot])

  const handleSelectCourt = (id: string, price: number) => {
    setSelectedCourt(id)
    setSelectedPrice(price)
  }

  const handleSubmit = async (e: React.FormEvent) => {
  e.preventDefault();

  try {
    // STEP 1: Create reservation
    const dataReservation = {
      date,
      courtId: Number(selectedCourt),
      timeSlotId: Number(selectedSlot),
    };
    const reservation = await createReservation(dataReservation);

    // STEP 2: Create payment
    const dataPayment = {
      reservationId: reservation.data.id,
      amount: selectedPrice,
    };
    const payment = await createPayment(dataPayment);

    // STEP 3: Ensure Snap is loaded
    if (typeof window.snap === "undefined") {
      alert("Midtrans Snap is not loaded yet. Please refresh.");
      return;
    }

    // STEP 4: Trigger Midtrans payment popup
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


  return (
    <div className="md:pt-28 pt-20 px-8 py-4 md:px-20 md:py-16 min-h-screen">
      <h1 className="text-center text-2xl font-bold mb-2">Book a Badminton Court</h1>
      <p className="mb-10 text-center">Choose your preferred date, time, and court.</p>
      <form onSubmit={handleSubmit} className="grid md:grid-cols-2 gap-4">
        {/* Date Input */}
        <div>
          <label className="block mb-1 font-medium text-gray-700">Select Date</label>
          <input
            type="date"
            value={date}
            onChange={(e) => setDate(e.target.value)}
            className="border p-2 rounded w-full"
          />
        </div>

        {/* Timeslot Dropdown */}
        {(
          <div>
            <label className="block mb-1 font-medium text-gray-700">Time Slot</label>
            <select
              value={selectedSlot}
              onChange={(e) => setSelectedSlot(e.target.value)}
              className="border p-2 rounded w-full"
            >
              <option value="">Select Time</option>
              {timeSlots.length > 0 ? (
                timeSlots.map(({id, startTime, endTime}) => (
                  <option key={id} value={id}>
                    {startTime}–{endTime}
                  </option>
                ))
              ) : (
                <option disabled>No slots available</option>
              )}
            </select>
          </div>
        )}

        {/* Court Select */}
        <div className="col-span-2">
          <label className="block mb-1 font-medium text-gray-700">Select Court</label>
          <div className="grid md:grid-cols-4 gap-4">
          {courts.length > 0 ? (courts.map(({id, name, type, imageUrl, price}) => (
            <div key={id}
            onClick={() => handleSelectCourt(id, price)}
            className={`${id === selectedCourt ? "bg-[var(--primary-light)]" : "bg-white"} md:p-4 p-2 rounded-lg grid md:gap-4 gap-2 shadow-[0_2px_10px_2px_rgba(0,0,0,0.1)]`}>
              <Image src={imageUrl} alt={name} width={500} height={250} className="w-full h-36 object-cover"/>
              <div className="grid gap-1">
                <h3 className="text-base/5 font-bold">{name}</h3>
                <p className="text-gray-500 text-sm">{type}</p>
              </div>
              <div className="grid gap-2 md:flex md:justify-between md:items-end">
                <p className="font-bold text-xl">{formatRupiah(price)}<span className="text-sm font-medium">/hour</span></p>
              </div>
            </div>
          ))) : (<p className="col-span-4 text-gray-500">{message}</p>)}
          </div>
        </div>

        <button
          type="submit"
          className="mt-8 flex justify-self-center col-span-2 w-fit py-2 px-4 font-bold bg-[var(--primary)] text-white rounded-lg hover:bg-green-700 disabled:bg-gray-400"
        >
          Pay Now
        </button>
      </form>
    </div>
  )
}
