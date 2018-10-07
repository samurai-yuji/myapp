package controllers

import (
	"github.com/revel/revel"
	"myapp/app/models"
	"strconv"
	"fmt"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
        greeting := "Aloha world"
	return c.Render(greeting)
}

func (c App) Hello() revel.Result {

	n1, _ := strconv.ParseUint(c.Params.Form.Get("number1"),10,32)
	n2, _ := strconv.ParseUint(c.Params.Form.Get("number2"),10,32)
	fmt.Println(n1,n2)

	data := &models.Table{
		Col1 : n1,
		Col2 : n2,
	}
	fmt.Println("AAA")

	if models.Db == nil {
		fmt.Println("BBB")
	}else{
		models.Db.Create(data)
	}
	return c.Render("OK")
}
