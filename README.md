# Routeify Cab Service Application 🚗

## Directory Structure

```
      routeify-backend/
      ├── controllers/
      │   ├── admin_controller.go
      │   ├── driver_controller.go
      │   └── user_controller.go
      ├── middleware/
      │   ├── auth.go
      │   ├── authorization.go
      │   └── error_handling.go
      ├── models/
      │   ├── booking.go
      │   ├── brand.go
      │   ├── cab.go
      │   ├── driver.go
      │   ├── location.go
      │   ├── payment.go
      │   ├── review.go
      │   ├── user.go
      │   └── vehicle.go
      ├── repository/
      │   └── repository.go
      ├── utils/
      │   └── token.go
      └── main.go
```

## Description
Routeify is a cab service application that allows users to book rides, track their cab in real-time, and make payments seamlessly. This application aims to provide a convenient and reliable transportation solution for users.

## Features
- User Registration: Users can create an account and login to the application.
- Ride Booking: Users can book a cab by providing their pickup and drop-off locations.
- Real-time Tracking: Users can track their cab in real-time on a map.
- Fare Estimation: Users can get an estimated fare for their ride before booking.
- Payment Integration: Users can make payments for their rides using various payment methods.
- Ride History: Users can view their past ride details and invoices.
- Driver Ratings: Users can rate and provide feedback for their drivers.

## Technologies Used
- Frontend: HTML, CSS, JavaScript, React
- Backend: Go Lang
- Database: MongoDB
- Mapping and Location Services: Google Maps API
- Payment Gateway: PhonePe



## Usage
1. Register a new user account or login with existing credentials.
2. Enter your pickup and drop-off locations to book a ride.
3. Track your cab in real-time on the map.
4. Make a payment for your ride using the available payment methods.
5. View your ride history and provide ratings and feedback for your drivers.

## Contributing
Contributions are welcome! If you have any ideas, suggestions, or bug reports, please open an issue or submit a pull request.

## License
This project is licensed under the [MIT License](LICENSE).

## Contact
For any inquiries or support, please contact us at support@routeify.com.
