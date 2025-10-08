

package requests


type RegisterRequest struct {
  Email  string `db:"email" json:"email"`
  FirstName  string `db:"first_name" json:"first_name"`
  LastName  string `db:"last_name" json:"last_name"`
  Mobile  string `db:"mobile" json:"mobile"`
  Password  string `db:"password" json:"password"`
}

type LoginRequest struct {
  Email    string `json:"email" binding:"required,email"`
  Password string `json:"password" binding:"required"`
}
