package controllers

import (
    "imoveis-back/database"
    "imoveis-back/models"
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
)

func GetLandById(w http.ResponseWriter, r *http.Request) {
    db := database.GetDatabase_SQL()
    vars := mux.Vars(r)
    id := vars["id"]
    var terrenos models.Terreno
    result := db.First(&terrenos, id)
    if result.Error != nil {
        w.WriteHeader(http.StatusNotFound)
        fmt.Fprintf(w, "Terreno com o ID %s não encontrado", id)
        return
    }
    json.NewEncoder(w).Encode(terrenos)
}

func GetAllLands(w http.ResponseWriter, r *http.Request) {
    db := database.GetDatabase_SQL()
    var terrenos []models.Terreno
    db.Find(&terrenos)
    json.NewEncoder(w).Encode(terrenos)
}

func InsertLand(w http.ResponseWriter, r *http.Request) {
    db := database.GetDatabase_SQL()
    var terrenos models.Terreno
    if err := json.NewDecoder(r.Body).Decode(&terrenos); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "Erro ao decodificar o corpo da requisição: %s", err)
        return
    }
    db.Create(&terrenos)
    w.WriteHeader(http.StatusCreated)
    fmt.Fprint(w, "Terreno criado com sucesso")
}

func UpdateLand(w http.ResponseWriter, r *http.Request) {
    db := database.GetDatabase_SQL()
    vars := mux.Vars(r)
    id := vars["id"]
    var terrenos models.Terreno
    if err := json.NewDecoder(r.Body).Decode(&terrenos); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "Erro ao decodificar o corpo da requisição: %s", err)
        return
    }
    db.Model(&terrenos).Where("id = ?", id).Updates(&terrenos)
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, "Terreno atualizado com sucesso")
}

func DeleteLand(w http.ResponseWriter, r *http.Request) {
    db := database.GetDatabase_SQL()
    vars := mux.Vars(r)
    id := vars["id"]
    var terrenos models.Terreno
    db.Delete(&terrenos, id)
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, "Terreno deletado com sucesso")
}
