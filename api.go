package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListAccounts(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, nil)
}
