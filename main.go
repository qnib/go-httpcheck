package main

import (
	"os"
	"github.com/codegangsta/cli"
	"github.com/qnib/go-httpcheck/lib"
)


func main() {
	app := cli.NewApp()
	app.Name = "ETC metrics collector based on qframe, inspired by qcollect,logstash and fullerite"
	app.Usage = "gcollect-metrics [options]"
	app.Version = "1.0.0"
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
		cli.StringFlag{
			Name:  "metrics-target",
			Value: "tasks.qframe-agent:11001",
			Usage: "TCP Server to send metrics to",
			EnvVar: "METRICS_TARGET",
		},
	}
	httpd := httpcheck.NewHttpd()
	app.Action = httpd.Run
	app.Run(os.Args)
}
