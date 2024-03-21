package controllers

import (
    "imoveis-back/database"
    "imoveis-back/models"
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
)

func GetHouseById(w http.ResponseWriter, r *http.Request) {
    db := database.GetDatabase_SQL()
    vars := mux.Vars(r)
    id := vars["id"]
    var house models.Casa
    result := db.First(&house, id)
    if result.Error != nil {
        w.WriteHeader(http.StatusNotFound)
        fmt.Fprintf(w, "Casa com o ID %s não encontrado", id)
        return
    }
    json.NewEncoder(w).Encode(house)
}

func GetAllHouses(w http.ResponseWriter, r *http.Request) {
    db := database.GetDatabase_SQL()
    var houses []models.Casa
    db.Find(&houses)
    json.NewEncoder(w).Encode(houses)
}

func InsertHouse(w http.ResponseWriter, r *http.Request) {
    db := database.GetDatabase_SQL()
    var houses models.Casa
    if err := json.NewDecoder(r.Body).Decode(&houses); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "Erro ao decodificar o corpo da requisição: %s", err)
        return
    }
    db.Create(&houses)
    w.WriteHeader(http.StatusCreated)
    fmt.Fprint(w, "Casa criada com sucesso")
}

func UpdateHouse(w http.ResponseWriter, r *http.Request) {
    db := database.GetDatabase_SQL()
    vars := mux.Vars(r)
    id := vars["id"]
    var houses models.Casa
    if err := json.NewDecoder(r.Body).Decode(&houses); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "Erro ao decodificar o corpo da requisição: %s", err)
        return
    }
    db.Model(&houses).Where("id = ?", id).Updates(&houses)
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, "Casa atualizada com sucesso")
}

func DeleteHouse(w http.ResponseWriter, r *http.Request) {
    db := database.GetDatabase_SQL()
    vars := mux.Vars(r)
    id := vars["id"]
    var houses models.Casa
    db.Delete(&houses, id)
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, "casa deletado com sucesso")
}
