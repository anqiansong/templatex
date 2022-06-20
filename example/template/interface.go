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
		"sayHi": func(arg ...any) (any, error) {
			var list []string
			for _, v := range arg {
				list = append(list, fmt.Sprintf("%v", v))
			}
			return fmt.Sprintf("Hi %s", strings.Join(list, " & ")), nil
		},
	}).Parse(`{{sayHi "Elton" 1}}`)
	if err != nil {
		log.Fatalln(err)
	}

	if err = t.Execute(os.Stdout, nil); err != nil {
		log.Fatalln(err)
	}

}
