package controller

import (
	"main/src/domain/entities"
	usecase "main/src/domain/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfileController interface {
	Fetch(context *gin.Context)
}

type profileController struct {
	profileUsecase usecase.ProfileUsecase
	context        *gin.Context
}

func NewProfileController(usecase usecase.ProfileUsecase) ProfileController {
	return &profileController{
		profileUsecase: usecase,
	}
}

func (pc *profileController) Fetch(context *gin.Context) {
	userID := pc.context.GetString("x-user-id")

	profile, err := pc.profileUsecase.GetProfileByID(pc.context, userID, http.DefaultClient.Timeout)
	if err != nil {
		pc.context.JSON(
			http.StatusInternalServerError,
			entities.ErrorResponse{
				Message: err.Error(),
			},
		)

		return
	}

	pc.context.JSON(
		http.StatusOK,
		profile,
	)
}
