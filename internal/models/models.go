package models




type BillingCode struct {
    ID     int64 `db:"id" json:"id"`
    Salary int   `db:"salary" json:"salary"`
}

type User struct {
    ID            int64        `db:"id" json:"id"`
    Email         string       `db:"email" json:"email"`
    FirstName     string       `db:"first_name" json:"first_name"`
    LastName      string       `db:"last_name" json:"last_name"`
    Mobile        string       `db:"mobile" json:"mobile"`
    Password      string       `db:"password" json:"password"`
		BillingCodeId int          `db:"billing_code_id" json:"billing_code_id"` 
		Entries 			*[]TimeEntry  `json:"entries"` 

}

type Activity struct {
    ID   int64  `db:"id" json:"id"`
    Name string `db:"name" json:"name"`
}

type TimeEntry struct {
	ID   						int64  `db:"id" json:"id"`
	UserID         	 int64      `db:"user_id" json:"user_id"`
	ActivityID      *int64     `db:"activity_id" json:"activity_id,omitempty"`
  Date       			string `json:"date" format:"2006-01-02"`
	StartTime 			string `db:"start_time" json:"start_time"`
	EndTime          string `db:"end_time" json:"end_time"`
	TotalHours	 float32   `db:"total_hours" json:"total_hours"`
}

type UserMonthHours struct {
	ID   						int64  `db:"id" json:"id"`
	UserID       int64 `db:"user_id" json:"user_id"`
	Year         int16 `db:"year" json:"year"`
	Month        int16 `db:"month" json:"month"`
	TotalHours	 float32   `db:"total_hours" json:"total_hours"`
}

type Day struct {
	ID   			int64  `db:"id" json:"id"`
	Date     	  	string `json:"date" format:"2006-01-02"`
	UserID         	int64      `db:"user_id" json:"user_id"`
	TotalHours	 	float32   `db:"total_hours" json:"total_hours"`
	TimeEntries     []TimeEntry
}
