"use client";

import { Testimonial } from "@/mocks/testimonial";
import { FormatQuote } from "@mui/icons-material";
import Image from "next/image";

export default function TestimonialCard({ testimonial }: {testimonial: Testimonial}) {
  return (
    <div className="min-w-60 md:w-full flex flex-col gap-6 justify-between p-4 rounded-lg shadow-[0_1px_2px_1px_rgba(0,0,0,0.1)] bg-white">
      <div className="space-y-2">
        <h3 className="font-semibold text-xl">{testimonial.title}</h3>
        <p>{testimonial.description}</p>
      </div>
      <div className="flex gap-2 text-sm">
        <Image src={testimonial.profile} alt={testimonial.name} width={500} height={250} className="rounded-full h-10 w-10 object-cover rounded-lg"/>
        <div className="space-y-1">
          <h4 className="font-semibold">{testimonial.name}</h4>
          <p className="text-sm text-gray-500">{testimonial.type}</p>
        </div>
      </div>
    </div>
  )
}