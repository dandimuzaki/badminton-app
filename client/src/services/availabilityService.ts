import api from "@/lib/api";

export async function getAvailableTimeslots(query: string) {
  console.log("query", query)
  const res = await api.get(`/available-timeslots?${query}`);
  console.log("res", res)
  return res.data;
}

export async function getAvailableCourts(query: URLSearchParams) {
  const res = await api.get(`/available-courts?${query}`);
  return res.data;
}
