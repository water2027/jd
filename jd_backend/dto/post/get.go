package post

type GetPostRequest struct {
	Limit  int `json:"limit"`
	LastId int `json:"id"`
}

type GetPostResponse []Post

func (gpr *GetPostRequest) Examine() error {
	if gpr.Limit <= 0 {
		gpr.Limit = 10
	}
	return nil
}


