package middleware

import (
	"saved-pictures-holder/dto"
	"saved-pictures-holder/utils"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func AuthRequired(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		protected_query := &dto.TokenHeader{}

		if err := c.ShouldBindHeader(&protected_query); err != nil {
			c.JSON(403, gin.H{
				"message": err.Error(),
			})
			return
		}

		user, err := utils.GetUserByJWT(db, protected_query.Token)

		if err != nil {
			c.JSON(403, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.Set("user", user)

	}
}
