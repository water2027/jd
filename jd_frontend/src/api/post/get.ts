import RequestHandler from '../useRequest'

export interface GetPostsListRequest {
  id: number
  limit?: number
}

export interface Post {
  id: number
  author: {
    id: number
    name: string
  }
  title: string
  content: string
  cover: string
  views: number
  likes: number
  comments: number
  created_at: string
  update_at: string
}

export interface GetPostRequest {
  id: number
}

export type GetPostsResponse = Post[]

export const getMaxId = async (): Promise<number> => {
  return RequestHandler.get<number>('/public/posts')
}

export const getPostsList = async (gpr: GetPostsListRequest): Promise<GetPostsResponse> => {
  return RequestHandler.post<GetPostsResponse>('/public/posts', gpr)
}

export const getPost = async (gpr: GetPostRequest): Promise<Post> => {
  return RequestHandler.post<Post>('/public/posts-by-id', gpr)
}
