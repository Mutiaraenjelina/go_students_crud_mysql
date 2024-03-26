/* NIM : 11422008
   Nama: Mutiara Enjelina
   Prodi: STr TRPL
*/

package models

import (
	"github.com/Mutiaraenjelina/go_students_crud_mysql/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Student struct {
	gorm.Model
	NIM      string `gorm:""json:"nim"`
	Name     string `json:"name"`
	IPK      string `json:"ipk"`
	Jurusan  string `json:"jurusan"`
	Angkatan string `json:"angkatan"`
	Status_Aktif string `json:"status_aktif"`
	Username string `json:"username"`
	Email_Akademik string `json:"email_akademik`
	Wali_Mahasiswa string `json:"wali_mahasiswa`
	Jalur_USM string `json:"jalur_usm"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Student{})
}
func (b *Student) CreateStudent() *Student {
	db.NewRecord(b)
	db.Create(&b)
	return b
}
func GetAllStudents() []Student {
	var Students []Student
	db.Find(&Students)
	return Students
}
func GetStudentbyId(nim int64) (*Student, *gorm.DB) {
	var getStudent Student
	db := db.Where("NIM=?", nim).Find(&getStudent)
	return &getStudent, db
}
func DeleteStudent(nim int64) Student {
	var student Student
	db.Where("NIM=?", nim).Delete(student)
	return student
}
