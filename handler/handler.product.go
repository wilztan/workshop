package handler

import (
	"fmt"
	"encoding/json"
	"log"
	"strconv"
	"net/http"
	"github.com/wilztan/workshop/models"
	"github.com/wilztan/workshop/resource"

	"github.com/gin-gonic/gin"
)

func GetProductsHandler(c *gin.Context)  {

	var (
		limit,offset int64
		err error
	)

	limitStr := c.DefaultQuery("limit","5")
	offsetStr := c.DefaultQuery("offset","0")
	searchQuery:= c.DefaultQuery("search","")

	limit,err = strconv.ParseInt(limitStr,10,64)
	if err !=nil{
		limit=5
	}
	offset,err = strconv.ParseInt(offsetStr,10,64)
	if err !=nil{
		limit=0
	}

	products, err := resource.GetAllProduct(limit, offset, searchQuery)
	if err!=nil{
		data:=map[string]interface{}{
			"status":"failed",
			"reason":err.Error(),
		}
		c.JSON(http.StatusInternalServerError, data)
		return
	}

	data:= map[string]interface{}{
		"status":"ok",
		"data":products,
	}
	c.JSON(http.StatusOK,data)
	return
}

func CreateProductHandler(c *gin.Context){
	name:= c.DefaultPostForm("product_name","")
	desc:= c.DefaultPostForm("product_description","")

	product, err := resource.CreateProduct(name,desc)
	if err!=nil{
		data:=map[string]interface{}{
			"status":"failed",
			"reason":err.Error(),
		}
		c.JSON(http.StatusInternalServerError, data)
	}

	data:= map[string]interface{}{
		"status":"ok",
		"data":product,
	}
	c.JSON(http.StatusOK,data)
	return
}

func GetNews(c *gin.Context){
	news := models.News{}

    req, err := http.NewRequest("GET", "https://newsapi.org/v2/everything?q=bitcoin&from=2019-02-28&sortBy=publishedAt&apiKey=19215a15a83247d2901b26d017f6d7e3", nil)
    if err != nil {
        //handle error
        fmt.Println(err)
    }

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer resp.Body.Close()

    if err := json.NewDecoder(resp.Body).Decode(&news); err != nil {
        log.Println(err)
    }

    data := map[string]interface{}{
        "news": news,
    }

    c.JSON(200, data)

    return
}