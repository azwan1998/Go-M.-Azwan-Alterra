package appController

import (
	"fmt"
	"net/http"

	"belajar-go-echo/appMiddleware"
	"belajar-go-echo/appModel"

	"github.com/labstack/echo/v4"
)

type LoginInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserController struct {
	model        appModel.UserModel
	profileModel appModel.ProfileModel
	jwtSecret    string
}

func NewUserController(jwtSecret string, m appModel.UserModel, profileModel appModel.ProfileModel) UserController {
	return UserController{
		jwtSecret:    jwtSecret,
		model:        m,
		profileModel: profileModel,
	}
}

func (pc UserController) Login(c echo.Context) error {
	loginInfo := LoginInfo{}
	c.Bind(&loginInfo)
	User, err := pc.model.GetByEmail(loginInfo.Email)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "cannot login")
	}

	if !appMiddleware.VerifyPassword(loginInfo.Password, User.Password) {
		return c.String(http.StatusUnauthorized, "invalid credentials")
	}

	token, err := appMiddleware.CreateToken(int(User.ID), User.Role, pc.jwtSecret)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "cannot create token")
	}
	User.Token = token

	User, err = pc.model.Edit(int(User.ID), User)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "cannot add token")
	}
	return c.JSON(http.StatusOK, User)
}

func (pc UserController) Add(c echo.Context) error {
	User := appModel.User{}
	c.Bind(&User)

	// Enkripsi kata sandi sebelum menyimpan ke database
	encryptedPassword, err := appMiddleware.EncryptPassword(User.Password)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "cannot add editor")
	}

	User.Password = encryptedPassword

	newUser, err := pc.model.Add(User)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "cannot add")
	}

	return c.JSON(http.StatusOK, newUser)
}

func (pc UserController) GetAll(c echo.Context) error {
	currentLoginUserId := appMiddleware.ExtractTokenUserId(c)
	fmt.Println("😸 Current user id: ", currentLoginUserId)
	allUsers, err := pc.model.GetAll()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "cannot get Users")
	}
	return c.JSON(http.StatusOK, allUsers)
}
