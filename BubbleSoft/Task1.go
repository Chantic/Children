package main

import "fmt"

type Person struct {
	Name   string
	Pocket float64
}

func main() {
	array := []int{1, 29, 3, 38, 12, 56, 7, 8, 312, 42, 2, 1, 7, 9}
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-1-i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}

		}
	}
	fmt.Println(array[0], array[len(array)-1])
	fmt.Println(array)
	Zakharkiv := Person{
		Name:   "Ilya",
		Pocket: 5555,
	}
	Toibaeva := Person{
		Name:   "Yana",
		Pocket: 35534,
	}
	Goncharova := Person{
		Name:   "Mari",
		Pocket: 5566,
	}
	Yalymova := Person{
		Name:   "K",
		Pocket: 5555,
	}
	Zhansugurov := Person{
		Name:   "Erik",
		Pocket: 55333,
	}
	Lebedev := Person{
		Name:   "Alex",
		Pocket: 4,
	}
	Zangegeh := Person{
		Name:   "Aziz",
		Pocket: 0.56,
	}
	//1 создать массив
	var Company []Person = []Person{Zakharkiv, Yalymova, Zhansugurov, Zangegeh, Lebedev, Toibaeva, Goncharova}
	for i := 0; i < len(Company)-1; i++ {
		for j := 0; j < len(Company)-1-i; j++ {
			if []rune(Company[1].Name)[0] > []rune(Company[2].Name)[0] {
				Company[j], Company[j+1] = Company[j+1], Company[j]
			}

		}
		// 1,9,4,3  0  0
		// 1 9 4 3  0  1
		// 1 9 4 3  0  2
		// 1 4 9 3  0  3

		// 1 4 3 9  1  0
		// 1,9,4,3  1  1
		// 1 9 4 3  1  2
		// 1 9 4 3  1  3

		// 1 4 9 3  2  0
		// 1 4 3 9  2  1
		// 1 3 4 9  2  2

	}
	fmt.Println(Company)
}
