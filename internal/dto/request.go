
package dto


type User struct {
  ID    int64  `db:"id" json:"id"`
  FullName  string `db:"full_name" json:"full_name"`
}
