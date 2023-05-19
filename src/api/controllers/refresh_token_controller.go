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
	RefreshToken() error
}

type refreshTokenController struct {
	refreshTokenUsecase usecase.RefreshTokenUsecase
	context             *gin.Context
}

func NewRefreshTokenUsecase(usecase usecase.RefreshTokenUsecase, context *gin.Context) RefreshTokenController {
	return &refreshTokenController{
		refreshTokenUsecase: usecase,
		context:             context,
	}
}

func (rtc *refreshTokenController) RefreshToken() error {
	var request entities.RefreshTokenRequest

	err := rtc.context.ShouldBind(&request)
	if err != nil {
		rtc.context.JSON(
			http.StatusBadRequest,
			entities.ErrorResponse{
				Message: err.Error(),
			},
		)

		return err
	}

	refreshTokenSecret := os.Getenv("SERVER_REFRESH_SECRET")
	refreshTokenExpiryHoursStr := os.Getenv("SERVER_REFRESH_SECRET_EXP")
	refreshTokenExpiryHours, err := strconv.Atoi(refreshTokenExpiryHoursStr)
	if err != nil {
		rtc.context.JSON(
			http.StatusInternalServerError,
			entities.ErrorResponse{
				Message: "Internal Error",
			},
		)

		return err
	}

	id, err := rtc.refreshTokenUsecase.ExtractIDFromToken(
		request.RefreshToken,
		refreshTokenSecret,
	)
	if err != nil {
		rtc.context.JSON(
			http.StatusUnauthorized,
			entities.ErrorResponse{
				Message: "User not found",
			},
		)

		return err
	}

	user, err := rtc.refreshTokenUsecase.GetUserByID(rtc.context, id)
	if err != nil {
		rtc.context.JSON(
			http.StatusUnauthorized,
			entities.ErrorResponse{
				Message: "User not found",
			},
		)

		return err
	}

	accessTokenSecret := os.Getenv("SERVER_SECRET")
	accessTokenExpiryHoursStr := os.Getenv("SERVER_SECRET_EXP")
	accessTokenExpiryHours, err := strconv.Atoi(accessTokenExpiryHoursStr)
	if err != nil {
		rtc.context.JSON(
			http.StatusInternalServerError,
			entities.ErrorResponse{
				Message: "Internal Error",
			},
		)

		return err
	}

	accessToken, err := rtc.refreshTokenUsecase.CreateAccessToken(
		&user,
		accessTokenSecret,
		accessTokenExpiryHours,
	)
	if err != nil {
		rtc.context.JSON(
			http.StatusInternalServerError,
			entities.ErrorResponse{
				Message: err.Error(),
			},
		)

		return err
	}

	refreshToken, err := rtc.refreshTokenUsecase.CreateRefreshToken(
		&user,
		refreshTokenSecret,
		refreshTokenExpiryHours,
	)
	if err != nil {
		rtc.context.JSON(
			http.StatusInternalServerError,
			entities.ErrorResponse{
				Message: err.Error(),
			},
		)

		return err
	}

	refreshTokenResponse := entities.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	rtc.context.JSON(
		http.StatusOK,
		refreshTokenResponse,
	)

	return nil
}
