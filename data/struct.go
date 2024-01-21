package data

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

var V []Perso

type AventuriersData struct {
	Aventuriers []Perso `json:"characters"`
}

