package api

import (
	"errors"
	"fmt"
	"net/http"

	"../api/dto"

	"../domain"
	"../domain/model"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

type GameHandler struct {
	GameService domain.Service
}

func NewGameHandler() domain.API {
	return &GameHandler{}
}

func (gh *GameHandler) Start(g *gin.Context) {
	players := []dto.Player{}
	errBind := g.BindJSON(&players)
	if errBind != nil {
		errs := fmt.Sprintf("ERRORS: %s", "Unprocessable Entity")
		responseError(g, http.StatusUnprocessableEntity, "Unprocessable Entity", errors.New(errs))
	}

	if len(players) == 0 {
		errs := fmt.Sprintf("ERRORS: %s", "No players sent")
		responseError(g, http.StatusBadRequest, "Bad Request", errors.New(errs))
		return
	}

	validate := validator.New()
	for _, v := range players {
		valerr := validate.Struct(v)
		if valerr != nil {
			errs := fmt.Sprintf("ERRORS: %v", "Bad Request")
			responseError(g, http.StatusBadRequest, "Bad Request", errors.New(errs))
			return
		}
	}

	aux := []model.Player{}
	for _, v := range players {
		aux = append(aux, &v.ToModel())
	}

	err := gh.GameService.Start(g, aux)
	if err != nil {
		errs := fmt.Sprintf("ERRORS: %s", err)
		responseError(g, http.StatusUnprocessableEntity, "Cannot start game", errors.New(errs))
	} else {
		g.JSON(http.StatusCreated, players)
	}
}

func responseError(g *gin.Context, code int, message string, err error) {
	fmt.Printf("%v\n", err)
	g.AbortWithStatusJSON(code, map[string]interface{}{"error": message})
}
