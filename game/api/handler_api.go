package api

import (
	"net/http"

	"github.com/fgarciaconejero/network-gaming/common"
	"github.com/fgarciaconejero/network-gaming/game/api/dto"
	"github.com/fgarciaconejero/network-gaming/game/service"
	"gopkg.in/go-playground/validator.v9"

	"github.com/fgarciaconejero/network-gaming/game/domain"
	"github.com/fgarciaconejero/network-gaming/game/domain/model"

	"github.com/gin-gonic/gin"
)

type GameHandler struct {
	GameService domain.Service
}

func NewGameHandler() domain.API {
	gs := service.NewGameService()
	return &GameHandler{GameService: gs}
}

func (gh *GameHandler) Start(g *gin.Context) {
	players := []dto.Player{}
	errBind := g.BindJSON(&players)
	if errBind != nil {
		g.AbortWithStatusJSON(http.StatusUnprocessableEntity, common.ErrResponse{
			Error:   "Unprocessable Entity",
			Message: "The body doesn't match up with expected format",
		})
		return
	}

	if len(players) == 0 {
		g.AbortWithStatusJSON(http.StatusBadRequest, common.ErrResponse{
			Error:   "Bad Request",
			Message: "No players were sent",
		})
		return
	}

	validate := validator.New()
	for _, v := range players {
		valerr := validate.Struct(v)
		if valerr != nil {
			g.AbortWithStatusJSON(http.StatusBadRequest, common.ErrResponse{
				Error:   "Bad Request",
				Message: valerr.Error(),
			})
			return
		}
	}

	aux := []model.Player{}
	for _, v := range players {
		v = sortNumbers(v)
		aux = append(aux, *v.ToModel())
	}

	err := gh.GameService.Start(g, aux)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusUnprocessableEntity, common.ErrResponse{
			Error:   "Unprocessable Entity",
			Message: err.Error(),
		})
		return
	} else {
		g.JSON(http.StatusOK, err)
	}
}

func sortNumbers(p dto.Player) dto.Player {
	if p.FirstNumber > p.SecondNumber {
		aux := p.SecondNumber
		p.SecondNumber = p.FirstNumber
		p.FirstNumber = aux
	}
	return p
}
