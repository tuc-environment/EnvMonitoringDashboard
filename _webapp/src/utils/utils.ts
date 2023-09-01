export const formatDisplayNumber = (num: number): string => {
  if (num < 1000) {
    return `${num}`
  } else if (num < 1000000) {
    return Intl.NumberFormat('en-US', { maximumFractionDigits: 2 }).format(num / 1000) + 'k'
  } else {
    return Intl.NumberFormat('en-US', { maximumFractionDigits: 2 }).format(num / 1000000) + 'm'
  }
}
