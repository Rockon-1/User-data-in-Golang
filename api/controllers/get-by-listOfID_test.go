package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestGetByListOfID(t *testing.T) {
	assert := require.New(t)
	tests := map[string]struct {
		c          int
		name       string
		id         string
		statusCode int
		response   string
	}{
		// TODO: Add test cases.
		"successullu run": {
			c:          6,
			name:       "successfully run name",
			id:         "2,2",
			statusCode: 200,
			response:   `{"ID":2,"First Name":"Sachin","City":"Uttarakhand","Phone":2345678902,"Height":5.9,"Married status":false}{"ID":2,"First Name":"Sachin","City":"Uttarakhand","Phone":2345678902,"Height":5.9,"Married status":false}`,
		},
		"Invalid ID Error": {
			c:          4,
			name:       "Invalid id error",
			id:         "0",
			statusCode: 400,
			response:   `"Please Provide Id and it should be Integer Type"`,
		},
		"Unsuccessfull run": {
			c:          5,
			name:       "unsuccessfull run",
			id:         "0",
			statusCode: 202,
			response:   `{"Id":0,"Message":"User does not exist"}`,
		},
	}

	for tCaseName, tt := range tests {
		w := httptest.NewRecorder()
		buf := new(bytes.Buffer)
		c, _ := gin.CreateTestContext(w)
		var err error
		c.Request, err = http.NewRequest("GET", "/", buf)
		assert.NoError(err, " for test Case ", tCaseName, "Test request should be created")

		q := c.Request.URL.Query()
		q.Add("id", tt.id)
		if tt.c == 4 {
			q.Del("id")
			q.Add("id", "invalid id")
		}
		c.Request.URL.RawQuery = q.Encode()
		var obj Interface = newMockControllerDescriber(tt.c)
		GetByListOfID(obj)(c)

		assert.Equal(tt.statusCode, w.Code)
		assert.Equal(tt.response, w.Body.String())
	}
}
