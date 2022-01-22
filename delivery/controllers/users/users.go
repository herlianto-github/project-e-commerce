package users

import (
	"net/http"
	"project-e-commerces/constants"
	"project-e-commerces/delivery/common"
	"project-e-commerces/entities"
	"project-e-commerces/repository/users"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type UsersController struct {
	Repo users.UserInterface
}

func NewUsersControllers(usrep users.UserInterface) *UsersController {
	return &UsersController{Repo: usrep}
}

func CreateTokenAuth(id uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userid"] = id
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.JWT_SECRET_KEY))
}

// POST /user/login
func (uscon UsersController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		loginFormat := LoginRequestFormat{}
		if err := c.Bind(&loginFormat); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}
		checkedUser, err := uscon.Repo.Login(loginFormat.Email, loginFormat.Password)

		if err != nil || checkedUser.ID != 0 {
			if loginFormat.Email != "" && loginFormat.Password != "" {
				token, _ := CreateTokenAuth(checkedUser.ID)

				return c.JSON(
					http.StatusOK, map[string]interface{}{
						"message": "Successful Operation",
						"token":   token,
					},
				)
			}
		}
		return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())

	}
}

// POST /user/register
func (uscon UsersController) PostUserCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		newUserReq := RegisterUserRequestFormat{}

		if err := c.Bind(&newUserReq); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		hash, _ := bcrypt.GenerateFromPassword([]byte(newUserReq.Password), 14)
		newUser := entities.User{
			Name:     newUserReq.Password,
			Email:    newUserReq.Email,
			Password: string(hash),
		}

		if _, err := uscon.Repo.Create(newUser); err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		} else {

			return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
		}

	}

}

// GET /users
func (uscon UsersController) GetUsersCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		user, err := uscon.Repo.GetAll()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}

		response := GetUsersResponseFormat{
			Message: "Successful Opration",
			Data:    user,
		}

		return c.JSON(http.StatusOK, response)
	}
}

// GET /users/:id
func (uscon UsersController) GetUserCtrl() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		user, err := uscon.Repo.Get(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success",
			"data":    user,
		})
	}

}

// PUT /users/:id
func (uscon UsersController) EditUserCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		updateUserReq := PutUserRequestFormat{}
		if err := c.Bind(&updateUserReq); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		updateUser := entities.User{
			Name:     updateUserReq.Name,
			Password: updateUserReq.Password,
		}

		if _, err := uscon.Repo.Update(updateUser, id); err != nil {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		}
		return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
	}

}

// DELETE /users/:id
func (uscon UsersController) DeleteUserCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		deletedUser, _ := uscon.Repo.Delete(id)

		if deletedUser.ID != 0 {
			return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
		} else {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		}

	}

}
