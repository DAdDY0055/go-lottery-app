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

type UserHandler struct {
	Db *gorm.DB
}

// func failOnError(err error) {
//     if err != nil {
//         log.Fatal("Error:", err)
//     }
// }

// 一覧表示
func (handler *UserHandler) GetAll(c *gin.Context) {
	var Users []models.User
	handler.Db.Find(&Users) // DBから全てのレコードを取得する
	c.HTML(http.StatusOK, "index.html", gin.H{"users": Users}) // index.htmlに全てのレコードを渡す
}

// 新規作成
func (handler *UserHandler) Create(c *gin.Context) {
	name, _ := c.GetPostForm("name") // index.htmlからnameを取得
	department, _ := c.GetPostForm("department") // index.htmlからdepartmentを取得
	handler.Db.Create(&models.User{Name: name, Department: department}) // レコードを挿入する
  c.Redirect(http.StatusMovedPermanently, "/")
}

// 編集画面
func (handler *UserHandler) Edit(c *gin.Context) {
	user := models.User{} // Task構造体の変数宣言
	id := c.Param("id")   // index.htmlからidを取得
	handler.Db.First(&user, id)  // idに一致するレコードを取得する
	c.HTML(http.StatusOK, "edit.html", gin.H{"user": user})
}

// 更新
func (handler *UserHandler) Update(c *gin.Context) {
	user := models.User{} // Task構造体の変数宣言
	id := c.Param("id")   // edit.htmlからidを取得
	name, _ := c.GetPostForm("name") // index.htmlからnameを取得
	department, _ := c.GetPostForm("department") // index.htmlからdepartmentを取得
	handler.Db.First(&user, id)      // idに一致するレコードを取得する
	user.Name = name       // nameを上書きする
	user.Department = department       // nameを上書きする
	handler.Db.Save(&user) // 指定のレコードを更新する
	c.Redirect(http.StatusMovedPermanently, "/")
}

// 削除
func (handler *UserHandler) Delete(c *gin.Context) {
	user := models.User{} // Task構造体の変数宣言
	id := c.Param("id")   // index.htmlからidを取得
	handler.Db.First(&user, id) // idに一致するレコードを取得する
	handler.Db.Delete(&user)    // 指定のレコードを削除する
	c.Redirect(http.StatusMovedPermanently, "/")
}

// CSV読み込み
func (handler *UserHandler) ReadCsv(c *gin.Context) {
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
		department := record[1]
		handler.Db.Create(&models.User{Name: name, Department: department}) // レコードを挿入する

        log.Printf("%#v", record)
    }
    c.Redirect(http.StatusMovedPermanently, "/")
}
