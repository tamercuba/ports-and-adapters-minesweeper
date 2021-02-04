package domain

// BoardSettings contains the game settings
type BoardSettings struct {
	Size  uint `json:"size"`
	Bombs uint `json:"bombs"`
}
