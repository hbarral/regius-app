package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
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

	bodyText, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(string(bodyText), "Regius") {
		reg.TakeScreenShot(testServer.URL+"/", "TestHome", 1500, 1000)
		t.Errorf("expected body to contain 'Regius', but got %s", string(bodyText))
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

func TestClicker(t *testing.T) {
	routes := getRoutes()

	testServer := httptest.NewTLSServer(routes)
	defer testServer.Close()

	page := reg.FetchPage(testServer.URL + "/tester")
	outputElement := reg.SelectElementByID(page, "output")
	button := reg.SelectElementByID(page, "clicker")

	testHTML, _ := outputElement.HTML()
	if strings.Contains(testHTML, "Clicked the button") {
		t.Errorf("expected output to be 0, but got %s", testHTML)
	}

	button.MustClick()
	testHTML, _ = outputElement.HTML()
	if !strings.Contains(testHTML, "Clicked the button") {
		t.Errorf("expected output to be 1, but got %s", testHTML)
	}
}
