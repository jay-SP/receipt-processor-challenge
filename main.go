// RECEIPT API
//  This is a sample receipts API. You can find out more about the API at https://github.com/jay-SP/receipt-processor-challenge
//  Schemes: http
//  Host: localhost:8080
//  BasePath: /
//  Version: 1.0.0
//  Contact: Jaya Surya
// <jaypagidi1@gmail.com>
//
//  Consumes:
//  - application/json
//
//  Produces:
//  - application/json
// swagger:meta

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jay-SP/receipt-processor-challenge/pkg/routes"
)

func main() {
	r := gin.Default()
	routes.RegisterReceiptRoutes(r)
	r.Run(":8080")
}
