package mv

import (
	"log"
	"strings"

	"github.com/labstack/echo"
)

const roleAdmin = "admin"

func RoleCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		val := ctx.Request().Header.Get("User-Role")
		if strings.EqualFold(val, roleAdmin) {
			log.Println("Red batton user detected")

		}
		err := next(ctx)
		if err != nil {
			return err
		}
		return nil
	}
}
