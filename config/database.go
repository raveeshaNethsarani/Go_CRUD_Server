//connct go application to postgresql database using gorm
package	config
import	(
"fmt"
"log"
"os"
"github.com/joho/godotenv"
"gorm.io/driver/postgres"
"gorm.io/gorm"
)
//store the database connection in a global variable
var	DB	*gorm.DB
func	ConnectDatabase()	{
//	Load	.env	file
err	:=	godotenv.Load()
if	err	!=	nil	{
log.Println("No	.env	file	found,	reading	from	system	env")
}
dsn	:=	fmt.Sprintf(
"host=%s	user=%s	password=%s	dbname=%s	port=%s	sslmode=disable",
os.Getenv("DB_HOST"),
os.Getenv("DB_USER"),
os.Getenv("DB_PASSWORD"),
os.Getenv("DB_NAME"),
os.Getenv("DB_PORT"),
)

//connnect to the database using gorm
db,	err	:=	gorm.Open(postgres.Open(dsn),	&gorm.Config{})
if	err	!=	nil	{
log.Fatal("Failed	to	connect	to	database:	",	err)
}
log.Println("Database	connected	successfully!")
DB	=	db
}