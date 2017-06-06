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
	fmt.Fprint(w, fmt.Sprintf("Welcome: %s\n", sAdd[0]))
	LogRequest(r, "/", now)
}

func ShowHealth(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	health := NewHealth("http", 200)
	json.NewEncoder(w).Encode(health)
	LogRequest(r, "/health", now)
}

func LogRequest(r *http.Request,ep string, start time.Time) {
	now := time.Now()
	dur := now.Sub(start).Nanoseconds()/1000000
	sAdd := strings.Split(r.RemoteAddr, ":")
	msg := fmt.Sprintf("request:+1|c client_ip=%s,endpoint=%s\n", sAdd[0], ep)
	fmt.Printf(msg)
	msg = fmt.Sprintf("duration:%d|ms client_ip=%s,endpoint=%s\n", dur, sAdd[0], ep)
	fmt.Printf(msg)
}


func ComputePi(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	num := strings.TrimPrefix(r.URL.Path, "/pi/")
	n, _ := strconv.Atoi(num)
	res := pi(n)
	//sAdd := strings.Split(r.RemoteAddr, ":")
	fmt.Fprint(w, fmt.Sprintf("Welcome: pi(%s)=%.f\n", num, res))
	LogRequest(r, "/pi", now)
}