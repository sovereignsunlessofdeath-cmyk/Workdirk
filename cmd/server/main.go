package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"workdirk/db"
	"workdirk/gateway/router"
	"workdirk/internal/handlers"
	"workdirk/internal/repository"
	"workdirk/internal/services"
)

// ============================================================================
// STANDALONE FUNCTION: main()
// ============================================================================
func main() {
	log.Println("Initializing Workdirk server with MySQL...")

	// 1. Environmental fallbacks for your local MySQL database connection string
	dbConnStr := os.Getenv("DATABASE_URL")
	if dbConnStr == "" {
		dbConnStr = "root:@tcp(127.0.0.1:3306)/workdirk?parseTime=true"
	}

	serverPort := os.Getenv("PORT")
	if serverPort == "" {
		serverPort = ":8080"
	}

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

	// 5. Setup Gateway Router routes
	mux := router.NewRouter()

	mux.HandleFunc("/api/v1/auth/register", userHandler.Register)
	mux.HandleFunc("/api/v1/auth/login", authHandler.Login)
	mux.HandleFunc("/api/v1/jobs/create", jobHandler.Create)
	mux.HandleFunc("/api/v1/payments/initiate", paymentHandler.InitializeEscrow)
	mux.HandleFunc("/api/v1/reviews/leave", reviewHandler.LeaveReview)
	mux.HandleFunc("/api/v1/sessions/check", sessionHandler.CheckSession)

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
