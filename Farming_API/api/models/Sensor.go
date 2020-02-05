package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type Sensor struct {
	Id          uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Humidity    float32   `gorm:"type:float" json:"humidity"`
	Temperature float32   `gorm:"type:float" json:"temperature"`
	Light       float32   `gorm:"type:float" json:"light"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (s *Sensor) Prepare() {
	s.Id = 0
	s.CreatedAt = time.Now()
	s.UpdatedAt = time.Now()
}

func (s *Sensor) Validate() error {
	if s.Humidity == 0 {
		return errors.New("try again")
	}
	if s.Temperature == 0 {
		return errors.New("try again")
	}
	if s.Light == 0 {
		return errors.New("try again")
	}
	return nil
}

func (s *Sensor) SaveSensor(db *gorm.DB) (*Sensor, error) {
	var err error
	//create & handler error

	err = db.Create(&s).Error
	if err != nil {
		return &Sensor{}, err
	}
	return s, nil
}

func (s *Sensor) FindAllSensors(db *gorm.DB) (*[]Sensor, error) {
	var err error
	sensor := []Sensor{}
	err = db.Debug().Model(&Sensor{}).Limit(100).Find(&sensor).Error
	if err != nil {
		return &[]Sensor{}, err
	}
	return &sensor, err
}

func (s *Sensor) FindSensorByDate(db *gorm.DB, CreatedAt time.Time) (*[]Sensor, error) {
	var err error
	sensors := []Sensor{}
	err = db.Debug().Model(&Sensor{}).Where("created_at > ?", CreatedAt).Find(&sensors).Error
	if err != nil {
		return &[]Sensor{}, err
	}
	return &sensors, err
}
