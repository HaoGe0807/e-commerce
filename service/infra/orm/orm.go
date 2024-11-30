package orm

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "gorm.io/driver/mysql"
	"log"
	"os"
)

var (
	//ignore mutex
	instance = map[string]*gorm.DB{}
	logger   = log.New(os.Stdout, "orm", log.LstdFlags)
)

func NewDB(name string, options ...Option) {
	op := defaultOption()
	for _, apply := range options {
		apply(&op)
	}

	dbUser := op.username + ":" + op.password
	connection := fmt.Sprintf("%s@tcp(%s)/%s?charset=%s&parseTime=true&loc=Local", dbUser, op.addr, op.dbname, op.charset)

	logger.Println("Connecting to Database:", connection)
	orm, err := gorm.Open("mysql", connection)
	if err != nil {
		log.Panic("Failed to connect to DB", err)
	}

	if os.Getenv("DB_DEBUG") == "1" {
		orm = orm.Debug()
	}

	orm.DB().SetMaxOpenConns((op.maxOpenConns)) //set max open connections
	orm.DB().SetMaxIdleConns((op.idleConns))    //allowed max idle connections, if mysql need connections < 20, the other open connections will be closed

	instance[name] = orm
	logger.Printf("Connet to database %s by %s successfully\n", connection, op.dbname)
}

func GetORM(name string) *gorm.DB {
	return instance[name]
}

func IsRecordNotFoundError(err error) bool {
	return gorm.IsRecordNotFoundError(err)
}
