package rest

const childrenRoutePath string = "/childrens"

func (fiberApp *App) initPublicChildrenRoutes() {
	v1 := fiberApp.apiUser.Group(childrenRoutePath + "/bind")

	v1.Post("/", ChildrenGreat()) // bind children : same as login system
}

func (fiberApp *App) initUserChildrenRoutes() {
	v1 := fiberApp.apiUser.Group(childrenRoutePath)

	v1.Use(parseUserAuthorizationHeader())
	v1.Get("/", ChildrenGreat())                // get all children
	v1.Get("/:children_id", ChildrenGreat())    // get children detail
	v1.Post("/", ChildrenGreat())               // add new children
	v1.Put("/:children_id", ChildrenGreat())    //edit existing children
	v1.Delete("/:children_id", ChildrenGreat()) // delete existing children
}
