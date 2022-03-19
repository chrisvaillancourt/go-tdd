package iteration

const defaultRepeatCount = 5

func Repeat(character string, count int) string {
	if count == 0 {
		count = defaultRepeatCount
	}
	var repeated string
	for i := 0; i < count; i++ {
		repeated += character
	}
	return repeated
}
