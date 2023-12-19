package bootstrap

import "gorm.io/gorm"

type Application struct {
	Env      *Env
	Postgree *gorm.DB
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Postgree = NewPostgreeDatabase(app.Env)

	return *app
}
