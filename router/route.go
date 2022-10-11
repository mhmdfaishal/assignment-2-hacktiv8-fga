package router

import (
	"assignment-2/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Route(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	orderController := &controllers.OrderController{DB: db}

	router.POST("/orders", orderController.CreateOrder)
	router.GET("/orders", orderController.GetOrder)
	router.PUT("/orders/:orderId", orderController.UpdateOrder)
	router.DELETE("/orders/:orderId", orderController.DeleteOrder)
	
	return router
}
