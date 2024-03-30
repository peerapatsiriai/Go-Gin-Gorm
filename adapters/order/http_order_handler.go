package orderadapter

import (
	"api/entities"
	orderusecase "api/usecases/order"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpOrderHandler struct {
	orderUseCase orderusecase.OrderUseCase
}

func NewHttpOrderHandler(useCase orderusecase.OrderUseCase) *HttpOrderHandler {
	return &HttpOrderHandler{orderUseCase: useCase}
}

func (h *HttpOrderHandler) Create(c *gin.Context) {
	var order entities.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := h.orderUseCase.Create(&order); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Order created successfully",
		"order":   order,
	})
}

func (h *HttpOrderHandler) GetAll(c *gin.Context) {
	orders, err := h.orderUseCase.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"results": &orders})
}
