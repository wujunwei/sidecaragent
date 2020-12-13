package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/wujunwei/sidecaragent/pkg/graceful"
	"github.com/wujunwei/sidecaragent/pkg/httpserver"
	"github.com/wujunwei/sidecaragent/pkg/util"
	"os"
)

var (
	listen          string
	appName         string
	callbackAddress string
	appSecret       string
	appClientID     string
	enableTLS       bool
	cert            string
	key             string
)

var log = util.Logger.Named("main")

func init() {
	flag.StringVar(&listen, "listen", ":80", "the host(ip) and port for listening")
	flag.StringVar(&appName, "name", "app", "the name of this application")
	flag.StringVar(&callbackAddress, "upstream", "127.0.0.1:8080", "the upstream of this agent")
	flag.StringVar(&appSecret, "secret", "", "the client secret of this application")
	flag.StringVar(&appClientID, "id", "", "the client id of this application")
	flag.BoolVar(&enableTLS, "enable-tls", false, "if this flag is set true, the server with run with tls.")
	flag.StringVar(&cert, "cert", "", "the path to the tls cert")
	flag.StringVar(&key, "key", "", "the path of the key to the cert")
}

func checkFlag() error {
	if appClientID == "" || appSecret == "" || appName == "" {
		return fmt.Errorf("app's client id , secret and name is empty string")
	}
	if enableTLS && (cert == "" || key == "") {
		return fmt.Errorf("if the tls is enabled, the cert and key can't be empty string")
	}
	return nil
}

func main() {
	flag.Parse()
	if err := checkFlag(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
	stop := graceful.SetupSignalHandler()
	server := httpserver.NewServer(listen)
	go func() {
		<-stop
		log.Infof("server start to stop.")
		_ = server.Shutdown(context.Background())
	}()
	var err error
	if enableTLS {
		err = server.ListenAndServeTLS(cert, key)
	} else {
		err = server.ListenAndServe()
	}
	if err != nil {
		log.Error(err)
	}
}
