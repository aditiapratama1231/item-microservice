package request

type Item struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"string"`
	IsFinish    bool   `json:"is_finish"`
}

type GetItemResponse struct {
	Data interface{} `json:"data"`
}

type CreateItemRequest struct {
	Data Item `json:"data"`
}

type ItemResponse struct {
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	StatusCode int32       `json:"status_code"`
	Err        bool        `json:"error,omitempty"`
}

type UpdateItemRequest struct {
	ID   string `json:"id"`
	Data Item   `json:"data"`
}

type ShowItemRequest struct {
	ID string `json:"id"`
}

type DeleteItemRequest struct {
	ID string `json:"id"`
}
