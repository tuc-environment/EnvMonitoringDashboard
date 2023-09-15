import { DateTime } from 'luxon'

export const formatDisplayNumber = (num: number): string => {
  if (num < 1000) {
    return `${num}`
  } else if (num < 1000000) {
    return Intl.NumberFormat('en-US', { maximumFractionDigits: 1 }).format(num / 1000) + 'k'
  } else {
    return Intl.NumberFormat('en-US', { maximumFractionDigits: 1 }).format(num / 1000000) + 'm'
  }
}

export const formatDatetime = (date: string | Date): string => {
  if (typeof date === 'string') {
    date = new Date(date)
  }
  return DateTime.fromISO(date.toISOString())
    .setLocale('zh-cn')
    .toLocaleString(DateTime.DATETIME_MED)
}
