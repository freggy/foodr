package main

import (
	"fmt"
	"math/rand"
)

type Generator struct {
	recipes      []Recipe
	maxRetries   int
	lookbackSize int
}

// Generate generates a list of recipes that caps at a certain threshold.
func (g Generator) Generate(threshold int) []Recipe {
	selected := make([]Recipe, 0)
	for {
		if len(selected) == threshold {
			break
		}

		i := rand.Intn(len(g.recipes))
		tries := 0
		pick := g.recipes[i]

		if contains(selected[len(selected)-g.lookbackSize:], pick) {
			if tries == g.maxRetries {
				fmt.Printf("could not find different recipe within %d tries. adding %s to the list", g.maxRetries, pick.Name)
				selected = append(selected, pick)
			}
			fmt.Printf("%s is already picked within in the last 3 days. trying again\n", pick.Name)
			tries++
			continue
		}
		selected = append(selected, pick)
	}
	return selected
}

func contains(recipes []Recipe, r Recipe) bool {
	for _, val := range recipes {
		if val.Name == r.Name {
			return true
		}
	}
	return false
}
