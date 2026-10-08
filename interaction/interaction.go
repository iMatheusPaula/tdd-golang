package interaction

const quantityToRepeat = 6

func Interaction(character string) string {
	var result string
	for i := 0; i < quantityToRepeat; i++ {
		result += character
	}
	return result
}
