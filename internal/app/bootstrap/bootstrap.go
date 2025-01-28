package bootstrap

import (
	"github.com/samber/do/v2"
	"online-shop-backend/internal/app"
	"online-shop-backend/pkg/logger"
)

type Bootstrap struct {
	inj *do.RootScope
}

func Website() {
	_, err := app.NewConfig()
	if err != nil {
		panic(err)
	}

	logger.InitLogger()
}
