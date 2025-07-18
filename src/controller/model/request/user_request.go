package request

type UserRequest struct {
	Email    string `json:"id" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,containsAny=#%&!@$"`
	Name     string `json:"name" binding:"required,min=3"`
	Age      int    `json:"age" binding:"required,min=1,max=150"`
}
