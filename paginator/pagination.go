package paginator

// Domain layer
type Pagination struct {
	Page     int
	PageSize int
	Sort     []SortOption
}

type SortOrder string

const (
	Asc  SortOrder = "asc"
	Desc SortOrder = "desc"
)

type SortOption struct {
	Field string
	Order SortOrder
}

type Filter struct {
	Field    string
	Operator string
	Value    interface{}
}

// Application layer
type PageResult[T any] struct {
	Data       []T   `json:"data"`       // Holds the list of items
	Page       int   `json:"page"`       // Current page number
	Size       int   `json:"size"`       // Number of items per page
	TotalPages int   `json:"totalPages"` // Total pages available
	TotalItems int64 `json:"totalItems"` // Total items in the dataset
}

func NewPageResult[T any](data []T, page, size int, totalItems int64) PageResult[T] {
	totalPages := int((totalItems + int64(size) - 1) / int64(size))
	return PageResult[T]{
		Data:       data,
		Page:       page,
		Size:       size,
		TotalPages: totalPages,
		TotalItems: totalItems,
	}
}
