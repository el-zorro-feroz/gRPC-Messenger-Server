package controller

import (
	"main/src/domain/entities"
	usecase "main/src/domain/usecases"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginController interface {
	Login(context *gin.Context)
}

type loginController struct {
	loginUsecase usecase.LoginUsecase
}

func NewLoginController(usecase usecase.LoginUsecase) LoginController {
	return &loginController{
		loginUsecase: usecase,
	}
}

func (lc *loginController) Login(context *gin.Context) {
	var request entities.LoginRequest

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

	user, err := lc.loginUsecase.GetUserByEmail(context, request.Email, http.DefaultClient.Timeout)
	if err != nil {
		context.JSON(
			http.StatusNotFound,
			entities.ErrorResponse{
				Message: "User Not Found",
			},
		)

		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		context.JSON(
			http.StatusUnauthorized,
			entities.ErrorResponse{
				Message: "Invalid Credentials",
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

	accessToken, err := lc.loginUsecase.CreateAccessToken(
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

	refreshToken, err := lc.loginUsecase.CreateRefreshToken(
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

	loginResponse := entities.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	context.JSON(
		http.StatusOK,
		loginResponse,
	)
}
