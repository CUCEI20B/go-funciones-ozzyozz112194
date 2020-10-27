package main

func Burbuja(numeros []int64) []int64 {

	for i := len(numeros); i > 0; i-- {
		for j := 1; j < i; j++ {
			if numeros[j-1] > numeros[j] {
				inter := numeros[j]
				numeros[j] = numeros[j-1]
				numeros[j-1] = inter
			}
		}
	}
	return numeros
}

func main() {

}
