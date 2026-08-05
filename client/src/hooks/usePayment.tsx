import { createPayment } from "@/services/paymentService";
import { PaymentRequest } from "@/types/payment";

export default function usePayment() {
  const handlePayment = async (dataPayment: PaymentRequest) => {
    const payment = await createPayment(dataPayment);

    if (typeof window.snap === "undefined") {
      alert("Midtrans Snap is not loaded yet. Please refresh.");
      return;
    }

    window.snap.pay(payment.snap_token, {
      onSuccess: async function (result) {
        console.log("✅ Payment success:", result);
        alert("Payment successful!");
      },
      onPending: function (result) {
        console.log("⏳ Payment pending:", result);
        alert("Payment pending. Please wait for confirmation.");
      },
      onError: function (result) {
        console.error("❌ Payment error:", result);
        alert("Payment failed. Please try again.");
      },
      onClose: function () {
        console.log("⚠️ Payment popup closed before finishing");
      },
    });
  }

  return {
    handlePayment
  }
}