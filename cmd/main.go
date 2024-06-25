package main

import (
	"fmt"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	r := gin.Default()
	r.POST("/receipts/process", PostReceipts)
	r.GET("/receipts/:id/points", GetPoints)
	r.Run("localhost:8080")
}

type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"` // Note that Total is a string here
	ID           string `json:"id"`    // To store the generated ID
	Points       int    `json:"points"`
}

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"` // Price is also a string
}

var receiptStorage = make(map[string]Receipt)

func PostReceipts(c *gin.Context) {
	var receipt Receipt

	//Bind JSON data to the receipt struct
	if err := c.BindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := validateReceipt(receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Generate a UUID
	receipt.ID = uuid.New().String()

	//Calculate Points
	points, err := calculatePoints(receipt)
	receipt.Points = points
	if err != nil {
		fmt.Errorf("error occurred while calculating points")
	}

	//store the receipt in the map
	receiptStorage[receipt.ID] = receipt

	c.JSON(http.StatusOK, gin.H{"id": receipt.ID})
}

func calculatePoints(receipt Receipt) (int, error) {
	points := 0

	// Rule 1: One point for every alphanumeric character in the retailer name.
	alphanumericRegex := regexp.MustCompile(`[a-zA-Z0-9]+`)
	matches := alphanumericRegex.FindAllString(receipt.Retailer, -1)
	points += len(strings.Join(matches, ""))

	// Rule 2: 50 points if the total is a round dollar amount with no cents.
	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		return 0, fmt.Errorf("error parsing total: %v", err)
	}
	if total == math.Trunc(total) {
		points += 50
	}

	// Rule 3: 25 points if the total is a multiple of 0.25.
	if int(total*100)%25 == 0 {
		points += 25
	}

	// Rule 4: 5 points for every two items on the receipt.
	points += (len(receipt.Items) / 2) * 5

	// Rule 5: If the trimmed length of the item description is a multiple of 3,
	// multiply the price by 0.2 and round up to the nearest integer.
	for _, item := range receipt.Items {
		trimmedDescription := strings.TrimSpace(item.ShortDescription)
		if len(trimmedDescription)%3 == 0 {
			price, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				return 0, fmt.Errorf("error parsing price: %v", err)
			}
			points += int(math.Ceil(price * 0.2))
		}
	}

	// Rule 6: 6 points if the day in the purchase date is odd.
	purchaseDate, err := time.Parse("2006-01-02", receipt.PurchaseDate)
	if err == nil && purchaseDate.Day()%2 == 1 {
		points += 6
	}

	// Rule 7: 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	purchaseTime, err := time.Parse("15:04", receipt.PurchaseTime)
	if err == nil && purchaseTime.Hour() >= 14 && purchaseTime.Hour() < 16 {
		points += 10
	}

	return points, err
}

func validateReceipt(receipt Receipt) error {

	// retailer validation
	if matched, _ := regexp.MatchString(`^[\w\s\-&]+$`, receipt.Retailer); !matched {
		return fmt.Errorf("invalid retailer name: %s", receipt.Retailer)
	}

	// purchaseDate validation
	if _, err := time.Parse("2006-01-02", receipt.PurchaseDate); err != nil {
		return fmt.Errorf("invalid purchase date format: %s", receipt.PurchaseDate)
	}

	// purchaseTime validation
	if _, err := time.Parse("15:04", receipt.PurchaseTime); err != nil {
		return fmt.Errorf("invalid purchase time format: %s", receipt.PurchaseTime)
	}

	// items validation
	if len(receipt.Items) < 1 {
		return fmt.Errorf("invalid item list, with %d items", len(receipt.Items))
	}
	for _, item := range receipt.Items {
		if matched, _ := regexp.MatchString(`^[\w\s\-]+$`, item.ShortDescription); !matched {
			return fmt.Errorf("invalid short description for item: %s", item.ShortDescription)
		}
		if matched, _ := regexp.MatchString(`^\d+\.\d{2}$`, item.Price); !matched {
			return fmt.Errorf("invalid price format for item: %s", item.Price)
		}
	}

	// total validation
	if matched, _ := regexp.MatchString(`^\d+\.\d{2}$`, receipt.Total); !matched {
		return fmt.Errorf("invalid total format: %s", receipt.Total)
	}

	return nil
}

func GetPoints(c *gin.Context) {
	receiptID := c.Param("id")
	receipt, exists := receiptStorage[receiptID]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "No receipt found for that id"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"points": receipt.Points})
}
