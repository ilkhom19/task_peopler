package main

import (
	"fmt"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"task_peopler/internal/config"
	"task_peopler/internal/repository/sqlite"
	"task_peopler/internal/service"
	"task_peopler/pkg"
	"task_peopler/pkg/database"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

// @title Peopler API documentation
// @version 1.0.1
// @host http://localhost:8080
// @BasePath
func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	// init db
	db, err := database.NewSQLite3Connection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	apiClient := pkg.NewAPIClient()
	personRepo := sqlite.NewPersonRepo(db)
	personService := service.NewPersonService(personRepo, apiClient)

	handler := http.NewHandler(
		usersService,
		timesheetService,
		orderService,
		recordService,
		dailyHoursService,
	)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	})
	routerWithCors := c.Handler(handler.InitRouter())

	// init & run server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: routerWithCors,
	}

	log.Info("Server has started...")

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
