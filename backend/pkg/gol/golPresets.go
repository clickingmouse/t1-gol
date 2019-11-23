package gol

var NewUserMessage = GolMessage{GolMsgType: "chat", Payload: "New User Joined"}

func InsertDummyData(b *[][]Cell) {
	(*b)[0][0].Breed("#ff0000")
	(*b)[1][1].Breed("#ff0000")
	(*b)[2][2].Breed("#ff0000")
	(*b)[3][3].Breed("#ff0000")
	(*b)[4][4].Breed("#ff0000")

}
