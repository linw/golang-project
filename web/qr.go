package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
	"math/rand"
)
type eat_data struct {
EatRand string
TotalEat []string
RandSeed int
}

var addr = flag.String("addr", ":8091", "http service address") // Q=17, R=18

var templ = template.Must(template.New("qr").Parse(templateStr))
var templ_eat = template.Must(template.New("eat").Parse(templatestr_eat))

func main() {
	flag.Parse()
	http.Handle("/", http.HandlerFunc(QR))
	http.Handle("/eat", http.HandlerFunc(EAT))
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal(`ListenAndServe:`, err)
	}
}

func enter(s ...string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		if len(res) > 0 {
			res = fmt.Sprintf("%v [%v]", res, s[i])
		} else {
			res = fmt.Sprintf("[%v]", s[i])
		}
	}
	fmt.Printf("[%v] [%v] %v\n", time.Now().Format("2006-01-02 15:04:05"), "entering:", res)
	return res
}

func leav(s string) {
	fmt.Printf("[%v] [%v] %v\n", time.Now().Format("2006-01-02 15:04:05"), "leaving:", s)
}

func QR(w http.ResponseWriter, req *http.Request) {
	defer leav(enter(fmt.Sprintf("s=%v", req.FormValue("s")), fmt.Sprintf("req=%v", req.Form.Encode())))
	templ.Execute(w, req.FormValue("s"))
}



func EAT(w http.ResponseWriter, req *http.Request) {
	defer leav(enter(fmt.Sprintf("req=%v", req.Form.Encode())))
	data := eat_data{
		"",
		[]string{"麦当劳","肯德基","金渝川菜","面向八方","G先生","西少爷","蒸功夫"},
		0,
	}
	data.RandSeed = rand.Int()%len(data.TotalEat)
	data.EatRand = data.TotalEat[data.RandSeed]
	templ_eat.Execute(w, data)
}

const templateStr = `
<html>
<head>
<title>QR Link Generator</title>
</head>
<body>
{{if .}}
<img src="http://api.qrserver.com/v1/create-qr-code/?size=150x150&data={{.}}" />
<br>
{{.}}
<br>
<br>
{{end}}
<form action="/" name=f method="GET">
<input maxLength=1024 size=70 name=s value="" title="Text to QR Encode">
<input type=submit value="Show QR" name=qr>
</form>
</body>
</html>
`

const templatestr_eat = `
<html>
<head>
<title>吃什么</title>
</head>
<body>
<br>
<tr><td>所有可选餐馆</td></tr>
{{range .TotalEat}}<div>{{ . }}</div>{{else}}<div><strong>no rows</strong></div>{{end}}

<br>
<br>
<table style="margin: 0px; padding: 0px;">
<tr><td >本次选出的餐馆</td><td>{{.EatRand}}<td></tr>
<tr><td>本次选出的餐馆所用随机数</td><td>{{.RandSeed}}<td></tr>
</table>
</body>
</html>
`
