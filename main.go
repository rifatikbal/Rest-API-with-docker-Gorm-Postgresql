package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

type STUDENT_INFO struct {
	S_ID     string `json:"S_ID"`
	Name     string `json:"Name"`
	Village  string `json:"Village"`
	Thana    string `json:"Thana"`
	District string `json:"District"`
}

func connect() {

	dsn := "host=localhost user=admin password=secret dbname=School port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("err")
	} else {
		fmt.Println("I am connected")
		fmt.Printf("check")
	}

	db = d

	db.AutoMigrate(&STUDENT_INFO{})
}

func Delete_Student_with_ID(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-Type", "application/json")

	var g_student, tm_student []STUDENT_INFO

	ID := chi.URLParam(r, "id")
	db.Where("S_ID = ?", ID).Delete(&g_student)

	err := db.Find(&tm_student).Error

	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(tm_student)

}

func Add_New_Student(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-Type", "application/json")

	var s_student []STUDENT_INFO

	var temp1student STUDENT_INFO

	err := json.NewDecoder(r.Body).Decode(&temp1student)

	db.Create(&temp1student)

	if err != nil {
		panic(err)
	}

	db.Find(&s_student)

	json.NewEncoder(w).Encode(s_student)
}

func Get_All_Students(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var g_student []STUDENT_INFO

	db.Find(&g_student)

	json.NewEncoder(w).Encode(g_student)
}

func Update_Student_With_ID(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-Type", "application/json")

	var g_student []STUDENT_INFO
	var p STUDENT_INFO

	ID := chi.URLParam(r, "id")

	var temp STUDENT_INFO
	err := json.NewDecoder(r.Body).Decode(&temp)
	if err != nil {
		fmt.Println(err)
	}

	db.Model(&p).Select("*").Where("S_ID=?", ID).Updates(STUDENT_INFO{S_ID: temp.S_ID, Name: temp.Name, Village: temp.Village, Thana: temp.Thana, District: temp.District})

	db.Find(&g_student)

	json.NewEncoder(w).Encode(g_student)
}

func Get_Student_With_ID(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-Type", "application/json")

	var g_student []STUDENT_INFO

	err := db.Find(&g_student)
	if err != nil {
		panic(err)
	}

	ID := chi.URLParam(r, "id")

	for _, item := range g_student {
		if item.S_ID == ID {

			var temp STUDENT_INFO
			temp.S_ID = item.S_ID
			temp.Name = item.Name
			temp.Village = item.Village
			temp.Thana = item.Thana
			temp.District = item.District

			json.NewEncoder(w).Encode(temp)
			return
		}
	}
}

func main() {
	connect()

	r := chi.NewRouter()

	r.Post("/student", Add_New_Student)

	r.Get("/student", Get_All_Students)

	r.Get("/student/{id}", Get_Student_With_ID)

	r.Put("/student/{id}", Update_Student_With_ID)

	r.Delete("/student/{id}", Delete_Student_with_ID)

	err := http.ListenAndServe(":8200", r)
	if err != nil {
		panic(err)
	}
}
