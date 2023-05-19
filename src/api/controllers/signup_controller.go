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
	Signup() error
}

type signupController struct {
	signupUsecase usecase.SignupUsecase
	context       *gin.Context
}

func NewSignupController(usecase usecase.SignupUsecase, context *gin.Context) SignupController {
	return &signupController{
		signupUsecase: usecase,
		context:       context,
	}
}

func (sc *signupController) Signup() error {
	var request entities.SignupRequest

	err := sc.context.ShouldBind(&request)
	if err != nil {
		sc.context.JSON(
			http.StatusBadRequest,
			entities.ErrorResponse{
				Message: err.Error(),
			},
		)

		return err
	}

	_, err = sc.signupUsecase.GetUserByEmail(sc.context, request.Email)
	if err == nil {
		sc.context.JSON(
			http.StatusConflict,
			entities.ErrorResponse{
				Message: "User already exists",
			},
		)

		return err
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		sc.context.JSON(
			http.StatusInternalServerError,
			entities.ErrorResponse{
				Message: err.Error(),
			},
		)

		return err
	}

	request.Password = string(encryptedPassword)

	user := entities.User{
		ID:       uuid.NewString(),
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	err = sc.signupUsecase.Signup(
		sc.context,
		&user,
	)
	if err != nil {
		sc.context.JSON(
			http.StatusInternalServerError,
			entities.ErrorResponse{
				Message: err.Error(),
			},
		)

		return err
	}

	accessTokenSecret := os.Getenv("SERVER_SECRET")
	accessTokenExpiryHoursStr := os.Getenv("SERVER_SECRET_EXP")
	accessTokenExpiryHours, err := strconv.Atoi(accessTokenExpiryHoursStr)
	if err != nil {
		sc.context.JSON(
			http.StatusInternalServerError,
			entities.ErrorResponse{
				Message: "Internal Error",
			},
		)

		return err
	}

	accessToken, err := sc.signupUsecase.CreateAccessToken(
		&user,
		accessTokenSecret,
		accessTokenExpiryHours,
	)
	if err != nil {
		sc.context.JSON(
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
		sc.context.JSON(
			http.StatusInternalServerError,
			entities.ErrorResponse{
				Message: "Internal Error",
			},
		)

		return err
	}

	refreshToken, err := sc.signupUsecase.CreateRefreshToken(
		&user,
		refreshTokenSecret,
		refreshTokenExpiryHours,
	)
	if err != nil {
		sc.context.JSON(
			http.StatusInternalServerError,
			entities.ErrorResponse{
				Message: err.Error(),
			},
		)

		return err
	}

	signupResponse := entities.SignupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	sc.context.JSON(
		http.StatusOK,
		signupResponse,
	)

	return nil
}
