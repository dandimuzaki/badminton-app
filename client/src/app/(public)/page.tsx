"use client";

import CourtCard from "@/components/CourtCard";
import SectionBadge from "@/components/home/section-badge";
import ServiceSection from "@/components/home/service";
import TestimonialCard from "@/components/home/testimonial-card";
import { Button } from "@/components/ui/button";
import { useAuth } from "@/context/AuthContext";
import { mockCourts } from "@/mocks/court";
import { testimonials } from "@/mocks/testimonial";
import { getCourts } from "@/services/courtService";
import { Court } from "@/types/court";
import { AlarmOn, Dashboard, LocalActivity, NearMe } from "@mui/icons-material";
import Image from "next/image";
import Link from "next/link";
import { useEffect, useState } from "react";

export default function Home() {
  const { token } = useAuth()
  const [courts, setCourts] = useState<Court[]>([])

  const fetchCourts = async () => {
    try {
      const res = await getCourts()
      setCourts(res)
    } catch (err: unknown) {
      console.error("Failed to fetch courts", err)
    }
  }
  useEffect(() => {
      fetchCourts()
  }, [])

  return (
    <>
      <section className="h-screen relative">
        <div className="top-0 z-0 absolute h-full w-full">
          <Image src={'/images/court-blue.jpg'} alt="ShuttleTime" width={900} height={1600} className="w-full h-full object-cover absolute top-0 scale-x-[-1]" />
          <div className="absolute z-2 bottom-0 h-full w-full bg-[linear-gradient(rgba(0,0,0,0.8),rgba(0,0,0,0),rgba(0,0,0,0),rgba(0,0,0,0.8))]"></div>
        </div>
        <div className="p-8 md:p-8 md:p-16 z-15 flex flex-col gap-4 md:justify-between relative w-full h-full justify-center items-center md:items-start text-center md:text-left">
          <div></div>
          <div className="md:w-[60%]">
            <h1 className="text-white text-3xl md:text-6xl/15 mb-4">Reserve Your Badminton<br/>Court in Just Seconds</h1>
            <p className="text-white text-base/5 md:text-lg mb-6">Book badminton courts anytime, anywhere — no waiting, no hassle. Just pick a court, grab your racket, and play your best game.</p>
            <div className="flex gap-2 items-center md:justify-start justify-center">
              <Link href="/reservations/new">
                <Button className="hover:bg-secondary hover:text-white py-2 text-secondary rounded-full bg-background px-4 text-base font-semibold">Book Now</Button>
              </Link>
              <Link href="/reservations/new">
                <Button className="hover:bg-primary hover:border-primary hover:text-secondary border border-white py-2 text-white rounded-full bg-transparent px-4 text-base font-semibold">Book Now</Button>
              </Link>
            </div>
          </div>
        </div>
      </section>
      <section className="p-8 md:p-16 bg-background">
        <div className="text-center text-2xl md:text-4xl/12 font-medium text-secondary wrap">
          At <p className="font-semibold inline"><span className="text-primary">Shuttle</span>Time</p>, 
          we aim to make badminton 
          <span className="inline-flex px-3 items-center align-middle relative">
            <img src={"/images/court-hero.jpg"} width={60} height={20} className="rounded md:rounded-lg w-16 md:w-24 md:h-10 h-6" />
          </span>
          accessible and enjoyable for everyone.
          Discover the best 
          <span className="inline-flex px-3 items-center align-middle relative">
            <img src={"/images/court-racket.jpg"} width={60} height={20} className="rounded md:rounded-lg w-16 md:w-24 md:h-10 h-6" />
          </span>
          courts available, tailored to your
          <span className="inline-flex px-3 items-center align-middle relative">
            <img src={"/images/indoor-cement.jpg"} width={60} height={20} className="rounded md:rounded-lg w-16 md:w-24 md:h-10 h-6" />
          </span>
          schedule and preferred location, ensuring a fantastic experience every time
        </div>
      </section>
      <ServiceSection/>
      <section id="about" className="p-8 md:p-16 flex flex-col items-center">
        <SectionBadge text="Our Courts"/>
        <h2 className="text-center text-2xl md:text-4xl font-bold tetx-secondary mb-4">Courts Ready For Game</h2>
        <p className="mb-8 text-center text-sm md:text-base" onClick={() => console.log(courts)}>Indoor or outdoor, wooden or synthetic, find the court that matches your style and start playing instantly.</p>
        <div className="grid lg:grid-cols-3 md:grid-cols-2 md:gap-6 gap-4">
          {courts?.map((c) => <CourtCard key={c.id} court={c} />)}
        </div>
      </section>
      <section className="p-8 md:p-16 flex flex-col items-center bg-white">
        <SectionBadge text="Our Values"/>
        <h1 className="text-center text-4xl font-bold mb-4">Play Smarter, Not Harder</h1>
        <p className="mb-8 text-center">From instant booking to smart scheduling, we make every game effortless</p>
            <div className="grid grid-cols-2 md:grid-cols-4 gap-6">
              <div className="w-full flex flex-col gap-2">
                <div className="text-5xl mb-2 text-secondary"><LocalActivity fontSize="inherit"/></div>
                <p className="font-semibold text-lg md:text-xl text-secondary">Seamless Booking</p>
                <p>Find, book, and confirm courts with just a few taps</p>
              </div>
              <div className="w-full flex flex-col gap-2">
                <div className="text-5xl mb-2 text-secondary"><NearMe fontSize="inherit"/></div>
                <p className="font-semibold text-lg md:text-xl text-secondary">Trusted Venues</p>
                <p>Verified courts with clear pricing and facility info</p>
              </div>
              <div className="w-full flex flex-col gap-2">
                <div className="text-5xl mb-2 text-secondary"><Dashboard fontSize="inherit"/></div>
                <p className="font-semibold text-lg md:text-xl text-secondary">Stay Organized</p>
                <p>Manage all your bookings in one simple dashboard</p>
              </div>
              <div className="w-full flex flex-col gap-2">
                <div className="text-5xl mb-2 text-secondary"><AlarmOn fontSize="inherit"/></div>
                <p className="font-semibold text-lg md:text-xl text-secondary">Game Ready</p>
                <p>Notifications and reminders keep your playtime on track</p>
              </div>
            </div>
      </section>
      <section className="relative">
        <div className="absolute top-0 z-0 h-full w-full">
          <Image src={'/images/court-hero.jpg'} alt="ShuttleTime" width={900} height={1600} className="w-full h-full object-cover scale-x-[-1]" />
          <div className="absolute z-2 bottom-0 h-full w-full bg-[linear-gradient(rgba(0,0,0,0),rgba(0,0,0,0.8))]"></div>
        </div>
        <div className="text-white relative z-20 text-center px-8 py-12 md:px-16 md:py-24 flex flex-col justify-center items-center gap-2 md:gap-4">
          <h2 className="text-2xl md:text-4xl font-semibold">Book your court today and enjoy the game!</h2>
          <p className="text-sm md:text-base">Secure your preferred time slot at our top-rated badminton courts.<br/>Fast, easy, and hassile-free booking</p>
          <Button className="w-fit rounded-full">Book Now</Button>          
        </div>
      </section>
      <section id="about" className="p-8 md:p-16 flex flex-col items-center">
        <SectionBadge text="Testimonial"/>
        <h2 className="text-center text-2xl md:text-4xl font-bold tetx-secondary mb-4">What Players Are Saying</h2>
        <p className="mb-8 text-sm md:text-base text-center">Sportive stories from players who’ve already stepped onto the court and loved every rally.</p>
          <div className="relative w-full h-full">
          <div className="overflow-hidden w-full h-full">
            <div className="px-6 md:p-0 flex flex-nowrap md:grid lg:grid-cols-3 md:grid-cols-2 md:gap-6 gap-4 w-full overflow-y-scroll">
              {testimonials.map((t) => <TestimonialCard key={t.id} testimonial={t} />)}
            </div>
          </div>
          <div className="md:hidden absolute top-0 left-[-1px] h-full w-8 bg-gradient-to-r from-background to-transparent"></div>
          <div className="md:hidden absolute top-0 right-[-1px] h-full w-8 bg-gradient-to-l from-background to-transparent"></div>
          </div>
      </section>
    </>
  );
}
