package test

import (
	_ "github.com/udistrital/oas_be_cancelacion_api/routers"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"

	"bytes"
	"encoding/json"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

// TestGet is a sample to run an endpoint test
func TestGetAllObject(t *testing.T) {
	r, _ := http.NewRequest("GET", "/oas_cancelacion_semestre/v1/object", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestGetAllObject", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

type Object struct {
	Id          int    `json:"Id"`
	Tipo        string `json:"Tipo"`
	Descripcion string `json:"Descripcion"`
}

// TestGet is a sample to run an endpoint test
func TestPostObject(t *testing.T) {
	// Create an instance of the Object struct.
	object := Object{
		Id:          1,
		Tipo:        "Este es un tipo de dato",
		Descripcion: "Esta es una descripción",
	}

	// Create JSON from the instance data.
	// ... Ignore errors.
	b, _ := json.Marshal(object)
	// Convert bytes to string.
	s := string(b)
	// Convert string to io.Reader
	bu := bytes.NewBuffer([]byte(s))

	r, _ := http.NewRequest("POST", "/oas_cancelacion_semestre/v1/object", bu)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestPostObject", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

func TestGetOneObject(t *testing.T) {
	r, _ := http.NewRequest("GET", "/oas_cancelacion_semestre/v1/object/1", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestGetOneObject", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

func TestPutObject(t *testing.T) {
	object := Object{
		Id:          1,
		Tipo:        "Este es un tipo de dato actualizado",
		Descripcion: "Esta es una descripción actualizada",
	}
	b, _ := json.Marshal(object)
	s := string(b)
	bu := bytes.NewBuffer([]byte(s))

	r, _ := http.NewRequest("PUT", "/oas_cancelacion_semestre/v1/object/1", bu)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestPutObject", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

func TestDeleteObject(t *testing.T) {
	r, _ := http.NewRequest("DELETE", "/oas_cancelacion_semestre/v1/object/1", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestDeleteObject", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}
