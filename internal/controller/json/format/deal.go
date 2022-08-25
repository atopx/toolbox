package format

import (
	"encoding/json"
	"net/http"
)

func (ctl *Controller) Deal() {
	result := make(map[string]interface{})
	if err := json.Unmarshal([]byte(ctl.param.JsonStr), &result); err != nil {
		ctl.NewErrorResponse(http.StatusBadRequest, err.Error())
		return
	}
	ctl.NewOkResponse(http.StatusOK, result)
}
