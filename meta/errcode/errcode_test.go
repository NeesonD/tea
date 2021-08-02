package errcode

import (
	"log"
	"testing"
)

func TestErrCode_String(t *testing.T) {
	log.Fatal(ERR_CODE_TIMEOUT, int(ERR_CODE_TIMEOUT))
}
