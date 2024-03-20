package main

import "fmt"

// книга это структура
type Books struct {
	// два поля автор и название
	Author string
	Name   string
}
type Rack struct {
	Books  []Books
	Number int
}

// библиотека, стелажи, книги
type Libery struct {
	Name   string
	Number int
	Racks  []Rack
}

func (r Libery) SortRack() {
	for _, rack := range r.Racks {
		rack.SortByName()
	}
}

// 1 создать метод
func (r Libery) SortLiberyByAuthor() {
	//	//2 пройтись по стелажам
	for i := 0; i < len(r.Racks); i++ {
		//3 пройтись по книгам
		for j := 0; j < len(r.Racks[i].Books)-1; j++ {

			for k := 0; k < len(r.Racks[i].Books)-1-j; k++ {
				//4 сортировать по авторам
				if r.Racks[i].Books[k].Author > r.Racks[i].Books[k+1].Author {
					//6 отсортировать по авторам
					r.Racks[i].Books[k].Author, r.Racks[i].Books[k+1].Author = r.Racks[i].Books[k+1].Author, r.Racks[i].Books[k].Author
				}
			}
		}

	}
}

// 1 создать функцию у полки по сортировке книг по названию
func (r Rack) SortByName() {
	for i := 0; i < len(r.Books)-1; i++ {
		for j := 0; j < len(r.Books)-1; j++ {
			if r.Books[i].Name > r.Books[i+1].Name {
				r.Books[i], r.Books[i+1] = r.Books[i+1], r.Books[i]
			}
		}
	}

}

// 2 иная функция для сортировки полки по автору
func (r Rack) SortByAuthor() {
	for i := 0; i < len(r.Books)-1; i++ {
		for j := 0; j < len(r.Books)-1; j++ {
			if r.Books[j].Author[0] > r.Books[j+1].Author[0] {
				r.Books[j], r.Books[j+1] = r.Books[j+1], r.Books[j]
			}
		}
	}
}
func main() {
	//стелаж с номерами
	One := Books{
		Author: "Suzanne Collins",
		Name:   "Mockingjay",
	}
	Two := Books{
		Author: "Ransom Riggs ",
		Name:   "Home for children",
	}
	Three := Books{
		Author: "Harper Lee ",
		Name:   "To Kill A Mockingbird",
	}
	var Cabins []Books = []Books{One, Two, Three}
	r := Rack{
		Books:  Cabins,
		Number: 0,
	}

	//создать библиотеку, с стелажами наполнеными книгами
	L := Libery{
		Name:   "Shantal",
		Number: 1,
		Racks:  []Rack{r},
	}
	fmt.Printf("%+v\n", L)
	//fmt.Printf("%+v \n", r)
	//r.SortByName()
	//fmt.Printf("%+v \n", r)
	//r.SortByAuthor()
	//fmt.Printf("%+v \n", r)
	//L.SortRack()
	//fmt.Printf("%+v,\n", L)
	L.SortLiberyByAuthor()
	fmt.Printf("%+v\n", L)
	//var Libery [] Books= [] Books {}

}
