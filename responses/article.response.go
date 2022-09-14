package responses

type ArticleListResponse []ArticleResponse

type ArticleResponse struct {
	Id          string
	CreatedAt   string
	UpdatedAt   string
	Title       string
	Description string
	Preview     string
	IsLiked     bool
}
