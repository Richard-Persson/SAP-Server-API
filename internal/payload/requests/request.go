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

type TimeEntryRequest struct {
	UserID         	 int64      `db:"user_id" json:"user_id"`
	ActivityID      *int64     `db:"activity_id" json:"activity_id,omitempty"`
	Date 						 string `db:"date" json:"date"`
	StartTime 			string `db:"start_time" json:"start_time"`
	EndTime          string `db:"end_time" json:"end_time"`
}

type UpdateTimeEntryRequest struct {
	ID   						int64  `db:"id" json:"id"`
	ActivityID      *int64     `db:"activity_id" json:"activity_id,omitempty"`
	Date 						 string `db:"date" json:"date"`
	StartTime 			string `db:"start_time" json:"start_time"`
	EndTime          string `db:"end_time" json:"end_time"`
}
