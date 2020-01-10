package controllers

import (
	"github.com/DAdDY0055/go-lottery-app/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	// "encoding/csv"
	// "io"
	// "log"
	// "os"
	"math/rand"
	"strconv"
	"time"
)

type LottelyHandler struct {
	Db *gorm.DB
}

// TODO: 超ファットコントローラーになっているのでモデルに処理を移す
// 抽選画面
// トップページ
func (handler *LottelyHandler) Top(c *gin.Context) {
	c.HTML(http.StatusOK, "lottery_index.html", gin.H{})
}

// おまけ商品ページ
func (handler *LottelyHandler) Next(c *gin.Context) {
	c.HTML(http.StatusOK, "lottery_next.html", gin.H{})
}

// 景品選択
// func (handler *LottelyHandler) WinnerAnnounce(c *gin.Context) {
// 	var WinPrizes []models.Prize
// 	var Users []models.User

// 	handler.Db.Find(&WinPrizes) // DBから全てのレコードを取得する
// 	handler.Db.Where("IsWin = ?", false).Find(&Users) // DBから当選していない全てのユーザーを取得する

//     rand.Seed(time.Now().UnixNano())
//     for i := range WinPrizes {
//         j := rand.Intn(i + 1)
//         WinPrizes[i], WinPrizes[j] = WinPrizes[j], WinPrizes[i]
// 	}

// 	WinPrize := WinPrizes[0]

// 	c.HTML(http.StatusOK, "lottery_prize.html", gin.H{"win": WinPrize})
// }

// 当選者選択(一つの景品対一つの商品)
func (handler *LottelyHandler) ChoiseOne(c *gin.Context) {
	var WinUsers []models.User
	winPrize := models.Prize{}

	id := c.Param("id") // edit.htmlからidを取得
	handler.Db.Where("id = ?", id).Find(&winPrize)

	// 対象の商品に当選者がいなかった場合、抽選を行う
	if winPrize.Winner == "" {
		handler.Db.Not("win = ?", "済").Find(&WinUsers) // DBから当選済み以外のユーザーを抽出する

		rand.Seed(time.Now().UnixNano())
		for i := range WinUsers {
			j := rand.Intn(i + 1)
			WinUsers[i], WinUsers[j] = WinUsers[j], WinUsers[i]
		}

		winUser := WinUsers[0]

		winUser.Win = "済"         // 当選したら当選済みにする
		handler.Db.Save(&winUser) // 指定のレコードを更新する

		winPrize.Winner = winUser.Name // 当選者名を入れる
		handler.Db.Save(&winPrize)     // 指定のレコードを更新する

		c.HTML(http.StatusOK, "lottery_winner.html", gin.H{"prize": winPrize})
	} else {
		c.HTML(http.StatusOK, "lottery_winner.html", gin.H{"prize": winPrize})
	}
}

// 10位〜14位の当選者選択
func (handler *LottelyHandler) ChoiseTwe(c *gin.Context) {
	var WinUsers []models.User
	var Prizes []models.Prize
	// まとめて型定義する方法があった気がする
	// prize   := models.Prize{}
	PrizeNumbers := []int{10, 11, 12, 13, 14}

	// 対象の商品に当選者がいなかった場合、抽選を行う
	for _, prizeNumber := range PrizeNumbers {
		prize := models.Prize{}
		handler.Db.Where("id = ?", prizeNumber).Find(&prize)

		if prize.Winner == "" {
			handler.Db.Not("win = ?", "済").Find(&WinUsers) // DBから当選済み以外のユーザーを抽出する

			rand.Seed(time.Now().UnixNano())
			for i := range WinUsers {
				j := rand.Intn(i + 1)
				WinUsers[i], WinUsers[j] = WinUsers[j], WinUsers[i]
			}

			winUser := models.User{} // 変数を初期化
			winUser = WinUsers[0]

			winUser.Win = "済"         // 当選したら当選済みにする
			handler.Db.Save(&winUser) // 指定のレコードを更新する

			prize.Winner = winUser.Name // 当選者名を入れる
			handler.Db.Save(&prize)     // 指定のレコードを更新する
		}
	}

	handler.Db.Where("id IN (?)", PrizeNumbers).Find(&Prizes)

	c.HTML(http.StatusOK, "lottery_winners.html", gin.H{"prizes": Prizes})
}

// 15位〜29位の当選者選択
func (handler *LottelyHandler) ChoiseTen(c *gin.Context) {
	var WinUsers []models.User
	var Prizes []models.Prize
	// まとめて型定義する方法があった気がする
	// prize   := models.Prize{}
	PrizeNumbers := []int{15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29}

	// 対象の商品に当選者がいなかった場合、抽選を行う
	for _, prizeNumber := range PrizeNumbers {
		prize := models.Prize{}
		handler.Db.Where("id = ?", prizeNumber).Find(&prize)

		if prize.Winner == "" {
			handler.Db.Not("win = ?", "済").Find(&WinUsers) // DBから当選済み以外のユーザーを抽出する

			rand.Seed(time.Now().UnixNano())
			for i := range WinUsers {
				j := rand.Intn(i + 1)
				WinUsers[i], WinUsers[j] = WinUsers[j], WinUsers[i]
			}

			winUser := models.User{} // 変数を初期化
			winUser = WinUsers[0]

			winUser.Win = "済"         // 当選したら当選済みにする
			handler.Db.Save(&winUser) // 指定のレコードを更新する

			prize.Winner = winUser.Name // 当選者名を入れる
			handler.Db.Save(&prize)     // 指定のレコードを更新する
		}
	}

	handler.Db.Where("id IN (?)", PrizeNumbers).Find(&Prizes)

	c.HTML(http.StatusOK, "lottery_winners.html", gin.H{"prizes": Prizes})
}

// 30位〜35位の当選者選択
func (handler *LottelyHandler) ChoiseThe(c *gin.Context) {
	var WinUsers []models.User
	var Prizes []models.Prize
	// まとめて型定義する方法があった気がする
	// prize   := models.Prize{}
	PrizeNumbers := []int{30, 31, 32, 33, 34, 35}

	// 対象の商品に当選者がいなかった場合、抽選を行う
	for _, prizeNumber := range PrizeNumbers {
		prize := models.Prize{}
		handler.Db.Where("id = ?", prizeNumber).Find(&prize)

		if prize.Winner == "" {
			handler.Db.Not("win = ?", "済").Find(&WinUsers) // DBから当選済み以外のユーザーを抽出する

			rand.Seed(time.Now().UnixNano())
			for i := range WinUsers {
				j := rand.Intn(i + 1)
				WinUsers[i], WinUsers[j] = WinUsers[j], WinUsers[i]
			}

			winUser := models.User{} // 変数を初期化
			winUser = WinUsers[0]

			winUser.Win = "済"         // 当選したら当選済みにする
			handler.Db.Save(&winUser) // 指定のレコードを更新する

			prize.Winner = winUser.Name // 当選者名を入れる
			handler.Db.Save(&prize)     // 指定のレコードを更新する
		}
	}

	handler.Db.Where("id IN (?)", PrizeNumbers).Find(&Prizes)

	c.HTML(http.StatusOK, "lottery_winners.html", gin.H{"prizes": Prizes})
}

// 36位〜40位の当選者選択
func (handler *LottelyHandler) ChoiseThe2(c *gin.Context) {
	var WinUsers []models.User
	var Prizes []models.Prize
	// まとめて型定義する方法があった気がする
	// prize   := models.Prize{}
	PrizeNumbers := []int{36, 37, 38, 39, 40}

	// 対象の商品に当選者がいなかった場合、抽選を行う
	for _, prizeNumber := range PrizeNumbers {
		prize := models.Prize{}
		handler.Db.Where("id = ?", prizeNumber).Find(&prize)

		if prize.Winner == "" {
			handler.Db.Not("win = ?", "済").Find(&WinUsers) // DBから当選済み以外のユーザーを抽出する

			rand.Seed(time.Now().UnixNano())
			for i := range WinUsers {
				j := rand.Intn(i + 1)
				WinUsers[i], WinUsers[j] = WinUsers[j], WinUsers[i]
			}

			winUser := models.User{} // 変数を初期化
			winUser = WinUsers[0]

			winUser.Win = "済"         // 当選したら当選済みにする
			handler.Db.Save(&winUser) // 指定のレコードを更新する

			prize.Winner = winUser.Name // 当選者名を入れる
			handler.Db.Save(&prize)     // 指定のレコードを更新する
		}
	}

	handler.Db.Where("id IN (?)", PrizeNumbers).Find(&Prizes)

	c.HTML(http.StatusOK, "lottery_winners.html", gin.H{"prizes": Prizes})
}

// 41位〜43位の当選者選択
func (handler *LottelyHandler) Choisefor(c *gin.Context) {
	var WinUsers []models.User
	var Prizes []models.Prize
	// まとめて型定義する方法があった気がする
	// prize   := models.Prize{}
	PrizeNumbers := []int{41, 42, 43}

	// 対象の商品に当選者がいなかった場合、抽選を行う
	for _, prizeNumber := range PrizeNumbers {
		prize := models.Prize{}
		handler.Db.Where("id = ?", prizeNumber).Find(&prize)

		if prize.Winner == "" {
			handler.Db.Not("win = ?", "済").Find(&WinUsers) // DBから当選済み以外のユーザーを抽出する

			rand.Seed(time.Now().UnixNano())
			for i := range WinUsers {
				j := rand.Intn(i + 1)
				WinUsers[i], WinUsers[j] = WinUsers[j], WinUsers[i]
			}

			winUser := models.User{} // 変数を初期化
			winUser = WinUsers[0]

			winUser.Win = "済"         // 当選したら当選済みにする
			handler.Db.Save(&winUser) // 指定のレコードを更新する

			prize.Winner = winUser.Name // 当選者名を入れる
			handler.Db.Save(&prize)     // 指定のレコードを更新する
		}
	}

	handler.Db.Where("id IN (?)", PrizeNumbers).Find(&Prizes)

	c.HTML(http.StatusOK, "lottery_winners.html", gin.H{"prizes": Prizes})
}

// あったで賞(46位〜49,51,53位)の当選者選択
func (handler *LottelyHandler) Choisearu(c *gin.Context) {
	var WinUsers []models.User
	var Prizes []models.Prize
	// まとめて型定義する方法があった気がする
	// prize   := models.Prize{}
	PrizeNumbers := []int{46, 47, 48, 49, 50, 51, 52, 53}

	// 対象の商品に当選者がいなかった場合、抽選を行う
	for _, prizeNumber := range PrizeNumbers {
		prize := models.Prize{}
		handler.Db.Where("id = ?", prizeNumber).Find(&prize)

		if prize.Winner == "" {
			handler.Db.Not("win = ?", "済").Find(&WinUsers) // DBから当選済み以外のユーザーを抽出する

			rand.Seed(time.Now().UnixNano())
			for i := range WinUsers {
				j := rand.Intn(i + 1)
				WinUsers[i], WinUsers[j] = WinUsers[j], WinUsers[i]
			}

			winUser := models.User{} // 変数を初期化
			winUser = WinUsers[0]

			winUser.Win = "済"         // 当選したら当選済みにする
			handler.Db.Save(&winUser) // 指定のレコードを更新する

			prize.Winner = winUser.Name // 当選者名を入れる
			handler.Db.Save(&prize)     // 指定のレコードを更新する
		}
	}

	handler.Db.Where("id IN (?)", PrizeNumbers).Find(&Prizes)

	c.HTML(http.StatusOK, "lottery_winners.html", gin.H{"prizes": Prizes})
}

// 参加者登録
// 一覧表示
func (handler *LottelyHandler) UserGetAll(c *gin.Context) {
	var Users []models.User
	handler.Db.Order("id asc").Find(&Users)                         // DBから全てのレコードを取得する
	c.HTML(http.StatusOK, "user_index.html", gin.H{"users": Users}) // user_index.htmlに全てのレコードを渡す
}

// 新規作成
func (handler *LottelyHandler) UserCreate(c *gin.Context) {
	stringID, _ := c.GetPostForm("id")                  // index.htmlからidを取得
	id, _ := strconv.Atoi(stringID)                     // idをStringからintに変換
	name, _ := c.GetPostForm("name")                    // user_index.htmlからnameを取得
	handler.Db.Create(&models.User{ID: id, Name: name}) // レコードを挿入する
	c.Redirect(http.StatusMovedPermanently, "/user")
}

// 編集画面
func (handler *LottelyHandler) UserEdit(c *gin.Context) {
	user := models.User{}       // Task構造体の変数宣言
	id := c.Param("id")         // user_index.htmlからidを取得
	handler.Db.First(&user, id) // idに一致するレコードを取得する
	c.HTML(http.StatusOK, "user_edit.html", gin.H{"user": user})
}

// 更新
func (handler *LottelyHandler) UserUpdate(c *gin.Context) {
	user := models.User{} // Task構造体の変数宣言

	stringID := c.Param("id")        // edit.htmlからidを取得
	id, _ := strconv.Atoi(stringID)  // idをStringからintに変換
	name, _ := c.GetPostForm("name") // index.htmlからnameを取得
	// department, _ := c.GetPostForm("department") // index.htmlからdepartmentを取得
	win, _ := c.GetPostForm("win") // index.htmlからdepartmentを取得

	handler.Db.First(&user, id) // idに一致するレコードを取得する

	user.ID = id           // idを上書きする
	user.Name = name       // nameを上書きする
	user.Win = win         // Winを上書きする
	handler.Db.Save(&user) // 指定のレコードを更新する
	c.Redirect(http.StatusMovedPermanently, "/user")
}

// 削除
func (handler *LottelyHandler) UserDelete(c *gin.Context) {
	user := models.User{}       // Task構造体の変数宣言
	id := c.Param("id")         // index.htmlからidを取得
	handler.Db.First(&user, id) // idに一致するレコードを取得する
	handler.Db.Delete(&user)    // 指定のレコードを削除する
	c.Redirect(http.StatusMovedPermanently, "/user")
}

// CSV読み込み
// func (handler *LottelyHandler) UserReadCsv(c *gin.Context) {
// 	ff := c.Param("csv")
// 	file, err := os.Open(ff)
//     failOnError(err)
//     defer file.Close()

//     reader := csv.NewReader(file)

//     for {
//         record, err := reader.Read() // 1行読み出す
//         if err == io.EOF {
//             break
//         } else {
//             failOnError(err)
//         }
// 		name := record[0]
// 		department := record[1]

//         log.Printf("%#v", record)
//     }
//     c.Redirect(http.StatusMovedPermanently, "/user")
// }

// 景品登録
// 一覧表示
func (handler *LottelyHandler) PrizeGetAll(c *gin.Context) {
	var Prizes []models.Prize
	handler.Db.Order("id asc").Find(&Prizes)                           // DBから全てのレコードを取得する
	c.HTML(http.StatusOK, "prize_index.html", gin.H{"prizes": Prizes}) // prize_index.htmlに全てのレコードを渡す
}

// 新規作成
func (handler *LottelyHandler) PrizeCreate(c *gin.Context) {
	stringID, _ := c.GetPostForm("id")                   // prize_index.htmlからidを取得
	id, _ := strconv.Atoi(stringID)                      // idをStringからintに変換
	name, _ := c.GetPostForm("name")                     // prize_index.htmlからnameを取得
	handler.Db.Create(&models.Prize{ID: id, Name: name}) // レコードを挿入する
	c.Redirect(http.StatusMovedPermanently, "/prize")
}

// 編集画面
func (handler *LottelyHandler) PrizeEdit(c *gin.Context) {
	prize := models.Prize{}      // Task構造体の変数宣言
	id := c.Param("id")          // prize_index.htmlからidを取得
	handler.Db.First(&prize, id) // idに一致するレコードを取得する
	c.HTML(http.StatusOK, "prize_edit.html", gin.H{"prize": prize})
}

// 更新
func (handler *LottelyHandler) PrizeUpdate(c *gin.Context) {
	prize := models.Prize{}          // Task構造体の変数宣言
	stringID := c.Param("id")        // edit.htmlからidを取得
	id, _ := strconv.Atoi(stringID)  // idをStringからintに変換
	name, _ := c.GetPostForm("name") // prize_edit.htmlからnameを取得
	win, _ := c.GetPostForm("win")   // prize_edit.htmlからwinを取得
	handler.Db.First(&prize, id)     // idに一致するレコードを取得する
	prize.ID = id                    // idを上書きする
	prize.Name = name                // nameを上書きする
	prize.Winner = win               // winを上書きする
	handler.Db.Save(&prize)          // 指定のレコードを更新する
	c.Redirect(http.StatusMovedPermanently, "/prize")
}

// 削除
func (handler *LottelyHandler) PrizeDelete(c *gin.Context) {
	prize := models.Prize{}      // Task構造体の変数宣言
	id := c.Param("id")          // prize_index.htmlからidを取得
	handler.Db.First(&prize, id) // idに一致するレコードを取得する
	handler.Db.Delete(&prize)    // 指定のレコードを削除する
	c.Redirect(http.StatusMovedPermanently, "/prize")
}

// CSV読み込み
// func (handler *LottelyHandler) PrizeReadCsv(c *gin.Context) {
// 	ff := c.Param("csv")
// 	file, err := os.Open(ff)
//     failOnError(err)
//     defer file.Close()

//     reader := csv.NewReader(file)

//     for {
//         record, err := reader.Read() // 1行読み出す
//         if err == io.EOF {
//             break
//         } else {
//             failOnError(err)
//         }
// 		name := record[0]
// 		handler.Db.Create(&models.Prize{Name: name}) // レコードを挿入する

//         log.Printf("%#v", record)
//     }
//     c.Redirect(http.StatusMovedPermanently, "/prize")
// }

// func failOnError(err error) {
//     if err != nil {
//         log.Fatal("Error:", err)
//     }
// }
