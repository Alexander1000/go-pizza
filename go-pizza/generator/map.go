package generator

import (
	"io"
	"math/rand"
	"time"

	"github.com/Alexander1000/go-pizza/go-pizza"
)

func GenerateMap(writer io.Writer, height int, width int) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			n := rand.Intn(len(go_pizza.IngredientList))
			writer.Write([]byte(go_pizza.IngredientList[n].Letter))
		}
		writer.Write([]byte{0x0A})
	}
}
