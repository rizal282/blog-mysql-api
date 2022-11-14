package controllers

import "github.com/rizal282/api-golang-mux/src/api/middlewares"

func (s *Server) initializeRoutes() {

	// home route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login route
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetAllUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUserByID)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	// Post route
	s.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(s.CreatePost)).Methods("POST")
	s.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(s.GetAllPosts)).Methods("GET")
	s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(s.GetPostByID)).Methods("GET")
	s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdatePost))).Methods("PUT")
	s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareAuthentication(s.DeletePost)).Methods("DELETE")

}