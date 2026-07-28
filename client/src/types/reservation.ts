import { Court } from "./court";
import { Timeslot } from "./timeslot";
import { Payment } from "./payment";
import { User } from "./user";

export interface Reservation {
  id: number,
  user_id: number,
  court_id: number,
  date: string,
  time_slot_id: number,
  status: string,

  user: User,
  court: Court,
  timeslot: Timeslot,
  payment: Payment
}