package main

import (
	"github.com/wilztan/workshop/handler"
	"github.com/wilztan/workshop/db"
	"github.com/gin-gonic/gin"
)


func main(){
	db.Init("postgres://bykmyqwxvtendy:10915adb0b7c211a4a995a53687fe167bdfc2c4cbbbdbae158822b1cdbec940f@ec2-23-21-136-232.compute-1.amazonaws.com:5432/d4isopfja1mjvv")
	r:=gin.Default()
	r.GET("/api/product",handler.GetProductsHandler)
	r.POST("/api/product",handler.CreateProductHandler)
	r.GET("/api/news",handler.GetNews)
	r.Run(":3030")
}