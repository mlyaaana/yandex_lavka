package app

import (
	"fmt"
	"os"

	"github.com/jinzhu/configor"

	courierrest "yandex-team.ru/bstask/api/courier/rest"
	orderrest "yandex-team.ru/bstask/api/order/rest"
	"yandex-team.ru/bstask/database/postgres"
	courierrepo "yandex-team.ru/bstask/repository/courier/postgres"
	orderrepo "yandex-team.ru/bstask/repository/order/postgres"
	"yandex-team.ru/bstask/server/rest"
	courierservice "yandex-team.ru/bstask/service/courier"
	orderservice "yandex-team.ru/bstask/service/order"
)

type App struct {
	config Config
	server *rest.Server
}

func New() (*App, error) {
	a := &App{}
	if err := a.loadConfig(); err != nil {
		return nil, err
	}

	db, err := postgres.New(&postgres.Params{
		Name:     a.config.Database.Name,
		Host:     a.config.Database.Host,
		Port:     a.config.Database.Port,
		User:     a.config.Database.User,
		Password: a.config.Database.Password,
	})
	if err != nil {
		return nil, err
	}

	courierRepository := courierrepo.NewRepository(
		&courierrepo.RepositoryParams{
			Database: db,
		},
	)
	courierService := courierservice.NewService(&courierservice.ServiceParams{
		CourierRepository: courierRepository,
	})
	courierController := courierrest.NewController(
		&courierrest.ControllerParams{
			Service: courierService,
		},
	)

	orderRepository := orderrepo.NewRepository(&orderrepo.RepositoryParams{
		Database: db,
	})
	orderService := orderservice.NewService(&orderservice.ServiceParams{
		OrderRepository:   orderRepository,
		CourierRepository: courierRepository,
	})
	orderController := orderrest.NewController(&orderrest.ControllerParams{
		Service: orderService,
	})

	a.server = rest.NewServer(
		&rest.ServerParams{Address: a.config.Http.Address},
	)
	courierController.Register(a.server)
	orderController.Register(a.server)

	return a, nil
}

func (a *App) Start() error {
	return a.server.Start()
}

func (a *App) loadConfig() error {
	file := os.Getenv("CONFIG_FILE")
	if file == "" {
		return fmt.Errorf("config file was not provided")
	}
	return configor.Load(&a.config, file)
}
