package gol

import (
	"fmt"

	"github.com/lucasb-eyer/go-colorful"
)

func GetRandomColor() string {

	return colorful.HappyColor().Hex()
}

//GetAverageColor of 3 colors with simple average
func GetAverageColor(parents *[]string) string {
	fmt.Println("***********************************************")
	fmt.Printf("calculating average of color %+v\n", *parents)
	//r, g, b := colorful.Hex("#FF0000").FastLinearRgb()

	//////////////////////////////////////////////////
	// test with red/blue

	c1, _ := colorful.Hex((*parents)[0])
	c2, _ := colorful.Hex((*parents)[1])
	c3, _ := colorful.Hex((*parents)[2])

	cNew, _ := colorful.Hex("#000000")
	cNew.R = (c1.R + c2.R + c3.R) / 3
	cNew.B = (c1.B + c2.B + c3.B) / 3
	cNew.G = (c1.G + c2.G + c3.G) / 3

	//cRedHex := "#FF0000"
	//cRedHex
	fmt.Printf("**%+v\n", cNew.Hex())
	return cNew.Hex()
}
