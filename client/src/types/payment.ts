export interface Payment {
  id: number,
  user_id: number,
  amount: number,
  status: string,
  transaction_id: string
}