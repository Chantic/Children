package main

import "fmt"

type Apple struct {
	Number int
}

type Box struct {
	Apples []Apple
	Number int
}
type Car struct {
	Box []Box
}

// метод сортировки номеров номеров яблок
func (r Box) SortApples() {

	for i := 0; i < len(r.Apples)-1; i++ {
		for j := 0; j < len(r.Apples)-1-i; j++ {
			if r.Apples[j].Number > r.Apples[j+1].Number {
				r.Apples[j].Number, r.Apples[j+1].Number = r.Apples[j+1].Number, r.Apples[j].Number
			}
		}
	}
}
func main() {
	one := Apple{
		44,
	}
	two := Apple{
		550,
	}
	three := Apple{
		3344,
	}
	four := Apple{
		606534,
	}
	five := Apple{
		3458095,
	}
	six := Apple{
		3249,
	}
	pecage := []Apple{two, one, three, five, four, six}
	// сортировка коробки с яблоками
	//for i := 0; i < len(pecage)-1; i++ {
	//	for j := 0; j < len(pecage)-1-i; j++ {
	//		if (pecage[j].Number) > (pecage[j+1].Number) {
	//			pecage[j].Number, pecage[j+1].Number = pecage[j+1].Number, pecage[j].Number
	//		}
	//	}
	//}
	//fmt.Printf("%+v\n", pecage)
	box2 := Box{
		pecage,
		5,
	}
	box3 := Box{
		pecage[:1],
		6,
	}
	box := Box{
		pecage,
		7,
	}
	car := Car{
		[]Box{box, box2, box3},
	}
	// цикл с сортировкой при использование метода
	//for i := 0; i < len(car.Box); i++ {
	//	fmt.Println(car.Box[i].Number)
	//
	//	for j := 0; j < len(car.Box[i].Apples); j++ {
	//		car.Box[i].SortApples()
	//	}
	//}

	//тройной цикл с I
	// пройтись циклом по машине с боксами
	for i := 0; i < len(car.Box); i++ {
		// пройтись машине с боксами и посмотреть яблоки
		for j := 0; j < len(car.Box[i].Apples)-1; j++ {
			// вернуться и опять пройтись по машине и посмотреть яблоки
			for k := 0; k < len(car.Box[i].Apples)-1-j; k++ {
				// условия для сортировки номеров яблок в машине с боксами
				//if car.Box[i].Apples[k].Number > car.Box[i].Apples[k+1].Number {
				//сортировка при выполнение условия
				//	car.Box[i].Apples[k].Number, car.Box[i].Apples[k+1].Number = car.Box[i].Apples[k+1].Number, car.Box[i].Apples[k].Number

			}

		}

	}
	// тройной цикл с range
	//первый цикл, пройтись по боксам в машине
	for _, box := range car.Box {
		//box.SortApples()
		// второй, перейти к яблокам
		for _, _ = range box.Apples {
			// третьий чтобы вернуться обратно в бокс
			for i, pointerApple := range box.Apples {
				//условия чтобы не выходил за цикл
				if i == len(box.Apples)-1 {
					continue
				}
				// условия чтобы сортировать
				if pointerApple.Number > box.Apples[i+1].Number {
					//смена при выполнение условия
					box.Apples[i], box.Apples[i+1] = box.Apples[i+1], box.Apples[i]
				}
			}

		}

	}
	fmt.Printf("%+v\n", car)

}
