package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jay-SP/receipt-processor-challenge/pkg/controllers"
)

var RegisterReceiptRoutes = func(r *gin.Engine) {
	r.POST("/receipts/process", controllers.PostReceipts)
	r.GET("/receipts/:id/points", controllers.GetPoints)
}
