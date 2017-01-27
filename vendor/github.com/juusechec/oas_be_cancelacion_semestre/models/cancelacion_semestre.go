package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type CancelacionSemestre struct {
	Id                   int                        `orm:"column(id);pk;auto"`
	IdTipo               *TipoCancelacionSemestre   `orm:"column(id_tipo);rel(fk)"`
	IdEstado             *EstadoCancelacionSemestre `orm:"column(id_estado);rel(fk);null"`
	IdEstudiante         int16                      `orm:"column(id_estudiante);null"`
	Uid                  int                        `orm:"column(uid);null"`
	Motivo               string                     `orm:"column(motivo);null"`
	Observaciones        string                     `orm:"column(observaciones);null"`
	NumFoliosAnexados    int16                      `orm:"column(num_folios_anexados);null"`
	MotivoEstadoAprobado string                     `orm:"column(motivo_estado_aprobado);null"`
	FechaSolicitud       time.Time                  `orm:"column(fecha_solicitud);type(timestamp without time zone);null"`
	FechaAprobacion      time.Time                  `orm:"column(fecha_aprobacion);type(timestamp without time zone);null"`
}

func (t *CancelacionSemestre) TableName() string {
	return "cancelacion_semestre"
}

func init() {
	orm.RegisterModel(new(CancelacionSemestre))
}

// AddCancelacionSemestre insert a new CancelacionSemestre into database and returns
// last inserted Id on success.
func AddCancelacionSemestre(m *CancelacionSemestre) (id int64, err error) {
	if m.IdEstado == nil {
			m.IdEstado, _ = GetEstadoCancelacionSemestreById(1) //Default 1 or m.IdEstado.Id
	}
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCancelacionSemestreById retrieves CancelacionSemestre by Id. Returns error if
// Id doesn't exist
func GetCancelacionSemestreById(id int) (v *CancelacionSemestre, err error) {
	o := orm.NewOrm()
	v = &CancelacionSemestre{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCancelacionSemestre retrieves all CancelacionSemestre matches certain condition. Returns empty list if
// no records exist
func GetAllCancelacionSemestre(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CancelacionSemestre))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []CancelacionSemestre
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateCancelacionSemestre updates CancelacionSemestre by Id and returns error if
// the record to be updated doesn't exist
func UpdateCancelacionSemestreById(m *CancelacionSemestre) (err error) {
	o := orm.NewOrm()
	v := CancelacionSemestre{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCancelacionSemestre deletes CancelacionSemestre by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCancelacionSemestre(id int) (err error) {
	o := orm.NewOrm()
	v := CancelacionSemestre{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&CancelacionSemestre{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
