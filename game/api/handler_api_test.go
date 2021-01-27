package api_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/fgarciaconejero/network-gaming/common"
	"github.com/fgarciaconejero/network-gaming/game/api"
	"github.com/fgarciaconejero/network-gaming/game/domain"
	"github.com/fgarciaconejero/network-gaming/game/domain/model"
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
	gh := &api.GameHandler{GameService: &GSMock{}}
	s := internal.NewServer(g, "8080").AddHandlers(gh)

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	body := `[{"id":"facu","first_number":2,"second_number":3}]`
	unsortedBody := `[{"id":"facu","first_number":4,"second_number":3}]`
	unprocessableBody := `[{"first_number":"asd","second_number":"3"}]`
	noLengthBody := `[]`
	validatorFailBody := `[{"first_number":2}]`

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
			request: request{r: common.MakeRequest(http.MethodPost, `/game/start`, body, headers)},
			want:    200,
		},
		{
			name:    "unsorted test",
			server:  server{s: s},
			request: request{r: common.MakeRequest(http.MethodPost, `/game/start`, unsortedBody, headers)},
			want:    200,
		},
		{
			name:    "Unprocessable entity test",
			server:  server{s: s},
			request: request{r: common.MakeRequest(http.MethodPost, `/game/start`, unprocessableBody, headers)},
			want:    400,
		},
		{
			name:    "Length == 0 test",
			server:  server{s: s},
			request: request{r: common.MakeRequest(http.MethodPost, `/game/start`, noLengthBody, headers)},
			want:    400,
		},
		{
			name:    "Validator fail test",
			server:  server{s: s},
			request: request{r: common.MakeRequest(http.MethodPost, `/game/start`, validatorFailBody, headers)},
			want:    400,
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

type GSMock struct {
}

func (gs *GSMock) Start(g context.Context, players []model.Player) string {
	return ""
}
