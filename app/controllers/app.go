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
	a := 1
	b := 2
	return c.Render(greeting,a,b)
}

func (c App) Hello() revel.Result {

	n1, _ := strconv.ParseUint(c.Params.Form.Get("number1"),10,32)
	n2, _ := strconv.ParseUint(c.Params.Form.Get("number2"),10,32)
	fmt.Println(n1,n2)

	data := &models.Table{
		Col1 : n1,
		Col2 : n2,
	}

	var myText string
	if models.Db == nil {
		myText = "DB connection is not established."
	}else{
		models.Db.Create(data)
		myText = fmt.Sprintf("Inserted %d, %d.",n1,n2)
	}

	return c.Render(myText)
}

func (c App) StoreSession() revel.Result {

	n1, _ := strconv.ParseUint(c.Params.Form.Get("number1"),10,32)
	fmt.Println(n1)

	data := &models.Table{}

	if models.Db == nil {
		return c.RenderError(nil)
	}else{
		models.Db.Where(&models.Table{Col1: n1}).First(data)
	}

	c.Session.Set("user1",data)
	fmt.Println(data.Col1,data.Col2)

	return c.Redirect("/")
}

func (c App) ReferSession() revel.Result {

	// Session.Get()
	/*
	sess, err := c.Session.Get("user1")
	data := sess.(map[string]interface{})
	if err == nil {
		n1 := data["Col1"]
		n2 := data["Col2"]
		flag := true
		return c.Render(flag,n1,n2)
	}else{
		flag := false
		return c.Render(flag)
	}
	*/

	// Session.GetInto()
	data := &models.Table{}
	_,  err := c.Session.GetInto("user1", data, false)
	if err == nil {
		n1 := data.Col1
		n2 := data.Col2
		flag := true
		return c.Render(flag,n1,n2)
	}else{
		flag := false
		return c.Render(flag)
	}

}
