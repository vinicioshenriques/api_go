package handlers

import (
	"api_go/src/repository/postgres/dbconfig"
	"api_go/src/repository/postgres/model"
	"database/sql"
	"encoding/json"
	"net/http"
)

// SaveUserHandler é um handler que salva um usuário
func SaveUserHandler(writer http.ResponseWriter, request *http.Request) {
	db, errDbOpen := sql.Open(dbconfig.PostgresDriver, dbconfig.DataSourceName)
	if errDbOpen != nil {
		http.Error(writer, "Erro ao conectar com o banco.", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var u model.User
	errDecode := json.NewDecoder(request.Body).Decode(&u)
	if errDecode != nil {
		http.Error(writer, "Erro ao decodificar o usuário.", http.StatusInternalServerError)
		return
	}

	_, errExec := db.Exec("INSERT INTO users (name, email) VALUES ($1, $2);", u.Name, u.Email)
	if errExec != nil {
		http.Error(writer, "Erro ao salvar o usuário.", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
}

// ListUsersHandler é um handler que retorna a lista de usuários
func ListUsersHandler(w http.ResponseWriter, r *http.Request) {
	db, errDbOpen := sql.Open(dbconfig.PostgresDriver, dbconfig.DataSourceName)
	if errDbOpen != nil {
		http.Error(w, "Erro ao conectar com o banco.", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, errQuery := db.Query("SELECT id, name, email FROM users;")
	if errQuery != nil {
		http.Error(w, "Erro ao buscar usuários.", http.StatusInternalServerError)
		return
	}

	var users []model.User
	for rows.Next() {
		var u model.User
		rows.Scan(&u.ID, &u.Name, &u.Email)
		users = append(users, u)
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
	db, errDbOpen := sql.Open(dbconfig.PostgresDriver, dbconfig.DataSourceName)
	if errDbOpen != nil {
		http.Error(w, "Erro ao conectar com o banco.", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var u model.User
	errDecode := json.NewDecoder(r.Body).Decode(&u)
	if errDecode != nil {
		http.Error(w, "Erro ao decodificar o usuário.", http.StatusInternalServerError)
		return
	}

	_, errExec := db.Exec("UPDATE users SET name = $1, email = $2 WHERE id = $3;", u.Name, u.Email, u.ID)
	if errExec != nil {
		http.Error(w, "Erro ao atualizar o usuário.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeleteUserHandler é um handler que deleta um usuário
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	db, errDbOpen := sql.Open(dbconfig.PostgresDriver, dbconfig.DataSourceName)
	if errDbOpen != nil {
		http.Error(w, "Erro ao conectar com o banco.", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID é obrigatório.", http.StatusBadRequest)
		return
	}

	_, errExec := db.Exec("DELETE FROM users WHERE id = $1;", id)
	if errExec != nil {
		http.Error(w, "Erro ao deletar o usuário.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
