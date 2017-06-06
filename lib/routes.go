package httpcheck

import (
	"time"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
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
	sAdd := strings.Split(r.RemoteAddr, ":")
	fmt.Fprint(w, fmt.Sprintf("Welcome: %s\n", sAdd[0]))
	msg := fmt.Sprintf("request:+1|c client_ip=%s,endpoint=index\n", sAdd[0])
	fmt.Println(msg)
}

func ShowHealth(w http.ResponseWriter, r *http.Request) {
	health := NewHealth("http", 200)
	json.NewEncoder(w).Encode(health)
}
