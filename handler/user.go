package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/handarudwiki/go-crowfunding/auth"
	"github.com/handarudwiki/go-crowfunding/helper"
	"github.com/handarudwiki/go-crowfunding/user"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService: userService}
}

func (h *userHandler) Register(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errorMessages := gin.H{"error": helper.FormatError(err)}

		response := helper.ApiResponse("Register akun failed", 400, "error", errorMessages)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newUser, err := h.userService.Register(input)

	if err != nil {
		response := helper.ApiResponse("Register failed", 500, "error", "Interneal server error")
		c.JSONP(http.StatusInternalServerError, response)
		return
	}

	jwtService := auth.NewJWTService()

	token, err := jwtService.GenerateToken(int(newUser.ID))
	if err != nil {
		errMessage := gin.H{"error": "internal server error"}
		response := helper.ApiResponse("Register failed", 400, "error", errMessage)

		c.JSON(http.StatusInternalServerError, response)
		return
	}
	formatterUser := user.FormatUser(newUser, token)

	response := helper.ApiResponse("Register done", 200, "Success", formatterUser)
	c.JSON(http.StatusOK, response)
	return
}

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errorMessage := gin.H{"error": helper.FormatError(err)}

		response := helper.ApiResponse("Login failed", 400, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	loggedInUser, err := h.userService.Login(input)

	if err != nil {
		response := helper.ApiResponse("Login failed", 403, "error", "email or password wrong")
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	jwtService := auth.NewJWTService()

	token, err := jwtService.GenerateToken(int(loggedInUser.ID))
	if err != nil {
		errMessage := gin.H{"error": err.Error()}
		response := helper.ApiResponse("Bad request", 400, "error", errMessage)

		c.JSON(http.StatusInternalServerError, response)
		return
	}
	formatterUser := user.FormatUser(loggedInUser, token)
	response := helper.ApiResponse("Login successful", 200, "Success", formatterUser)
	c.JSON(http.StatusOK, response)
	return
}

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errMessage := gin.H{"error": helper.FormatError(err)}
		response := helper.ApiResponse("Bad request", 400, "error", errMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	isEmailAvailable, err := h.userService.CheckEmailAvailability(input)

	if err != nil {
		errMessage := gin.H{"error": "internal server error"}
		response := helper.ApiResponse("Bad request", 400, "error", errMessage)

		c.JSON(http.StatusInternalServerError, response)
		return
	}

	data := gin.H{"is_available": isEmailAvailable}

	metaMessage := "user has beeen registered"

	if isEmailAvailable {
		metaMessage = "email is available"
	}

	response := helper.ApiResponse(metaMessage, 200, "success", data)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")

	userId := int(c.MustGet("currentUser").(user.User).ID)

	if err != nil {
		data := gin.H{"is_uploaded": "false"}
		response := helper.ApiResponse("upload avatar failed", 400, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// path := fmt.Sprintf("/images/%d-%s", userId, file.Filename)

	path := "images/" + "-" + strconv.Itoa(userId) + file.Filename

	err = c.SaveUploadedFile(file, path)

	if err != nil {
		data := gin.H{"is_uploaded": "false"}
		response := helper.ApiResponse(err.Error(), 500, "error", data)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	_, err = h.userService.SaveAvatar(userId, path)

	if err != nil {
		data := gin.H{"is_uploaded": "false"}
		response := helper.ApiResponse(err.Error(), 500, "error", data)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	data := gin.H{"is_uploaded": "true"}
	response := helper.ApiResponse("upload avatar success", 200, "success", data)
	c.JSON(http.StatusOK, response)
	return
}
