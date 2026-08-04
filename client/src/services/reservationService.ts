import api from "@/lib/api";
import { ReservationRequest } from "@/types/reservation";

export async function createReservation(data: ReservationRequest) {
  const res = await api.post("/reservations", data);
  return res.data;
}

export async function getReservations() {
  const res = await api.get("/reservations");
  return res.data;
}