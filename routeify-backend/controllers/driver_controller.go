package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/samriddha-basu-cloud/Routeify/repository"
)

func GetDriverBookings(c *gin.Context) {
	// Implement logic to fetch bookings for the driver
	// Example:
	driverID := getUserIdFromToken(c.Request.Header.Get("Authorization"))
	bookings, err := repository.GetDriverBookings(driverID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch driver bookings"})
		return
	}

	c.JSON(http.StatusOK, bookings)
}
