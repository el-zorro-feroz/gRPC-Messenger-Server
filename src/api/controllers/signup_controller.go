package controller

import (
	"main/src/domain/entities"
	usecase "main/src/domain/usecases"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type SignupController interface {
	Signup(context *gin.Context)
}

type signupController struct {
	signupUsecase usecase.SignupUsecase
}

func NewSignupController(usecase usecase.SignupUsecase) SignupController {
	return &signupController{
		signupUsecase: usecase,
	}
}

func (sc *signupController) Signup(context *gin.Context) {
	var request entities.SignupRequest

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

	_, err = sc.signupUsecase.GetUserByEmail(context, request.Email, http.DefaultClient.Timeout)
	if err == nil {
		context.JSON(
			http.StatusConflict,
			entities.ErrorResponse{
				Message: "User already exists",
			},
		)

		return
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
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

	request.Password = string(encryptedPassword)

	user := entities.User{
		ID:       uuid.NewString(),
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	err = sc.signupUsecase.Signup(context, &user, http.DefaultClient.Timeout)
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			entities.ErrorResponse{
				Message: err.Error(),
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

	accessToken, err := sc.signupUsecase.CreateAccessToken(
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

	refreshToken, err := sc.signupUsecase.CreateRefreshToken(
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

	signupResponse := entities.SignupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	context.JSON(
		http.StatusOK,
		signupResponse,
	)
}
