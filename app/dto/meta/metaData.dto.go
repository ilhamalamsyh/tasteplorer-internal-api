package metadata_dto

type MetaData struct {
	CurrentPage int `json:"currentPage"`
	PageSize    int `json:"pageSize"`
	Total       int `json:"total"`
	TotalPage   int `json:"totalPage"`
}
