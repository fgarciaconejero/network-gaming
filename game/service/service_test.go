package service_test

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/fgarciaconejero/network-gaming/game/domain"
	"github.com/fgarciaconejero/network-gaming/game/domain/model"
	service "github.com/fgarciaconejero/network-gaming/game/service"
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

func TestGameService_Start(t *testing.T) {
	type fields struct {
		gr domain.Repository
	}
	type args struct {
		ctx context.Context
		p   []model.Player
	}
	tests := []struct {
		name    string
		args    args
		fields  fields
		want    *service.GameService
		wantErr bool
	}{
		{
			name: "first test",
			fields: fields{
				gr: MockRepository{},
			},
			args: args{
				ctx: context.Background(),
				p:   []model.Player{},
			},
			want:    &service.GameService{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &service.GameService{
				GameRepository: tt.fields.gr,
			}
			err := k.Start(tt.args.ctx, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("GameService.Start() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
		})
	}
}

type MockRepository struct {
	wantErr bool
}

func (mr MockRepository) AddPoints(id string, points int) error {
	if mr.wantErr {
		return errors.New("Error while adding points")
	}
	return nil
}
