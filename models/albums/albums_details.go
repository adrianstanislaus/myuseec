package albums

type Album_detail struct {
	Id           int    `gorm:"primaryKey" json:"id"`
	Title        string `json:"title"`
	Album_type   string `json:"album_type"`
	About        string `json:"about"`
	Album_art    string `json:"album_art"`
	Release_date string `json:"release_date"`
}
