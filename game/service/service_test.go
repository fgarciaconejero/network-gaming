package service_test

import (
	"context"
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
		want    string
		wantErr bool
	}{
		{
			name: "on the spot test",
			fields: fields{
				gr: MockRepository{},
			},
			args: args{
				ctx: context.Background(),
				p: []model.Player{
					{
						ID:           "facu",
						FirstNumber:  5,
						SecondNumber: 9,
					},
				},
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "in between test",
			fields: fields{
				gr: MockRepository{},
			},
			args: args{
				ctx: context.Background(),
				p: []model.Player{
					{
						ID:           "facu",
						FirstNumber:  4,
						SecondNumber: 8,
					},
				},
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "out of bounds test",
			fields: fields{
				gr: MockRepository{},
			},
			args: args{
				ctx: context.Background(),
				p: []model.Player{
					{
						ID:           "facu",
						FirstNumber:  1,
						SecondNumber: 2,
					},
				},
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &service.GameService{
				GameRepository: tt.fields.gr,
			}
			got := k.Start(tt.args.ctx, tt.args.p)
			if got != tt.want {
				t.Errorf("GameService.Start() got = %v, want = %v", got, tt.want)
				return
			}
		})
	}
}

type MockRepository struct {
	wantErr bool
}

func (mr MockRepository) AddPoints(id string, points int) map[string]int {
	return map[string]int{}
}

func (mr MockRepository) GenerateRandomNumber() int {
	return 5
}
