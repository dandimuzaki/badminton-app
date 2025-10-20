export const formatRupiah = (value: number): string => {
  return new Intl.NumberFormat("id-ID", {
    style: "currency",
    currency: "IDR",
    minimumFractionDigits: 0,
  }).format(value)
}

export const formatDate = (value: string): string => {
  const date = new Date(value)
  return date.toLocaleDateString('en-US', {
    day: 'numeric',
    month: 'short',
    year: 'numeric'
  })
}

export const getFirstName = (value: string): string => {
  const arr = value.split(" ")
  return arr[0]
}