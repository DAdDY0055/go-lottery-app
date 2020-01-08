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
    "math/rand"
	"time"
)

type LottelyHandler struct {
	Db *gorm.DB
}

// 抽選画面
// トップページ
func (handler *LottelyHandler) Top(c *gin.Context) {
	c.HTML(http.StatusOK, "lottery_index.html", gin.H{})
}

// 当選発表
func (handler *LottelyHandler) WinnerAnnounce(c *gin.Context) {
	var WinPrizes []models.Prize
	var Users []models.User

	handler.Db.Find(&WinPrizes) // DBから全てのレコードを取得する
	handler.Db.Where("IsWin = ?", false).Find(&Users) // DBから全てのレコードを取得する

    rand.Seed(time.Now().UnixNano())
    for i := range WinPrizes {
        j := rand.Intn(i + 1)
        WinPrizes[i], WinPrizes[j] = WinPrizes[j], WinPrizes[i]
	}

	WinPrize := WinPrizes[0]

	c.HTML(http.StatusOK, "lottery_prize.html", gin.H{"win": WinPrize})
}

// 当選者選択
func (handler *LottelyHandler) ChoiseUser(c *gin.Context) {
	var WinUsers []models.User
	// var WinPrize []models.Prize

	// 対象の商品に当選者がいなかった場合、抽選を行う
		handler.Db.Not("win = ?", "sumi").Find(&WinUsers) // DBから全てのレコードを取得する

		rand.Seed(time.Now().UnixNano())
		for i := range WinUsers {
			j := rand.Intn(i + 1)
			WinUsers[i], WinUsers[j] = WinUsers[j], WinUsers[i]
		}

		winUser := WinUsers[0]

		winUser.Win = "sumi"       // 当選したら当選済みにする
		handler.Db.Save(&winUser)  // 指定のレコードを更新する

	c.HTML(http.StatusOK, "lottery_winner.html", gin.H{"winner": winUser})
}


// 参加者登録
// 一覧表示
func (handler *LottelyHandler) UserGetAll(c *gin.Context) {
	var Users []models.User
	handler.Db.Find(&Users) // DBから全てのレコードを取得する
	c.HTML(http.StatusOK, "user_index.html", gin.H{"users": Users}) // user_index.htmlに全てのレコードを渡す
}

// 新規作成
func (handler *LottelyHandler) UserCreate(c *gin.Context) {
	name, _ := c.GetPostForm("name") // user_index.htmlからnameを取得
	department, _ := c.GetPostForm("department") // index.htmlからdepartmentを取得
	handler.Db.Create(&models.User{Name: name, Department: department}) // レコードを挿入する
  c.Redirect(http.StatusMovedPermanently, "/user")
}

// 編集画面
func (handler *LottelyHandler) UserEdit(c *gin.Context) {
	user := models.User{} // Task構造体の変数宣言
	id := c.Param("id")   // user_index.htmlからidを取得
	handler.Db.First(&user, id)  // idに一致するレコードを取得する
	c.HTML(http.StatusOK, "user_edit.html", gin.H{"user": user})
}

// 更新
func (handler *LottelyHandler) UserUpdate(c *gin.Context) {
	user := models.User{} // Task構造体の変数宣言
	id := c.Param("id")   // edit.htmlからidを取得
	name, _ := c.GetPostForm("name") // index.htmlからnameを取得
	department, _ := c.GetPostForm("department") // index.htmlからdepartmentを取得
	// isWin, _ := c.GetPostForm("isWin") // index.htmlからdepartmentを取得
	handler.Db.First(&user, id)      // idに一致するレコードを取得する
	user.Name = name       // nameを上書きする
	user.Department = department  // departmentを上書きする
	// user.IsWin = isWin       // isWinを上書きする→Stringと判定されダメらしい
	handler.Db.Save(&user) // 指定のレコードを更新する
	c.Redirect(http.StatusMovedPermanently, "/user")
}

// 削除
func (handler *LottelyHandler) UserDelete(c *gin.Context) {
	user := models.User{} // Task構造体の変数宣言
	id := c.Param("id")   // index.htmlからidを取得
	handler.Db.First(&user, id) // idに一致するレコードを取得する
	handler.Db.Delete(&user)    // 指定のレコードを削除する
	c.Redirect(http.StatusMovedPermanently, "/user")
}

// CSV読み込み
func (handler *LottelyHandler) UserReadCsv(c *gin.Context) {
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
    c.Redirect(http.StatusMovedPermanently, "/user")
}

// 景品登録
// 一覧表示
func (handler *LottelyHandler) PrizeGetAll(c *gin.Context) {
	var Prizes []models.Prize
	handler.Db.Find(&Prizes) // DBから全てのレコードを取得する
	c.HTML(http.StatusOK, "prize_index.html", gin.H{"prizes": Prizes}) // prize_index.htmlに全てのレコードを渡す
}

// 新規作成
func (handler *LottelyHandler) PrizeCreate(c *gin.Context) {
	name, _ := c.GetPostForm("name") // prize_index.htmlからnameを取得
	handler.Db.Create(&models.Prize{Name: name}) // レコードを挿入する
  c.Redirect(http.StatusMovedPermanently, "/prize")
}

// 編集画面
func (handler *LottelyHandler) PrizeEdit(c *gin.Context) {
	prize := models.Prize{} // Task構造体の変数宣言
	id := c.Param("id")   // prize_index.htmlからidを取得
	handler.Db.First(&prize, id)  // idに一致するレコードを取得する
	c.HTML(http.StatusOK, "prize_edit.html", gin.H{"prize": prize})
}

// 更新
func (handler *LottelyHandler) PrizeUpdate(c *gin.Context) {
	prize := models.Prize{} // Task構造体の変数宣言
	id := c.Param("id")   // edit.htmlからidを取得
	name, _ := c.GetPostForm("name") // prize_edit.htmlからnameを取得
	win, _ := c.GetPostForm("win") // prize_edit.htmlからwinを取得
	handler.Db.First(&prize, id)      // idに一致するレコードを取得する
	prize.Name = name       // nameを上書きする
	prize.WinUserName = win // winを上書きする
	handler.Db.Save(&prize) // 指定のレコードを更新する
	c.Redirect(http.StatusMovedPermanently, "/prize")
}

// 削除
func (handler *LottelyHandler) PrizeDelete(c *gin.Context) {
	prize := models.Prize{} // Task構造体の変数宣言
	id := c.Param("id")   // prize_index.htmlからidを取得
	handler.Db.First(&prize, id) // idに一致するレコードを取得する
	handler.Db.Delete(&prize)    // 指定のレコードを削除する
	c.Redirect(http.StatusMovedPermanently, "/prize")
}

// CSV読み込み
func (handler *LottelyHandler) PrizeReadCsv(c *gin.Context) {
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
    c.Redirect(http.StatusMovedPermanently, "/prize")
}

// func failOnError(err error) {
//     if err != nil {
//         log.Fatal("Error:", err)
//     }
// }
