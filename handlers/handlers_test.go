package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHome(t *testing.T) {
	routes := getRoutes()
	testServer := httptest.NewServer(routes)
	defer testServer.Close()

	response, err := testServer.Client().Get(testServer.URL + "/")
	if err != nil {
		t.Log(err)
		t.Fatal(err)
	}

	if response.StatusCode != 200 {
		t.Errorf("expected status code 200, but got %d", response.StatusCode)
	}
}

func TestHomeWithSession(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	responseRecorder := httptest.NewRecorder()
	reg.Session.Put(ctx, "test_key", "test_value")
	h := http.HandlerFunc(testHandlers.Home)
	h.ServeHTTP(responseRecorder, req)

	if responseRecorder.Code != 200 {
		t.Errorf("expected status code 200, but got %d", responseRecorder.Code)
	}

	if reg.Session.GetString(ctx, "test_key") != "test_value" {
		t.Errorf("session values do not match, want %v, got %v", "test_value", reg.Session.GetString(ctx, "test_key"))
	}
}
