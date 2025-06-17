package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/prem0x01/hospital/internal/config"
	"github.com/prem0x01/hospital/internal/database"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, pls create one in root directory of this project !")
	}

	cfg := config.Load()

	db, err := database.Initialize(cfg.DBUrl)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err := database.RunMigrations(db); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	userRepo := repo.NewUserRepository(db)
	patientRepo := repo.NewPatientRepository(db)
	appointmentRepo := repo.NewAppointmentRepository(db)

	authService := services.NewAuthService(userRepo, cfg.JWTSecret)
	patientService := services.NewPatientService(patientRepo)
	appointmentService := services.NewAppointmentService(appointmentRepo, patientRepo)

	authHandler := handlers.NewAuthHandler(authService)
	patientHandler := handlers.NewPatientHandler(patientService)
	appointmentHandler := handlers.NewAppointmentHandler(appointmentService)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.Static("/static", "./web/static")
	router.LoadHTMLGlob("web/static/*.html")

	api := router.Group("/api/v1")
	{

		auth := api.Group("/auth")
		{
			auth.POST("/login", authHandler.Login)
			auth.POST("/register", authHandler.Register)
		}

		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware(cfg.JWTSecret))
		{
			patients := protected.Group("/patients")
			{
				patients.GET("", patientHandler.GetPatients)
				patients.POST("", patientHandler.CreatePatient)
				patients.GET("/:id", patientHandler.GetPatient)
				patients.PUT("/:id", patientHandler.UpdatePatient)
				patients.DELETE("/:id", patientHandler.DeletePatient)
			}

			appointments := protected.Group("/appointments")
			{
				appointments.GET("", appointmentHandler.GetAppointments)
				appointments.POST("", appointmentHandler.CreateAppointment)
				appointments.GET("/:id", appointmentHandler.GetAppointment)
				appointments.PUT("/:id", appointmentHandler.UpdateAppointment)
				appointments.DELETE("/:id", appointmentHandler.DeleteAppointment)
			}

			protected.GET("/dashboard/stats", handlers.GetDashboardStats(patientRepo, appointmentRepo))
		}
	}

	// frontend routes
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	router.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login.html", nil)
	})
	router.GET("/receptionist", func(c *gin.Context) {
		c.HTML(200, "receptionist.html", nil)
	})
	router.GET("/doctor", func(c *gin.Context) {
		c.HTML(200, "doctor.html", nil)
	})

	log.Printf("Server starting on port %s", cfg.Port)
	log.Fatal(router.Run(":" + cfg.Port))

}
