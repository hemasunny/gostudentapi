package pkg

import(
	"time"
)
type User struct{
	ID int64
    Name string
	Email string
	Password_hash string
	Role string
	Created_at time.Time
}