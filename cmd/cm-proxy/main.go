package main

import (
	"flag"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	log "github.com/ngaut/logging"
	"github.com/wandoulabs/cm/config"
	"github.com/wandoulabs/cm/proxy"
)

var configFile = flag.String("config", "./etc/cfg.json", "cm config file")

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

	log.SetLevelByString(cfg.LogLevel)

	log.CrashLog("./cm-proxy.dump")

	var svr *proxy.Server
	svr, err = proxy.NewServer(*configFile)
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
		log.Infof("Got signal [%d] to exit.", sig)
		svr.Close()
		os.Exit(0)
	}()

	go svr.Run()

	http.HandleFunc("/api/reload", svr.HandleReload)
	//todo: using configuration
	http.ListenAndServe(":8888", nil)
}
