package handler

import (
	"fmt"
	"my-driver/banana"
	"my-driver/log"
	"my-driver/model"
	"my-driver/model/req"
	"my-driver/repository"
	"my-driver/security"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	uuid "github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserRepo repository.UserRepo
}

func (u *UserHandler)HandleSignUp(c echo.Context) error{
	req := req.ReqSignUp{}

	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message: err.Error(),
			Data: nil,
		})
	}

	//check validate
	validate := validator.New()
	
	if err := req.Validate(validate); err != nil {
		
		result := make(map[string]string)

		// Phân tách chuỗi đầu vào thành các cặp key-value
		pairs := strings.Split(err.Error(), "\n")
		for _, pair := range pairs {
			// Tách cặp key-value
			parts := strings.SplitN(strings.TrimSpace(pair), ": ", 2)

			// Kiểm tra xem cặp có đúng dạng không
			if len(parts) == 2 {
				// Thêm vào map
				result[parts[0]] = parts[1]
			}
		}

		
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Err: "err:form_validation_failed",
			Data: result,
		})
	}

	hash := security.HashAndSalt([]byte(req.Password))
	role := model.MEMBER.String()
	userId, err := uuid.NewUUID()

	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusForbidden, model.Response{
			StatusCode: http.StatusForbidden,
			Message: err.Error(),
			Data: nil,
		})
	}

	user := model.User{
		UserId: 	userId.String(),
		FullName: 	req.FullName,
		Email: 		req.Email,
		Password: 	hash,
		Role: 		role,
		Token: 		"",
	}

	user, err = u.UserRepo.SaveUser(c.Request().Context(), user)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message: err.Error(),
			Err: "err:user_already_exists",
		})
	}

	//gen token
	token, err := security.GenToken(user)
	if err != nil {
		log.Error(err.Error());
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
			Data: nil,
		})
	}

	user.Token = token

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message: "Success",
		Data: user,
	})
}

func (u *UserHandler)HandleSignIn(c echo.Context) error{
	req := req.ReqSignIn{}
	//check param
	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	validate := validator.New()
	//check validate
	
	if err := req.Validate(validate); err != nil {
		
		result := make(map[string]string)

		// Phân tách chuỗi đầu vào thành các cặp key-value
		pairs := strings.Split(err.Error(), "\n")
		for _, pair := range pairs {
			// Tách cặp key-value
			parts := strings.SplitN(strings.TrimSpace(pair), ": ", 2)

			// Kiểm tra xem cặp có đúng dạng không
			if len(parts) == 2 {
				// Thêm vào map
				result[parts[0]] = parts[1]
			}
		}

		
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Err: "err:form_validation_failed",
			Data: result,
		})
	}

	//check login
	user, err := u.UserRepo.CheckLogin(c.Request().Context(), req)
	if err != nil {
		if err == banana.UserNotFoundText {
			return c.JSON(http.StatusNotFound, model.Response{
				StatusCode: http.StatusNotFound,
				Message: banana.UserNotFoundText.Error(),
				Err: banana.ErrUserNotFound.Error(),
			})
		}
		log.Error(err.Error())
		return c.JSON(http.StatusUnauthorized, model.Response{
			StatusCode: http.StatusUnauthorized,
			Message: err.Error(),
			Data: nil,
		})
	}

	//check pass
	isTheSame := security.ComparePasswords(user.Password, []byte(req.Password))
	if !isTheSame {
		return c.JSON(http.StatusUnauthorized, model.Response{
			StatusCode: http.StatusUnauthorized,
			Message: banana.WrongPasswordText.Error(),
			Err: banana.ErrWrongPassword.Error(),
		})
	}

	//gen token
	token, err := security.GenToken(user)
	if err != nil {
		log.Error(err.Error());
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
			Data: nil,
		})
	}

	user.Token = token

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message: "Success",
		Data: user,
	})
}

func (u *UserHandler)HandleProfile(c echo.Context) error{
	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)
	
	user, err := u.UserRepo.SelectUserById(c.Request().Context(), claims.UserId)

	fmt.Println("============", user)

	if err != nil {
		if err == banana.UserNotFoundText {
			return c.JSON(http.StatusNotFound, model.Response{
				StatusCode: http.StatusNotFound,
				Message: banana.UserNotFoundText.Error(),
				Err: banana.ErrUserNotFound.Error(),
			})
		}

		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
			Data: nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message: "Success",
		Data: user,
	}) 
}

func (u *UserHandler)HandleUpdateProfile(c echo.Context) error{
	req := req.ReqUpdateProfile{}

	//check param
	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	validate := validator.New()
	//check validate
	
	if err := req.Validate(validate); err != nil {
		
		result := make(map[string]string)

		// Phân tách chuỗi đầu vào thành các cặp key-value
		pairs := strings.Split(err.Error(), "\n")
		for _, pair := range pairs {
			// Tách cặp key-value
			parts := strings.SplitN(strings.TrimSpace(pair), ": ", 2)

			// Kiểm tra xem cặp có đúng dạng không
			if len(parts) == 2 {
				// Thêm vào map
				result[parts[0]] = parts[1]
			}
		}

		
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Err: "err:form_validation_failed",
			Data: result,
		})
	}

	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)

	user := model.UserResponse{
		UserId: claims.UserId,
		FullName: req.FullName,
		Email: req.Email,
	}
	
	user, err := u.UserRepo.UpdateUser(c.Request().Context(), user)

	if err != nil {
		if err == banana.UserNotFoundText {
			return c.JSON(http.StatusNotFound, model.Response{
				StatusCode: http.StatusNotFound,
				Message: banana.UserNotFoundText.Error(),
				Err: banana.ErrUserNotFound.Error(),
			})
		}

		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
			Data: nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message: "Success",
		Data: user,
	}) 
}