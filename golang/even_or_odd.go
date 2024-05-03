package kata

func EvenOrOdd(number int) string {
	if number&1 == 1 {
		return "Odd"
	} else {
		return "Even"
	}
}
