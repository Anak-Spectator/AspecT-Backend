package rest

func (fiberApp *App) InitUserAccountRoutes() {
	v1 := fiberApp.apiV1.Group("/account")
	v1.Get("/")
}
