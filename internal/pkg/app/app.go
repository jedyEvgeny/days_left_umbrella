package app

import (
	"fmt"
	"log"

	"github.com/jedyEvgeny/days_left_umbrella/internal/app/endpoint"
	"github.com/jedyEvgeny/days_left_umbrella/internal/app/mv"
	"github.com/jedyEvgeny/days_left_umbrella/internal/app/service"
	"github.com/labstack/echo"
)

type App struct {
	e    *endpoint.Endpoint
	s    *service.Service
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}
	a.s = service.New()
	a.e = endpoint.New(a.s)
	a.echo = echo.New()
	a.echo.Use(mv.RoleCheck)
	a.echo.GET("/status", a.e.Status)
	return a, nil
}

func (a *App) Run() error {
	fmt.Println("cервер запущен")
	err := a.echo.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
