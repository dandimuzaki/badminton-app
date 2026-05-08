export interface Testimonial {
  id: number,
  title: string,
  description: string,
  name: string,
  type: string,
  profile: string
}

export const testimonials: Testimonial[] = [
  {
    id: 1,
    title: "Booking courts has never been this easy",
    description: "I reserved a slot in under a minute. The court was spotless and the lighting was perfect even for night games.",
    name: "Arif Pratama",
    type: "Weekend Player",
    profile: "/images/man-1.jpg"
  },
  {
    id: 2,
    title: "Perfect place for serious training",
    description: "The flooring quality and lighting really make a difference. Our team trains here every week now.",
    name: "Kevin Santoso",
    type: "Amateur Athlete",
    profile: "/images/man-2.jpg"
  },
  {
    id: 3,
    title: "Great for playing with friends!",
    description: "We booked a court for the weekend and had so much fun. The facilities are clean and the staff is super helpful.",
    name: "Nadia Putri",
    type: "Casual Player",
    profile: "/images/woman-1.jpg"
  },
  {
    id: 4,
    title: "Smooth experience from start to finish",
    description: "From booking to playing, everything just works. No hassle, no confusion, just good badminton.",
    name: "Rizky Maulana",
    type: "Office Worker",
    profile: "/images/man-3.jpg"
  },
  {
    id: 5,
    title: "Perfect for after-work games!",
    description: "Booking a court after office hours is super convenient. The lighting is excellent, and the whole place feels well-managed and welcoming.",
    name: "Andi Saputra",
    type: "Marketing Executive",
    profile: "/images/man-4.jpg"
  },
  {
    id: 6,
    title: "Feels like a professional court",
    description: "The court quality really surprised me. The flooring, spacing, and overall setup make every match feel more serious and enjoyable.",
    name: "Siti Mentari",
    type: "Amateur Player",
    profile: "/images/woman-2.jpg"
  },
]