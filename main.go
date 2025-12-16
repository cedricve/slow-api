package main

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type HealthResponse struct {
	Status      string `json:"status"`
	Timestamp   string `json:"timestamp"`
	Delay       int    `json:"delay_ms"`
	ServiceName string `json:"service_name,omitempty"`
}

func main() {
	// Load environment variables from .env.local
	if err := godotenv.Load(".env.local"); err != nil {
		log.Printf("Warning: Error loading .env.local file: %v", err)
	}

	// Get timeout values from environment variables
	minTimeout := getEnvAsInt("MIN_TIMEOUT", 1000)
	maxTimeout := getEnvAsInt("MAX_TIMEOUT", 10000)

	// Create Gin router
	router := gin.Default()

	// Helper function to create instance endpoint handler
	createInstanceHandler := func(serviceName string, min, max int) gin.HandlerFunc {
		return func(c *gin.Context) {
			// Calculate random delay between min and max timeout
			var delay int
			if max > min {
				delay = min + rand.Intn(max-min+1)
			} else {
				delay = min
			}

			// Sleep for the specified delay
			time.Sleep(time.Duration(delay) * time.Millisecond)

			// Return health response
			response := HealthResponse{
				Status:      "ok",
				Timestamp:   time.Now().Format(time.RFC3339),
				Delay:       delay,
				ServiceName: serviceName,
			}

			c.JSON(200, response)
		}
	}

	// Instance endpoints with different timeout ranges
	router.GET("/instance1", createInstanceHandler("instance1", 500, 1000))
	router.GET("/instance2", createInstanceHandler("instance2", 1000, 2000))
	router.GET("/instance3", createInstanceHandler("instance3", 2000, 4000))

	// Start server
	port := getEnv("PORT", "8081")
	log.Printf("Starting server on port %s with timeout range %d-%d ms", port, minTimeout, maxTimeout)

	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}

// getEnv retrieves an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvAsInt retrieves an environment variable as int or returns a default value
func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
