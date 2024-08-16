package account

import (
	"setup_go/database"
	modelsPg "setup_go/models/pg"
	"setup_go/utility"
	"time"

	"github.com/gofiber/fiber/v2"
)

type responseGetAcccount struct {
	Username    string `json:"username"`
	FirstNameTh string `json:"first_name_th"`
}

// GetAuthAgora
// @Summary ดึงข้อมูล appid cert
// @Description ดึงข้อมูล appid cert
// @Security ApiKeyAuth
// @Tags App Store
// @Accept json
// @Produce json
// @Success 200 {object} models.ResponseSuccess{data=resGetAuthAgora}
// @Failure 401 {object} models.ResponseError401
// @Failure 500 {object} models.ResponseError
// @Router /v1/get-auth-agora [get]
func GetAccount(c *fiber.Ctx) error {
	var result *modelsPg.Account
	database.CachingCtx().Get("account_data", &result)

	if result == nil { // without cache

		var accountData modelsPg.Account
		db := database.DBConn
		err := db.Where("").Find(&accountData).Error
		if err != nil {
			return utility.ResponseError(c, fiber.StatusBadRequest, err.Error())
		}

		// store in redis
		err = database.CachingCtx().Set("account_data", accountData, time.Minute*5)
		if err != nil {
			return utility.ResponseError(c, fiber.StatusBadRequest, err.Error())
		}
		return utility.ResponseSuccess(c, accountData.Username)
	} else { // have cache

		return utility.ResponseSuccess(c, responseGetAcccount{FirstNameTh: result.FirstNameTh, Username: result.Username})
	}
}
