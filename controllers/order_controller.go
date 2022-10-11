package controllers

import (
	"assignment-2/models"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type OrderController struct {
	DB *gorm.DB
}

func (r *OrderController) CreateOrder(c *gin.Context) {
	
	var (
		order  models.Order
		result gin.H
	)

	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = json.Unmarshal(jsonData, &order)
	if err != nil {
		log.Fatal(err)
		return
	}

	r.DB.Create(&order)
	if order.Items != nil {
		for _, item := range order.Items {
			r.DB.Create(&item)
		}
	} else {
		log.Fatal("No items in order")
		return
	}

	result = gin.H{
		"message": "Order created successfully",
		"result":  order,
	}

	c.JSON(200, result)
}

func (r *OrderController) GetOrder(c *gin.Context) {

	var (
		order  []models.Order
		result gin.H
	)

	r.DB.Find(&order)

	result = gin.H{
		"message": "Get Order successfully",
		"result":  order,
	}

	c.JSON(200, result)

}

func (r *OrderController) UpdateOrder(c *gin.Context) {
	
	var (
		result gin.H
	)

	orderID := c.Param("orderId")

	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	order := models.Order{}

	r.DB.First(&order, orderID)

	if order.OrderID == 0 {
		result = gin.H{
			"message": "Order not found",
		}
		c.JSON(404, result)
		return
	}

	orderData := models.Order{}
	err = json.Unmarshal(jsonData, &orderData)
	if err != nil {
		log.Fatal(err)
		return
	}

	r.DB.Model(&order).Updates(map[string]interface{}{
		"customer_name": orderData.CustomerName,
		"ordered_at":    orderData.OrderedAt,
	})

	if orderData.Items != nil {
		for _, dataItem := range orderData.Items {

			item := models.Item{}
			r.DB.Where("line_item_id = ?", dataItem.LineItemID).First(&item)

			r.DB.Model(&item).Updates(map[string]interface{}{
				"itemCode":    dataItem.ItemCode,
				"description": dataItem.Description,
				"quantity":    dataItem.Quantity,
			})
		}
		order.Items = orderData.Items
	}

	result = gin.H{
		"message": "Order updated successfully",
		"result":  order,
	}

	c.JSON(200, result)
}

func (r *OrderController) DeleteOrder(c *gin.Context) {
	
	var (
		result gin.H
	)

	orderID := c.Param("orderId")

	order := models.Order{}

	r.DB.First(&order, orderID)

	if order.OrderID == 0 {
		result = gin.H{
			"message": "Order not found",
		}
		c.JSON(404, result)
		return
	}

	r.DB.Delete(&order)

	result = gin.H{
		"message": "Order deleted successfully",
	}

	c.JSON(200, result)
}
