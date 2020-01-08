package router

import (
	"github.com/gin-gonic/gin"
	"github.com/DAdDY0055/go-lottery-app/controllers"
	"github.com/DAdDY0055/go-lottery-app/db"
)

func Router() {
	router := gin.Default()

	// globパターンに一致するHTMLファイルをロードしHTML Rendererに関連付ける
	router.LoadHTMLGlob("templates/*/*.html")

	// UserHandler構造体に紐付けたCRUDメソッドを呼び出す
	handler := controllers.LottelyHandler{
		db.Get(),
	}

	// 抽選
	router.GET("/", handler.Top)
	router.GET("/lottery", handler.WinnerAnnounce)
	// router.GET("/lottery/aru", handler.ChoisePrize)
	// router.GET("/lottery/kj", handler.ChoisePrize)
	// router.GET("/lottery/yamamura", handler.ChoisePrize)
	// router.GET("/lottery/kogishi", handler.ChoisePrize)
	router.GET("/winner", handler.ChoiseUser)

	// ユーザー
	router.GET("/user", handler.UserGetAll)
	router.POST("/user", handler.UserCreate)
	router.GET("/user/:id", handler.UserEdit)
	router.POST("/user/update/:id", handler.UserUpdate)
	router.POST("/user/delete/:id", handler.UserDelete)
	router.POST("/user/csv", handler.UserReadCsv)

	// 商品
	router.GET("/prize", handler.PrizeGetAll)
	router.POST("/prize", handler.PrizeCreate)
	router.GET("/prize/:id", handler.PrizeEdit)
	router.POST("/prize/update/:id", handler.PrizeUpdate)
	router.POST("/prize/delete/:id", handler.PrizeDelete)
	router.POST("/prize/csv", handler.PrizeReadCsv)

  // Routerをhttp.Serverに接続し、HTTPリクエストのリスニングとサービスを開始する
  router.Run()
}
