package handler

import (
	"fmt"
	"gokafkaclient/client"
	"gokafkaclient/internal/model"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type INottifyUser interface {
	WriteMessage(context *gin.Context)
}

type notifyUser struct {
	kafkaWriter client.Iproducer
}

func NewNotifyUser(prod client.Iproducer) INottifyUser {
	return &notifyUser{
		kafkaWriter: prod,
	}
}

func (n *notifyUser) WriteMessage(c *gin.Context) {
	defer c.Request.Body.Close()

	// log := model.Logs{}

	d, err := ioutil.ReadAll(c.Request.Body)

	// if err := c.ShouldBindBodyWith(&log, binding.JSON);
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	fmt.Println(string(d))
	// d, _ := json.Marshal(log)

	err = n.kafkaWriter.Write(c, d)

	if err != nil {
		writeResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	writeResponse(c, http.StatusCreated, "Message posted successfully")
}

func writeResponse(c *gin.Context, code int, msg string) {
	c.JSON(code,
		model.Response{
			StatusCode: code,
			Message:    msg,
		})
}
