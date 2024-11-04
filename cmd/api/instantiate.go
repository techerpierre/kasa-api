package main

import (
	"github.com/gin-gonic/gin"
	"github.com/techerpierre/kasa-api/internal/application/handlers"
	"github.com/techerpierre/kasa-api/internal/domain/services"
	"github.com/techerpierre/kasa-api/internal/infrastructure/api"
	"github.com/techerpierre/kasa-api/internal/infrastructure/repositories"
	db "github.com/techerpierre/kasa-api/models"
)

func Instanciate(app *gin.Engine, prisma *db.PrismaClient) {
	userRepository := repositories.CreateUserRepository(prisma)
	accommodationRepository := repositories.CreateAccommodationRepository(prisma)
	authorizationRepository := repositories.CreateAuthorizationsRepository(prisma)
	bookingRepository := repositories.CreateBookingRepository(prisma)
	commentRepository := repositories.CreateCommentRepository(prisma)
	ratingRepository := repositories.CreateRatingRepository(prisma)
	passwordRepository := repositories.CreatePasswordRepository()
	jwtRepository := repositories.CreateJwtRepository()

	passwordService := services.CreatePasswordService(passwordRepository)
	jwtService := services.CreateJwtService(jwtRepository)
	userService := services.CreateUserService(userRepository, passwordService)
	accommodationService := services.CreateAccomodationService(accommodationRepository)
	authorizationService := services.CreateAuthorizationsService(authorizationRepository)
	bookingService := services.CreateBookingService(bookingRepository)
	commentService := services.CreateCommentService(commentRepository)
	ratingService := services.CreateRatingService(ratingRepository)
	authService := services.CreateAuthService(userService, passwordService, jwtService)

	userAPI := api.CreateUserAPI(userService)
	accommodationAPI := api.CreateAccommodationAPI(accommodationService)
	authorizationAPI := api.CreateAuthorizationAPI(authorizationService)
	bookingAPI := api.CreateBookingAPI(bookingService)
	commentAPI := api.CreateCommentAPI(commentService)
	ratingAPI := api.CreateRatingAPI(ratingService)
	authApi := api.CreateAuthApi(authService)

	handlers.CreateUserHTTPHandler(app, userAPI).RegisterRoutes()
	handlers.CreateAccomodationHTTPHandler(app, accommodationAPI).RegisterRoutes()
	handlers.CreateAuthorizationHTTPHandler(app, authorizationAPI).RegisterRoutes()
	handlers.CreateBookingHTTPHandler(app, bookingAPI).RegisterRoutes()
	handlers.CreateCommentHTTPHandler(app, commentAPI).RegisterRoutes()
	handlers.CreateRatingHTTPHandler(app, ratingAPI).RegisterRoutes()
	handlers.CreateAuthHTTPHandler(app, authApi).RegisterRoutes()
}
