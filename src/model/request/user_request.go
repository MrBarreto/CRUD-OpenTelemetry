package request

type UserRequest struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
}
