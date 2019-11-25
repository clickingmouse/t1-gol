package gol

// Cell for Game of Life
type Cell struct {
	X             int      `json:"x"`
	Y             int      `json:"y"`
	Alive         bool     `json:"status"` // alive or dead
	ColorHex      string   `json:"colorHex"`
	PreviousLife  bool     `json:"previousLife"`
	Parents       []string `json:"parentsColors"`
	PreviousColor string   `json:"previousColorHex"`
	// Now          bool
	// Next         bool
	// ThisGen      string
	// NextGen      string
	// Retired      bool
}

// for insert dummydatafunc
func (c *Cell) Breed(color string) {
	c.Alive = true
	c.ColorHex = color
}

//SavePreviousLife method
func (c *Cell) SavePreviousLife() {
	c.PreviousLife = c.Alive
	c.PreviousColor = c.ColorHex
}

//RetireCell method
func (c *Cell) RetireCell() {
	//	c.Retired = c.Alive
}

func (c *Cell) KillCell() {
	c.Alive = false
	c.ColorHex = "dead"
}
