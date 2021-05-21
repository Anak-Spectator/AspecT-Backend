package rest

func (fiberApp *App) initChildrenRoutes() {
	v1 := fiberApp.apiChildren.Group("/v1/")

	v1.Get("/", ChildrenGreat())
}
