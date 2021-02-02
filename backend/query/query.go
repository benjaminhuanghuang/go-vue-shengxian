package query

type ListQuery struct {
	Limit int64 `json:"limit"`
	page  int64 `json:"page"`
}
