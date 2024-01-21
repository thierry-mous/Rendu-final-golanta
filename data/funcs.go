package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

var Datajson = "data.json"

func ChargePerso() AventuriersData {

	var aventuriersData AventuriersData

	jsonData, _ := ioutil.ReadFile("data.json")

	json.Unmarshal(jsonData, &aventuriersData)
	fmt.Println("1")
	return aventuriersData
}

func GenerateID() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(9000) + 1000

}

func ModifyChara(updatedChara Perso) error {

	character := ChargePerso()
	

	for i, chara := range character.Aventuriers {
		if chara.ID == updatedChara.ID {
			character.Aventuriers[i] = updatedChara
		}
	}

	if err := ChangeChara(character.Aventuriers); err != nil {
		return err
	}

	return nil
}

func ChangeChara(character []Perso) error {

	data, err := json.MarshalIndent(character, "   ", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(Datajson, data, 0666)
	if err != nil {
		return err
	}

	return nil
}

func GetCharacter(id int) (Perso, error) {
	characters := ChargePerso()

	for _, character := range characters.Aventuriers {
		if character.ID == id {
			return character, nil
		}
	}

	return Perso{}, nil
}

func SupprimerVParID(id int, Vchara *[]Perso) bool {
	index := -1
	for i, a := range *Vchara {
		if a.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		return false
	}

	*Vchara = append((*Vchara)[:index], (*Vchara)[index+1:]...)

	return true
}
