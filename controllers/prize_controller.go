package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/DAdDY0055/go-lottery-app/models"
	"github.com/jinzhu/gorm"
	"net/http"
	"encoding/csv"
    "io"
    "log"
    "os"
)

type Handler struct {
	Db *gorm.DB
}

// 一覧表示
func (handler *Handler) GetAll(c *gin.Context) {
	var Prizes []models.Prize
	handler.Db.Find(&Prizes) // DBから全てのレコードを取得する
	c.HTML(http.StatusOK, "index.html", gin.H{"prizes": Prizes}) // index.htmlに全てのレコードを渡す
}

// 新規作成
func (handler *Handler) Create(c *gin.Context) {
	name, _ := c.GetPostForm("name") // index.htmlからnameを取得
	handler.Db.Create(&models.Prize{Name: name}) // レコードを挿入する
  c.Redirect(http.StatusMovedPermanently, "/")
}

// 編集画面
func (handler *Handler) Edit(c *gin.Context) {
	prize := models.Prize{} // Task構造体の変数宣言
	id := c.Param("id")   // index.htmlからidを取得
	handler.Db.First(&prize, id)  // idに一致するレコードを取得する
	c.HTML(http.StatusOK, "edit.html", gin.H{"prize": prize})
}

// 更新
func (handler *Handler) Update(c *gin.Context) {
	prize := models.Prize{} // Task構造体の変数宣言
	id := c.Param("id")   // edit.htmlからidを取得
	name, _ := c.GetPostForm("name") // index.htmlからnameを取得
	handler.Db.First(&prize, id)      // idに一致するレコードを取得する
	prize.Name = name       // nameを上書きする
	handler.Db.Save(&prize) // 指定のレコードを更新する
	c.Redirect(http.StatusMovedPermanently, "/")
}

// 削除
func (handler *Handler) Delete(c *gin.Context) {
	prize := models.Prize{} // Task構造体の変数宣言
	id := c.Param("id")   // index.htmlからidを取得
	handler.Db.First(&prize, id) // idに一致するレコードを取得する
	handler.Db.Delete(&prize)    // 指定のレコードを削除する
	c.Redirect(http.StatusMovedPermanently, "/")
}

// CSV読み込み
func (handler *Handler) ReadCsv(c *gin.Context) {
	ff := c.Param("csv")
	file, err := os.Open(ff)
    failOnError(err)
    defer file.Close()

    reader := csv.NewReader(file)

    for {
        record, err := reader.Read() // 1行読み出す
        if err == io.EOF {
            break
        } else {
            failOnError(err)
        }
		name := record[0]
		handler.Db.Create(&models.Prize{Name: name}) // レコードを挿入する

        log.Printf("%#v", record)
    }
    c.Redirect(http.StatusMovedPermanently, "/")
}

func failOnError(err error) {
    if err != nil {
        log.Fatal("Error:", err)
    }
}
