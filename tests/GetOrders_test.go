package tests

import (
	"github.com/Baraulia/COURIER_SERVICE/Controllers"
	"github.com/Baraulia/COURIER_SERVICE/db"
	_ "github.com/lib/pq"
	"net/http"
	"net/http/httptest"
	"testing"
)

var Orders []db.Order

func TestGetOrdersHandler(t *testing.T) {
	db.ConnectDB()
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/orders", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Controllers.GetOrders)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}
