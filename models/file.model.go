package models

type File struct {
	tableName     struct{}  `pg:"files"`
	Id            string    `json:"id"`
	Created_at    string    `json:"createdAt"`
	Updated_at    string    `json:"updatedAt"`
	Original_name string    `json:"originalName"`
	Mimetype      string    `json:"mimetype"`
	Size          int       `json:"size"`
	Bucket        string    `json:"bucket"`
	Key           string    `json:"key"`
	Articles      []Article `json:"-" pg:"rel:has-many"`
}
