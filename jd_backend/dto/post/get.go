package post

type GetPostRequest struct {
	Id int `json:"id"`
}

type GetPostResponse Post

func (gpr *GetPostRequest) Examine() error {
	return nil
}
