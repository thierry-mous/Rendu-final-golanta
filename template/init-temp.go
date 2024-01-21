package template

import (
	"fmt"
	"html/template"
	"os"
)

var Temp *template.Template

func InitTemplate() {
	temp, errTemp := template.ParseGlob("./template/*.html")
	if errTemp != nil {
		fmt.Printf("Erreur template: %v", errTemp.Error())
		os.Exit(1)
	}
	Temp = temp
}
