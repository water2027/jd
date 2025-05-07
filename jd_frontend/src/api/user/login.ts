import RequestHandler from '../useRequest'

export interface LoginRequest {
  telephone: string
  password: string
}

export interface LoginResponse {
  id: number
  token: string
  name: string
}

export const login = async (lr: LoginRequest) => {
  return RequestHandler.post<LoginResponse>('/public/login', lr)
}
