package controller

import (
	"main/src/domain/entities"
	usecase "main/src/domain/usecases"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RefreshTokenController interface {
	RefreshToken(context *gin.Context)
}

type refreshTokenController struct {
	refreshTokenUsecase usecase.RefreshTokenUsecase
}

func NewRefreshTokenUsecase(usecase usecase.RefreshTokenUsecase) RefreshTokenController {
	return &refreshTokenController{
		refreshTokenUsecase: usecase,
	}
}

func (rtc *refreshTokenController) RefreshToken(context *gin.Context) {
	var request entities.RefreshTokenRequest

	err := context.ShouldBind(&request)
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			entities.ErrorResponse{
				Message: err.Error(),
			},
		)

		return
	}

	refreshTokenSecret := os.Getenv("SERVER_REFRESH_SECRET")
	refreshTokenExpiryHoursStr := os.Getenv("SERVER_REFRESH_SECRET_EXP")
	refreshTokenExpiryHours, err := strconv.Atoi(refreshTokenExpiryHoursStr)
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			entities.ErrorResponse{
				Message: "Internal Error",
			},
		)

		return
	}

	id, err := rtc.refreshTokenUsecase.ExtractIDFromToken(
		request.RefreshToken,
		refreshTokenSecret,
	)
	if err != nil {
		context.JSON(
			http.StatusUnauthorized,
			entities.ErrorResponse{
				Message: "User not found",
			},
		)

		return
	}

	user, err := rtc.refreshTokenUsecase.GetUserByID(context, id, http.DefaultClient.Timeout)
	if err != nil {
		context.JSON(
			http.StatusUnauthorized,
			entities.ErrorResponse{
				Message: "User not found",
			},
		)

		return
	}

	accessTokenSecret := os.Getenv("SERVER_SECRET")
	accessTokenExpiryHoursStr := os.Getenv("SERVER_SECRET_EXP")
	accessTokenExpiryHours, err := strconv.Atoi(accessTokenExpiryHoursStr)
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			entities.ErrorResponse{
				Message: "Internal Error",
			},
		)

		return
	}

	accessToken, err := rtc.refreshTokenUsecase.CreateAccessToken(
		&user,
		accessTokenSecret,
		accessTokenExpiryHours,
	)
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			entities.ErrorResponse{
				Message: err.Error(),
			},
		)

		return
	}

	refreshToken, err := rtc.refreshTokenUsecase.CreateRefreshToken(
		&user,
		refreshTokenSecret,
		refreshTokenExpiryHours,
	)
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			entities.ErrorResponse{
				Message: err.Error(),
			},
		)

		return
	}

	refreshTokenResponse := entities.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	context.JSON(
		http.StatusOK,
		refreshTokenResponse,
	)
}
