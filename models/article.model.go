package models

type Article struct {
	tableName   struct{} `pg:"articles"`
	Id          string   `json:"id"`
	Created_at  string   `json:"createdAt"`
	Updated_at  string   `json:"updatedAt"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	User_id     string   `json:"userId"`
	File_id     string   `json:"fileId"`
	File        File     `json:"file" pg:"rel:has-one"`
}
