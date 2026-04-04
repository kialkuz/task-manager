package dto

type TaskRequest struct {
	ID      string `json:"id"`
	Date    string `json:"date" validate:"omitempty,taskdate"`
	Title   string `json:"title" validate:"required"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}
