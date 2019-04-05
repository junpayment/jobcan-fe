package main

import (
	"jobcan-fe/jobcan"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Message struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	r := gin.Default()
	r.AppEngine = true
	h := r.Group("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"result": "ok"})
	})
	h.GET("/_ah/health")

	j := r.Group("/jobcan")
	j.POST("/touch", func(c *gin.Context) {
		msg := &Message{}
		err := c.BindJSON(msg)
		if err != nil {
			_ = c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		err = jobcan.Touch(msg.Email, msg.Password)
		if err != nil {
			_ = c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"result": "ok"})
	})

	err := r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}
