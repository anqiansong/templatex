package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {
	type Person struct {
		Name string
		Age  int
	}
	var person = Person{
		Name: "Elton",
		Age:  12,
	}
	t, err := template.New("x").Funcs(template.FuncMap{
		"sayHi": func(arg ...any) (any, error) {
			data := arg[0].(Person)
			return fmt.Sprintf("Hi %s", data.Name), nil
		},
	}).Parse(`{{sayHi .data}}`)
	if err != nil {
		log.Fatalln(err)
	}

	if err = t.Execute(os.Stdout, map[string]interface{}{
		"data": person,
	}); err != nil {
		log.Fatalln(err)
	}

}
