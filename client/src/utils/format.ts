export const formatRupiah = (value: number): string => {
  return new Intl.NumberFormat("id-ID", {
    style: "currency",
    currency: "IDR",
    minimumFractionDigits: 0,
  }).format(value)
}

export function formatDate(date?: Date): string {
  if (!date) return ""
  const year = date.getFullYear()
  const month = String(date.getMonth() + 1).padStart(2, "0")
  const day = String(date.getDate()).padStart(2, "0")

  return `${year}-${month}-${day}`
}

export function formatDateStr(dateStr: string): string {
  const date = new Date(dateStr)
  return date.toLocaleDateString(
    'en-US',
    {
      year: 'numeric',
      month: 'long',
      day: 'numeric'
    }
  )
}

export const getFirstName = (value: string): string => {
  const arr = value.split(" ")
  return arr[0]
}

export const capitalize = (str: string): string => {
  return str.charAt(0).toUpperCase() + str.slice(1);
}