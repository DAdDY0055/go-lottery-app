package router

import (
	"github.com/gin-gonic/gin"
	"github.com/DAdDY0055/go-lottery-app/controllers"
	"github.com/DAdDY0055/go-lottery-app/db"
)

func Router() {
	router := gin.Default()

	// globパターンに一致するHTMLファイルをロードしHTML Rendererに関連付ける
	router.LoadHTMLGlob("templates/*.html")

	// TaskHandler構造体に紐付けたCRUDメソッドを呼び出す
	handler := controllers.TaskHandler{
		db.Get(),
	}

	router.GET("/", handler.GetAll)
	router.POST("/", handler.Create)
	router.GET("/:id", handler.Edit)
	router.POST("/update/:id", handler.Update)
	router.POST("/delete/:id", handler.Delete)

  // Routerをhttp.Serverに接続し、HTTPリクエストのリスニングとサービスを開始する
  router.Run()
}
