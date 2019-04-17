package main

import (
	"jobcan-fe/jobcan"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Message struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	SlackName string `json:"slack_name"`
}

func main() {
	r := gin.Default()
	r.AppEngine = true
	h := r.Group("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"result": "ok"})
	})
	h.GET("/_ah/health")

	j := r.Group("/jobcan")
	j.Use(msg)
	j.POST("/touch", func(c *gin.Context) {
		tmp, _ := c.Get("message")
		msg := tmp.(Message)
		err := jobcan.Touch(msg.Email, msg.Password)
		if err != nil {
			_ = c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"result": "ok"})
	})
	j.POST("/checkin", func(c *gin.Context) {
		tmp, _ := c.Get("message")
		msg := tmp.(*Message)
		err := jobcan.Touch(msg.Email, msg.Password, jobcan.CheckIn)
		if err != nil {
			_ = c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"result": "ok"})
	})
	j.POST("/checkout", func(c *gin.Context) {
		tmp, _ := c.Get("message")
		msg := tmp.(*Message)
		err := jobcan.Touch(msg.Email, msg.Password, jobcan.CheckOut)
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

func msg(c *gin.Context) {
	msg := &Message{}
	err := c.BindJSON(msg)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.Set("message", msg)
	c.Next()
}
