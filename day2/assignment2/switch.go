package assignment2

func SwitchStatment(para int) string {
	i := para
	switch i % 2 {
	case 0:
		return "even"
	default:
		return "odd"
	}
}
