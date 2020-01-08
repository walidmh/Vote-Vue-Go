package controllers

import "github.com/Muhammad-Tounsi/Vote-Go-Vue/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/user", middlewares.SetMiddlewareJSON(s.GetLoggedUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	//votes routes
	s.Router.HandleFunc("/votes", middlewares.SetMiddlewareJSON(s.Createvote)).Methods("POST")
	s.Router.HandleFunc("/votes", middlewares.SetMiddlewareJSON(s.Getvotes)).Methods("GET")
	s.Router.HandleFunc("/votes/{id}", middlewares.SetMiddlewareJSON(s.Getvote)).Methods("GET")
	s.Router.HandleFunc("/vote/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.Addvote))).Methods("PUT")
	s.Router.HandleFunc("/votes/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.Updatevote))).Methods("PUT")
	s.Router.HandleFunc("/vote/{id}", middlewares.SetMiddlewareAuthentication(s.Deleteuservote)).Methods("DELETE")
	s.Router.HandleFunc("/votes/{id}", middlewares.SetMiddlewareAuthentication(s.Deletevote)).Methods("DELETE")
}
