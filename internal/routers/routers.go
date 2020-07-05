package routers

import (
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"server_simple/internal/middleware"
	"server_simple/internal/routers/v1/client"
	"server_simple/internal/routers/v1/puclic"
	"server_simple/internal/routers/v1/ws"
)

func Init() *gin.Engine {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type,Uid",
		ExposedHeaders:  "",
		Credentials:     true,
		ValidateHeaders: false,
	}))

	apiV1Client := r.Group("v1/client") //.Use(middleware.NewMiddleware().VerifyToken())
	{
		apiV1Client.GET("/blog/list", client.BlogController.GetList)
		apiV1Client.GET("/blog/detail", client.BlogController.GetDetail)
		apiV1Client.GET("/blog/types", client.BlogController.GetTypes)
		apiV1Client.GET("/blog/ranking", client.BlogController.GetRanking)
		apiV1Client.GET("/blog/recommend", client.BlogController.GetRecommend)
		apiV1Client.GET("/blog/comment/list", client.BlogController.GetCommentList)
		apiV1Client.POST("/blog/comment/submit", client.BlogController.SubmitComment)
		apiV1Client.POST("/blog/good/make", client.BlogController.MakeGood)
		apiV1Client.GET("/blog/room/list", client.BlogController.GetRoomList)
	}

	apiV1Public := r.Group("v1/public")
	{
		apiV1Public.GET("netfile/read", puclic.PublicController.ReadNetFile)
	}

	apiV1Ws := r.Group("v1/ws").Use(middleware.Ws.HttpUpgrader())
	{
		apiV1Ws.GET("/chatroom", ws.Ws.ChatRoom)
		apiV1Ws.GET("/gdraw", ws.Ws.Gdraw)
	}

	apiV1OpenApiGdraw := r.Group("v1/gdraw/openapi") //.Use(middleware.NewMiddleware().VerifyToken())
	{
		apiV1OpenApiGdraw.POST("/user/auth", client.GdrawUserController.UserAuth)
		apiV1OpenApiGdraw.POST("/user/regist", client.GdrawUserController.UserRegist)
		apiV1OpenApiGdraw.POST("/user/login", client.GdrawUserController.Login)
		apiV1OpenApiGdraw.GET("/vcode/get", client.GdrawVcodeController.Get)
		apiV1OpenApiGdraw.GET("/token/verify", client.GdrawUserController.TokenVerify)
	}

	apiV1ApiGdraw := r.Group("v1/gdraw/api") //.Use(services.GdrawUserer.VerifyToken())
	{
		apiV1ApiGdraw.GET("/room/list", client.GdrawController.GetRoomList)
	}
	return r
}
