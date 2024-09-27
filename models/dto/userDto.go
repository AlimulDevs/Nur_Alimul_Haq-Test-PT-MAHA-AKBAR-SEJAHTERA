package dto

type UserResponse struct {
	Name     string `json:"name"`
	Age      int16  `json:"age"`
	JobTitle string `json:"job_title"`
	Company  string `json:"company"`
}

type UserRequest struct {
	Name     string `json:"name"`
	Age      int16  `json:"age"`
	JobTitle string `json:"job_title"`
	Company  string `json:"company"`
}
