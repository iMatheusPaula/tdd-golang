package interaction

func Interaction(character string, quantityToRepeat int) string {
	var result string
	for i := 0; i < quantityToRepeat; i++ {
		result += character
	}
	return result
}
