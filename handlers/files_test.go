package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUnitHandleCreateFile(t *testing.T) {
	Convey("Successful photo file creation", t, func() {
		b := []byte(`{"id":"2","type":"test","amount":150}`)
		req, _ := http.NewRequest("POST", "/create", bytes.NewReader(b))
		w := httptest.NewRecorder()
		HandleCreateFile(w, req)
		So(w.Code, ShouldEqual, 200)
	})

	Convey("No Type Provided", t, func() {
		b := []byte(`{"id":"2","amount":150}`)
		req, _ := http.NewRequest("POST", "/test", bytes.NewReader(b))
		w := httptest.NewRecorder()
		HandleCreateFile(w, req)
		So(w.Code, ShouldEqual, 400)
	})

	Convey("Request Body Invalid", t, func() {
		b := []byte(`{"redirect_uri":"invalid", "reference":"invalid", "resource": "invalid", "state": "invalid"}`)
		req := httptest.NewRequest("POST", "/test", bytes.NewReader(b))
		w := httptest.NewRecorder()
		HandleCreateFile(w, req)
		So(w.Code, ShouldEqual, 400)
	})
}
