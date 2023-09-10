package delivery

import (
	"final-project-enigma-clean/config"
	"final-project-enigma-clean/delivery/controller"
	"final-project-enigma-clean/manager"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Server struct {
	um  manager.UsecaseManager
	gin *gin.Engine
	log *logrus.Logger
}

func (s *Server) initMiddlewares() {

}

func (s *Server) initControllers() {
	controller.NewUserDetailsController(s.um.UDetailsUC(), s.gin).Route()
}

func (s *Server) Run() {
	s.initMiddlewares()
	s.initControllers()
	s.gin.Run(":3000")
}

func NewServer() *Server {

	cfg, err := config.NewDbConfig()
	if err != nil {
		fmt.Printf("Failed on db constructor %v", err.Error())
	}
	//define constructor dari infra
	im, err := manager.NewInfraManager(cfg)
	if err != nil {
		fmt.Printf("Failed on infra constructor %v", err.Error())
	}

	rm := manager.NewRepoManager(im)
	um := manager.NewUsecaseManager(rm)

	//gin serv
	g := gin.Default()

	//init log
	log := logrus.New()
	return &Server{
		um:  um,
		gin: g,
		log: log,
	}
}
