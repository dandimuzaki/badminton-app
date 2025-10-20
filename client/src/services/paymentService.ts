import api from "@/lib/api";

export async function createPayment(data: object) {
  const res = await api.post("/payments/create", data);
  return res.data;
}