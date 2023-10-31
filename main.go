package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Location struct {
	ID        string  `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

var locations = map[string]Location{
	"0": {
		ID:        "0",
		Latitude:  0,
		Longitude: 0,
	},
	"1": {
		ID:        "1",
		Latitude:  1,
		Longitude: 1,
	},
}

func getAllLocations(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, locations)
}

func getLocationById(c *gin.Context) {
	id := c.Param("id")

	if location, exists := locations[id]; exists {
		c.IndentedJSON(http.StatusOK, location)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "location not found"})
}

func addLocation(c *gin.Context) {
	var newLocation Location

	if err := c.BindJSON(&newLocation); err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	locations[newLocation.ID] = newLocation
	c.IndentedJSON(http.StatusCreated, newLocation)
}

func main() {
	router := gin.Default()
	router.GET("/locations", getAllLocations)
	router.GET("/locations/:id", getLocationById)
	router.POST("/locations", addLocation)

	router.Run("localhost:8080")
}
