package models

import (
	"os"

	entity "github.com/edenriquez/load-balancer-proxy-go/pkg/proxy/entity/"
	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/native"
)

// SetUpDB creates the connection
func SetUpDB() {
	globalConnection := mysql.New(
		"tcp",
		"",
		os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
	)
	err := globalConnection.Connect()
	if err != nil {
		panic(err)
	}

	_, _, err = globalConnection.Query("CREATE TABLE IF NOT EXISTS products (id INT NOT NULL AUTO_INCREMENT, UNIQUE KEY id (id), name varchar(11), price double, sku varchar(255), image text)")
	if err != nil {
		panic("ERROR CREATING DATABASE" + err.Error())
	}
}

type repo struct {
	connection *mysql.Conn
}

func (r *repo) Find(id string) (*entity.Proxy, error) {
	result := entity.Proxy{}
	session := r.connection.Connect()
	coll := session.DB(os.Getenv("MONGODB_DATABASE")).C("user")
	err := coll.Find(bson.M{"_id": id}).One(&result)
	if err != nil {
		return nil, err
	}
}
