package viewmodel

type LoginForm struct {
	Room string `form:"room"`
	Name string `form:"name"`
}

type LoginQueryParams struct {
	LoginForm
	Valid string `form:"valid"`
}
