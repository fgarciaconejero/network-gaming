package api_test

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/fgarciaconejero/network-gaming/common"
	"github.com/fgarciaconejero/network-gaming/game/api"
	"github.com/fgarciaconejero/network-gaming/game/domain"
	"github.com/fgarciaconejero/network-gaming/internal"
	"github.com/gin-gonic/gin"
)

func TestGameHandler_NewGameHandler(t *testing.T) {
	tests := []struct {
		name string
		want domain.API
	}{
		{
			name: "first test",
			want: &api.GameHandler{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := api.NewGameHandler(); !(reflect.TypeOf(got) == reflect.TypeOf(tt.want)) {
				t.Errorf("NewGameHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestGameHandler_Start(t *testing.T) {
	gin.SetMode(gin.TestMode)
	g := gin.Default()
	g.Use(gin.Logger())
	s := internal.NewServer(g, "8080").AddHandlers()

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	body := `{
		"first_number": "2",
		"second_number": "3"
	}`

	type server struct {
		s *internal.SRV
	}
	type request struct {
		r *http.Request
	}
	tests := []struct {
		name    string
		server  server
		request request
		want    int
	}{
		{
			name:    "first test",
			server:  server{s: s},
			request: request{r: common.MakeRequest(http.MethodPost, "game/start", body, headers)},
			want:    200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := httptest.NewRecorder()
			tt.server.s.ServeHTTP(resp, tt.request.r)
			if got := resp.Code; !(got == tt.want) {
				t.Errorf("GameHandler_Start() = %v, want %v", got, tt.want)
			}
		})
	}
}
