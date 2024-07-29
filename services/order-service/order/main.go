package main

import (
	"fmt"
	"log"
	"context"
	"os"
	"net/http"
	"order-service/ent"
	_ "github.com/lib/pq"
	"github.com/gin-gonic/gin"
	"strconv"
	// "github.com/labstack/echo/v4"
)

type Server struct {
	db   *ent.Client
	http *gin.Engine
}
var svr Server
var client *ent.Client

func main() {
	client, err := initDb()
	if err != nil {
		log.Fatalf("failed to initialize db: %v", err)
		return 
	}
	
    if client == nil {
        log.Fatal("Database client is nil")
    }

    svr.db = client

    fmt.Println("Database client is connected")
    fmt.Printf("client = %v\n", client)
    defer client.Close()

	r := gin.Default()
    svr.http = r
	r.GET("/", getHello)
	r.POST("/orders", createOrder)
	r.GET("/orders/:id", getOrder)
	r.PUT("/orders/:id", updateOrder)

	r.Run(":1323")
}

// Initialize the database	
func initDb() (*ent.Client, error) {
	dburl := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"))	 	
	log.Printf("connecting to db: %s", dburl)
	client, err := ent.Open("postgres", dburl)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}	
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}	
	return client, nil				
}

func getHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}

//post input params: {"user_id": 5, "product_id": 25, "quantity":20, "status":"OrderCreated"}
func createOrder(c *gin.Context) {
    var input struct {
        UserID    int    `json:"user_id"`
        ProductID int    `json:"product_id"`
        Quantity  int    `json:"quantity"`
        Status    string `json:"status"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	log.Printf("input: %v", input)
	log.Printf("Input Data: UserID=%d, ProductID=%d, Quantity=%d, Status=%s", input.UserID, input.ProductID, input.Quantity, input.Status)
    fmt.Printf("UserID = %T\n", input.UserID)

    if svr.db == nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database client is not initialized"})
        return
    }

    fmt.Printf("client = %v\n", svr.db)
	order, err := svr.db.Order.
        Create().
        SetUserID(input.UserID).
        SetProductID(input.ProductID).
        SetQuantity(input.Quantity).
        SetStatus(input.Status).
        Save(context.Background())
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, order)
}

func getOrder(c *gin.Context) {
    ids := c.Param("id")
	id, err := strconv.Atoi(ids)
    if err != nil {
        // ... handle error
        panic(err)
    }
    order, err := svr.db.Order.Get(context.Background(), id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
        return
    }
    c.JSON(http.StatusOK, order)
}

func updateOrder(c *gin.Context) {
    ids := c.Param("id")
	id, err := strconv.Atoi(ids)
    if err != nil {
        // ... handle error
        panic(err)
    }
    order, err := svr.db.Order.Get(context.Background(), id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
        return
    }

    var input struct {
        ProductID int    `json:"product_id"`
        Quantity  int    `json:"quantity"`
        Status    string `json:"status"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    order, err = order.Update().
        SetProductID(input.ProductID).
        SetQuantity(input.Quantity).
        SetStatus(input.Status).
        Save(context.Background())
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, order)
}