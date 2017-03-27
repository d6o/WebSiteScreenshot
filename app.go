package main

type App struct {
	Config Config
}

func NewApp(Config Config) *App {
	return &App{
		Config: Config,
	}
}
