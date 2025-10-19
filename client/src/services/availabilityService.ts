import api from "@/lib/api";

export async function getAvailableTimeslots(query: URLSearchParams) {
  const res = await api.get(`/available-timeslots?${query}`);
  return res.data;
}

export async function getAvailableCourts(query: URLSearchParams) {
  const res = await api.get(`/available-courts?${query}`);
  return res.data;
}
