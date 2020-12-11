package test

import (
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"

	_ "github.com/hoangmirs/go-scraper/routers"

	log "github.com/astaxie/beego/core/logs"
	beego "github.com/astaxie/beego/server/web"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

// TestBeego is a sample to run an endpoint test
func TestBeego(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	log.Trace("TestBeego: Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status code is 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The result is not empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}
