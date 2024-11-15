package rotationalcipher

func RotationalCipher(plain string, shiftKey int) string {
<<<<<<< Updated upstream
	panic("Please implement the RotationalCipher function")
=======
	var result string

	for _, c := range plain {
		switch {
		case c >= 'a' && c <= 'z':
			result += string((c-'a'+rune(shiftKey))%26 + 'a')
		case c >= 'A' && c <= 'Z':
			result += string((c-'A'+rune(shiftKey))%26 + 'A')
		default:
			result += string(c)
		}
	}

	return result
>>>>>>> Stashed changes
}
