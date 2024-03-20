package main

import "fmt"

func main() {
	//BanIlya(10)
	//IlyaAsshole(3)
	a := []int{3, 5, 6, 213, 566, 34, 2, 13, 6}
	QuickSort2(a, 0, len(a)-1)
	fmt.Println(a)
}
func BanIlya(n int) {
	fmt.Println("I'm tired of thing that i have to do", n)
	if n <= 1 {
		return
	}
	BanIlya(n - 1)

	//функции которые вызывают саму себя это рекурсивыные
}
func IlyaAsshole(n int) {
	for i := 0; i < n; i++ {

		fmt.Println("FuckIlya")
	}

}
func QuickSort(a []int) {
	pivot := a[len(a)-1]
	wall := 0
	for i := 0; i < len(a)-1; i++ {
		//wall= a[i]
		if pivot >= a[i] {
			a[wall], a[i] = a[i], a[wall]
			wall++
		}

	}
	a[wall], a[len(a)-1] = a[len(a)-1], a[wall]

}

//3,6,213,566,2,13,6
// wall and i
// w 0 i 0 = 3 3 done pivot 6..... 3,6,213,566,2,13,6
// w 1 i 1 = 6 6 done pivot 6..... 3,6,213,566,2,13,6
// w 2 i 2 = 213 213 false pivot 6 ..... 3,6,213,566,2,13,6
// w 2 i 3 = 213 566 false pivot 6 ..... 3,6,213,566,2,13,6
// w 2 i 4 = 213 2 done pivot 6 ...... 3 6 2 566 213 13 6
// w 3 i 5 = 566 13 false pivot 6.....3 6 2 566 213 13 6
// w 3 i 6 = 566 6 done pivot 6 ......3 6 2 6 213 13 566
//
// sub a 1 = 3 6 2
// w 0 i 0 = 3 3 false pivot 2 .....3 6 2
// w 0 i 1 = 3 6 false pivot 2 ...... 3 6 2
// w 0 i 2 = 3 2 done pivot 2 ....... 2 6 3
// sub a 1 2 =  6 3
// w 0 i 0 = 6 6 false pivot 3 ......6 3
// w 0 i 1 = 6 3 done pivot 3...... 3 6
// w 0 i 0  =  2 6 done pivot 3 .....
// sub a 2 = 213 13 566
// nothing change
// sub a 2 1 = 213 13
// w 0 i 0 = 213 213 false pivot 13 .... 213 13
// w 0 i 1 = 213 13 done pivot 13 ........ 13 213

func QuickSort2(a []int, low int, high int) {
	if low > high {
		return
	}
	pivot := a[high]
	wall := low
	for i := low; i < high-1; i++ {
		//wall= a[i]
		if pivot >= a[i] {
			a[wall], a[i] = a[i], a[wall]
			wall++
		}
	}
	a[wall], a[high] = a[high], a[wall]
	QuickSort2(a, low, wall-1)
	QuickSort2(a, wall+1, high)
}
