package httpcheck

import (
	"fmt"
	"strings"
	"sort"
)

type Dimensions struct {
	Map map[string]string
}

func NewDimension() Dimensions {
	return Dimensions{
		Map: map[string]string{},
	}
}

func NewPreDimension(dim map[string]string) Dimensions {
	return Dimensions{
		Map: dim,
	}
}

func (d *Dimensions) Add(key,val string) {
	d.Map[key] = val
 }

func (d *Dimensions) String() string {
	res := []string{}
	keys := []string{}
	for k, _ := range d.Map {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		res = append(res, fmt.Sprintf("%s=%s", k, d.Map[k]))
	}
	return strings.Join(res, ",")
}
