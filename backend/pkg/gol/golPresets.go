package gol

var NewUserMessage = GolMessage{GolMsgType: "chat", Payload: "New User Joined"}

func InsertDummyData(b *[][]Cell) {
	(*b)[0][0].Breed("#ff0000")
	(*b)[1][1].Breed("#ff0000")
	(*b)[2][2].Breed("#ff0000")
	(*b)[3][3].Breed("#ff0000")
	(*b)[4][4].Breed("#ff0000")
}

func PreloadBlinker(b *[][]Cell, cHex string) {
	(*b)[1][2].Breed(cHex)
	(*b)[2][2].Breed(cHex)
	(*b)[3][2].Breed(cHex)
}

func PreloadToad(b *[][]Cell, cHex string) {
	(*b)[2][2].Breed(cHex)
	(*b)[2][3].Breed(cHex)
	(*b)[2][4].Breed(cHex)
	(*b)[3][1].Breed(cHex)
	(*b)[3][2].Breed(cHex)
	(*b)[3][3].Breed(cHex)

}
func PreloadBeacon(b *[][]Cell, cHex string) {
	(*b)[1][1].Breed(cHex)
	(*b)[1][2].Breed(cHex)
	(*b)[2][1].Breed(cHex)

	(*b)[4][4].Breed(cHex)
	(*b)[4][3].Breed(cHex)
	(*b)[3][4].Breed(cHex)

}

// func  PreloadBlinker(b *[][]Cell, cHex string){
// 	(*b)[0][0].Breed(CHex)
// 	(*b)[1][1].Breed(CHex)
// 	(*b)[2][2].Breed(CHex)
// 	(*b)[3][3].Breed(CHex)
// 	(*b)[4][4].Breed(CHex)
// }
