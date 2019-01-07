/*
Main package for goexpose binary.

Goexpose provides several command line arguments such as:
* config - configuration file
* format - format of configuration file (json, yaml), default is json

*/
package cmd

import (
	"fmt"
	"flag"
	"os"
	"github.com/golang/glog"
	"github.com/fm0l/goexpose/internal"
)

func Execute() {
	var configVar string
	formatVar := flag.String("format", "json", "Configuration file format. (json, yaml)")
	configTmp := flag.String("config", "", "Configuration full path of file")
	// Parse command line flags
	flag.Parse()

	// Configuration filename derived from format if not specified
	if *configTmp == "" {
		configVar = fmt.Sprintf("config.%s", *formatVar)
	} else {
		configVar = *configTmp
	}

	var (
		config *goexpose.Config
		server *goexpose.Server
		err    error
	)

	// read config file
	// if config, err = goexpose.NewConfigFromFilename(*configVar, *formatVar); err != nil {
	if config, err = goexpose.NewConfigFromFilename(configVar, *formatVar); err != nil {
		glog.Errorf("config error: %v", err)
		os.Exit(1)
	}

	// change working directory to config directory
	if err = os.Chdir(config.Directory); err != nil {
		glog.Errorf("config error: %v", err)
		os.Exit(1)
	}

	if server, err = goexpose.NewServer(config); err != nil {
		glog.Errorf("server error: %v", err)
		os.Exit(1)
	}

	if err = server.Run(); err != nil {
		glog.Errorf("server run error: %v", err)
		os.Exit(1)
	}
}
