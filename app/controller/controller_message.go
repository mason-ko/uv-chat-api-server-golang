package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"uv-chat-api-server-golang/domain"
	"uv-chat-api-server-golang/domain/message"
	"uv-chat-api-server-golang/internal/appctx"
)

type messageController struct {
	service domain.Service
}

// 메시지 생성
func (m *messageController) Create(ctx *gin.Context) {
	var messageDTO message.Message
	if err := ctx.ShouldBindJSON(&messageDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := m.service.MessageService().Create(appctx.NewContext(ctx), messageDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Message created successfully"})
}

// 메시지 삭제
func (m *messageController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = m.service.MessageService().Delete(appctx.NewContext(ctx), uint(idInt))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Message deleted successfully"})
}

// 단일 메시지 조회
func (m *messageController) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	message, err := m.service.MessageService().Get(appctx.NewContext(ctx), uint(idInt))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, message)
}

// 메시지 목록 조회
func (m *messageController) GetList(ctx *gin.Context) {
	messages, err := m.service.MessageService().GetList(appctx.NewContext(ctx), message.GetListParam{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, messages)
}

func newMessageController(service domain.Service) message.Controller {
	return &messageController{
		service: service,
	}
}
