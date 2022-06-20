package main

type Ingredient struct {
	Name    string  `json:"name"`
	Amount  float64 `json:"amount"`
	IsFluid bool    `json:"isFluid"`
}

type Recipe struct {
	Name        string       `json:"name"`
	Ingredients []Ingredient `json:"ingredients"`
}
