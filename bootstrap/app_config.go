package bootstrap

import "gorm.io/gorm"

type Application struct {
	Env      *Env
	Database *gorm.DB
}

func App() Application {
	app := new(Application)
	app.Env = NewEnv()
	app.Database = NewPostgresDatabase(app.Env)

	return *app
}
