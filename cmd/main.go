package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/wujunwei/sidecaragent/pkg/graceful"
	"github.com/wujunwei/sidecaragent/pkg/httpserver"
	"log"
	"net/http"
	"os"
)

var (
	address         string
	innerAddress    string
	appName         string
	callbackAddress string
	appSecret       string
	appClientID     string
	enableTLS       bool
	cert            string
	key             string
)

func init() {
	flag.StringVar(&address, "address", ":80", "the host(ip) and port for the forward proxy")
	flag.StringVar(&innerAddress, "innerAddress", "127.0.0.1:808", "the host(ip) and port for inner service")
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
		log.Println(err)
		os.Exit(1)
	}
	stop := graceful.SetupSignalHandler()
	forwardProxy := httpserver.NewForwardProxy(address)
	reverseProxy := httpserver.NewReverseProxy(innerAddress)
	var err error
	go func() {
		if enableTLS {
			err = forwardProxy.ListenAndServeTLS(cert, key)
		} else {
			err = forwardProxy.ListenAndServe()
		}
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	go func() {
		err = reverseProxy.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	<-stop
	log.Println("server start to stop.")
	_ = reverseProxy.Shutdown(context.Background())
	_ = forwardProxy.Shutdown(context.Background())
}
