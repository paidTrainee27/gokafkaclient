package main

import (
	"fmt"
	"gokafkaclient/client"
	"gokafkaclient/internal/handler"
	"gokafkaclient/util"
	"log"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"

	"github.com/spf13/viper"
)

var (
	zAddr net.Addr
	// groupId string
	topic    string
	portHttp string
)

func init() {
	config, err := util.LoadConfig("../../.env")
	if err != nil {
		log.Fatal("Unable to load config file ", err)
	}
	bServer := fmt.Sprint(config.KafkaServer)
	zAddr = kafka.TCP(bServer)
	topic = fmt.Sprint(config.KafkaTopic)
	portHttp = fmt.Sprintf(":%s", viper.Get("PORT"))
}

func main() {
	// ctx := context.Background()

	kproducer := client.NewProducer(zAddr, topic)
	notifyUser := handler.NewNotifyUser(kproducer)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/uploadlogs", notifyUser.WriteMessage)

	r.Run(portHttp)
	//TODO: graceful shutdown
}
