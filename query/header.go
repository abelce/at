package query

import (
	"net/http"
	"strconv"
	"strings"
)

// 获取operator id
func GetOperatorID(r *http.Request) string {
	return r.Header.Get("operatorID")
}

func GetQueryValues(r *http.Request) (filter, sort string, offset, limit uint64, err error) {
	ps := r.URL.Query()
	filter = ps.Get("filter")
	sort = ps.Get("sort")
	offsetStr := ps.Get("page[offset]")
	limitStr := ps.Get("page[limit]")
	limit = 10
	if strings.TrimSpace(offsetStr) != "" {
		offset, err = strconv.ParseUint(offsetStr, 10, 32)
		if err != nil {
			return
		}
	}
	if strings.TrimSpace(limitStr) != "" {
		limit, err = strconv.ParseUint(limitStr, 10, 32)
		if err != nil {
			return
		}
	}
	// 如果有operatorID就加上
	operatorID := GetOperatorID(r)
	if operatorID != "" {
		if filter == "" {
			filter = "operatorID eq '" + operatorID + "'"
		} else {
			filter = filter + " and operatorID eq '" + operatorID + "'"
		}
	}
	return
}
