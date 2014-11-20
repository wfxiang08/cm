package main

import (
	"flag"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	log "github.com/ngaut/logging"
	"github.com/wandoulabs/cm/config"
	"github.com/wandoulabs/cm/proxy"
)

var configFile *string = flag.String("config", "/etc/mixer.conf", "mixer proxy config file")
var logLevel *string = flag.String("log-level", "", "log level [debug|info|warn|error], default error")

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.Parse()

	if len(*configFile) == 0 {
		log.Error("must use a config file")
		return
	}

	cfg, err := config.ParseConfigFile(*configFile)
	if err != nil {
		log.Error(err.Error())
		return
	}

	var svr *proxy.Server
	svr, err = proxy.NewServer(cfg)
	if err != nil {
		log.Error(err.Error())
		return
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		sig := <-sc
		log.Info("Got signal [%d] to exit.", sig)
		svr.Close()
	}()

	svr.Run()
}
