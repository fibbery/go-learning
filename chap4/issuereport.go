package main

import (
	"fmt"
	"go-learning/chap4/geohash"
	"time"
)

const templ = `{{.TotalCount}} issues:
{{range.Items}}---------------------------------------
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title | printf "%.64s" }}
Age: {{.CreateAt | daysAgo }} days
{{end}}
`

func dayAgo(t time.Time) int {
	return int(time.Since(t).Hours()) / 24
}

func main() {
	str, area := geohash.Encode(31.1932993, 121.43960190000007, 6)
	fmt.Printf("this str %s, and area is %v", str, area)
}
