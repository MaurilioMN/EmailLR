package endpoint

import (
	internalerrors "campaing/internal/internalErrors"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HendlerError_endpoit_InternalError(t *testing.T) {
	assert := assert.New(t)

	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return nil, 0, internalerrors.Errinternal
	}
	handlerFunc := HandlerError(endpoint)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusInternalServerError, res.Code)
	assert.Contains(res.Body.String(), internalerrors.Errinternal.Error())

}

func Test_HendlerError_endpoit_DomaninError(t *testing.T) {

	assert := assert.New(t)

	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return nil, 0, errors.New("Domain error")
	}
	handlerFunc := HandlerError(endpoint)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusBadRequest, res.Code)
	assert.Contains(res.Body.String(), "Domain error")

}

func Test_HendlerError_endpoit_Status_Obj(t *testing.T) {

	assert := assert.New(t)

	type bodyforTest struct {
		Id int
	}

	objExpectd := bodyforTest{Id: 2}

	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return objExpectd, 201, nil
	}
	handlerFunc := HandlerError(endpoint)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusCreated, res.Code)
	objReturned := bodyforTest{}
	json.Unmarshal(res.Body.Bytes(), &objReturned)
	assert.Equal(objExpectd, objReturned)

}
