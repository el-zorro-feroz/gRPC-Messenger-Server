package controller

import (
	"main/src/domain/entities"
	usecase "main/src/domain/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfileController interface {
	Fetch() error
}

type profileController struct {
	profileUsecase usecase.ProfileUsecase
	context        *gin.Context
}

func NewProfileController(usecase usecase.ProfileUsecase, context *gin.Context) ProfileController {
	return &profileController{
		profileUsecase: usecase,
		context:        context,
	}
}

func (pc *profileController) Fetch() error {
	userID := pc.context.GetString("x-user-id")

	profile, err := pc.profileUsecase.GetProfileByID(pc.context, userID)
	if err != nil {
		pc.context.JSON(
			http.StatusInternalServerError,
			entities.ErrorResponse{
				Message: err.Error(),
			},
		)

		return err
	}

	pc.context.JSON(
		http.StatusOK,
		profile,
	)

	return nil
}
