package server

import (
	"enigma-lms/config"
	"enigma-lms/controller"
	"enigma-lms/manager"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	ucManager manager.UseCaseManager
	engine    *gin.Engine
	host      string
	apiCfg    config.ApiConfig
}

func (s *Server) setupControllers() {
	rg := s.engine.Group("/api/v1")
	controller.NewUserController(s.ucManager.UserUseCase(), rg, s.apiCfg).Route()
	controller.NewCourseController(s.ucManager.CourseUseCase(), rg).Route()
	controller.NewEnrollmentController(s.ucManager.EnrollmentUseCase(), rg).Route()
}

func (s *Server) Run() {
	s.setupControllers()
	if err := s.engine.Run(s.host); err != nil {
		panic(err)
	}
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	infraManager, _ := manager.NewInfraManager(cfg)

	repoManager := manager.NewRepoManager(infraManager)
	ucManager := manager.NewUseCaseManager(repoManager)

	engine := gin.Default()
	apiHost := fmt.Sprintf(":%s", cfg.ApiConfig.ApiPort)

	return &Server{
		ucManager: ucManager,
		engine:    engine,
		host:      apiHost,
		apiCfg:    cfg.ApiConfig,
	}
}
