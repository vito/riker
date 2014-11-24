package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"
)

var rikers = []string{
	"http://i.imgur.com/IObN0tJ.gif",
	"http://i.imgur.com/DlmpR08.gif",
	"http://i.imgur.com/E9xIncq.gif",
	"http://i.imgur.com/hCcuFRm.gif",
	"http://i.imgur.com/Axc9V1j.gif",
	"http://i.imgur.com/Je5Rdv3.jpg",
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Riker here...")
		out := "<dl>"
		for _, env := range os.Environ() {
			kv := strings.Split(env, "=")
			out += "<dt>" + kv[0] + "</dt>"
			out += "<dd>" + kv[1] + "</dd>"
		}
		out += "</dl>"
		randomizer := rand.New(rand.NewSource(time.Now().UnixNano()))
		styledTemplate.Execute(w, Content{
			Envs:  `<div class="envs">` + out + `</div>`,
			Riker: rikers[randomizer.Intn(len(rikers))],
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Riker here, listening on port %s\n", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

var styledTemplate = template.Must(template.New("riker").Parse(`
<html>
<head>
<style>
body {
    font-family: "helveticaneue-light";
    font-size: 16px;
    color: #333;
    position: absolute;
    margin:0;
    width:100%;
    height:100%;
}

dt {
  color:#777;
}

img {
    margin:10px
}

.envs {
  margin:10px
}

</style>
</head>
<body>
<img src="{{.Riker}}">
{{.Envs}}
</body>
</html>
`))

type Content struct {
	Envs  string
	Riker string
}
