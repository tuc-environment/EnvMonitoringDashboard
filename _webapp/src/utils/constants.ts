import { SensorSampleMethod } from '@/http-client'

export const airOptionNames = [
  '空气温度',
  '空气湿度',
  '大气压力',
  '天空短波辐射',
  '地球反射短波辐射',
  '天空长波辐射',
  '地球长波辐射',
  '四分量辐射仪传感器温度',
  '反照率',
  '净辐射',
  '风速',
  '风向',
  '雨量'
]
export const soilOptionNames = ['土壤水分', '土壤温度', '土壤电导率']
export const allOptionNames = airOptionNames.concat(soilOptionNames)
export const sensorSamplingMethods: SensorSampleMethod[] = [
  SensorSampleMethod.sampling,
  SensorSampleMethod.average,
  SensorSampleMethod.total,
  SensorSampleMethod.max,
  SensorSampleMethod.min
]
