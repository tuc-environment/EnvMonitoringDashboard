import axios from 'axios'

interface Response<T> {
  code: number
  error: string
  payload: T
  status: string
}

interface Base {
  id: number
  createdAt?: Date
  updatedAt?: Date
  deletedAt?: Date
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

export interface Sensor extends Base {
  stationID: number
  position?: SensorPosition
  tag?: string
  name?: string
  group?: string
  unit?: string
}

export interface DataRecord extends Base {
  sensorID: number
  value?: number
  time?: Date
}

class HttpClient {
  public baseUrl: string = (function (): string {
    if (import.meta.env.VITE_API_ENDPOINT) return import.meta.env.VITE_API_ENDPOINT
    return import.meta.env.DEV
      ? 'http://localhost:8080/api'
      : 'https://tuc-env-monitoring-dashboard.vercel.app/api'
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
      return resp?.data
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
      alert('Registration successful. Please login to continue.')
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

  // stations

  public async getStations(): Promise<Response<Station[]> | null> {
    const resp = await this.get<Station[]>('/stations')
    return resp
  }

  public async upsertStation(station: Station): Promise<Response<Station> | null> {
    const resp = await this.post<Station>('/stations')
    return resp
  }

  // sensors

  public async getSensors(stationID: string): Promise<Response<Sensor[]> | null> {
    const resp = await this.get<Sensor[]>(`/sensors?station_id=${stationID}`)
    return resp
  }

  public async upsertSensor(sensor: Sensor): Promise<Response<Sensor> | null> {
    const resp = await this.post<Sensor>('/sensors', sensor)
    return resp
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

  public async getRecords(
    sensorIDs?: number[],
    startTime?: Date,
    endTime?: Date,
    offset?: number,
    limit?: number
  ): Promise<Response<DataRecord[]> | null> {
    const ret = []
    if (sensorIDs) {
      const str = sensorIDs.map((sensorID) => sensorID.toString()).join(',')
      ret.push('sensor_ids=' + encodeURIComponent(str))
    }
    if (startTime) {
      ret.push('start_time=' + encodeURIComponent(startTime.toISOString()))
    }
    if (endTime) {
      ret.push('end_time=' + encodeURIComponent(endTime.toISOString()))
    }
    if (offset != undefined) {
      ret.push(`offset=${offset}`)
    }
    if (limit != undefined) {
      ret.push(`limit=${limit}`)
    }
    const resp = await this.get<DataRecord[]>(`/records?${ret.join('&')}`)
    return resp
  }
}

export default new HttpClient()
