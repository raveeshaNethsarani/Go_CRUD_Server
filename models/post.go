package	models
import	"time"

type Post struct {
	ID 			 uint   	`json:"id" gorm:"primaryKey"`
	Image		 string 	`json:"image" gorm:"not null"`
	Title 		 string 	`json:"title" gorm:"not null"`
	Content		 string 	`json:"content" gorm:"not null"`
	CreatedAt 	 time.Time 	`json:"created_at"`
	UpdatedAt 	 time.Time 	`json:"updated_at"`
}
