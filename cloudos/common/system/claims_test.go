package system_test

import (
	"cloudos/common/system"
	"testing"
)

func TestUnSignClaims(t *testing.T) {
	got, err := system.UnSignClaims("Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODgxMzEyNjMsImlhdCI6MTY3NzMzMTI2MywiVXNlcklkIjo0LCJSb2xlSWQiOjF9.aY2UaHys3gTDLs8y5wWEbjswNMUQGnOi8Ivs2Wcc_Gs")
	if err != nil {
		t.Errorf("UnSignClaims() error = %v", err.Error())
		return
	}
	t.Logf("%+v", got)
}
