package contracts

type ApiErrorResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}
