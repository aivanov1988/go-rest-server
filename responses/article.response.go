package responses

type ArticleListResponse []ArticleResponse

type ArticleResponse struct {
	Id          string
	Title       string
	Description string
}
