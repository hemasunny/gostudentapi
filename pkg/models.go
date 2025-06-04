package pkg

import(
	"time"
)
type Student struct{
	ID int
    Name string
	Email string
	Password_hash string
	Role string
	Created_at time.Time
}