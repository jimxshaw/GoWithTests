package iteration

func Repeat(character string, howManyTimes int) string {
	if howManyTimes <= 0 || howManyTimes > 1000 {
		return "howManyTimes can only be between 1 and 1000"
	}

	var repeated string
	for i := 0; i < howManyTimes; i++ {
		repeated += character
	}

	return repeated
}
