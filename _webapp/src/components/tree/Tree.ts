import type { Sensor, Station } from '@/httpclient'

export interface Tree {
  label: string
  children: Tree[]
  sensor: Sensor | null
  station: Station | null
}
