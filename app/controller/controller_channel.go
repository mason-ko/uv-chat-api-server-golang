package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"uv-chat-api-server-golang/domain"
	"uv-chat-api-server-golang/domain/channel"
	"uv-chat-api-server-golang/internal/appctx"
)

type channelController struct {
	service domain.Service
}

// 메시지 생성
func (m *channelController) Create(ctx *gin.Context) {
	var messageDTO channel.Channel
	if err := ctx.ShouldBindJSON(&messageDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := m.service.ChannelService().Create(appctx.NewContext(ctx), messageDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Message created successfully"})
}

// 메시지 삭제
func (m *channelController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = m.service.ChannelService().Delete(appctx.NewContext(ctx), uint(idInt))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Message deleted successfully"})
}

// 단일 메시지 조회
func (m *channelController) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	message, err := m.service.ChannelService().Get(appctx.NewContext(ctx), uint(idInt))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, message)
}

// 메시지 목록 조회
func (m *channelController) GetList(ctx *gin.Context) {
	messages, err := m.service.ChannelService().GetList(appctx.NewContext(ctx), channel.GetListParam{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, messages)
}

func newChannelController(service domain.Service) channel.Controller {
	return &channelController{
		service: service,
	}
}
