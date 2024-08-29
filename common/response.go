package common

type CommonResponse struct {
	Data interface{}
}

type BaseResponse[T any] struct {
	Data    T           `json:"data"`
	Message string      `json:"message"`
	Meta    interface{} `json:"meta,omitempty"`
	Cause   interface{} `json:"cause,omitempty"`
	Status  bool        `json:"status"`
}

type PageMetaResponse struct {
	Current      int   `json:"current"`
	TotalPages   int   `json:"totalPages"`
	PerPage      int   `json:"perPage"`
	TotalRecords int64 `json:"totalRecords"`
}

func NewBaseResponse[T interface{}](message string, data T, meta interface{}) *BaseResponse[T] {
	return &BaseResponse[T]{
		Message: message,
		Data:    data,
		Meta:    meta,
		Status:  true,
	}
}

func NewErrorResponse(message string, cause interface{}) *BaseResponse[any] {
	return &BaseResponse[any]{
		Message: message,
		Data:    nil,
		Status:  false,
		Cause:   cause,
	}
}
