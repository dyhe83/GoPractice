package gin_restful

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetPing(t *testing.T) {
	// Create a gin engine
	r := gin.Default()
	r.GET("/ping", getPing)

	// Create a request
	req, _ := http.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Check the response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message": "pong"}`, w.Body.String())
}

func TestGetAlbums(t *testing.T) {
	// Create a gin engine
	r := gin.Default()
	r.GET("/albums", getAlbums)

	// Create a request
	req, _ := http.NewRequest("GET", "/albums", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Check the response
	assert.Equal(t, http.StatusOK, w.Code)
	// Add your own assertions based on the albums data structure
	// assert.JSONEq(t, `expected JSON string`, w.Body.String())
}

func TestGetAlbumById(t *testing.T) {
	// Create a gin engine
	r := gin.Default()
	r.GET("/albums/:id", getAlbumById)

	// Create a request
	req, _ := http.NewRequest("GET", "/albums/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Check the response
	assert.Equal(t, http.StatusOK, w.Code)
	// Add your own assertions based on the albums data structure
	// assert.JSONEq(t, `expected JSON string`, w.Body.String())
}

func TestAddAlbum(t *testing.T) {
	// Create a gin engine
	r := gin.Default()
	r.POST("/albums", addAlbum)

	// Create a request
	reqBody := strings.NewReader(`{"id": "4", "title": "New Album", "artist": "New Artist", "price": 9.99}`)
	req, _ := http.NewRequest("POST", "/albums", reqBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Check the response
	assert.Equal(t, http.StatusCreated, w.Code)
	// Add your own assertions based on the expected response
	// assert.JSONEq(t, `expected JSON string`, w.Body.String())
}

func TestInvalidRoute(t *testing.T) {
	// Create a gin engine
	r := gin.Default()

	// Create a request for an invalid route
	req, _ := http.NewRequest("GET", "/invalid", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Check the response
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, "404 page not found", w.Body.String())
}

func TestGetAlbumByIdNotFound(t *testing.T) {
	// Create a gin engine
	r := gin.Default()
	r.GET("/albums/:id", getAlbumById)

	// Create a request for a non-existent album ID
	req, _ := http.NewRequest("GET", "/albums/999", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Check the response
	assert.Equal(t, http.StatusNotFound, w.Code)
	// assert.Equal(t, "404 page not found", w.Body.String())
}

func TestGetAlbumByIdSuccess(t *testing.T) {
	// Create a gin engine
	r := gin.Default()
	r.GET("/albums/:id", getAlbumById)

	// Create a request for an existing album ID
	req, _ := http.NewRequest("GET", "/albums/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Check the response
	assert.Equal(t, http.StatusOK, w.Code)
	// Add your own assertions based on the albums data structure
	// assert.JSONEq(t, `expected JSON string`, w.Body.String())
}

func TestAddAlbumInvalidInput(t *testing.T) {
	// Create a gin engine
	r := gin.Default()
	r.POST("/albums", addAlbum)

	// Create a request with invalid input
	reqBody := strings.NewReader(`{"id": "4", "title": "", "artist": "New Artist", "price": 9.99}`)
	req, _ := http.NewRequest("POST", "/albums", reqBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Check the response
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "Invalid input: title is required", w.Body.String())
}

func TestAddAlbumDuplicateID(t *testing.T) {
	// Create a gin engine
	r := gin.Default()
	r.POST("/albums", addAlbum)

	// Create a request with duplicate ID
	reqBody := strings.NewReader(`{"id": "1", "title": "Duplicate Album", "artist": "Duplicate Artist", "price": 9.99}`)
	req, _ := http.NewRequest("POST", "/albums", reqBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Check the response
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "Invalid input: album with ID 1 already exists", w.Body.String())
}

func TestAddAlbumInvalidJSONFormat(t *testing.T) {
	// Create a gin engine
	r := gin.Default()
	r.POST("/albums", addAlbum)

	// Create a request with invalid JSON format
	reqBody := strings.NewReader(`{"id": "4", "title": "New Album", "artist": "New Artist", "price": 9.99`)
	req, _ := http.NewRequest("POST", "/albums", reqBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Check the response
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
