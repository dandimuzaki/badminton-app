import api from "@/lib/api";
import { CourtQuery } from "@/types/court";
import { TimeslotQuery } from "@/types/timeslot";

export async function getAvailableTimeslots(params: TimeslotQuery) {
  const res = await api.get(`/timeslots/available`, {params});
  return res.data;
}

export async function getAvailableCourts(params: CourtQuery) {
  const res = await api.get(`/courts/available`, {params});
  return res.data;
}
