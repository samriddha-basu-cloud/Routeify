package main

import (
	"github.com/gin-gonic/gin"
	"github.com/samriddha-basu-cloud/Routeify/routeify-backend/controllers"
	"github.com/samriddha-basu-cloud/Routeify/routeify-backend/repository"
)

func main() {
	// Initialize Gin router
	r := gin.Default()

	// Connect to MongoDB
	repository.ConnectMongoDB("mongodb://localhost:27017")

	// Define routes
	defineRoutes(r)

	// Start the server
	r.Run(":8080")
}

func defineRoutes(r *gin.Engine) {
	// Routes for user authentication
	r.POST("/register", controllers.RegisterUser)
	r.POST("/login", controllers.Login)

	// Routes for admin actions
	admin := r.Group("/admin")
	admin.Use(authMiddleware) // Apply authentication middleware for admin routes
	{
		admin.GET("/users", controllers.GetUsers)
		admin.POST("/users", controllers.AddUser)
		admin.PUT("/users/:id", controllers.UpdateUser)
		admin.DELETE("/users/:id", controllers.DeleteUser)

		admin.POST("/brands", controllers.AddBrand)
		admin.GET("/brands", controllers.GetBrands)
		admin.PUT("/brands/:id", controllers.UpdateBrand)
		admin.DELETE("/brands/:id", controllers.DeleteBrand)

		admin.POST("/cabs", controllers.AddCab)
		admin.GET("/cabs", controllers.GetCabs)
		admin.PUT("/cabs/:id", controllers.UpdateCab)
		admin.DELETE("/cabs/:id", controllers.DeleteCab)

		admin.GET("/bookings", controllers.GetBookings)
	}

	// Routes for user actions
	user := r.Group("/user")
	user.Use(authMiddleware) // Apply authentication middleware for user routes
	{
		user.GET("/profile", controllers.GetUserProfile)
		user.POST("/book-cab", controllers.BookCab)
		user.GET("/booking-history", controllers.GetBookingHistory)
	}

	// Routes for driver actions
	driver := r.Group("/driver")
	driver.Use(authMiddleware) // Apply authentication middleware for driver routes
	{
		driver.GET("/bookings", controllers.GetDriverBookings)
	}
}

func authMiddleware(c *gin.Context) {
	// Implement authentication logic here
	// Verify JWT token and extract user information
	// Example:
	// if !isValidToken(c.Request.Header.Get("Authorization")) {
	//     c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	//     c.Abort()
	//     return
	// }
	c.Next()
}
