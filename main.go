package main

import (
	"os"
	"github.com/codegangsta/cli"
	"github.com/qnib/go-httpcheck/lib"
)


func main() {
	app := cli.NewApp()
	app.Description = "ETC metrics collector based on qframe, inspired by qcollect,logstash and fullerite"
	app.Name = "go-httpcheck"
	app.Usage = "gcollect-metrics [options]"
	app.Version = "1.1.3"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "0.0.0.0",
			Usage: "Bind host for the webserver.",
			EnvVar: "LISTEN_HOST",
		},
		cli.IntFlag{
			Name:  "port",
			Value: 8080,
			Usage: "Listen port for the webserver.",
			EnvVar: "LISTEN_PORT",
		},
		cli.BoolFlag{
			Name:  "log-stdout-disable",
			Usage: "Disable logging to stdout",
			EnvVar: "LOG_STDOUT_DISABLE",
		},
		cli.BoolFlag{
			Name:  "log-tcp",
			Usage: "Log to TCP server",
			EnvVar: "LOG_TCP",
		},
		cli.StringFlag{
			Name:  "log-tcp-target",
			Value: "tasks.qframe-agent.11001",
			Usage: "TCP Server to send logs to",
			EnvVar: "LOG_TCP_TARGET",
		},
	}
	httpd := httpcheck.NewHttpd()
	app.Action = httpd.Run
	app.Run(os.Args)
}
