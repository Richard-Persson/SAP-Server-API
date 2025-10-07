
package models

type User struct {
  ID    int64  `db:"id" json:"id"`
  Email  string `db:"email" json:"email"`
  FirstName  string `db:"first_name" json:"first_name"`
  LastName  string `db:"last_name" json:"last_name"`
  Mobile  int `db:"mobile" json:"mobile"`
  Password  string `db:"password" json:"password"`
}
