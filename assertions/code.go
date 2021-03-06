package assertions

import (
	"fmt"
	"net/http"
)

func init() {
	asserters["code"] = &Code{}
}

type Code struct{}

func (ca *Code) Assert(r *http.Response, args map[string]interface{}) error {
	c, ok := args["code"]
	if !ok {
		return fmt.Errorf("Misconfigured, code required")
	}
	code, ok := c.(int64)
	if !ok {
		return fmt.Errorf("Code must be an int %T", c)
	}
	if int64(r.StatusCode) != code {
		return fmt.Errorf("code not equal")
	}
	return nil
}
