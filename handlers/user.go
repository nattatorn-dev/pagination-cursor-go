// handlers/user.go
package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nattatorn-dev/pagination-cursor-go/ent"
	"github.com/nattatorn-dev/pagination-cursor-go/utils"
)

// GetUsers handles the request to retrieve users with cursor-based pagination and dynamic sorting.
func GetUsers(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		cursor := c.Query("cursor")
		limitStr := c.Query("limit")
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			limit = 10 // Default limit
		}

		sortFields := utils.ParseSortFields(c)
		query := client.User.Query()

		if cursor != "" {
			decodedCursor, err := utils.DecodeCursor(cursor)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid cursor"})
				return
			}
			for k := range decodedCursor {
				sortFields[k] = "asc"
			}
			query = utils.BuildCursorQuery(query, decodedCursor, sortFields)
		}

		query = utils.BuildOrderQuery(query, sortFields)

		users, err := query.Limit(limit).All(context.Background())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var nextCursor string
		if len(users) > 0 {
			cursorMap := make(map[string]interface{})
			lastUser := users[len(users)-1]
			for field := range sortFields {
				switch field {
				case "id":
					cursorMap["id"] = lastUser.ID
				case "salary":
					cursorMap["salary"] = lastUser.Salary
				}
			}
			nextCursor = utils.EncodeCursor(cursorMap)
		}

		c.JSON(http.StatusOK, gin.H{
			"data":        users,
			"next_cursor": nextCursor,
		})
	}
}

func parseLimit(c *gin.Context) int {
	limitStr := c.Query("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10 // Default limit
	}
	return limit
}

func parseCursor(c *gin.Context) (map[string]string, error) {
	cursor := c.Query("cursor")
	if cursor == "" {
		return nil, nil
	}

	decodedCursor, err := utils.DecodeCursor(cursor)
	if err != nil {
		return nil, err
	}
	return decodedCursor, nil
}

func generateNextCursor(users []*ent.User, sortFields map[string]string) string {
	if len(users) == 0 {
		return ""
	}

	lastUser := users[len(users)-1]
	cursorMap := make(map[string]interface{})

	for field := range sortFields {
		switch field {
		case "id":
			cursorMap["id"] = lastUser.ID
		case "salary":
			cursorMap["salary"] = lastUser.Salary
		}
	}

	return utils.EncodeCursor(cursorMap)
}
