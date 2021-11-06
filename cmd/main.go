package main

import (
	"Paste-Echo/config"
	_ "Paste-Echo/docs"
	"Paste-Echo/moudle/webserver"
	modules "Paste-Echo/pkg/moudles"
	"errors"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
	"os"
	"os/signal"
	"runtime/pprof"
	"syscall"
	"time"
)

var(
	rootCmd = cli.NewApp()
)

func init(){

	rootCmd.Name = "paste-go"
	rootCmd.UsageText = "app"
	rootCmd.Version = "v1.0.0"

	rootCmd.Commands = []cli.Command{
		server,
	}

	rootCmd.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "config,c",
			Usage:  "config file name",
			EnvVar: "CONFIG_FILE",
		},
	}
	rootCmd.Before = func(ctx *cli.Context) error {
		file := ctx.String("c")
		if file!=""{

			viper.SetConfigFile(file)
		}else{

			viper.SetConfigFile("../config.yaml")

		}

		go vipWatch()

		if err:=viper.ReadInConfig();err!=nil{
			log.Error(err)
		}

		if err:=config.InitMysql();err!=nil{
			log.Error(err)
		}

		if err:=config.InitRedisPoll();err!=nil{
			log.Error(err)
		}

		log.Infof(
			`
		server.addr:%s
		server.autoTls:%s
		server.Department:%s `,
			viper.Get("server.addr"),
			viper.Get("server.autoTls"),
			viper.Get("server.staticDir"), )
		return nil
	}


}

var server = cli.Command{
	Name:      "server",
	ShortName: "s",
	Aliases:   nil,
	Usage:     "server [option]",
	UsageText: "server manage",
	Before:    nil,
	Action:    startServer,
}


func startServer(ctx *cli.Context)error{

	go webserver.Start(
		viper.GetString("server.addr"),
		viper.GetBool("server.autoTls"),
		viper.GetString("server.cert"),
		viper.GetString("server.key"),
	)

	ch:=make(chan os.Signal)

	signal.Notify(ch,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGKILL)

	select {
	case <-ch:
		config.CloseMysql()
		config.CloseRedisPoll()
	case <-modules.GlobalShutdown:
	}
	// wait for shutdown to complete, panic after timeout
	time.Sleep(5 * time.Second)
	fmt.Println("===== TAKING TOO LONG FOR SHUTDOWN - PRINTING STACK TRACES =====")
	pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)

	return errors.New("server has been shutdown!")

}

// 更新配置
func vipWatch() {
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		config.SetDebug(viper.GetBool("mysql.debug"))
	})
}

func main(){


	if err := rootCmd.Run(os.Args); err != nil {
		log.Error(err)
		os.Exit(1)
	}

}
