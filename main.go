package main

import (
	"encoding/json"
	"fmt"
	"math"      // Add this
	"net/http"
	"strconv"   // Add this
	"strings"   // Add this
	"sync"
	"github.com/google/uuid"
)

type Receipt struct {
	Retailer     string  `json:"retailer"`
	PurchaseDate string  `json:"purchaseDate"`
	PurchaseTime string  `json:"purchaseTime"`
	Total        string  `json:"total"`
	Items        []Item  `json:"items"`
}

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price           string `json:"price"`
}

type Response struct {
	ID string `json:"id"`
}

type PointsResponse struct {
	Points int `json:"points"`
}

var receiptStore = make(map[string]Receipt)
var storeLock = sync.Mutex{}

func processReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt Receipt
	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		http.Error(w, "Invalid receipt format", http.StatusBadRequest)
		return
	}

	receiptID := uuid.New().String()

	storeLock.Lock()
	receiptStore[receiptID] = receipt
	storeLock.Unlock()

	response := Response{ID: receiptID}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func getPoints(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Path[len("/receipts/"):] // Extract full path after "/receipts/"

    // Fix: Remove trailing "/points"
    if idx := len(id) - len("/points"); idx > 0 {
        id = id[:idx]
    }

    fmt.Println("Corrected Requested ID:", id) // Debugging

    storeLock.Lock()
    receipt, exists := receiptStore[id]
    storeLock.Unlock()

    fmt.Println("Stored Receipts:", receiptStore) // Debugging

    if !exists {
        fmt.Println("No receipt found for ID:", id) // Debugging
        http.Error(w, "No receipt found", http.StatusNotFound)
        return
    }

    points := calculatePoints(receipt)
    response := PointsResponse{Points: points}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func calculatePoints(receipt Receipt) int {
    points := 0

    // Rule 1: Retailer Name
    alphanumericCount := 0
    for _, char := range receipt.Retailer {
        if ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z') || ('0' <= char && char <= '9') {
            alphanumericCount++
        }
    }
    fmt.Println("Retailer Name Points:", alphanumericCount)
    points += alphanumericCount

    // Rule 2: Round Total
    if strings.HasSuffix(receipt.Total, ".00") {
        fmt.Println("Round Total Points: 50")
        points += 50
    }

    // Rule 3: Multiple of 0.25
    totalValue, _ := strconv.ParseFloat(receipt.Total, 64)
    if math.Mod(totalValue, 0.25) == 0 {
        fmt.Println("Total Multiple of 0.25 Points: 25")
        points += 25
    }

    // Rule 4: 5 Points for Every 2 Items
    fmt.Println("Items Points:", (len(receipt.Items)/2)*5)
    points += (len(receipt.Items) / 2) * 5

    // Rule 5: Item Description Multiple of 3
    for _, item := range receipt.Items {
        trimmedLength := len(strings.TrimSpace(item.ShortDescription))
        if trimmedLength%3 == 0 {
            itemPoints := int(math.Ceil(parsePrice(item.Price) * 0.2))
            fmt.Println("Item:", item.ShortDescription, "Points:", itemPoints)
            points += itemPoints
        }
    }

    // Rule 6: Odd Day
    day, _ := strconv.Atoi(receipt.PurchaseDate[len(receipt.PurchaseDate)-2:])
    if day%2 != 0 {
        fmt.Println("Odd Day Points: 6")
        points += 6
    }

    // Rule 7: Time Between 2 PM - 4 PM
    purchaseHour, _ := strconv.Atoi(receipt.PurchaseTime[:2])
    if purchaseHour >= 14 && purchaseHour < 16 {
        fmt.Println("Time Bonus Points: 10")
        points += 10
    }

    fmt.Println("Final Points:", points)
    return points
}

func parsePrice(price string) float64 {
    value, err := strconv.ParseFloat(price, 64)
    if err != nil {
        fmt.Println("Error parsing price:", price)
        return 0.0
    }
    return value
}

func main() {
	http.HandleFunc("/receipts/process", processReceipt)
	http.HandleFunc("/receipts/", getPoints)
	http.ListenAndServe(":8080", nil)
}