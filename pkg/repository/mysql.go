package repository

import (
	"fmt"
	"os"

	"github.com/edenriquez/load-balancer-proxy-go/api/utils"
	"github.com/edenriquez/load-balancer-proxy-go/pkg/entity"
	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/native"
)

const (
	tableQuery = `
CREATE TABLE IF NOT EXISTS proxy 
(
	id INT NOT NULL AUTO_INCREMENT, 
	UNIQUE KEY id (id), 
	domain varchar(250),
	UNIQUE (domain),
	weight INT, 
	priority INT
)`
	registriesSeed = `
INSERT IGNORE INTO proxy SET domain="alpha", weight=1, priority=3;
INSERT IGNORE INTO proxy SET domain="omega", weight=3, priority=2;
INSERT IGNORE INTO proxy SET domain="delta", weight=4, priority=1;`
)

//MysqlRepository mysql conn
type MysqlRepository struct {
	connection mysql.Conn
}

//NewMysqlRepository create new repository
func NewMysqlRepository() *MysqlRepository {
	conn := mysql.New(
		"tcp",
		"",
		os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
	)
	return &MysqlRepository{
		connection: conn,
	}
}

//Find a bookmark
func (r *MysqlRepository) Find(service string) (*entity.Proxy, error) {
	q := fmt.Sprintf(
		`
		SELECT * FROM proxy 
		where domain="%s"
		`,
		utils.SanitizeSQLParam(service),
	)

	r.connection.Connect()
	row, _, err := r.connection.QueryFirst(q)
	r.connection.Close()

	if len(row) == 0 {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &entity.Proxy{
		Domain:   row.Str(1),
		Weight:   row.Int(2),
		Priority: row.Int(3),
	}, nil
}

// Migrate should create proxy table if not exists
func (r *MysqlRepository) Migrate() {
	r.connection.Connect()

	_, _, err := r.connection.Query(tableQuery)
	_, _, err = r.connection.Query(registriesSeed)
	if err != nil {
		panic("Error executing migration" + err.Error())
	}
	defer r.connection.Close()
}
