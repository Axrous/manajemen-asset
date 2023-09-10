package manager

import (
	"database/sql"
	"final-project-enigma-clean/config"
	"fmt"
	_ "github.com/lib/pq"
)

type InfraManager interface {
	Connect() *sql.DB
}

type infraManager struct {
	db  *sql.DB
	cfg *config.Config
}

func (i *infraManager) Connect() *sql.DB {
	return i.db
}

func (i *infraManager) initdb() error {
	//init dsn in here
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		i.cfg.DbConfig.Host,
		i.cfg.DbConfig.Port,
		i.cfg.DbConfig.User,
		i.cfg.DbConfig.Password,
		i.cfg.DbConfig.DbName,
	)
	db, err := sql.Open(i.cfg.DbConfig.DbDriver, dsn)
	if err != nil {
		return fmt.Errorf("Failed to open sql %v ", err.Error())
	}

	i.db = db
	return nil
}

// constructor
func NewInfraManager(cfg *config.Config) (InfraManager, error) {
	connect := &infraManager{
		cfg: cfg,
	}

	//define initdbmethod
	if err := connect.initdb(); err != nil {
		return nil, fmt.Errorf("Failed init db %v", err.Error())
	}
	return connect, nil
}
