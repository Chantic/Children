package main

import "fmt"

type Country struct {
	Name      string
	Districts []District
}

type District struct {
	Name   string
	Humans []Human
}

type Human struct {
	Name string
}

func (r Country) SortPeopleByName() {
	for i := 0; i < len(r.Districts); i++ {

		for j := 0; j < len(r.Districts[i].Humans)-1; j++ {

			for k := 0; k < len(r.Districts[i].Humans)-1-j; k++ {
				if r.Districts[i].Humans[k].Name[0] > r.Districts[i].Humans[k+1].Name[0] {

					r.Districts[i].Humans[k].Name, r.Districts[i].Humans[k+1].Name = r.Districts[i].Humans[k+1].Name, r.Districts[i].Humans[k].Name
				}
			}
		}
	}

}
func (r District) SortPeopleByName() {
	for i := 0; i < len(r.Humans)-1; i++ {
		for j := 0; j < len(r.Humans)-1; j++ {
			if r.Humans[j].Name > r.Humans[j+1].Name {

				r.Humans[j].Name, r.Humans[j+1].Name = r.Humans[j+1].Name, r.Humans[j].Name
			}

		}
	}

}
func main() {

	America := Country{
		Name: "America",
	}
	capitol := District{
		Name: "capitol",
	}

	d1 := District{
		Name: "luxury items",
		Humans: []Human{
			{Name: "Katniss Everdeen"},
			{Name: "Peeta Mellark"},
			{Name: "Gale Hawthorne"},
		},
	}

	d2 := District{
		Name: "masonry and defense ",
		Humans: []Human{
			{Name: "Effie Trinket"},
			{Name: "Haymitch Abernathy"},
		},
	}

	d3 := District{
		Name: "general electronics ",
		Humans: []Human{
			{Name: "Cato"},
			{Name: "Cloves"},
		},
	}

	d4 := District{
		Name: "fishing ",
		Humans: []Human{

			{Name: "Finnick Odair"},
			{Name: "Annie Cresta"},
		},
	}

	d5 := District{
		Name: "power/electricity",
		Humans: []Human{
			{Name: "Rue"},
			{Name: "Thresh"},
		},
	}

	d6 := District{
		Name: "transportation ",
		Humans: []Human{
			{Name: "Johanna Mason"},
			{Name: "Blight"},
		},
	}

	d7 := District{
		Name: "lumber ",
		Humans: []Human{
			{Name: "Beetee Latier"},
			{Name: "Wiress"},
		},
	}

	d8 := District{
		Name: "textiles ",
		Humans: []Human{
			{Name: "Cecelia"},
			{Name: "Woof"},
		},
	}

	d9 := District{
		Name: "grain ",
		Humans: []Human{
			{Name: "Cashmere"},
			{Name: "Gloss"},
		},
	}

	d10 := District{
		Name: "livestock",
		Humans: []Human{
			{Name: "Brutus"},
			{Name: "Enobaria"},
		},
	}

	d11 := District{
		Name: "agriculture ",
		Humans: []Human{
			{Name: "Chaff"},
			{Name: "Seeder"},
		},
	}

	d12 := District{
		Name: "coal mining ",
		Humans: []Human{
			{Name: "Cinna"},
			{Name: "Portia"},
		},
	}

	d13 := District{
		Name: "charge of nuclear weaponry ",
		Humans: []Human{
			{Name: "President Alma Coin"},
			{Name: "Plutarch Heavensbee"},
		},
	}

	// Append Districts to the Country
	America.Districts = append(America.Districts, capitol, d1, d2, d3, d4, d5, d6, d7, d8, d9, d10, d11, d12, d13)
	//fmt.Println(America.Districts)
	//fmt.Println(America.Name)
	//for _, dictrict := range America.Districts {
	//	fmt.Println(dictrict.Name)
	//	for _, v := range dictrict.Humans {
	//		fmt.Println(v.Name)
	//
	//	}
	//}
	//for i := 0; i < len(America.Districts); i++ {
	//	fmt.Println("District", America.Districts[i].Name)
	//	for j := 0; j < len(America.Districts[i].Humans); j++ {
	//		fmt.Println("Name", America.Districts[i].Humans[j].Name)
	//		if America.Districts[i].Humans[j].Name[0] == 'B' {
	//			fmt.Println("Name B", America.Districts[i].Humans[j].Name)
	//		}
	//	}
	//
	//}
	//for i := 0; i < len(America.Districts)-1; i++ {
	//	for j := 0; j < len(America.Districts)-1-i; j++ {
	//		if America.Districts[j].Name[0] > America.Districts[j+1].Name[0] {
	//			America.Districts[j].Name, America.Districts[j+1].Name = America.Districts[j+1].Name, America.Districts[j].Name
	//		}
	//
	//	}
	//}
	//for i := 0; i < len(America.Districts); i++ {
	//
	//	for j := 0; j < len(America.Districts[i].Humans)-1; j++ {
	//
	//		for k := 0; k < len(America.Districts[i].Humans)-1-j; k++ {
	//
	//			//if America.Districts[i].Humans[k].Name[0] > America.Districts[i].Humans[k+1].Name[0] {
	//
	//			//America.Districts[i].Humans[k].Name, America.Districts[i].Humans[k+1].Name = America.Districts[i].Humans[k+1].Name, America.Districts[i].Humans[k].Name
	//		}
	//	}
	//}
	//	for _, dictrict := range America.Districts {
	//		for _, _ = range dictrict.Humans {
	//			for i, pointerHumans := range dictrict.Humans {
	//if i == len(America.Districts)-1{
	//	continue
	//}
	//
	//
	//					if pointerHumans.Name > dictrict.Humans[i+1].Name {
	//						dictrict.Humans[i], dictrict.Humans[i+1] = dictrict.Humans[i+1], dictrict.Humans[i]
	//			}
	//		}
	//
	//	}
	//
	//}
	for _, district := range America.Districts {
		district.SortPeopleByName()
	}
	fmt.Printf("%+v\n", America.Districts)
}

//America.Districts

//fmt.Printf("%+v\n", America.Districts)
//fmt.Println(America.Districts[0])
