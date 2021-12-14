package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "root:admin@tcp(127.0.0.1:3306)/player"

type Player struct {
	gorm.Model
	UserName string `json:"username"`
	NickName string `json:"nickname"`
	Grade    int    `json:"grade"`
	Gp       int    `json:"gp"`
	Power    int    `json:"power"`
	State    bool   `json:"state"`
}

func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	DB.AutoMigrate(&Player{})
}

func GetPlayers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var players []Player
	DB.Find(&players)
	json.NewEncoder(w).Encode(players)
}

func GetPlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var player []Player
	DB.First(&player, params["id"])
	json.NewEncoder(w).Encode(player)
}

func AddPlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var player Player
	json.NewDecoder(r.Body).Decode(&player)
	DB.Create(&player)
	json.NewEncoder(w).Encode(player)
}

func DeletePlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var player []Player
	DB.Delete(&player, params["id"])
	json.NewEncoder(w).Encode("Player deleted is successfuly! ")

}

func UpdatePlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var player []Player
	DB.First(&player, params["id"])
	json.NewDecoder(r.Body).Decode(&player)
	DB.Save(&player)
	json.NewEncoder(w).Encode(player)
}
