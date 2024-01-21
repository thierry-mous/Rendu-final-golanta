package controller

import (
	"encoding/json"
	"fmt"
	"golanta/data"
	"golanta/template"
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func IndexPage(w http.ResponseWriter, r *http.Request) {
	template.Temp.ExecuteTemplate(w, "index", nil)
}

func PersoPage(w http.ResponseWriter, r *http.Request) {
	fileData, fileErr := os.ReadFile("data.json")
	if fileErr != nil {
		http.Error(w, fmt.Sprintf("Erreur de lecture du fichier : %s", fileErr), http.StatusInternalServerError)
		return
	}

	var dataPerso data.AventuriersData

	errJson := json.Unmarshal(fileData, &dataPerso)
	if errJson != nil {
		http.Error(w, fmt.Sprintf("Erreur de lecture du fichier : %s", errJson), http.StatusInternalServerError)
		return
	}
	fmt.Println(dataPerso)
	fmt.Println("2")
	template.Temp.ExecuteTemplate(w, "perso", dataPerso)
}

func CreatePage(w http.ResponseWriter, r *http.Request) {
	template.Temp.ExecuteTemplate(w, "create", nil)
}

func Modify(w http.ResponseWriter, r *http.Request) {
	var aventurierData data.AventuriersData
	queryParams := r.URL.Query()

	idParam := queryParams.Get("id")

	id, _ := strconv.Atoi(idParam)

	jsonData, _ := ioutil.ReadFile("data.json")

	json.Unmarshal(jsonData, &aventurierData)

	var aventurierRecherche data.Perso
	for _, aventurier := range aventurierData.Aventuriers {
		if aventurier.ID == id {
			aventurierRecherche = aventurier
			break
		}
	}

	if aventurierRecherche.ID == 0 {
		fmt.Println("Aventurier non trouvé avec l'ID:", id)
		http.Error(w, "Aventurier non trouvé", http.StatusNotFound)
		return
	}

	template.Temp.ExecuteTemplate(w, "modify", aventurierRecherche)
}

func FormJson(w http.ResponseWriter, r *http.Request) {

	dataJson := "data.json"

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "erreur", http.StatusInternalServerError)
		return
	}

	age := r.FormValue("age")      //string
	ageInt, _ := strconv.Atoi(age) //int

	id := data.GenerateID()

	form := data.Perso{
		ID:          id,
		Nom:         r.FormValue("nom"), //string
		Age:         ageInt,             //int
		Sexe:        r.FormValue("sexe"),
		Affiliation: r.FormValue("affiliation"),
		SkinColor:   r.FormValue("skincolor"),
		Hair:        r.FormValue("hair"),
		HairColor:   r.FormValue("haircolor"),
		Beard:       r.FormValue("beard"),
		Cyber:       r.FormValue("cyber"),
		Upgrade:     r.FormValue("upgrade"),
	}

	dataForm := data.ChargePerso()

	dataForm.Aventuriers = append(dataForm.Aventuriers, form)

	dataWrite,_ := json.Marshal(dataForm)


	errWriteFile := os.WriteFile(dataJson, dataWrite, fs.FileMode(0644))
	if errWriteFile != nil {
		http.Error(w, fmt.Sprintf("erreur écriture fichier: %v", errWriteFile), http.StatusInternalServerError)
		return
	}

	fmt.Println("données ajoutées au json")
	http.Redirect(w, r, "http://localhost:8080/perso", http.StatusSeeOther)
}

func DeletePerso(w http.ResponseWriter, r *http.Request) {
	jsonData, _ := ioutil.ReadFile("data.json")
	queryParams := r.URL.Query()

	idParam := queryParams.Get("id")

	idToDelete, _ := strconv.Atoi(idParam)

	type AventuriersData struct {
		Vchara []data.Perso `json:"characters"`
	}

	var aventurierData AventuriersData
	json.Unmarshal(jsonData, &aventurierData)

	if data.SupprimerVParID(idToDelete, &aventurierData.Vchara) {

		jsonUpdated, _ := json.MarshalIndent(aventurierData, "", "  ")

		ioutil.WriteFile("data.json", jsonUpdated, os.ModePerm)

		http.Redirect(w, r, "/perso", http.StatusSeeOther)
	} else {
		fmt.Printf("personnage avec ID %d non trouvé.\n", idToDelete)
	}

}

func ModifyCharaHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	CharaId := r.FormValue("character")
	fmt.Println(CharaId)
	CharaIdInt, _ := strconv.Atoi(CharaId)
	updatedChara, err := data.GetCharacter(CharaIdInt)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la récupération du personnage : %s", err), http.StatusInternalServerError)
		return
	}

	fmt.Println(updatedChara.Nom)

	w.WriteHeader(http.StatusOK)

	// Utiliser updatedChara dans votre logique de modification ici

	template.Temp.ExecuteTemplate(w, "modify", updatedChara)
}


func SubmitModif(w http.ResponseWriter, r *http.Request) {
	idStr := r.PostFormValue("id")
	id, _ := strconv.Atoi(idStr)

	data, _ := ioutil.ReadFile("data.json")

	type Perso struct {
		ID          int    `json:"ID"`
		Nom         string `json:"nom"`
		Age         int    `json:"age"`
		Sexe        string `json:"sexe"`
		Affiliation string `json:"affiliation"`
		SkinColor   string `json:"skincolor"`
		Hair        string `json:"hair"`
		HairColor   string `json:"haircolor"`
		Beard       string `json:"beard"`
		Cyber       string `json:"cyber"`
		Upgrade     string `json:"upgrade"`
	}

	type AventuriersData struct {
		Aventuriers []Perso `json:"characters"`
	}
	
	var aventuriersData AventuriersData
	json.Unmarshal(data, &aventuriersData)
	fmt.Println(aventuriersData)
	index := -1
	for i, aventurier := range aventuriersData.Aventuriers {
		fmt.Println(aventurier.ID)
		fmt.Println(id)
		if aventurier.ID == id {
			index = i
			break
		}
	}

	aventuriersData.Aventuriers[index].Nom = r.PostFormValue("nom")
	aventuriersData.Aventuriers[index].Age, _ = strconv.Atoi(r.PostFormValue("age"))
	aventuriersData.Aventuriers[index].Sexe = r.PostFormValue("sexe")
	aventuriersData.Aventuriers[index].Affiliation = r.PostFormValue("affiliation")
	aventuriersData.Aventuriers[index].SkinColor = r.PostFormValue("skincolor")
	aventuriersData.Aventuriers[index].Hair = r.PostFormValue("hair")
	aventuriersData.Aventuriers[index].HairColor = r.PostFormValue("haircolor")
	aventuriersData.Aventuriers[index].Beard = r.PostFormValue("beard")
	aventuriersData.Aventuriers[index].Cyber = r.PostFormValue("cyber")
	aventuriersData.Aventuriers[index].Upgrade = r.PostFormValue("upgrade")
	nouvellesDonneesJSON, _ := json.MarshalIndent(aventuriersData, "", "  ")

	ioutil.WriteFile("data.json", nouvellesDonneesJSON, 0644)

	http.Redirect(w, r, "/perso", http.StatusSeeOther)
}

