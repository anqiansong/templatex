package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/anqiansong/templatex"
)

func main() {
	t, err := templatex.New("x", templatex.WithBinary("./custom")).Funcs(template.FuncMap{
		"join": func(arg ...string) (string, error) {
			return fmt.Sprintf("Hi %s", strings.Join(arg[0:], " & ")), nil
		},
	}).Parse(`{{join "Elton" "Kam" "&&"}}
Hello {{join "A" "B" ","}}
`)
	if err != nil {
		log.Fatalln(err)
	}

	if err = t.Execute(os.Stdout, nil); err != nil {
		log.Fatalln(err)
	}

}
