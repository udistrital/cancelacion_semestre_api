package endpointTests

//https://github.com/goinggo/beego-mgo/blob/master/test/endpointTests/buoyEndpoints_test.go
import (
	"net/http"
	"net/http/httptest"
	"testing"

	"encoding/json"

	"github.com/astaxie/beego"
	log "github.com/goinggo/tracelog"
	. "github.com/smartystreets/goconvey/convey"
	//"github.com/juusechec/oas_be_cancelacion_semestre"
)

func TestTipoCancelacionSemestre(t *testing.T) {
	r, _ := http.NewRequest("GET", "oas_cancelacion_semestre//v1/tipo_cancelacion_semestre", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	log.Trace("testing", "TestTipoCancelacionSemestre", "Code[%d]\n%s", w.Code, w.Body.String())

	var response struct {
		Id          int    `json:"id"`
		Tipo        string `json:"tipo"`
		Descripcion string `json:"descripcion"`
	}

	json.Unmarshal(w.Body.Bytes(), &response)

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
		Convey("There Should Be A Result For Station 42002", func() {
			So(response.Id, ShouldEqual, "42002")
		})
	})
}
