package core_test

import (
	"strconv"
	"testing"

	"github.com/retitle/go-sdk/v5/core"
	"github.com/stretchr/testify/assert"
)

func TestPagination(t *testing.T) {
	var (
		pagination            core.PageParams
		qParams               map[string]string
		expectedLimit         int
		expectedStartingAfter string
	)

	ttests := []struct {
		name    string
		arrange func()
		act     func()
		assert  func()
	}{
		{
			name: "Should get limit and StartedAfter",
			arrange: func() {
				expectedStartingAfter = "some sub"
				expectedLimit = 50
				pagination = core.NewPageParamsWith(expectedLimit, expectedStartingAfter)

			},
			act: func() {
				qParams = pagination.GetQueryParams()
			},
			assert: func() {
				q1, ok1 := qParams["starting_after"]
				q2, ok2 := qParams["limit"]
				expectedLimitStr := strconv.Itoa(expectedLimit)
				assert.True(t, ok1)
				assert.True(t, ok2)
				assert.Equal(t, expectedStartingAfter, q1)
				assert.Equal(t, expectedLimitStr, q2)
			},
		},
	}

	for _, tt := range ttests {
		t.Run(tt.name, func(t *testing.T) {
			tt.arrange()

			tt.act()

			tt.assert()
		})
	}
}

func TestWithPageParams(t *testing.T) {
	var (
		requestOptions        *core.RequestOptionsImpl
		requestOptionFunc     core.RequestOption
		qParams               map[string]string
		expectedLimit         int
		expectedStartingAfter string
	)

	ttests := []struct {
		name    string
		arrange func()
		act     func()
		assert  func()
	}{
		{
			name: "Should set limit and StartedAfter as a query params in requestOptions",
			arrange: func() {
				expectedStartingAfter = "some sub"
				expectedLimit = 50
				requestOptions = &core.RequestOptionsImpl{}

			},
			act: func() {
				requestOptionFunc = core.WithPageParams(expectedLimit, expectedStartingAfter)
				requestOptionFunc(requestOptions)
			},
			assert: func() {
				qParams = requestOptions.GetQParams()
				q1, ok1 := qParams["starting_after"]
				q2, ok2 := qParams["limit"]
				expectedLimitStr := strconv.Itoa(expectedLimit)
				assert.True(t, ok1)
				assert.True(t, ok2)
				assert.Equal(t, expectedStartingAfter, q1)
				assert.Equal(t, expectedLimitStr, q2)
			},
		},
	}

	for _, tt := range ttests {
		t.Run(tt.name, func(t *testing.T) {
			tt.arrange()

			tt.act()

			tt.assert()
		})
	}
}

func TestWithPage(t *testing.T) {
	var (
		requestOptions    *core.RequestOptionsImpl
		requestOptionFunc core.RequestOption
		qParams           map[string]string
	)

	ttests := []struct {
		name    string
		arrange func()
		act     func()
		assert  func()
	}{
		{
			name: "Should not set values for limit and StartedAfter",
			arrange: func() {
				requestOptions = &core.RequestOptionsImpl{}

			},
			act: func() {
				requestOptionFunc = core.WithPage(nil)
				requestOptionFunc(requestOptions)
			},
			assert: func() {
				qParams = requestOptions.GetQParams()
				_, ok1 := qParams["starting_after"]
				_, ok2 := qParams["limit"]
				assert.False(t, ok1)
				assert.False(t, ok2)
			},
		},
	}

	for _, tt := range ttests {
		t.Run(tt.name, func(t *testing.T) {
			tt.arrange()

			tt.act()

			tt.assert()
		})
	}
}
