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
	// Rota para deletar um usuário
	mux.HandleFunc("DELETE /users", handlers.DeleteUserHandler)
	// Rota para atualizar um usuário
	mux.HandleFunc("PUT /users", handlers.UpdateUserHandler)

	return mux
}
