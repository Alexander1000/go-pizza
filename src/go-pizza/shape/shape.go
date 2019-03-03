package shape

type Shape struct {
	Height int
	Width int
}

// Generate - возвращает все возможные варианты фигур для которых будет соблюдено
// min - минимальная площадь, max - максимальная площадь
func Generate(min int, max int) []Shape {
	result := make([]Shape, 0)
	for i := 1; i <= max; i++ {
		for j := 1; j <= max; j++ {
			if i * j >= min && i * j <= max {
				result = append(result, Shape{Height: i, Width: j})
			}
		}
	}
	return result
}
