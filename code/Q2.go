package code

import "fmt"

//Your code goes here
type Rectangle struct {
	Width  int
	Height int
}

func (rect Rectangle) Area() int {
	return rect.Width * rect.Height
}

// Do not change the code in the main function
func RunQ2() {

	rect := Rectangle{Width: 10, Height: 5}
	fmt.Println("Area:", rect.Area())

}
