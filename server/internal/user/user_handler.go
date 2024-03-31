package user

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nhan10132020/chatapp/server/internal/validator"
)

type Handler struct {
	Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		Service: s,
	}
}

func (h *Handler) CreateUser(c *gin.Context) {
	var u CreateUserReq
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	v := validator.New()

	if ValidateCreateUser(v, &u); !v.Valid() {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": v.Errors,
		})
		return
	}

	res, err := h.Service.CreateUser(c.Request.Context(), &u)
	if err != nil {
		switch {
		case err.Error() == "pq: duplicate key value violates unique constraint \"users_email_key\"":
			v.AddError("email", "has already been registered")
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": v.Errors,
			})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) Login(c *gin.Context) {
	var user_req LoginUserReq
	if err := c.ShouldBindJSON(&user_req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	v := validator.New()

	if ValidateLogin(v, &user_req); !v.Valid() {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": v.Errors,
		})
		return
	}

	user_res, err := h.Service.Login(c.Request.Context(), &user_req)
	if err != nil {
		switch {
		case err.Error() == "invalid authentication credentials" || errors.Is(err, sql.ErrNoRows):
			v.AddError("authenticated", "Invalid email or password")
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": v.Errors,
			})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}
	}

	c.SetCookie("jwt", user_res.accessToken, 3600, "/", "localhost", false, true)

	c.JSON(http.StatusOK, user_res)
}

func (h *Handler) Logout(c *gin.Context) {
	// remove jwt in cookie
	c.SetCookie("jwt", "", -1, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "logout successful",
	})
}
