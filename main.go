package main

import(
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomeHandler(c *gin.Context){
	c.JSON(http.StatusOK, "HOME PAGE")
}
func main(){
	  router := gin.Default()

	  router.GET("/",HomeHandler )

	  router.Run()//by default port 8080
}