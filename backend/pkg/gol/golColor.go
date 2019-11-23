package gol

import (
	"github.com/lucasb-eyer/go-colorful"
)

func GetRandomColor() string {

	return colorful.HappyColor().Hex()
}
