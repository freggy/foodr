package food

import (
	"testing"
)

func TestGeneratorLookbackSize(t *testing.T) {
	recipes := []Recipe{
		{
			Name: "1",
			Ingredients: []Ingredient{
				{
					Name:    "ingredient1",
					Amount:  3,
					IsFluid: false,
				},
			},
		},
		{
			Name: "2",
			Ingredients: []Ingredient{
				{
					Name:    "ingredient1",
					Amount:  3,
					IsFluid: false,
				},
			},
		},
		{
			Name: "3",
			Ingredients: []Ingredient{
				{
					Name:    "ingredient1",
					Amount:  3,
					IsFluid: false,
				},
			},
		},
		{
			Name: "4",
			Ingredients: []Ingredient{
				{
					Name:    "ingredient1",
					Amount:  3,
					IsFluid: false,
				},
			},
		},
		{
			Name: "5",
			Ingredients: []Ingredient{
				{
					Name:    "ingredient1",
					Amount:  3,
					IsFluid: false,
				},
			},
		},
		{
			Name: "6",
			Ingredients: []Ingredient{
				{
					Name:    "ingredient1",
					Amount:  3,
					IsFluid: false,
				},
			},
		},
	}

	g := Generator{
		maxRetries:   2,
		lookbackSize: 3,
		recipes:      recipes,
		seed:         int64(5),
	}

	gen := g.Generate(8)

	// with give configuration this is a correct sequence
	// that should always be produced by the generator
	expected := []string{"1", "5", "2", "4", "5", "5", "3", "1"}

	for i, val := range gen {
		if val.Name != expected[i] {
			t.Fatalf("expected %s found %s", expected[i], val.Name)
		}
	}
}
