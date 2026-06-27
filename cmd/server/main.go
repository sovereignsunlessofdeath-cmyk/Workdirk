package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	"workdirk/db"
	"workdirk/gateway/router"
	"workdirk/internal/handlers"
	"workdirk/internal/repository"
	"workdirk/internal/services"
)

func main() {
	log.Println("Initializing Workdirk server with MySQL...")

	// 1. Force Go to load the .env file from the current directory first
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found, reading from system environment variables")
	}

	// Fetch the database URL from the environment
	dbConnStr := os.Getenv("DATABASE_URL")
	if dbConnStr == "" {
		log.Fatal("❌ Critical Error: DATABASE_URL environment variable is not set in your .env file!")
	}

	// Set up server port fallback
	serverPort := os.Getenv("PORT")
	if serverPort == "" {
		serverPort = ":8080"
	}

	log.Println("✅ Environment loaded successfully. Running configurations...")

	// 2. Open MySQL database connection pool
	dbConn, err := db.InitDB(dbConnStr)
	if err != nil {
		log.Fatalf("❌ Critical failure during database connection: %v", err)
	}
	defer func() {
		log.Println("Closing MySQL database connection pool...")
		dbConn.Close()
	}()

	// 3. Setup and run migrations
	migrator := db.NewMigrationEngine(dbConn)
	if err := migrator.RunSafeMigrations("./db/migrations"); err != nil {
		log.Fatalf("❌ Critical failure during schema migration: %v", err)
	}

	// 4. Dependency Injection across all 6 layers
	userRepo := repository.NewUserRepository(dbConn)
	authRepo := repository.NewAuthRepository(dbConn)
	jobRepo := repository.NewJobRepository(dbConn)
	paymentRepo := repository.NewPaymentRepository(dbConn)
	reviewRepo := repository.NewReviewRepository(dbConn)
	sessionRepo := repository.NewSessionRepository(dbConn)

	userService := services.NewUserService(userRepo)
	authService := services.NewAuthService(authRepo)
	jobService := services.NewJobService(jobRepo)
	paymentService := services.NewPaymentService(paymentRepo)
	reviewService := services.NewReviewService(reviewRepo)
	sessionService := services.NewSessionService(sessionRepo)

	userHandler := handlers.NewUserHandler(userService)
	authHandler := handlers.NewAuthHandler(authService, sessionService)
	jobHandler := handlers.NewJobHandler(jobService)
	paymentHandler := handlers.NewPaymentHandler(paymentService)
	reviewHandler := handlers.NewReviewHandler(reviewService)
	sessionHandler := handlers.NewSessionHandler(sessionService)

	// ========================================================================
	// 5. ROUTER & ROUTE DEFINITIONS
	// ========================================================================
	mux := router.NewRouter()

	// API Endpoints
	mux.HandleFunc("/api/v1/auth/register", userHandler.Register)
	mux.HandleFunc("/api/v1/auth/login", authHandler.Login)
	mux.HandleFunc("/api/v1/jobs/create", jobHandler.Create)
	mux.HandleFunc("/api/v1/payments/initiate", paymentHandler.InitializeEscrow)
	mux.HandleFunc("/api/v1/reviews/leave", reviewHandler.LeaveReview)
	mux.HandleFunc("/api/v1/sessions/check", sessionHandler.CheckSession)

	// Authentication Endpoints
	mux.HandleFunc("/login", authHandler.Login)
	mux.HandleFunc("/logout", authHandler.Logout)
	mux.HandleFunc("/forgot-password", authHandler.ForgotPassword)

	// ========================================================================
	// 6. NETWORKING & ENGINE STARTUP
	// ========================================================================
	server := &http.Server{
		Addr:         serverPort,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("📡 Workdirk Gateway router online! Listening smoothly on port %s", serverPort)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("❌ Server crashed unexpectedly: %v", err)
	}
}