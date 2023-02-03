package dblog

import (
	"encoding/json"
	"net/http"
	"strings"
)

func (ctl *Controller) Deal() {
	dblog := make(map[string]interface{})
	if err := json.Unmarshal([]byte(ctl.param.JsonStr), &dblog); err != nil {
		ctl.NewErrorResponse(http.StatusBadRequest, err.Error())
		return
	}
	result := Reply{Message: dblog["message"].(string)}
	temp := strings.Split(result.Message, "] ")
	if len(temp) > 1 {
		result.Sql = strings.TrimSpace(temp[len(temp)-1])

	}
	ctl.NewOkResponse(http.StatusOK, result)
}
