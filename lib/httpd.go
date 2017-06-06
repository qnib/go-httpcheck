package httpcheck

import (
	"github.com/gorilla/mux"
	"github.com/codegangsta/cli"
	"github.com/zpatrick/go-config"
	"fmt"
	"log"
	"net/http"
	"net"
	"strings"
	"encoding/json"
	"strconv"
	"time"
)


type Httpd struct {
	Cfg *config.Config
	Metrics map[string]Metric
}

func NewHttpd() Httpd {
	return Httpd{
		Metrics: map[string]Metric{},
	}
}

func (h *Httpd) Run(ctx *cli.Context) {
	cfg := config.NewConfig([]config.Provider{})
	cfg.Providers = append(cfg.Providers, config.NewCLI(ctx, false))
	h.Cfg = cfg

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", h.Index)
	router.HandleFunc("/health", h.ShowHealth)
	router.HandleFunc("/pi", h.ComputePi)
	router.HandleFunc("/pi/{{num}}", h.ComputePi)
	port := ctx.Int("port")
	host := ctx.String("host")
	addr := fmt.Sprintf("%s:%d", host, port)
	l, err := net.Listen("tcp4", addr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Start serving on %s", addr)
	log.Fatal(http.Serve(l, router))
}

func (h *Httpd) Index(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	sAdd := strings.Split(r.RemoteAddr, ":")
	dim := NewDimension()
	dim.Add("client_id", sAdd[0])
	dim.Add("client_port", sAdd[1])
	fmt.Fprint(w, fmt.Sprintf("Welcome: %s\n", sAdd[0]))
	h.LoqRequest(r, "/", now, dim)


}

func (h *Httpd) ShowHealth(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	sAdd := strings.Split(r.RemoteAddr, ":")
	dim := NewDimension()
	dim.Add("client_id", sAdd[0])
	dim.Add("client_port", sAdd[1])
	health := NewHealth("http", 200)
	json.NewEncoder(w).Encode(health)
	if sAdd[0] != "127.0.0.1" {
		h.LoqRequest(r, "/health", now, dim)
	}
}

func (h *Httpd) AssembleLogLines(r *http.Request,ep string, start time.Time, dim Dimensions) (res []string) {
	now := time.Now()
	dur := now.Sub(start).Nanoseconds()/1000000
	res = append(res, fmt.Sprintf("request:+1|c %s", dim.String()))
	res = append(res, fmt.Sprintf("duration:%d|ms %s", dur, dim.String()))
	return res
}

func (h *Httpd) LoqRequest(r *http.Request,ep string, start time.Time, dim Dimensions) {
	now := time.Now()
	msgs := []string{}
	dur := now.Sub(start).Nanoseconds() / 1000000
	msgs = append(msgs, fmt.Sprintf("request:+1|c %s", dim.String()))
	msgs = append(msgs, fmt.Sprintf("duration:%d|ms %s", dur, dim.String()))
	lstdout, _ := h.Cfg.Bool("log-stdout-disabled")
	ltcp, _ := h.Cfg.Bool("log-tcp")
	if ltcp {
		addr, _ := h.Cfg.String("log-tcp-target")
		for _, msg := range msgs {
			conn, _ := net.Dial("tcp", addr)
			fmt.Fprintf(conn, msg+"\n")
			conn.Close()
		}
	}
	if !lstdout {
		fmt.Println(strings.Join(msgs, "\n"))
	}
}

func (h *Httpd) ComputePi(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	sAdd := strings.Split(r.RemoteAddr, ":")
	dim := NewDimension()
	dim.Add("client_id", sAdd[0])
	dim.Add("client_port", sAdd[1])
	num := strings.TrimPrefix(r.URL.Path, "/pi/")
	if num == "/pi" || num == "" {
		num = "9999"
	}
	n, _ := strconv.Atoi(num)
	res := pi(n)
	fmt.Fprint(w, fmt.Sprintf("Welcome: pi(%s)=%f\n", num, res))
	h.LoqRequest(r, "/pi", now, dim)
}