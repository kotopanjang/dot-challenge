package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"dot-rahadian-ardya-kotopanjang/middleware"
	"dot-rahadian-ardya-kotopanjang/model"
	"dot-rahadian-ardya-kotopanjang/pkg/redis"
	"dot-rahadian-ardya-kotopanjang/services/project/handler"
	"dot-rahadian-ardya-kotopanjang/services/project/repository"
	"dot-rahadian-ardya-kotopanjang/services/project/usecase"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	config *viper.Viper
)

func init() {
	config = viper.NewWithOptions(
		viper.EnvKeyReplacer(
			strings.NewReplacer(".", "_"),
		),
	)
	config.SetConfigFile("config")
	config.SetConfigType("ini")
	config.AutomaticEnv()
	if err := config.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// initiate database
	mysqlUri := config.GetString("mysql.uri")
	db, err := gorm.Open(mysql.Open(mysqlUri), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// inititate redis
	redisURL := config.GetString("redis.url")
	redisPassword := config.GetString("redis.password")
	rds := redis.NewRedis(redisURL, redisPassword)

	// initiate auto migration
	if err := db.AutoMigrate(&model.Project{}, &model.Member{}); err != nil {
		log.Fatal(err)
	}

	// initiate logger
	logger := logrus.New()

	// initiate mux router
	router := mux.NewRouter()

	// initiate health check
	router.HandleFunc("/", healthCheck)

	// set cors
	httpHandler := middleware.CORS(router)

	// initiate Project
	projectRepo := repository.NewProjectRepository(db, rds, logger)
	projectUsecase := usecase.NewProjectUsecase(logger, projectRepo)
	handler.NewProjectHandler(router, logger, projectUsecase)

	apiPort := config.GetString("api.port")
	srv := &http.Server{
		Handler:      httpHandler,
		Addr:         fmt.Sprintf("127.0.0.1:%v", apiPort),
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	fmt.Println("Starting app on http://localhost:" + apiPort)

	log.Fatal(srv.ListenAndServe())
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	if _, err := io.WriteString(w, `ok`); err != nil {
		fmt.Println("error")
	}
}
