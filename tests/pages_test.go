package tests

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestHomePage(t *testing.T) {
// 	baseUrl := "http://127.0.0.1:3000"

// 	var (
// 		resp *http.Response
// 		err  error
// 	)

// 	resp, err = http.Get(baseUrl + "/")

// 	assert.NoError(t, err, "有错误发生")
// 	assert.Equal(t, 200, resp.StatusCode, "状态码应该为200")
// }

func TestAllPages(t *testing.T) {
	baseUrl := "http://127.0.0.1:3000"

	var tests = []struct {
		method  string
		url     string
		expeced int
	}{
		{"GET", "/", 200},
		{"GET", "/about", 200},
		{"GET", "/notfound", 200},
		{"GET", "/articles", 200},
		{"GET", "/articles/create", 200},
		{"GET", "/articles/2", 200},
		{"GET", "/articles/5/edit", 200},
		{"POST", "/articles/5", 200},
		{"POST", "/articles", 200},
		{"POST", "/articles/1/delete", 200},
	}

	for _, test := range tests {
		t.Logf("当前请求URL %v \n", test.url)
		var (
			resp *http.Response
			err  error
		)

		switch {
		case test.method == "POST":
			data := make(map[string][]string)
			resp, err = http.PostForm(baseUrl+test.url, data)
		default:
			resp, err = http.Get(baseUrl + test.url)
		}
		assert.NoError(t, err, "请求"+test.url+"时报错")
		assert.Equal(t, test.expeced, resp.StatusCode, test.url+"应返回状态码"+strconv.Itoa(test.expeced))
	}
}
