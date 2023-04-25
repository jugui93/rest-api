package models
import "gorm.io/gorm"

type Fact struct {
	gorm.Model
	Question string `json:"question" gorm:"text;not null;default:null`
	Answer string `json:"answer" gorm:"text;not null;default:null`
	A string `json:"a" gorm:"text;not null;default:null`
	B string `json:"b" gorm:"text;not null;default:null`
	C string `json:"c" gorm:"text;not null;default:null`
	D string `json:"d" gorm:"text;not null;default:null`
	Level uint `json:"level" gorm:"integer;not null;default:null`
}