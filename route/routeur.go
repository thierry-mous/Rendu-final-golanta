package route

import (
	"fmt"
	"golanta/controller"
	"log"
	"net/http"
	"os"
)

func InitServ() {
	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/asset"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	http.HandleFunc("/index", controller.IndexPage)
	http.HandleFunc("/perso", controller.PersoPage)
	http.HandleFunc("/treatment", controller.FormJson)
	http.HandleFunc("/create", controller.CreatePage)
	http.HandleFunc("/delete", controller.DeletePerso)
	http.HandleFunc("/modify", controller.Modify)
	http.HandleFunc("/treatment_modif", controller.SubmitModif)

	//Init serv
	log.Println(" Serveur lancé !")
	fmt.Println("http://localhost:8080/index")
	http.ListenAndServe("localhost:8080", nil)
}
