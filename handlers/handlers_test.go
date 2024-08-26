package handlers

import (
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
