package defaultcontroller

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"net/http"

	"net/http/httptest"

	"gotest.tools/v3/assert"
)

func TestDaultController(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:50051/v1", nil)
	rec := httptest.NewRecorder()

	dftCtl := NewDefaultController()
	dftCtl.GetHome(rec, req)

	assert.Equal(t, rec.Code, http.StatusOK)

	bodyResp, err := ioutil.ReadAll(rec.Body)
	assert.NilError(t, err, "Failed to read body response.")

	var respData map[string]interface{}
	err = json.Unmarshal(bodyResp, &respData)
	assert.NilError(t, err, "Failed to decode body response.")
	assert.Equal(t, respData["success"], true)
	assert.Equal(t, respData["message"], "News Api v1")
}
