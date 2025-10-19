import { useState } from "react";
import axios from "axios";

export default function PaymentPage() {
  const [loading, setLoading] = useState(false);

  const handlePay = async () => {
    setLoading(true);
    try {
      const res = await axios.post("http://localhost:8080/api/payment", {
        reservationId: 123,
        amount: 50000,
        customerName: "Dandi",
        customerEmail: "dandi@example.com",
      });

      const { snapToken } = res.data;

      window.snap.pay(snapToken, {
        onSuccess: function(result: any) {
          alert("Payment success!");
          console.log(result);
        },
        onPending: function(result: any) {
          alert("Waiting for payment...");
          console.log(result);
        },
        onError: function(result: any) {
          alert("Payment failed!");
          console.log(result);
        },
      });
    } catch (err) {
      console.error(err);
      alert("Failed to initiate payment");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="p-6">
      <h2 className="text-xl font-semibold mb-4">Mock Payment</h2>
      <button
        onClick={handlePay}
        disabled={loading}
        className="bg-green-500 text-white px-4 py-2 rounded"
      >
        {loading ? "Processing..." : "Pay with Midtrans"}
      </button>
    </div>
  );
}
