package httpcheck

import (
	"time"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"strconv"
)

type Health struct {
	Name      string
	Status    int
	Time      time.Time
}

func NewHealth(n string, s int) Health {
	return Health{
		Name: n,
		Status: s,
		Time: time.Now(),
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	sAdd := strings.Split(r.RemoteAddr, ":")
	dim := NewDimension()
	dim.Add("client_id", sAdd[0])
	dim.Add("client_port", sAdd[1])
	fmt.Fprint(w, fmt.Sprintf("Welcome: %s\n", sAdd[0]))
	LogRequest(r, "/", now, dim)
}

func ShowHealth(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	sAdd := strings.Split(r.RemoteAddr, ":")
	dim := NewDimension()
	dim.Add("client_id", sAdd[0])
	dim.Add("client_port", sAdd[1])
	health := NewHealth("http", 200)
	json.NewEncoder(w).Encode(health)
	LogRequest(r, "/health", now, dim)
}

func LogRequest(r *http.Request,ep string, start time.Time, dim Dimensions) {
	now := time.Now()
	dur := now.Sub(start).Nanoseconds()/1000000
	msg := fmt.Sprintf("request:+1|c %s\n", dim.String())
	fmt.Printf(msg)
	msg = fmt.Sprintf("duration:%d|ms %s\n", dur, dim.String())
	fmt.Printf(msg)
	msg = fmt.Sprintf("duration:%d|ms %s\n", dur, dim.String())
	fmt.Printf(msg)
}


func ComputePi(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	sAdd := strings.Split(r.RemoteAddr, ":")
	dim := NewDimension()
	dim.Add("client_id", sAdd[0])
	dim.Add("client_port", sAdd[1])
	num := strings.TrimPrefix(r.URL.Path, "/pi/")
	if num == "/pi" || num == "" {
		num = "9999"
	}
	dim.Add("pi_num", num)
	n, _ := strconv.Atoi(num)
	res := pi(n)
	fmt.Fprint(w, fmt.Sprintf("Welcome: pi(%s)=%f\n", num, res))
	LogRequest(r, "/pi", now, dim)
}