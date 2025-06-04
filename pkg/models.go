package pkg

import(
	"time"
)
type User struct{
	ID int
    Name string
	Email string
	Password_hash string
	Role string
	Created_at time.Time
}