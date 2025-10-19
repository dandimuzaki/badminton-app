import api from "@/lib/api";

export async function getCourts() {
  const res = await api.get("/courts");
  return res.data;
}