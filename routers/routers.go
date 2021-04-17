package routers

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/vipindasvg/courserawebapp/controllers"
)

const (
	versionpref = "/api/v1"
)

func InitRoutes() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.Logger.SetLevel(log.DEBUG)
	e.Logger.SetOutput(os.Stdout)
	if l, ok := e.Logger.(*log.Logger); ok {
		l.SetHeader("${time_rfc3339} ${level}")
	}

	e.POST(versionpref+"/savecourses", controllers.SaveCourses)
	e.GET(versionpref+"/getcourses", controllers.GetCourses)

	return e
}
