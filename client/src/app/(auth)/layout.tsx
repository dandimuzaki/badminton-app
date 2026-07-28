'use client'

export default function AuthLayout({ children }: { children: React.ReactNode }) {
  return (
    <main className="py-16 min-h-screen bg-primary-dark flex flex-col items-center justify-center">
      {children}
    </main>
  );
}
