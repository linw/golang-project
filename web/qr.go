package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

var addr = flag.String("addr", ":8091", "http service address") // Q=17, R=18

var templ = template.Must(template.New("qr").Parse(templateStr))

func main() {
	flag.Parse()
	http.Handle("/", http.HandlerFunc(QR))
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
