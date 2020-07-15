package request

type AddSong struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Link        string `json:"link" binding:"required"`
}

type UpdateSong struct {
	Target      int    `json:"target" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Link        string `json:"link" binding:"required"`
}

type AddList struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Public      *bool  `json:"public" binding:"required"`
}

type UpdateList struct {
	Target      int    `json:"target" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Public      *bool  `json:"public" binding:"required"`
}

type AddTag struct {
	Name string `json:"name" binding:"required"`
}

type UpdateTag struct {
	Name   string `json:"name" binding:"required"`
	Target int    `json:"target" binding:"required"`
}

type Bridge struct {
	Item   int `json:"item" binding:"required"`
	Target int `json:"target" binding:"required"`
}
