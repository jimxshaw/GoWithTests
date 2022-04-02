package iteration

func Repeat(character string, howManyTimes int) string {
	var repeated string
	for i := 0; i < howManyTimes; i++ {
		repeated += character
	}

	return repeated
}
