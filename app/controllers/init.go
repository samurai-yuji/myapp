package controllers

import (
    "github.com/revel/revel"
    "myapp/app/models"
)

func init() {
	revel.TemplateFuncs["my_eq"] = func(a, b interface{}) bool {
		return a == 0  || a == b
	}
	revel.TemplateFuncs["my_disp"] = func(v interface{}) string {
		if v == 1 {
			return "OK"
		}else{
			return "NG"
		}
	}
    revel.OnAppStart(models.InitDB)
}
