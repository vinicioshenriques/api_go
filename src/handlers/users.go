package handlers

import (
	"api_go/src/repository/postgres/dbconfig"
	"api_go/src/repository/postgres/model"
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
)

var db *gorm.DB

func init() {
	var err error
	db, err = dbconfig.InitDB()
	if err != nil {
		panic("failed to connect database")
	}
}

// SaveUserHandler é um handler que salva um usuário
func SaveUserHandler(writer http.ResponseWriter, request *http.Request) {
	var u model.User
	errDecode := json.NewDecoder(request.Body).Decode(&u)
	if errDecode != nil {
		http.Error(writer, "Erro ao decodificar o usuário.", http.StatusInternalServerError)
		return
	}

	result := db.Create(&u)
	if result.Error != nil {
		http.Error(writer, "Erro ao salvar o usuário.", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
}

// ListUsersHandler é um handler que retorna a lista de usuários
func ListUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []model.User
	result := db.Find(&users)
	if result.Error != nil {
		http.Error(w, "Erro ao buscar usuários.", http.StatusInternalServerError)
		return
	}

	// Retornar os usuários em JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	errEncode := json.NewEncoder(w).Encode(users)
	if errEncode != nil {
		http.Error(w, "Erro ao buscar usuários.", http.StatusInternalServerError)
		return
	}
}

// UpdateUserHandler é um handler que atualiza um usuário
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var u model.User
	errDecode := json.NewDecoder(r.Body).Decode(&u)
	if errDecode != nil {
		http.Error(w, "Erro ao decodificar o usuário.", http.StatusInternalServerError)
		return
	}

	result := db.Save(&u)
	if result.Error != nil {
		http.Error(w, "Erro ao atualizar o usuário.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeleteUserHandler é um handler que deleta um usuário
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID é obrigatório.", http.StatusBadRequest)
		return
	}

	result := db.Delete(&model.User{}, id)
	if result.Error != nil {
		http.Error(w, "Erro ao deletar o usuário.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
