package webserver

import (
	"Paste-Echo/app/router"
	modules "Paste-Echo/pkg/moudles"
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"time"
)

var (
	webServer  *modules.Module
	echoGlobal *echo.Echo
)
func Start(addr string, autoTls bool, cert string, key string) {
	webServer = modules.Register("webServer", 1)

	echoGlobal = echo.New()

	echoGlobal.HideBanner = true
	echoGlobal.Debug = true
	//全局中间件
	echoGlobal.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		Skipper:           nil,
		StackSize:         1024,
		DisableStackAll:   false,
		DisablePrintStack: false,
	}))

	echoGlobal.Use(middleware.Gzip())
	echoGlobal.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: nil,
		Format: `{"time":"${time_rfc3339}","id":"${id}","method":"${method}","uri":"${uri}",` +
			`"status":${status},"bytes_in":${bytes_in},"bytes_out":${bytes_out},"remote_ip":"${remote_ip}"}` + "\n",
		CustomTimeFormat: "2021-11-06 15:04:05",
		Output:           nil,
	}))

	router.Run(echoGlobal)

	var erch = make(chan error, 1)
	if autoTls {
		go func() {
			erch <- echoGlobal.StartAutoTLS(addr)
		}()
	} else if cert != "" && key != "" {
		go func() {
			erch <- echoGlobal.StartTLS(addr, cert, key)
		}()
	} else {
		go func() {
			erch <- echoGlobal.Start(addr)
		}()
	}

	for {
		select {
		case err := <-erch:
			log.Fatal(err)
		case <-webServer.Stop:
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			echoGlobal.Shutdown(ctx)
			webServer.StopComplete()
		}
	}
}