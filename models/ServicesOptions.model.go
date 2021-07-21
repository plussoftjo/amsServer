// Package models ..
package models

import (
	"github.com/jinzhu/gorm"
)

// ServicesOptions ..
type ServicesOptions struct {
	Title         string      `json:"title"`
	ServicesID    uint        `json:"servicesID"`
	SubServicesID uint        `json:"subServicesID"`
	Services      Services    `json:"services"`
	SubServices   SubServices `json:"subServices"`
	gorm.Model
}
