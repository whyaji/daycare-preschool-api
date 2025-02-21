package utils

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetPageAndLimitFromQuery(c *fiber.Ctx) (page, limit int) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}

	limit, err = strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 10
	}

	return page, limit
}
