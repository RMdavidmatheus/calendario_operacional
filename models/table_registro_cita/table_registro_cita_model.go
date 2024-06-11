package table_registro_cita

import (
	"calendarioOperacional/models/table_cita"
	"time"
)

type TableRegistroCita struct {
	Id           string               `json:"id_registro_cita" gorm:"type:uuid;default:uuid_generate_v4();primarykey"`
	IdCita       string               `json:"id_cita" gorm:"type:text;not null; unique"`
	CitaRelation table_cita.TableCita `gorm:"foreignKey:IdCita;references:Id"`
	StateCita    string               `json:"state_cita" gorm:"type:text;not null"`
	IsActive     bool                 `json:"is_active" gorm:"type:boolean;default:true; not null"`
	CreatedAt    time.Time            `json:"created_at"`
	UpdatedAt    *time.Time           `json:"updated_at" gorm:"type:timestamp;default:null"`
	DeletedAt    *time.Time           `json:"deleted_at" gorm:"type:timestamp;default:null"`
}
