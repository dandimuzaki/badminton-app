"use client"

import { Calendar } from "@/components/ui/calendar"
import useBooking from "@/hooks/useBooking"
import { formatRupiah } from "@/utils/format"
import Image from "next/image"

export default function BookPage() {
  const {
    date, setDate, timeSlots, courts, selectedCourt, 
    selectedSlot, setSelectedSlot,  
    error, handleSelectCourt, handleSubmit
  } = useBooking()

  return (
    <div className="md:pt-28 pt-20 px-8 py-4 md:px-20 md:py-16 min-h-screen">
      <h1 className="text-center text-2xl font-bold mb-2">Book a Badminton Court</h1>
      <p className="mb-10 text-center">Choose your preferred date, time, and court.</p>
      <form onSubmit={handleSubmit} className="grid md:grid-cols-[2fr_3fr] gap-4">
        <div className="space-y-8">
          {/* Date Input */}
          <div>
            <label className="block mb-1 font-medium text-gray-700">Select Date</label>
            <Calendar
              mode="single"
              selected={date}
              onSelect={setDate}
              className="rounded-lg border"
              captionLayout="dropdown"
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
                {timeSlots?.length > 0 ? (
                  timeSlots.map((t) => (
                    <option key={t.id} value={t.id}>
                      {t.start_time}–{t.end_time}
                    </option>
                  ))
                ) : (
                  <option disabled>No slots available</option>
                )}
              </select>
            </div>
          )}
        </div>

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
          ))) : (<p className="col-span-4 text-gray-500">{error}</p>)}
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
