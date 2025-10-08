package models

import (
	"time"
)

type User struct {
  ID    int64  `db:"id" json:"id"`
  Email  string `db:"email" json:"email"`
  FirstName  string `db:"first_name" json:"first_name"`
  LastName  string `db:"last_name" json:"last_name"`
  Mobile  int `db:"mobile" json:"mobile"`
  Password  string `db:"password" json:"password"`
	BillingCode BillingCode `db:"billing_code" json:"billing_code"`
	Months []Month
}


type Month struct {

  ID    int64  `db:"id" json:"id"`
	Year  uint16 `db:"year" json:"year"`
	Month uint8 `db:"month" json:"month"`
	User_id uint64`db:"user_id" json:"user_id"`
	Total_hours uint64`db:"total_hours" json:"total_hours"`
	Days []Day

}


type Day struct {

  ID    int64  `db:"id" json:"id"`
	Date  time.Time `db:"date" json:"date"`
	Month_id uint8 `db:"month_id" json:"month_id"`
	User_id uint64`db:"user_id" json:"user_id"`
	Total_hours uint64`db:"total_hours" json:"total_hours"`

}

type TimeEntry struct {

  ID    int64  `db:"id" json:"id"`
	User_id uint64 `db:"user_id" json:"user_id"`
	Day_id uint8 `db:"day_id" json:"day_id"`
	Activity_id uint8 `db:"activity_id" json:"activity_id"`
	Total_hours uint64`db:"total_hours" json:"total_hours"`
	Date  time.Time `db:"date" json:"date"`

}

type Activity struct {
  ID    int64  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}


type BillingCode struct {
  ID    int8  `db:"id" json:"id"`
	Number int64 `db:"number" json:"number"`
}
