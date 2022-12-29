package responses

type ArticleListResponse []ArticleResponse

type ArticleResponse struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	FileName    string `json:"file.originalName"`
}
