import api from "@/lib/api";

export async function createReservation(data: object) {
  const res = await api.post("/reservations", data);
  return res.data;
}

export async function getReservations() {
  const res = await api.get("/reservations");
  return res.data;
}