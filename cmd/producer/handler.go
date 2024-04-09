package producer

import (
	"distributed-queue/client"
	"distributed-queue/internal/boot"
	"distributed-queue/internal/entity"
	"distributed-queue/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ProducerHandlerFunc(ctx *gin.Context) {
	topic := ctx.Param("topic")
	var request map[string]interface{}

	if err := ctx.BindJSON(&request); err != nil {
		// Handle client error
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":  "BAD_REQUEST",
			"error": "Invalid request payload",
		})
		return
	}
	producer := client.NewProducer().Broker(boot.GetAppConfig().Hostname, boot.GetAppConfig().Port).Topic(topic).Build()
	payload, err := createPayload(topic, request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":  "BAD_REQUEST",
			"error": "Invalid request payload",
		})
	}
	if err = producer.Publish(payload); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":  "INTERNAL_SERVER_ERROR",
			"error": "Something went wrong!",
		})
	}
	ctx.JSON(http.StatusCreated, map[string]interface{}{
		"status": "SUCCESS",
	})
}

func createPayload(topic string, request map[string]interface{}) (entity.IPayload, error) {
	payload := entity.NewPayload(topic, request["key"].(string), request["data"], utils.String)
	return payload, nil
}
