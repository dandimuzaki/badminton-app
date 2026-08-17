export interface Court {
  ID: number,
  name: string,
  image_url: string,
  type: string,
  description: string,
  location: string,
  price: number
}

export interface CourtQuery {
  date: string,
  time_slot_id: number
}