package rest

func (fiberApp *App) InitAuthRoutes() {
	v1 := fiberApp.apiUser.Group("/auth")
	v1.Post("/login", userLogin(fiberApp.service.accSvc.accAppSvc))
	v1.Post("/register", userRegister(fiberApp.service.ssSvc.accAppSvc))
}
