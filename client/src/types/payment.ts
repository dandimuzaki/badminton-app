export interface Payment {
  id: number,
  user_id: number,
  amount: number,
  status: string,
  transaction_id: string
}

export interface PaymentRequest {
  reservation_id: number,
  amount: number
}