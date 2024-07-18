// Endpoint формирует ответ клиенту
package endpoint

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type Service interface {
	DaysLeft() int
}

type Endpoint struct {
	s Service
}

func New(svc Service) *Endpoint {
	return &Endpoint{
		s: svc,
	}
}

func (e *Endpoint) Status(ctx echo.Context) error { //Status==Handler
	days := e.s.DaysLeft()
	s := fmt.Sprintf("Количество дней: %d", days)
	err := ctx.String(http.StatusOK, s)
	if err != nil {
		return err
	}
	return nil
}
