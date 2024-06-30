package controllers

import (
	"database/sql"
	"encoding/json"
	"go-crud-api/config"
	"go-crud-api/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GET
func GetVendedor(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT idvendedor, cpf, logradouro, numero, telefone, perc_comissao, idcidade FROM vendedor")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var vendedores []models.Vendedor
	for rows.Next() {
		var vendedor models.Vendedor
		if err := rows.Scan(&vendedor.ID, &vendedor.Nome, &vendedor.CPF, &vendedor.Logradouro, &vendedor.Numero, &vendedor.Bairro, &vendedor.CEP, &vendedor.Telefone, &vendedor.Perc_comissao, &vendedor.IdCidade); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		vendedores = append(vendedores, vendedor)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vendedores)
}

//GET ID

func GetVendedort(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var vendedor models.Vendedor
	err = db.QueryRow("SELECT id, nome, cpf, lougradouro, numero, bairro, cep, telefone, perc_comissao, idcidade FROM users WHERE id = ?", id).Scan(&vendedor.ID, &vendedor.Nome, &vendedor.CPF, &vendedor.Logradouro, &vendedor.Numero, &vendedor.Bairro, &vendedor.CEP, &vendedor.Telefone, &vendedor.Perc_comissao, &vendedor.IdCidade)
	if err == sql.ErrNoRows {
		http.Error(w, "Vendedor not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vendedor)
}

//POST

func CreateVendedor(w http.ResponseWriter, r *http.Request) {
	var vendedor models.Vendedor
	if err := json.NewDecoder(r.Body).Decode(&vendedor); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO vendedor (nome, cpf, lougradouro, numero, bairro, cep, telefone, perc_comissao, idcidade) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?,)", &vendedor.ID, &vendedor.Nome, &vendedor.CPF, &vendedor.Logradouro, &vendedor.Numero, &vendedor.Bairro, &vendedor.CEP, &vendedor.Telefone, &vendedor.Perc_comissao, &vendedor.IdCidade)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	vendedor.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vendedor)
}

// PUT
func UpdateVendedor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var vendedor models.User
	if err := json.NewDecoder(r.Body).Decode(&vendedor); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec("UPDATE vendedor SET nome = ?, cpf = ?, logradouro = ?, numero = ?, bairro = ?, cep = ?, telefone = ?,  perc_comissao = ?, idcidade = ? WHERE idvendedor = ?", &vendedor.ID, &vendedor.Nome, &vendedor.CPF, &vendedor.Logradouro, &vendedor.Numero, &vendedor.Bairro, &vendedor.CEP, &vendedor.Telefone, &vendedor.Perc_comissao, &vendedor.IdCidade, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	vendedor.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vendedor)
}

// DELETE
func DeleteVendedor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid vendedor ID", http.StatusBadRequest)
		return
	}

	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM users WHERE idvendedor = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
