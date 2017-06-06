package httpcheck

import (
	"github.com/gorilla/mux"
	"github.com/codegangsta/cli"
	"fmt"
	"log"
	"net/http"
	"net"
)


type Httpd struct {
	Metrics map[string]Metric
}

func NewHttpd() Httpd {
	return Httpd{
		Metrics: map[string]Metric{},
	}
}

func (h *Httpd) Run(ctx *cli.Context) {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/health", ShowHealth)
	port := ctx.Int("port")
	host := ctx.String("host")
	addr := fmt.Sprintf("%s:%d", host, port)
	l, err := net.Listen("tcp4", addr)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.Serve(l, router))
}

/*func (h *Httpd) SendRequestMetric(ip string) {
	m, ok := h.Metrics[ip]
	if !ok {
		m := NewMetric(ip)
		h.Metrics[ip] = m
	}
	h.Metrics[ip].Increment()
}

func (h *Httpd) SubmitRequestMetrics() {

}
*/