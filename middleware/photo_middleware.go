package middleware

import (
	"net/http"
	"strconv"

	"final-project/model/domain"
	"final-project/model/http/response"
	"final-project/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func PhotoMiddleware(photoService service.PhotoService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			photo domain.Photo
			err   error
		)

		photoID, _ := strconv.Atoi(ctx.Param("id"))
		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))

		if photo, err = photoService.GetOne(uint(photoID)); err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, response.ErrorResponse{
				Code:   http.StatusBadRequest,
				Status: "Bad Request",
				Errors: gin.H{
					"message": "Photo not found",
				},
			})

			return
		}

		if photo.UserID != userID {
			ctx.AbortWithStatusJSON(http.StatusForbidden, response.ErrorResponse{
				Code:   http.StatusForbidden,
				Status: "Forbidden",
				Errors: gin.H{
					"message": "You don't have permission",
				},
			})

			return
		}
	}
}
