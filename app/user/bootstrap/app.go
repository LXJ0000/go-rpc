package bootstrap

import (
	"github.com/LXJ0000/go-rpc/orm"
	"github.com/LXJ0000/go-rpc/utils/snowflake"
)

type Application struct {
	Env *Env
	Orm orm.Database
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Orm = NewOrmDatabase(app.Env)
	//log.Init(app.Env.AppEnv) // TODO
	snowflake.Init(app.Env.Server.SnowFlakeStartTime, app.Env.Server.SnowFlakeMachine)

	return *app
}
