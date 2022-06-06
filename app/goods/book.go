package goods

type Book struct {
	ID     int    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
