import RequestHandler from '../useRequest'

export interface CreatePostsRequest {
  title: string
  content: string
}

export const createPost = async (gpr: CreatePostsRequest): Promise<null> => {
  return RequestHandler.post<null>('/post/create', gpr)
}
