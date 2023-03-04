package http

import (
	"log"
	"mckp/vesta/database"
	"mckp/vesta/usecases"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateCommunity(ctx *gin.Context) {
	var newCommunity usecases.NewCommunity

	if err := ctx.ShouldBindJSON(&newCommunity); err != nil {
		Error(ctx, http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	user, err := usecases.SaveNewCommunity(newCommunity)

	if err != nil {
		msg := err.Error()

		if strings.Contains(msg, "duplicate key value violates unique constraint \"communities_slug_key\"") {
			Error(ctx, http.StatusBadRequest, gin.H{
				"mesasge": "Slug already in use. Please modify your request and try again.",
			})
			return
		}

		if strings.Contains(msg, "pq: ") {
			log.Printf("[WARNING] There was an issue with the database: %s", msg)

			Error(ctx, http.StatusInternalServerError, gin.H{
				"message": "Internal Server Error. Please try again later.",
			})

			return
		}

		Error(ctx, http.StatusInternalServerError, gin.H{
			"message": msg,
		})

		return
	}

	Response(ctx, http.StatusCreated, user)
}

func CreateUser(ctx *gin.Context) {
	var newUser usecases.NewUser

	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		Error(ctx, http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	user, err := usecases.SaveNewUser(newUser)

	if err != nil {
		msg := err.Error()

		if strings.Contains(msg, "duplicate key value violates unique constraint \"users_email_key\"") {
			Error(ctx, http.StatusBadRequest, gin.H{
				"mesasge": "Email already in use. Please modify your request and try again.",
			})
			return
		}

		if strings.Contains(msg, "pq: ") {
			log.Printf("[WARNING] There was an issue with the database: %s", msg)

			Error(ctx, http.StatusInternalServerError, gin.H{
				"message": "Internal Server Error. Please try again later.",
			})

			return
		}

		Error(ctx, http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	Response(ctx, http.StatusCreated, user)

}

func Healthcheck(ctx *gin.Context) {
	healthy, err := database.IsHealthy()

	if !healthy {
		log.Printf("System not healthy due to Databse issue: %s", err)

		Error(ctx, http.StatusInternalServerError, gin.H{
			"healthy": false,
		})

		return
	}

	Response(ctx, http.StatusOK, gin.H{
		"healthy": true,
	})
}

type AddUserToCommunityRequest struct {
	UserId string `json:"user_id"`
}

func AddUserToCommunity(ctx *gin.Context) {
	var request AddUserToCommunityRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		Error(ctx, http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	communityId := ctx.Param("id")

	err := usecases.AddUserToCommunity(request.UserId, communityId)

	if err != nil {
		msg := err.Error()

		if strings.Contains(msg, "duplicate key value violates unique constraint \"user_communities_user_id_community_id_key\"") {
			Error(ctx, http.StatusBadRequest, gin.H{
				"message": "User already inside of Community. Please modify your request before trying again.",
			})

			return
		}

		if strings.Contains(msg, "pq: ") {
			log.Printf("[WARNING] There was an issue with the database: %s", msg)

			Error(ctx, http.StatusInternalServerError, gin.H{
				"message": "Internal Server Error. Please try again later.",
			})

			return
		}

		Error(ctx, http.StatusInternalServerError, gin.H{
			"message": msg,
		})

		return
	}

	Response(ctx, http.StatusCreated, gin.H{
		"complete": true,
	})
}

func GetUserCommunities(ctx *gin.Context) {
	userId := ctx.Param("id")

	communities, err := usecases.GetCommunitiesOfUser(userId)

	if err != nil {
		msg := err.Error()

		if strings.Contains(msg, "pq: ") {
			log.Printf("[WARNING] There was an issue with the database: %s", msg)

			Error(ctx, http.StatusInternalServerError, gin.H{
				"message": "Internal Server Error. Please try again later.",
			})

			return
		}

		Error(ctx, http.StatusInternalServerError, gin.H{
			"message": msg,
		})

		return
	}

	Response(ctx, http.StatusOK, communities)
}

func GetUsersInCommunity(ctx *gin.Context) {
	communityId := ctx.Param("id")

	users, err := usecases.GetUsersInCommunity(communityId)

	if err != nil {
		msg := err.Error()

		if strings.Contains(msg, "pq: ") {
			log.Printf("[WARNING] There was an issue with the database: %s", msg)

			Error(ctx, http.StatusInternalServerError, gin.H{
				"message": "Internal Server Error. Please try again later.",
			})

			return
		}

		Error(ctx, http.StatusInternalServerError, gin.H{
			"message": msg,
		})

		return
	}

	Response(ctx, http.StatusOK, users)
}

func RemoveUserFromCommunity(ctx *gin.Context) {
	communityId := ctx.Param("id")
	userId := ctx.Param("user_id")

	err := usecases.DeleteUserFromCommunity(communityId, userId)

	if err != nil {
		msg := err.Error()

		if strings.Contains(msg, "pq: ") {
			log.Printf("[WARNING] There was an issue with the database: %s", msg)

			Error(ctx, http.StatusInternalServerError, gin.H{
				"message": "Internal Server Error. Please try again later.",
			})

			return
		}

		Error(ctx, http.StatusInternalServerError, gin.H{
			"message": msg,
		})

		return
	}

	Response(ctx, http.StatusOK, gin.H{
		"removed": true,
	})
}

func GetUser(ctx *gin.Context) {
	userId := ctx.Param("id")

	user, err := usecases.GetUserById(userId)

	if err != nil {
		msg := err.Error()

		if strings.Contains(msg, "pq: ") {
			log.Printf("[WARNING] There was an issue with the database: %s", msg)

			Error(ctx, http.StatusInternalServerError, gin.H{
				"message": "Internal Server Error. Please try again later.",
			})

			return
		}

		Error(ctx, http.StatusInternalServerError, gin.H{
			"message": msg,
		})

		return
	}

	Response(ctx, http.StatusOK, user)
}

func GetCommunities(ctx *gin.Context) {
	limitArg := ctx.Query("limit")

	if limitArg == "" {
		limitArg = "100"
	}

	limit, err := strconv.ParseInt(limitArg, 10, 64)

	if err != nil {
		msg := err.Error()

		Error(ctx, http.StatusInternalServerError, gin.H{
			"message": msg,
		})

		return
	}
	offsetArg := ctx.Query("offset")

	if offsetArg == "" {
		offsetArg = "0"
	}

	offset, err := strconv.ParseInt(offsetArg, 10, 64)

	if err != nil {
		msg := err.Error()

		Error(ctx, http.StatusInternalServerError, gin.H{
			"message": msg,
		})

		return
	}

	communities, err := usecases.GetCommunities(limit, offset)

	if err != nil {
		msg := err.Error()

		Error(ctx, http.StatusInternalServerError, gin.H{
			"message": msg,
		})

		return
	}

	Response(ctx, http.StatusOK, communities)
}

func GetCommunity(ctx *gin.Context) {
	communityId := ctx.Param("id")

	community, err := usecases.GetCommunityById(communityId)

	if err != nil {
		msg := err.Error()

		if strings.Contains(msg, "pq: ") {
			log.Printf("[WARNING] There was an issue with the database: %s", msg)

			Error(ctx, http.StatusInternalServerError, gin.H{
				"message": "Internal Server Error. Please try again later.",
			})

			return
		}

		Error(ctx, http.StatusInternalServerError, gin.H{
			"message": msg,
		})

		return
	}

	Response(ctx, http.StatusOK, community)
}
