package models

import (
    "github.com/jinzhu/gorm"
)

type Table struct {
    gorm.Model
    Col1 uint64
    Col2 uint64
}
