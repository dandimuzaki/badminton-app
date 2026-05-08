"use client"

import Image from "next/image"
import { useState } from "react"
import { Button } from "../ui/button"
import SectionBadge from "./section-badge"

export default function ServiceSection() {
  const data = [
    {
      id: 1,
      name: "Clean & Well-Maintained Courts",
      description: "Enjoy consistently clean courts with high-quality flooring, ensuring safe movement and smooth gameplay every session.",
      image: "/images/clean-court.jpg",
      isActive: true
    },
    {
      id: 2,
      name: "Professional-Grade Lighting",
      description: "Bright, evenly distributed lighting that eliminates shadows and keeps visibility sharp, perfect for both casual and competitive matches.",
      image: "/images/professional-lighting.jpg",
      isActive: false
    },
    {
      id: 3,
      name: "Changing Rooms & Showers",
      description: "Freshen up before or after your game with clean, comfortable changing areas and well-maintained shower facilities.",
      image: "/images/changing-room.jpg",
      isActive: false
    },
    {
      id: 4,
      name: "Equipment Rental",
      description: "No gear? No problem. Rent rackets and shuttlecocks on-site and jump straight into the action.",
      image: "/images/rental-equipment.jpg",
      isActive: false
    },
    {
      id: 5,
      name: "Comfortable Waiting & Rest Areas",
      description: "Relax between matches in cozy seating areas designed for players and spectators alike.",
      image: "/images/family-court.jpg",
      isActive: false
    }
  ]

  const [services, setServices] = useState(data)

  const handleOpenCard = (id: number) => {
    setServices((prev) => prev.map((s) => ({
      ...s,
      isActive: s.id === id
    })))
  }

  return (
    <section className="p-16 bg-white">
      <div className="grid grid-cols-2 gap-4 items-end mb-8">
        <div>
          <SectionBadge text="Our Services"/>
          <h2 className="text-4xl font-bold tetx-secondary">Explore Our Incredible Amenities and Services</h2>
        </div>
        <p>
          Crafted for comfort, safety, and peak performance, every court is thoughtfully equipped to elevate your game. Whether you're here for a quick rally or a serious match, everything is designed to keep you focused on what matters most, playing your best.
        </p>
      </div>
      
      <div className="flex gap-4 md:flex-row flex-col">
        {services.map((s) => (
          <div 
            key={s.id} 
            onClick={() => handleOpenCard(s.id)} 
            className={`cursor-pointer transition-all duration-500 ease-in-out w-full ${
              s.isActive ? 'md:flex-[3]' : 'md:flex-1'
            }`}
          >
            {s.isActive ? (
              <div className="bg-secondary p-4 md:p-6 rounded-lg flex flex-col md:flex-row gap-4 w-full h-full transition-all duration-500">
  
                {/* Image */}
                <div className="relative w-full md:w-[60%] h-48 md:h-auto overflow-hidden rounded-md">
                  <Image 
                    src={s.image} 
                    alt={s.name} 
                    fill
                    className="object-cover scale-x-[-1]" 
                  />
                  <div className="absolute inset-0 bg-gradient-to-t from-black/80 to-transparent"></div>
                </div>

                {/* Text */}
                <div className="text-white flex flex-col justify-end">
                  <h3 className="text-xl md:text-3xl font-semibold mb-2">{s.name}</h3>
                  <p className="mb-4 text-sm md:text-base">{s.description}</p>
                  <Button className="w-fit rounded-full">Book Now</Button>
                </div>

              </div>
            ) : (
              <div className="relative rounded-lg w-full h-40 md:h-full overflow-hidden">
                <Image 
                  src={s.image} 
                  alt={s.name} 
                  fill
                  className="object-cover scale-x-[-1]" 
                />

                <div className="absolute inset-0 bg-black/50 flex items-center justify-center md:block">
                  <h3 className="text-white text-center md:text-left md:absolute md:rotate-[-90deg] md:bottom-10 md:left-4 text-sm md:text-xl font-semibold">
                    {s.name}
                  </h3>
                </div>
              </div>
            )}
          </div>
        ))}
      </div>
    </section>
  )
}