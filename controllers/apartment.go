package controllers

import (
    "imoveis-back/database"
    "imoveis-back/models"
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
)

func GetApartamentById(w http.ResponseWriter, r *http.Request) {
    db := database.GetDatabase_SQL()
    vars := mux.Vars(r)
    id := vars["id"]
    var apartamento models.Apartamentos
    result := db.First(&apartamento, id)
    if result.Error != nil {
        w.WriteHeader(http.StatusNotFound)
        fmt.Fprintf(w, "Apartamento com o ID %s não encontrado", id)
        return
    }
    json.NewEncoder(w).Encode(apartamento)
}

func GetAllApartaments(w http.ResponseWriter, r *http.Request) {
    db := database.GetDatabase_SQL()
    var apartamentos []models.Apartamentos
    db.Find(&apartamentos)
    json.NewEncoder(w).Encode(apartamentos)
}

func InsertApartament(w http.ResponseWriter, r *http.Request) {
    db := database.GetDatabase_SQL()
    var apartamento models.Apartamentos
    if err := json.NewDecoder(r.Body).Decode(&apartamento); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "Erro ao decodificar o corpo da requisição: %s", err)
        return
    }
    db.Create(&apartamento)
    w.WriteHeader(http.StatusCreated)
    fmt.Fprint(w, "Apartamento criado com sucesso")
}

func UpdateApartament(w http.ResponseWriter, r *http.Request) {
    db := database.GetDatabase_SQL()
    vars := mux.Vars(r)
    id := vars["id"]
    var apartamento models.Apartamentos
    if err := json.NewDecoder(r.Body).Decode(&apartamento); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "Erro ao decodificar o corpo da requisição: %s", err)
        return
    }
    db.Model(&apartamento).Where("id = ?", id).Updates(&apartamento)
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, "Apartamento atualizado com sucesso")
}

func DeleteApartament(w http.ResponseWriter, r *http.Request) {
    db := database.GetDatabase_SQL()
    vars := mux.Vars(r)
    id := vars["id"]
    var apartamento models.Apartamentos
    db.Delete(&apartamento, id)
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, "Apartamento deletado com sucesso")
}
