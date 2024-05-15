package router

import (
	"api_go/src/handlers"
	"net/http"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux()
	// Rota para listar usuários
	mux.HandleFunc("/users", handlers.ListUsersHandler)
	// Rota para salvar um usuário
	mux.HandleFunc("POST /users", handlers.SaveUserHandler)

	return mux
}
