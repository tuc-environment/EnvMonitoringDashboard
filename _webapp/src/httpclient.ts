import axios from 'axios'
import md5 from 'md5'

interface Response<T> {
  code: number
  error: string
  payload: T
  status: string
}

class HttpClient {
  public baseUrl: string = 'http://localhost:8080/api'

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

  public async get<T>(url: string): Promise<Response<T> | null> {
    url = this.absoluteUrl(url)
    const resp = await axios.get<Response<T>>(url, { headers: this._headers() })
    return resp?.data
  }

  public async post<T>(url: string, data: any = null): Promise<Response<T> | null> {
    url = this.absoluteUrl(url)
    try {
      const resp = await axios.post<Response<T>>(url, data, { headers: this._headers() })
      return resp?.data
    } catch (err: any) {
      return err?.response.data
    }
  }

  public async put<T>(url: string, data: any = null): Promise<Response<T> | null> {
    url = this.absoluteUrl(url)
    const resp = await axios.put<Response<T>>(url, data, { headers: this._headers() })
    return resp?.data
  }

  public async delete<T>(url: string): Promise<Response<T> | null> {
    url = this.absoluteUrl(url)
    const resp = await axios.delete<Response<T>>(url, { headers: this._headers() })
    return resp?.data
  }

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
}

export default new HttpClient()
