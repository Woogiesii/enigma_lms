package config

import (
	"database/sql"
	"enigma-lms/controller"
	"enigma-lms/repository"
	"enigma-lms/usecase"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	uc     usecase.UserUseCase
	engine *gin.Engine
	host   string
}

func (s *Server) setupControllers() {
	rg := s.engine.Group("/api/v1")
	controller.NewUserController(s.uc, rg).Route()
}

func (s *Server) Run() {
	s.setupControllers()
	if err := s.engine.Run(s.host); err != nil {
		panic(err)
	}
}

func NewServer() *Server {
	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "root"
	dbName := "enigma_lms_db"
	driver := "postgres"

	dbconf := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := sql.Open(driver, dbconf)
	if err != nil {
		panic(err)
	}

	userRepo := repository.NewUserRepository(db)
	userUc := usecase.NewUserUseCase(userRepo)

	engine := gin.Default()
	apiHost := fmt.Sprintf(":%s", "8081")

	return &Server{
		uc:     userUc,
		engine: engine,
		host:   apiHost,
	}
}
