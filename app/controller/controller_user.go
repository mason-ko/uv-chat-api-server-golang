package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"uv-chat-api-server-golang/domain"
	"uv-chat-api-server-golang/domain/user"
	"uv-chat-api-server-golang/internal/appctx"
)

type userController struct {
	service domain.Service
}

// 메시지 생성
func (m *userController) Create(ctx *gin.Context) {
	var messageDTO user.User
	if err := ctx.ShouldBindJSON(&messageDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := m.service.UserService().Create(appctx.NewContext(ctx), messageDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Message created successfully"})
}

// 메시지 삭제
func (m *userController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = m.service.UserService().Delete(appctx.NewContext(ctx), uint(idInt))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Message deleted successfully"})
}

// 단일 메시지 조회
func (m *userController) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	message, err := m.service.UserService().Get(appctx.NewContext(ctx), uint(idInt))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, message)
}

// 메시지 목록 조회
func (m *userController) GetList(ctx *gin.Context) {
	messages, err := m.service.UserService().GetList(appctx.NewContext(ctx), user.GetListParam{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, messages)
}

func newUserController(service domain.Service) user.Controller {
	return &userController{
		service: service,
	}
}
