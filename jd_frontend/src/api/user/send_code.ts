import RequestHandler from '../useRequest'

export interface SendCodeRequest {
  telephone: string
}

export const sendCode = async (lr: SendCodeRequest) => {
  return RequestHandler.post<null>('/public/login', lr)
}
