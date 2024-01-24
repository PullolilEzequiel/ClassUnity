package classroom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"gotest.tools/assert"
)

func startApp() *gin.Engine {
	r := gin.Default()
	ImportClassRoomRoute(r)
	return r
}
func TestFindAllClassrooms(t *testing.T) {
	router := startApp()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/classrooms", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	doc, err := goquery.NewDocumentFromReader(w.Body)
	assert.NilError(t, err)
	assert.Equal(t, 3, doc.Find(".classroom").Length())
}
