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
	minTimeout := getEnvAsInt("MIN_TIMEOUT", 4000)
	maxTimeout := getEnvAsInt("MAX_TIMEOUT", 4000)

	// Create Gin router
	router := gin.Default()

	// Health endpoint
	router.GET("/health", func(c *gin.Context) {
		// Get service name from query param
		serviceName := c.Query("service")

		// Get min/max timeout from query params, fallback to env vars
		min := minTimeout
		max := maxTimeout

		if minParam := c.Query("min"); minParam != "" {
			if val, err := strconv.Atoi(minParam); err == nil {
				min = val
			}
		}

		if maxParam := c.Query("max"); maxParam != "" {
			if val, err := strconv.Atoi(maxParam); err == nil {
				max = val
			}
		}

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
	})

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
