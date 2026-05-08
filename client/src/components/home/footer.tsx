"use client"
import { Audiotrack, Email, Instagram, LinkedIn, LocationOn, Twitter, WhatsApp, X } from '@mui/icons-material';
import { Button } from '../ui/button';
import Link from 'next/link';

const Footer = () => {
  return (
      <div className='relative md:grid gap-4 px-16 py-8 grid-cols-[2fr_1fr_1fr]'>
        <div className=''>
          <div className='flex justify-center items-center w-48'><img src={"/images/shuttle-time-logo-trp.png"} className="w-full" /></div>
          <p className='text-sm/5'>Book badminton courts anytime, anywhere — no waiting, no hassle. Just pick a court, grab your racket, and play your best game.</p>
          <div className='mt-4 flex gap-2'>
            <Button className="aspect-square rounded-full h-full" ><Instagram fontSize='small'/></Button>
            <Button className="aspect-square rounded-full h-full"><X fontSize='small'/></Button>
            <Button className="aspect-square rounded-full h-full"><Audiotrack fontSize='small'/></Button>
            <Button className="aspect-square rounded-full h-full"><LinkedIn fontSize='small'/></Button>
          </div>
        </div>
        <div className='space-y-2'>
          <p className='font-bold text-secondary'>Quick Links</p>
          <div className="text-sm flex flex-col gap-1">
            <Link href={"/"} className='hover:text-secondary'>Home</Link>
            <Link href={"/products"} className='hover:text-secondary'>Browse Products</Link>
            <Link href={"/orders"} className='hover:text-secondary'>Track Your Order</Link>
            <Link href={"/about"} className='hover:text-secondary'>About Us</Link>
          </div>
        </div>
        <div className='space-y-2'>
          <p className='text-secondary font-bold'>Get in Touch</p>
          <a className='flex gap-2 text-sm cursor-pointer items-center'><WhatsApp />+62 853-2409-1088</a>
          <a className='flex gap-2 text-sm cursor-pointer items-center'><Email />agripacul@itb.lpik.org</a>
          <a className='flex gap-2 text-sm cursor-pointer'><LocationOn/>Jl. Cisintok Kadumulya, Cihanjuang, Kec. Parongpong, Kabupaten Bandung Barat</a>
        </div>
      </div>
  );
};

export default Footer;