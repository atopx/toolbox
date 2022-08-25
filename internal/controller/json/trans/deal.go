package trans

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (ctl *Controller) Deal() {
	jsonStr, err := strconv.Unquote(fmt.Sprintf("\"%s\"", ctl.param.JsonStr))
	if err != nil {
		ctl.NewErrorResponse(http.StatusBadRequest, err.Error())
		return
	}
	result := make(map[string]interface{})
	if err = json.Unmarshal([]byte(jsonStr), &result); err != nil {
		ctl.NewErrorResponse(http.StatusBadRequest, err.Error())
		return
	}
	ctl.NewOkResponse(http.StatusOK, result)
}
