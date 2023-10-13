package model

type Todo struct {
	ID        string `json:"_id,omitempty" bson:"_id,omitempty"`
	Aid       string `json:"aid"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
