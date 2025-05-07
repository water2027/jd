import RequestHandler from '../useRequest'

export interface RegisterRequest {
  name: string
  telephone: string
  password: string
  password2: string
  vcode: string
}

export interface RegisterResponse {
  id: number
  token: string
  name: string
}

export const register = async (rr: RegisterRequest) => {
  return RequestHandler.post<RegisterResponse>('/public/login', rr)
}
