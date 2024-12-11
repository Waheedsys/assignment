package assignment4

/*  input:
   myMap := map[string][]int{
		"A": {1, 2, 3},
		"B": {4, 3, 2},
		"C": {8, 2, 6},
	}
	output:
 	"A":{6}
	"B":{9}
	"C":{16}
*/

func SumValuesByKey(m map[string][]int) (s map[string]int, err error) {
	result := make(map[string]int)
	for key, values := range m {
		sum := 0
		for _, value := range values {
			sum += value
		}
		result[key] = sum
	}
	return result, nil
}
