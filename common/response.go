package common

type GeneralResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Meta    interface{} `json:"meta,omitempty"`
	Cause   interface{} `json:"cause,omitempty"`
}

type PageMetaResponse struct {
	Current      int   `json:"current"`
	TotalPages   int   `json:"totalPages"`
	PerPage      int   `json:"perPage"`
	TotalRecords int64 `json:"totalRecords"`
}

const (
	SuccessMessage = "success"
)
