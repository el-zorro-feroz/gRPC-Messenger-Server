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
	Login() error
}

type loginController struct {
	loginUsecase usecase.LoginUsecase
	context      *gin.Context
}

func NewLoginController(usecase usecase.LoginUsecase, context *gin.Context) LoginController {
	return &loginController{
		loginUsecase: usecase,
		context:      context,
	}
}

func (lc *loginController) Login() error {
	var request entities.LoginRequest

	err := lc.context.ShouldBind(&request)
	if err != nil {
		lc.context.JSON(
			http.StatusBadRequest,
			entities.ErrorResponse{
				Message: err.Error(),
			},
		)

		return err
	}

	user, err := lc.loginUsecase.GetUserByEmail(lc.context, request.Email)
	if err != nil {
		lc.context.JSON(
			http.StatusNotFound,
			entities.ErrorResponse{
				Message: "User Not Found",
			},
		)

		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		lc.context.JSON(
			http.StatusUnauthorized,
			entities.ErrorResponse{
				Message: "Invalid Credentials",
			},
		)

		return err
	}

	accessTokenSecret := os.Getenv("SERVER_SECRET")
	accessTokenExpiryHoursStr := os.Getenv("SERVER_SECRET_EXP")
	accessTokenExpiryHours, err := strconv.Atoi(accessTokenExpiryHoursStr)
	if err != nil {
		lc.context.JSON(
			http.StatusInternalServerError,
			entities.ErrorResponse{
				Message: "Internal Error",
			},
		)

		return err
	}

	accessToken, err := lc.loginUsecase.CreateAccessToken(
		&user,
		accessTokenSecret,
		accessTokenExpiryHours,
	)
	if err != nil {
		lc.context.JSON(
			http.StatusInternalServerError,
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
		lc.context.JSON(
			http.StatusInternalServerError,
			entities.ErrorResponse{
				Message: "Internal Error",
			},
		)

		return err
	}

	refreshToken, err := lc.loginUsecase.CreateRefreshToken(
		&user,
		refreshTokenSecret,
		refreshTokenExpiryHours,
	)
	if err != nil {
		lc.context.JSON(
			http.StatusInternalServerError,
			entities.ErrorResponse{
				Message: err.Error(),
			},
		)

		return err
	}

	loginResponse := entities.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	lc.context.JSON(
		http.StatusOK,
		loginResponse,
	)

	return nil
}
