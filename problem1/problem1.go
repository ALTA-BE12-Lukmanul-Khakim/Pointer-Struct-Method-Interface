package main

import "fmt"

func Swap(a, b *int) (int, int) {
	tempa := *a //memasukkan data
	tempb := *b
	*b = tempa //proses dibalik
	*a = tempb
	aa := int(*a)
	bb := int(*b)

	return aa, bb

}

func main() {
	a := 40
	b := 20

	Swap(&a, &b) //pemanggilan nilai
	fmt.Println(a, b)
}
