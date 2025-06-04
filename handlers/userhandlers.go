package handlers

import (
	"net/http"
	"text/template/parse"
	"time"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hemasunny/gostudentapi/pkg"
	"github.com/rs/xid"
)

var users []pkg.User
func NewUserHandler(c *gin.Context)  {
	
	var user pkg.User

	if err:= c.ShouldBindJSON(&user); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"message":"Input data not readable"})
	}

	//go get "github.com/rs/xid"
	var err error
	if user.ID,err = strconv.ParseInt( xid.New().String(), 10, 64); err != nil{
         c.JSON(http.StatusInternalServerError, gin.H{"message":"unexpected error occured"})
	}

	user.Created_at= time.Now()
	users = append(users, user)

	
}

func GetUsers(c *gin.Context)  {

	if len(users)>0{
		c.JSON(http.StatusOK, users)
	}else{
		c.JSON(http.StatusNotFound, gin.H{"message":"no data exists"})
	}
	
}