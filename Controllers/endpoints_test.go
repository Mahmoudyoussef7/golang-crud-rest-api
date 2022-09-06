package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetProducts(t *testing.T) {
	r, err := http.NewRequest("GET", "https://localhost:8084/api/products/", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(GetProducts)
	handler.ServeHTTP(w, r)
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[
		{
		  "id": 793288815576317953,
		  "name": "test-product",
		  "price": 100,
		  "description": "random-description"
		},
		{
		  "id": 793289485879705601,
		  "name": "test-product",
		  "price": 100,
		  "description": "rddson"
		},
		{
		  "id": 793289501790044161,
		  "name": "tedsaoduct",
		  "price": 100,
		  "description": "rddson"
		},
		{
		  "id": 793289514409558017,
		  "name": "tedsaodsdssdsadduct",
		  "price": 100,
		  "description": "rddds"
		},
		{
		  "id": 793292193912651777,
		  "name": "test-product",
		  "price": 100,
		  "description": "random-description"
		}
	  ]`
	if w.Body.String() != expected {
		t.Errorf("handler returned unexpected body : got %v want %v", w.Body.String(), expected)
	}
}

func TestGetProductById(t *testing.T) {
	r, err := http.NewRequest("GET", "/api/products/", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := r.URL.Query()
	q.Add("id", "793288815576317953")

	r.URL.RawQuery = q.Encode()
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(GetProductById)
	handler.ServeHTTP(w, r)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"id": 793288815576317953,"name": "test-product","price": 100,"description": "random-description"}`
	if w.Body.String() != expected {
		t.Errorf("handler returned wrong status code: got %v want %v", w.Body.String(), expected)
	}
}

func TestGetProductByIdNotFound(t *testing.T) {
	r, err := http.NewRequest("GET", "/api/products/", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := r.URL.Query()
	q.Add("id", "793288815576317953")

	r.URL.RawQuery = q.Encode()
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(GetProductById)
	handler.ServeHTTP(w, r)

	if status := w.Code; status == http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestCreateProduct(t *testing.T) {
	var jsonStr = []byte(`{"id": 793288815576318653,"name": "test-product","price": 100,"description": "random-description"}`)
	r, err := http.NewRequest("POST", "/api/products/", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateProduct)
	handler.ServeHTTP(w, r)
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"id": 793288815576318653,"name": "test-product","price": 100,"description": "random-description"}`
	if w.Body.String() != expected {
		t.Errorf("handler returned wrong status code: got %v want %v", w.Body.String(), expected)
	}
}

func TestUpdateProduct(t *testing.T) {
	var jsonStr = []byte(`{"id": 793288815576317953,"name": "asdafdsdfgdsfg","price": 100,"description": "rasgdjhgklhghdsafgsdahfdjfsghtvjthfdygeion"}`)

	r, err := http.NewRequest("PUT", "/api/products/", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	q := r.URL.Query()
	q.Add("id", "793288815576317953")
	r.URL.RawQuery = q.Encode()
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateProduct)
	handler.ServeHTTP(w, r)
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"id": 793288815576317953,"name": "asdafdsdfgdsfg","price": 100,"description": "rasgdjhgklhghdsafgsdahfdjfsghtvjthfdygeion"}`
	if w.Body.String() != expected {
		t.Errorf("handler returned wrong status code: got %v want %v", w.Body.String(), expected)
	}
}

func TestDeleteProduct(t *testing.T) {
	r, err := http.NewRequest("DELETE", "/api/products/", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := r.URL.Query()
	q.Add("id", "793288815576317953")
	r.URL.RawQuery = q.Encode()
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteProduct)
	handler.ServeHTTP(w, r)
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	expected := `{"id": 793288815576317953,"name": "asdafdsdfgdsfg","price": 100,"description": "rasgdjhgklhghdsafgsdahfdjfsghtvjthfdygeion"}`
	if w.Body.String() != expected {
		t.Errorf("handler returned wrong status code: got %v want %v", w.Body.String(), expected)
	}
}
