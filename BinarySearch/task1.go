package main

import "fmt"

type Student struct {
	LastName string
}
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
		"Fast",
	}
	// первая задача создать группу с учениками
	group := []Student{Ilya, Aurora, Cerdic, Lana, Scott, Timo, Natali}
	fmt.Println("before", group)
	index, exist := binarysearch1(group, "Klammer")
	if exist == false {
		fmt.Println("not exist ")
	} else {
		fmt.Println(group[index], index)
	}
	fmt.Println(group)
	//ниже использую бинарный поиск на массиве интовых значений
	//a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	//binarysearch(a, 12)
	//target, exist := binarysearch(a, 12)
	//if exist == false {
	//	fmt.Println("not exist")
	//} else {
	//	fmt.Println(target)
	//}
	//

	//group1 := []Student{Ilya, Aurora}
	//group2 := []Student{Timo, Lana}
	//group3 := []Student{Natali, Scott, Cerdic}
	//	a := Class{
	//		1,
	//		group1,
	//	}
	//	fmt.Println(a)
	//	b := Class{
	//		2,
	//		group2}
	//
	//	fmt.Println(b)
}

// 1 2 3 4 5
// 1 5
// 2 5
// 3 5
// 4 5
// 5 5 done
//
// binarysearch output index return
func binarysearch(array []int, target int) (int, bool) {

	low := 0
	high := len(array) - 1
	for low <= high {

		midllindex := (high + low) / 2
		if array[midllindex] == target {
			return midllindex, true
		}
		if array[midllindex] < target {
			low = midllindex + 1
		}
		if array[midllindex] > target {
			high = midllindex - 1
		}
		fmt.Println(low, high, midllindex)
	}
	return 0, false
} //отсортировать массив студентов в рчунчую или другой функцией
// и использовать с массивом студентов бинарный поиск
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

// binarysearch1, поиск  искомого элемента, возращает индекс массива, true найден false не найден
func binarysearch1(array []Student, searchValue string) (int, bool) {
	//сортировка массива
	QuickSort3(array, 0, len(array)-1)
	//поиск искомомго элемента
	low := 0
	high := len(array) - 1
	for low <= high {

		midllindex := (high + low) / 2
		if array[midllindex].LastName == searchValue {
			return midllindex, true
		}
		if array[midllindex].LastName < searchValue {
			low = midllindex + 1
		}
		if array[midllindex].LastName > searchValue {
			high = midllindex - 1
		}
		fmt.Println(low, high, midllindex)
	}
	return 0, false
}
