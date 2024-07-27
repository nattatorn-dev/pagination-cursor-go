// utils/query.go
package utils

import (
	"strconv"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/gin-gonic/gin"
	"github.com/nattatorn-dev/pagination-cursor-go/ent"
	"github.com/nattatorn-dev/pagination-cursor-go/ent/user"
)

func BuildOrderQuery(query *ent.UserQuery, sortFields map[string]string) *ent.UserQuery {
	for field, order := range sortFields {
		switch field {
		case "id":
			if order == "asc" {
				query = query.Order(ent.Asc(user.FieldID))
			} else {
				query = query.Order(ent.Desc(user.FieldID))
			}
		case "salary":
			if order == "asc" {
				query = query.Order(ent.Asc(user.FieldSalary))
			} else {
				query = query.Order(ent.Desc(user.FieldSalary))
			}
		}
	}
	return query
}

func BuildCursorQuery(query *ent.UserQuery, cursorMap map[string]string, sortFields map[string]string) *ent.UserQuery {
	query = query.Where(func(s *sql.Selector) {
		conditions := make([]*sql.Predicate, 0, len(sortFields))
		for field := range sortFields {
			value := cursorMap[field]
			switch field {
			case "id":
				id, _ := strconv.Atoi(value)
				conditions = append(conditions, sql.GT(s.C(user.FieldID), id))
			case "salary":
				salary, _ := strconv.ParseFloat(value, 64)
				conditions = append(conditions, sql.GT(s.C(user.FieldSalary), salary))
			}
		}

		s.Where(sql.And(conditions...))
	})
	return query
}

func ParseSortFields(c *gin.Context) map[string]string {
	sortFields := c.Query("sort")
	if sortFields == "" {
		return map[string]string{"id": "asc"}
	}
	fieldList := strings.Split(sortFields, ",")
	sortMap := make(map[string]string)
	for _, field := range fieldList {
		sortMap[field] = "asc"
	}
	return sortMap
}
