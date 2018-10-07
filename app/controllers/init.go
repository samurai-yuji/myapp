package controllers

import (
    "github.com/revel/revel"
    "myapp/app/models"
)

func init() {
    revel.OnAppStart(models.InitDB)
}
