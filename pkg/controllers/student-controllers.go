/* NIM : 11422008
   Nama: Mutiara Enjelina
   Prodi: STr TRPL
*/

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Mutiaraenjelina/go_students_crud_mysql/pkg/models"
	"github.com/Mutiaraenjelina/go_students_crud_mysql/pkg/utils"
	"github.com/gorilla/mux"
)

var NewStudent models.Student

func GetStudent(w http.ResponseWriter, r *http.Request) {
	newStudents := models.GetAllStudents()
	res, _ := json.Marshal(newStudents)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetStudentById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentId := vars["studentId"]
	NIM, err := strconv.ParseInt(studentId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	studentDetails, _ := models.GetStudentbyId(NIM)
	res, _ := json.Marshal(studentDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreateStudent(w http.ResponseWriter, r *http.Request) {
	CreateStudent := &models.Student{}
	utils.ParseBody(r, CreateStudent)
	b := CreateStudent.CreateStudent()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentId := vars["studentId"]
	NIM, err := strconv.ParseInt(studentId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	student := models.DeleteStudent(NIM)
	res, _ := json.Marshal(student)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	var updateStudent = &models.Student{}
	utils.ParseBody(r, updateStudent)
	vars := mux.Vars(r)
	studentId := vars["studentId"]
	NIM, err := strconv.ParseInt(studentId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	studentDetails, db := models.GetStudentbyId(NIM)
	if updateStudent.Name != "" {
		studentDetails.Name = updateStudent.Name
	}
	if updateStudent.IPK != "" {
		studentDetails.IPK = updateStudent.IPK
	}
	if updateStudent.Jurusan != "" {
		studentDetails.Jurusan = updateStudent.Jurusan
	}
	if updateStudent.Angkatan != "" {
		studentDetails.Angkatan = updateStudent.Angkatan
	}
	if updateStudent.Status_Aktif != "" {
		studentDetails.Status_Aktif = updateStudent.Status_Aktif
	}
	if updateStudent.Email_Akademik != "" {
		studentDetails.Email_Akademik = updateStudent.Email_Akademik
	}
	if updateStudent.Wali_Mahasiswa != "" {
		studentDetails.Wali_Mahasiswa = updateStudent.Wali_Mahasiswa
	}
	if updateStudent.Jalur_USM != "" {
		studentDetails.Jalur_USM = updateStudent.Jalur_USM
	}
	db.Save(&studentDetails)
	res, _ := json.Marshal(studentDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}