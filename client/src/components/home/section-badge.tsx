"use client"

export default function SectionBadge ({text}: {text: string}) {
  return (
    <div className="rounded-full text-secondary bg-secondary/20 px-3 py-1 w-fit mb-4">
      <p className="uppercase text-xs md:text-sm">{text}</p>
    </div>
  )
}