package controller

type Controller struct {
	User interface{ User }
	Link interface{ Link }
}
