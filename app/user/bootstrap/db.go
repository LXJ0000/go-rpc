package bootstrap

import (
	"github.com/LXJ0000/go-rpc/app/user/internal/domain"
	"github.com/LXJ0000/go-rpc/orm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func NewOrmDatabase(env *Env) orm.Database {

	db, err := gorm.Open(mysql.Open(env.MySQL.DSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}

	if err = db.AutoMigrate(
		&domain.User{},
	); err != nil {
		log.Fatal(err)
	}

	database := orm.NewDatabase(db)

	return database
}
