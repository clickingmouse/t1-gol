package gol

// Cell for Game of Life
type Cell struct {
	X            int      `json:"x"`
	Y            int      `json:"y"`
	Alive        bool     `json:"status"` // alive or dead
	ColorHex     string   `json:"colorHex"`
	PreviousLife bool     `json:"previousLife"`
	Parents      []string `json:"parentsColors"`
	// Now          bool
	// Next         bool
	// ThisGen      string
	// NextGen      string
	// Retired      bool
}
