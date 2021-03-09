package main

import "fmt"

func neighboursList(i, width int) (list []int) {
	area := width * width

	//corners
	switch i {
	case 0:
		list = append(list, 1)
		list = append(list, width)
		list = append(list, width+1)
		return
	case width - 1:
		list = append(list, width-2)
		list = append(list, 2*(width-1))
		list = append(list, (2*width)-1)
		return
	case area - width:
		list = append(list, area-width)
		list = append(list, area-2*width)
		list = append(list, area-width+1)
		return
	case area - 1:
		list = append(list, area-1-width)
		list = append(list, area-2-width)
		list = append(list, area-2)
		return
	}

	//top line
	if 0 < i && i < width-1 {
		list = append(list, i-1)
		list = append(list, i+1)
		list = append(list, i+width)
		list = append(list, i+width-1)
		list = append(list, i+width+1)
		return
	}

	//bottom line
	if area-width < i && i < area-1 {
		list = append(list, i-1)
		list = append(list, i+1)
		list = append(list, i-width)
		list = append(list, i-width-1)
		list = append(list, i-width+1)
		return
	}

	//left column
	if i%width == 0 {
		list = append(list, i-width)
		list = append(list, i+width)
		list = append(list, i+1)
		list = append(list, i-width+1)
		list = append(list, i+width+1)
		return
	}

	//right column
	if i%width == width-1 {
		list = append(list, i-width)
		list = append(list, i+width)
		list = append(list, i-1)
		list = append(list, i-width-1)
		list = append(list, i+width-1)
		return
	}

	//all remaining
	list = append(list, i-width)
	list = append(list, i+width)
	list = append(list, i-1)
	list = append(list, i+1)
	list = append(list, i-width+1)
	list = append(list, i+width+1)
	list = append(list, i-width-1)
	list = append(list, i+width-1)
	return
}

func testUnitMap() {

	var width int = 5
	var tab []int

	area := width * width

	for i := 0; i < area; i++ {
		neighbours := neighboursList(i, width)
		size := len(neighbours)

		if size == 3 {
			if i == 0 || i == width-1 || i == area-width || i == area-1 {
				fmt.Println("PASSED index", i, "list", neighbours)
				tab = append(tab, size)
			} else {
				fmt.Println("FAILED index", i)
			}
		}
		if size == 5 {
			if (i > 0 && i < width-1) || (i > area-width && i < area-1) {
				fmt.Println("PASSED index", i, "list", neighbours)
				tab = append(tab, size)
			} else {
				if i%width == 0 || i%width == width-1 {
					fmt.Println("PASSED index", i, "list", neighbours)
					tab = append(tab, size)
				} else {
					fmt.Println("FAILED index", i)
				}
			}
		}
		if size == 8 {
			if i%width != 0 || i%width != width-1 {
				fmt.Println("PASSED index", i, "list", neighbours)
				tab = append(tab, size)
			} else {
				fmt.Println("FAILED index", i)
			}

		}

	}

	fmt.Println("\nDisplay number of neighbours for each cell\n")
	for i := 0; i < area; i++ {
		if i%width < width-1 {
			fmt.Printf("%d ", tab[i])
		} else {
			fmt.Println(tab[i])
		}
	}
}
