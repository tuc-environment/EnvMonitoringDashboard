import axios from 'axios'

const HEADER_KEY_TOTAL_COUNT = 'x-total-count'
interface Response<T> {
  code: number
  error: string
  payload: T
  total: number | undefined
  status: string
}

interface Base {
  id: number
  created_at?: Date
  updated_at?: Date
  deleted_at?: Date
}

export interface Account extends Base {
  username: string
  token: string
}

export interface Station extends Base {
  name?: string
  lat?: number
  lng?: number
  altitude?: number
}

export enum SensorPosition {
  up = 'up',
  middle = 'middle',
  down = 'down'
}

export const getPositionName = (position: SensorPosition | undefined): string => {
  if (position) {
    switch (position) {
      case SensorPosition.up:
        return '板外点'
      case SensorPosition.middle:
        return '板间点'
      case SensorPosition.down:
        return '板下点'
    }
  }
  return ''
}

export enum SensorSampleMethod {
  sampling = 'sampling',
  average = 'average',
  total = 'total',
  max = 'max',
  min = 'min'
}

export const getSampleMethodDiplayText = (position: SensorSampleMethod | undefined): string => {
  if (position) {
    switch (position) {
      case SensorSampleMethod.sampling:
        return '采样值'
      case SensorSampleMethod.average:
        return '平均值'
      case SensorSampleMethod.total:
        return '总计值'
      case SensorSampleMethod.max:
        return '最大值'
      case SensorSampleMethod.min:
        return '最小值'
    }
  }
  return ''
}

export interface Sensor extends Base {
  station_id: number
  position?: SensorPosition
  tag?: string
  name?: string
  group?: string
  unit?: string
  sensor_code?: string
  sensor_report_code?: string
  sample_method?: SensorSampleMethod
  visible_in_dashboard?: boolean
}

export const getSensorDisplayText = (sensor: Sensor, stationName?: string): string => {
  let displayText = ''
  if (sensor.id) {
    displayText += `${sensor.id}. `
  }
  if (stationName) {
    displayText += `${stationName}: `
  }
  if (sensor.position) {
    displayText += `${getPositionName(sensor.position)}-`
  }
  if (sensor.group) {
    displayText += `${sensor.group}-`
  }
  if (sensor.name) {
    displayText += sensor.name
  }
  if (sensor.tag) {
    displayText += `(${sensor.tag})`
  }
  if (sensor.unit) {
    displayText += `-${sensor.unit}`
  }
  return displayText
}

export interface DataRecord extends Base {
  sensor_id: number
  value?: number
  time?: Date
  record_index?: number
}

export interface StationPrediction {
  down_soil_temp_shallow?: number
  down_soil_temp_deep?: number
  down_soil_water_content_shallow?: number
  down_soil_water_content_deep?: number
  down_temp?: number
  down_humidity?: number
  middle_soil_temp_shallow?: number
  middle_soil_temp_deep?: number
  middle_soil_water_content_shallow?: number
  middle_soil_water_content_deep?: number
  middle_temp?: number
  middle_humidity?: number
}

class HttpClient {
  public baseUrl: string = (function (): string {
    if (import.meta.env.VITE_API_ENDPOINT) return import.meta.env.VITE_API_ENDPOINT
    return import.meta.env.DEV ? 'http://localhost:8080/api' : '/api'
  })()

  private _token: string = ''

  public get token(): string {
    if (this._token) return this._token
    const token = localStorage.getItem('token')
    if (token) this.token = token
    return this._token
  }

  public set token(token: string) {
    this._token = token
    localStorage.setItem('token', token)
  }

  public absoluteUrl(url: string): string {
    if (url.startsWith('/')) {
      return this.baseUrl + url
    }
    return url
  }

  private _headers(): { [key: string]: string } {
    const headers: { [key: string]: string } = {}
    if (this._token) {
      headers['Authorization'] = `${this._token}`
    }
    return headers
  }

  // base
  private async get<T>(url: string): Promise<Response<T> | null> {
    url = this.absoluteUrl(url)
    try {
      const resp = await axios.get<Response<T>>(url, { headers: this._headers() })
      const count = resp.headers[HEADER_KEY_TOTAL_COUNT]
      console.log(JSON.stringify(resp.headers))
      return {
        ...resp.data,
        total: count ? parseInt(count) : undefined
      }
    } catch (err: any) {
      return err?.response.data
    }
  }

  private async post<T>(url: string, data: any = null): Promise<Response<T> | null> {
    url = this.absoluteUrl(url)
    try {
      const resp = await axios.post<Response<T>>(url, data, { headers: this._headers() })
      return resp?.data
    } catch (err: any) {
      return err?.response.data
    }
  }

  private async put<T>(url: string, data: any = null): Promise<Response<T> | null> {
    url = this.absoluteUrl(url)
    try {
      const resp = await axios.put<Response<T>>(url, data, { headers: this._headers() })
      return resp?.data
    } catch (err: any) {
      return err?.response.data
    }
  }

  private async delete<T>(url: string): Promise<Response<T> | null> {
    url = this.absoluteUrl(url)
    try {
      const resp = await axios.delete<Response<T>>(url, { headers: this._headers() })
      return resp?.data
    } catch (err: any) {
      return err?.response.data
    }
  }

  // auth

  public isAuthorized(): boolean {
    if (!this.token) return false
    return true
  }

  public async register(
    username: string,
    password: string
  ): Promise<Response<{ token: string }> | null> {
    const resp = await this.post<{ token: string }>('/register', { username, password: password })
    if (resp?.code === 200) {
      resp.payload.token && (this.token = resp.payload.token)
    }
    return resp
  }

  public async login(
    username: string,
    password: string
  ): Promise<Response<{ token: string }> | null> {
    const resp = await this.post<{ token: string }>('/login', { username, password: password })
    if (resp?.code === 200) {
      resp.payload.token && (this.token = resp.payload.token)
    }
    return resp
  }

  public async logout() {
    this.token = ''
  }

  public async getAccount(): Promise<Response<Account> | null> {
    const resp = await this.get<Account>('/account')
    return resp
  }

  public async regenrateToken(): Promise<Response<Account> | null> {
    const resp = await this.post<Account>('/account/regenrateToken')
    if (resp?.payload.token) {
      this.token = resp?.payload.token
    }
    return resp
  }

  public async changePassword(newPassword: string): Promise<Response<Account> | null> {
    const resp = await this.post<Account>('/account/changePassword', { new_password: newPassword })
    return resp
  }

  // stations

  public async getStations(params?: {
    offset?: number
    limit?: number
  }): Promise<Response<Station[]> | null> {
    const ret = []
    if (params?.offset != undefined) {
      ret.push(`offset=${params.offset}`)
    }
    if (params?.limit != undefined) {
      ret.push(`limit=${params.limit}`)
    }
    const resp = await this.get<Station[]>(`/stations?${ret.join('&')}`)
    return resp
  }

  public async upsertStation(params: {
    id?: number
    name?: string
    lat?: number
    lng?: number
    altitude?: number
  }): Promise<Response<Station> | null> {
    const resp = await this.post<Station>('/stations', params)
    return resp
  }

  public async deleteStation(stationId: number): Promise<number> {
    const resp = await this.delete<void>(`/stations/${stationId}`)
    return stationId
  }

  public async predictStation(params: {
    lat?: number
    lng?: number
    temp?: number
    humidity?: number
    barometric_pressure?: number
    soil_temp_shallow?: number
    soil_temp_deep?: number
    soil_water_content_shallow?: number
    soil_water_content_deep?: number
    soil_electrical_conductivity?: number
  }): Promise<Response<StationPrediction> | null> {
    const ret = []
    if (params?.lat != undefined) {
      ret.push(`lat=${params.lat}`)
    }
    if (params?.lng != undefined) {
      ret.push(`lng=${params.lng}`)
    }
    if (params?.temp != undefined) {
      ret.push(`temp=${params.temp}`)
    }
    if (params?.humidity != undefined) {
      ret.push(`humidity=${params.humidity}`)
    }
    if (params?.barometric_pressure != undefined) {
      ret.push(`barometric_pressure=${params.barometric_pressure}`)
    }
    if (params?.soil_temp_shallow != undefined) {
      ret.push(`soil_temp_shallow=${params.soil_temp_shallow}`)
    }
    if (params?.soil_temp_deep != undefined) {
      ret.push(`soil_temp_deep=${params.soil_temp_deep}`)
    }
    if (params?.soil_water_content_shallow != undefined) {
      ret.push(`soil_water_content_shallow=${params.soil_water_content_shallow}`)
    }
    if (params?.soil_water_content_deep != undefined) {
      ret.push(`soil_water_content_deep=${params.soil_water_content_deep}`)
    }
    if (params?.soil_electrical_conductivity != undefined) {
      ret.push(`soil_electrical_conductivity=${params.soil_electrical_conductivity}`)
    }
    const resp = await this.get<StationPrediction>(`/stations/predict?${ret.join('&')}`)
    return resp
  }

  // sensors

  public async getSensors(params?: {
    stationID?: number
    visibleInDashboard?: boolean
    offset?: number
    limit?: number
  }): Promise<Response<Sensor[]> | null> {
    const ret = []
    if (params?.offset) {
      ret.push(`offset=${params.offset}`)
    }
    if (params?.limit) {
      ret.push(`limit=${params.limit}`)
    }
    if (params?.stationID) {
      ret.push(`station_id=${params.stationID}`)
    }
    if (params?.visibleInDashboard) {
      ret.push(`visible_in_dashboard=${params.visibleInDashboard}`)
    }
    const resp = await this.get<Sensor[]>(`/sensors?${ret.join('&')}`)
    return resp
  }

  public async upsertSensor(params: {
    id?: number
    station_id: number
    position?: SensorPosition
    tag?: string
    name?: string
    group?: string
    unit?: string
  }): Promise<Response<Sensor> | null> {
    const resp = await this.post<Sensor>('/sensors', params)
    return resp
  }

  public async deleteSensor(sensorId: number): Promise<number> {
    const resp = await this.delete<void>(`/sensors/${sensorId}`)
    return sensorId
  }

  // records

  public async downloadTemplate(): Promise<string | null> {
    const url = this.absoluteUrl('/records/template')
    const resp = await axios.get<string>(url, { headers: this._headers() })
    return resp.data
  }

  public async uploadCSVRecords(data: any): Promise<Response<any> | null> {
    const resp = await this.post<any>('/records/upload', data)
    return resp
  }

  public async getRecords(params?: {
    sensorIDs?: number[]
    startTime?: Date
    endTime?: Date
    beforeCreatedAt?: Date
    afterCreatedAt?: Date
    offset?: number
    limit?: number
  }): Promise<Response<DataRecord[]> | null> {
    const ret = []
    if (params?.sensorIDs) {
      const str = params?.sensorIDs.map((sensorID) => sensorID.toString()).join(',')
      ret.push('sensor_ids=' + encodeURIComponent(str))
    }
    if (params?.startTime) {
      ret.push('start_time=' + encodeURIComponent(params.startTime.toISOString()))
    }
    if (params?.endTime) {
      ret.push('end_time=' + encodeURIComponent(params.endTime.toISOString()))
    }
    if (params?.beforeCreatedAt) {
      ret.push('before_created_at=' + encodeURIComponent(params.beforeCreatedAt.toISOString()))
    }
    if (params?.afterCreatedAt) {
      ret.push('after_created_at=' + encodeURIComponent(params.afterCreatedAt.toISOString()))
    }
    if (params?.offset) {
      ret.push(`offset=${params.offset}`)
    }
    if (params?.limit) {
      ret.push(`limit=${params.limit}`)
    }
    const resp = await this.get<DataRecord[]>(`/records?${ret.join('&')}`)
    return resp
  }

  public async updateRecords(records: DataRecord[]) {
    const resp = await this.post<void>('/records', records)
    return resp
  }
}

export default new HttpClient()
