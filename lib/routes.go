package httpcheck

import (
	"time"
	/*"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"strconv"
	*/
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

