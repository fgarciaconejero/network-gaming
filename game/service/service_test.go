package service_test

import (
	"reflect"
	"testing"

	"github.com/fgarciaconejero/network-gaming/game/domain"
	"github.com/fgarciaconejero/network-gaming/game/service"
)

func TestGameService_NewService(t *testing.T) {
	tests := []struct {
		name string
		want domain.Service
	}{
		{
			name: "first test",
			want: &service.GameService{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := service.NewGameService(); !(reflect.TypeOf(got) == reflect.TypeOf(tt.want)) {
				t.Errorf("NewGameService() = %v, want %v", got, tt.want)
			}
		})
	}
}
