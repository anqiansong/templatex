package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

func main() {
	t, err := template.New("x").Funcs(template.FuncMap{
		"sayHi": func(arg ...string) (string, error) {
			return fmt.Sprintf("Hi %s", strings.Join(arg[0:], " & ")), nil
		},
	}).Parse(`{{sayHi "Elton" "Kam"}}`)
	if err != nil {
		log.Fatalln(err)
	}

	if err = t.Execute(os.Stdout, nil); err != nil {
		log.Fatalln(err)
	}

}
