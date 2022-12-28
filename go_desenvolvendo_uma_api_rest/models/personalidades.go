package models

type Personalidade struct {
	Id       int    `Json:"id"`
	Nome     string `Json:"nome"`
	Historia string `Json:"historia"`
}

var Personalidades []Personalidade
