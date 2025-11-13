package app

import (
	"gocarts/internal/box"
	"gocarts/internal/repository/postgresql"
)

func Run() error {
	envBox, err := box.New()
	if err != nil {
		return err
	}

	//logger

	postgre := postgresql.NewDB(envBox.PostgreSql)

	//service
	_ = postgre

	return nil

}
