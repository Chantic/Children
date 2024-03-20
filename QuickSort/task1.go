package main

import "fmt"

// создать тип учеников
type Student struct {
	LastName string
}

// создать тип класс
type Class struct {
	Number   int
	Students []Student
}

func main() {
	Ilya := Student{
		"Ivanov",
	}
	Aurora := Student{
		"Favre",
	}
	Timo := Student{
		"Klammer",
	}
	Lana := Student{
		"Fcl",
	}
	Natali := Student{
		"Grueidy",
	}
	Scott := Student{
		"Chenier",
	}
	Cerdic := Student{
		"Ast",
	}
	// первая задача создать группу с учениками
	group := []Student{Ilya, Aurora, Cerdic, Lana, Scott, Timo, Natali}
	fmt.Println("before", group)
	QuickSort3(group, 0, len(group)-1)
	fmt.Println("after", group)
	////отсортировать их
	//pivot := group[len(group)-1]
	//wall := 0
	//for i := 0; i < len(group)-1; i++ {
	//	//wall= a[i]
	//	if pivot.LastName >= group[i].LastName {
	//		group[wall], group[i] = group[i], group[wall]
	//		wall++
	//	}
	//}
	//group[wall], group[len(group)-1] = group[len(group)-1], group[wall]
	//fmt.Println(group)
	//// создать группы с фамилиями учеников
	group1 := []Student{Ilya, Aurora}
	group2 := []Student{Timo, Lana}
	//group3 := []Student{Natali, Scott, Cerdic}
	a := Class{
		1,
		group1,
	}
	QuickSort3(group1, 0, len(group1)-1)
	//pivota := group[len(group1)-1]
	//walla := 0
	//for i := 0; i < len(group1)-1; i++ {
	//	//wall= a[i]
	//	if pivota.LastName >= group1[i].LastName {
	//		group1[walla], group1[i] = group1[i], group1[walla]
	//		walla++
	//	}
	//}
	fmt.Println(a)
	b := Class{
		2,
		group2,
		//}
		//pivotb := group[len(group2)-1]
		//wallb := 0
		//for i := 0; i < len(group2)-1; i++ {
		//	//wall= a[i]
		//	if pivotb.LastName >= group2[i].LastName {
		//		group2[wallb], group2[i] = group2[i], group2[wallb]
		//		wallb++
		//	}
	}
	QuickSort3(group2, 0, len(group2)-1)
	fmt.Println(b)
	//задача_2: у тебя есть классы(1, 2) остсортируй у каждого класа группы по фамилии ученика.
	//1 класс имеет множество груп(у группы есть литерал "a, b,c" и массив учеников)
}
func QuickSort3(array []Student, low int, high int) {
	if low > high {
		return
	}
	pivot := array[high]
	low2 := low
	for i := low; i < high-1; i++ {
		if pivot.LastName[0] >= array[i].LastName[0] {
			array[low2], array[i] = array[i], array[low2]
			low2++
		}

	}
	array[low2], array[high] = array[high], array[low2]
	QuickSort3(array, low, low2-1)
	QuickSort3(array, low2+1, high)
}
