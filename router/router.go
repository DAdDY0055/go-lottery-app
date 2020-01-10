package router

import (
	"github.com/DAdDY0055/go-lottery-app/controllers"
	"github.com/DAdDY0055/go-lottery-app/db"
	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()

	router.Static("/assets", "./assets")
	// globパターンに一致するHTMLファイルをロードしHTML Rendererに関連付ける
	router.LoadHTMLGlob("templates/*/*.html")

	// UserHandler構造体に紐付けたCRUDメソッドを呼び出す
	handler := controllers.LottelyHandler{
		db.Get(),
	}

	// 抽選
	router.GET("/", handler.Top)
	router.GET("/next", handler.Next)
	router.GET("/winner/one/:id", handler.ChoiseOne) // 一つの景品発表の場合はこれでID指定
	router.GET("/winner/ten", handler.ChoiseTen)
	router.GET("/winner/twe", handler.ChoiseTwe)
	router.GET("/winner/the", handler.ChoiseThe)
	router.GET("/winner/the2", handler.ChoiseThe2)
	router.GET("/winner/for", handler.Choisefor)
	router.GET("/winner/aru", handler.Choisearu)

	// ユーザー
	router.GET("/user", handler.UserGetAll)
	router.POST("/user", handler.UserCreate)
	router.GET("/user/:id", handler.UserEdit)
	router.POST("/user/update/:id", handler.UserUpdate)
	router.POST("/user/delete/:id", handler.UserDelete)
	// router.POST("/user/csv", handler.UserReadCsv)

	// 商品
	router.GET("/prize", handler.PrizeGetAll)
	router.POST("/prize", handler.PrizeCreate)
	router.GET("/prize/:id", handler.PrizeEdit)
	router.POST("/prize/update/:id", handler.PrizeUpdate)
	router.POST("/prize/delete/:id", handler.PrizeDelete)
	// router.POST("/prize/csv", handler.PrizeReadCsv)

	// Routerをhttp.Serverに接続し、HTTPリクエストのリスニングとサービスを開始する
	router.Run()
}
