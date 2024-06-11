package table_cita

import "time"

type TableCita struct {
	Id                string     `json:"id_cita" gorm:"type:uuid;default:uuid_generate_v4();primarykey"`
	DateMeeting       time.Time  `json:"date_meeting" gorm:"type:timestamp;default:null"`
	AddressMeeting    string     `json:"address_meeting" gorm:"type:text; not null"`
	PersonNameMeeting string     `json:"person_name_meeting" gorm:"type:text; not null"`
	IsActive          bool       `json:"is_active" gorm:"type:boolean;default:true; not null"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         *time.Time `json:"updated_at" gorm:"type:timestamp;default:null"`
	DeletedAt         *time.Time `json:"deleted_at" gorm:"type:timestamp;default:null"`
}
