package products

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	repo := NewRepository()
	service := NewService(repo)
	handler := NewHandler(service)
	router := gin.Default()

	pr := router.Group("/products")
	pr.GET("", handler.GetProducts)
	return router
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application-json")
	return req, httptest.NewRecorder()
}

func Test_GetProducts(t *testing.T) {
	router := createServer()

	req, rr := createRequestTest(http.MethodGet, "/products?seller_id=FEX112AC", "")

	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	//Fail code 400
	req, rr = createRequestTest(http.MethodGet, "/products", "")
	router.ServeHTTP(rr, req)
	assert.Equal(t, 400, rr.Code)

}
