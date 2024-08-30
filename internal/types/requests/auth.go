package requests

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email,gte=2,lte=40"`
	Password string `json:"password" binding:"required,gte=6,lte=40"`
	// Role string `json:"role" binding:"oneof= 'ADMIN' 'USER'"`
}
