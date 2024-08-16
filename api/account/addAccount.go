package account

import (
	"setup_go/database"
	"setup_go/logs"
	modelsPg "setup_go/models/pg"
	"setup_go/utility"

	"github.com/gofiber/fiber/v2"
)

// AddAccount
// @Summary API Login Username
// @Description  API Login Username
// @Security ApiKeyAuth
// @Tags Login
// @Accept json
// @Produce json
// @Param JSON body requestAddAccount true "ใส่ UserName & Password "
// @Success 200 {object} models.ResponseSuccess{data=modelsPg.Account}
// @Failure 400 {object} models.ResponseError400
// @Failure 401 {object} models.ResponseError401
// @Failure 500 {object} models.ResponseError
// @Router /v1/login_by_username [post]
func AddAccount(c *fiber.Ctx) error {
	logs.Info("AddAccount")
	db := database.DBConn

	reqBody := new(modelsPg.Account)
	if err := c.BodyParser(reqBody); err != nil {
		logs.Error(err)
		return utility.ResponseError(c, fiber.StatusBadRequest, err.Error())
	}

	if reqBody.Username == "" || reqBody.FirstNameEng == "" || reqBody.FirstNameTh == "" {
		return utility.ResponseError(c, fiber.StatusBadRequest, "parameter_missing")
	}

	if err := db.Where("username = ?", reqBody.Username).FirstOrCreate(&reqBody).Error; err != nil {
		logs.Error(err)
		return utility.ResponseError(c, fiber.StatusOK, "create_fail")
	}

	return utility.ResponseSuccess(c, reqBody)
}
