package main

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/hemasunny/gostudentapi/handlers"
)

func HomeHandler(c *gin.Context){
	c.JSON(http.StatusOK, "HOME PAGE")
}
func main(){
	  router := gin.Default()

	  router.GET("/",HomeHandler )

	  router.GET("/student",handlers.NewUserHandler )

	  router.Run()//by default port 8080
}