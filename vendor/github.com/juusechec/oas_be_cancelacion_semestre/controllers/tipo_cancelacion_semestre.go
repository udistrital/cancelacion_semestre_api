package controllers

import (
	"encoding/json"
	"errors"
	"github.com/juusechec/oas_be_cancelacion_semestre/components"
	"github.com/juusechec/oas_be_cancelacion_semestre/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/juusechec/jwt-beego"
)

// oprations for TipoCancelacionSemestre
type TipoCancelacionSemestreController struct {
	beego.Controller
}

func (c *TipoCancelacionSemestreController) Prepare() {
	//tokenString := c.Ctx.Input.Query("tokenString")
	tokenString := c.Ctx.GetCookie("tokenString")

	et := jwtbeego.EasyToken{}
	valido, _ := et.ValidateToken(tokenString)

	if !valido {
		c.Ctx.Output.SetStatus(401)
		c.Data["json"] = "Permission Deny"
		c.ServeJSON()
	}
	return
}

func (c *TipoCancelacionSemestreController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description create TipoCancelacionSemestre
// @Param	body		body 	models.TipoCancelacionSemestre	true		"body for TipoCancelacionSemestre content"
// @Success 201 {int} models.TipoCancelacionSemestre
// @Failure 403 body is empty
// @router / [post]
func (c *TipoCancelacionSemestreController) Post() {
	_, header, err := c.GetFile("archivo")
	if err == nil {
		c.Ctx.Output.SetStatus(201)
		var nombre = "/tmp/" + header.Filename
		c.SaveToFile("archivo", nombre)

		notificador := components.Notificador{
			From: "juusechec@correo.udistrital.edu.co",
			To:   "naturalmentejorge@gmail.com",
		}
		notificador.Enviar()

		c.Data["json"] = "{'filename': '" + nombre + "'}"
		c.ServeJSON()
		return
	}
	var v models.TipoCancelacionSemestre
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddTipoCancelacionSemestre(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// @Title Get
// @Description get TipoCancelacionSemestre by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.TipoCancelacionSemestre
// @Failure 403 :id is empty
// @router /:id [get]
func (c *TipoCancelacionSemestreController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetTipoCancelacionSemestreById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// @Title Get All
// @Description get TipoCancelacionSemestre
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.TipoCancelacionSemestre
// @Failure 403
// @router / [get]
func (c *TipoCancelacionSemestreController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query map[string]string = make(map[string]string)
	var limit int64 = 10
	var offset int64 = 0

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllTipoCancelacionSemestre(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// @Title Update
// @Description update the TipoCancelacionSemestre
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.TipoCancelacionSemestre	true		"body for TipoCancelacionSemestre content"
// @Success 200 {object} models.TipoCancelacionSemestre
// @Failure 403 :id is not int
// @router /:id [put]
func (c *TipoCancelacionSemestreController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.TipoCancelacionSemestre{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateTipoCancelacionSemestreById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// @Title Delete
// @Description delete the TipoCancelacionSemestre
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *TipoCancelacionSemestreController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteTipoCancelacionSemestre(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
