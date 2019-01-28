package main

import (
	"log"
	"os"
	"time"
	"text/template"
	"math"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("index.gohtml"))
}

func monthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}

func double(x int) int {
	return x + x
}

func square(x int) float64 {
	return math.Pow(float64(x),2)
}

func sqRoot(x float64) float64{
	return math.Sqrt(x)
}

var fm = template.FuncMap{
	"fdateMDY": monthDayYear,
	"fdbl": double,
	"fsq": square,
	"fsqrt": sqRoot,
}

//time.Now() was used for the date

func main(){
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", 3)
	if err != nil {
		log.Fatal(err)
	}
}