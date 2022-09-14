package main

import "fmt"

type Student struct {
	name  []string
	score []int
}

func (s Student) Average() float64 {
	var sum int
	var result float64
	//var konv float64
	for i := 0; i < len(s.score); i++ {
		sum = sum + s.score[i]
	}
	result = float64(sum) / float64(len(s.score))

	return result
}

func (s Student) Min() (min int, name string) {
	min = s.score[0]
	for i := 0; i < len(s.score); i++ {
		if min > s.score[i] {
			min = s.score[i]
			name = s.name[i]
		}
	}
	return min, name
}

func (s Student) Max() (max int, name string) {
	max = s.score[0]
	for j := 0; j < len(s.score); j++ {
		if max < s.score[j] {
			max = s.score[j]
			name = s.name[j]
		}
	}
	return max, name

}

func main() {
	var a = Student{}

	for i := 0; i < 6; i++ {
		var name string
		fmt.Print("Input " + fmt.Sprint(i) + " Student's Name : ")
		fmt.Scan(&name)
		a.name = append(a.name, name)

		var score int
		fmt.Print("Input " + name + " Score : ")
		fmt.Scan(&score)
		a.score = append(a.score, score)
	}

	fmt.Println("\n\nAverage Score Students is", a.Average())
	scoreMax, nameMax := a.Max()
	fmt.Println("Max Score Students is : "+nameMax+" (", scoreMax, ")")
	scoreMin, nameMin := a.Min()
	fmt.Println("Min Score Students is : "+nameMin+" (", scoreMin, ")")
}
