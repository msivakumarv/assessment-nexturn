package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	db "blog-management/config"
	"blog-management/controller"
	"blog-management/repository"
	"blog-management/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := db.InitializeDatabase(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	userRepo := repository.NewUserRepository(db.GetDB())

	authService := service.NewAuthService(userRepo)

	blogRepo := repository.NewBlogRepository(db.GetDB())
	blogService := service.NewBlogService(blogRepo)
	blogController := controller.NewBlogController(blogService)

	authController := controller.NewAuthController(authService)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.POST("/register", authController.Register)
	r.POST("/login", authController.Login)

	protected := r.Group("/blogs")
	protected.Use(controller.AuthMiddleware([]byte("your-jwt-secret-key")))
	{
		protected.POST("/", blogController.CreateBlog)
		protected.GET("/:id", blogController.GetBlog)
		protected.GET("/", blogController.GetAllBlogs)
		protected.PUT("/:id", blogController.UpdateBlog)
		protected.DELETE("/:id", blogController.DeleteBlog)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" //
	}

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
