package code

import "fmt"

func swap(x, y *int) {

	*x = *x + *y
	*y = *x - *y
	*x = *x - *y

}

func RunQ1() {
	x := 10
	y := 20

	swap(&x, &y)
	fmt.Println(x, y)
}
