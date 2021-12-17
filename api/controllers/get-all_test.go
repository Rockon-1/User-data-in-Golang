package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestGetall(t *testing.T) {
	assert := require.New(t)
	tests := map[string]struct {
		c          int
		name       string
		statusCode int
		response   string
	}{
		// TODO: Add test cases.
		"successullu run": {
			c:          1,
			name:       "successfully run name",
			statusCode: 200,
			response:   `[{"ID":1,"First Name":"steve","City":"LA","Phone":1234567890,"Height":5.8,"Married status":true},{"ID":2,"First Name":"Sachin","City":"Uttarakhand","Phone":2345678902,"Height":5.9,"Married status":false},{"ID":3,"First Name":"Virat","City":"Mumbai","Phone":7778889901,"Height":8.4,"Married status":true}]`,
		},
		"Unsuccessfull run": {
			c:          2,
			name:       "unsuccessfull run",
			statusCode: 500,
			response:   `"Some Issue Happened !!! Sorry for Inconvenigence"`,
		},
	}

	for tCaseName, tt := range tests {
		w := httptest.NewRecorder()
		buf := new(bytes.Buffer)
		c, _ := gin.CreateTestContext(w)
		var err error
		c.Request, err = http.NewRequest("GET", "/", buf)
		assert.NoError(err, " for test Case ", tCaseName, "Test request should be created")

		// t.Run(tt.name, func(t *testing.T) {
		var obj Interface = newMockControllerDescriber(tt.c)
		Getall(obj)(c)
		//	})
		assert.Equal(tt.statusCode, w.Code)
		assert.Equal(tt.response, w.Body.String())
	}
}
