package glide_test

import (
	"fmt"
	"io"
	"net/http"
	"testing"

	glide "github.com/retitle/go-sdk/v3"
	"github.com/retitle/go-sdk/v3/core"
	"github.com/retitle/go-sdk/v3/tests_utils"
	"github.com/stretchr/testify/assert"
)

func TestFoldersGetDetail(t *testing.T) {
	var (
		stringReadCloser    io.ReadCloser
		errStringReadCloser io.ReadCloser
		folderResource      glide.FoldersResource
	)
	stringReadCloser = tests_utils.ParseStructToIoReadCloser(&glide.Folder{
		Id:    "FOLDER ID",
		Kind:  "AGENT",
		Title: "THE MOST IMPORTANT FOLDER",
	})
	err := core.ErrorObject{
		Message: "ERROR SENDING EMAIL",
		Object:  "ERROR OBJECT EMAIL",
	}
	errStringReadCloser = tests_utils.ParseStructToIoReadCloser(&err)
	folderId := "23"
	folderTrxId := "10"
	url := fmt.Sprintf("https://api.glide.com/transactions/%s/folders/%s", folderTrxId, folderId)
	ttests := []tests_utils.GlideExternalApiTestCase[glide.Folder]{
		{
			Name: "Should send email",
			Arrange: func(client glide.Client) {
				folderResource = glide.GetFoldersResource(client)
			},
			Act: func(client glide.Client) (*glide.Folder, error) {
				return folderResource.GetDetail(folderTrxId, folderId)
			},
			ExpectedRequest:        tests_utils.MakeRequest(http.MethodGet, url, nil),
			MockResponse:           &http.Response{StatusCode: http.StatusOK, Body: stringReadCloser},
			ErrorInsteadOfResponse: nil,
			Assert:                 nil,
		},
		{
			Name: "Should not send email, some error happen",
			Arrange: func(client glide.Client) {
				folderResource = glide.GetFoldersResource(client)

			},
			Act: func(client glide.Client) (*glide.Folder, error) {
				return folderResource.GetDetail(folderTrxId, folderId)
			},
			ExpectedRequest:        tests_utils.MakeRequest(http.MethodGet, url, nil),
			MockResponse:           &http.Response{StatusCode: http.StatusBadRequest, Body: errStringReadCloser},
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.Folder, err error) {
				assert.NotNil(t, err)
				e := err.(*core.ApiErrorImpl)
				assert.Equal(t, e.StatusCode, http.StatusBadRequest)
				assert.Nil(t, response)
			},
		},
	}

	tests_utils.RunGlideExternalApiTestCases(t, ttests)
}

func TestFoldersGetMulti(t *testing.T) {
	var (
		stringReadCloser    io.ReadCloser
		errStringReadCloser io.ReadCloser
		folderResource      glide.FoldersResource
	)
	stringReadCloser = tests_utils.ParseStructToIoReadCloser(&glide.Folder{
		Id:    "FOLDER ID",
		Kind:  "AGENT",
		Title: "THE MOST IMPORTANT FOLDER",
	})
	err := core.ErrorObject{
		Message: "ERROR SENDING EMAIL",
		Object:  "ERROR OBJECT EMAIL",
	}
	errStringReadCloser = tests_utils.ParseStructToIoReadCloser(&err)
	folderId := "23"
	folderTrxId := "10"
	url := fmt.Sprintf("https://api.glide.com/transactions/%s/folders?ids=%s", folderTrxId, folderId)
	ttests := []tests_utils.GlideExternalApiTestCase[glide.FolderList]{
		{
			Name: "Should send email",
			Arrange: func(client glide.Client) {
				folderResource = glide.GetFoldersResource(client)
			},
			Act: func(client glide.Client) (*glide.FolderList, error) {
				return folderResource.GetMulti(folderTrxId, []string{folderId})
			},
			ExpectedRequest:        tests_utils.MakeRequest(http.MethodGet, url, nil),
			MockResponse:           &http.Response{StatusCode: http.StatusOK, Body: stringReadCloser},
			ErrorInsteadOfResponse: nil,
			Assert:                 nil,
		},
		{
			Name: "Should not send email, some error happen",
			Arrange: func(client glide.Client) {
				folderResource = glide.GetFoldersResource(client)

			},
			Act: func(client glide.Client) (*glide.FolderList, error) {
				return folderResource.GetMulti(folderTrxId, []string{folderId})
			},
			ExpectedRequest:        tests_utils.MakeRequest(http.MethodGet, url, nil),
			MockResponse:           &http.Response{StatusCode: http.StatusBadRequest, Body: errStringReadCloser},
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.FolderList, err error) {
				assert.NotNil(t, err)
				e := err.(*core.ApiErrorImpl)
				assert.Equal(t, e.StatusCode, http.StatusBadRequest)
				assert.Nil(t, response)
			},
		},
	}

	tests_utils.RunGlideExternalApiTestCases(t, ttests)
}

func TestFoldersList(t *testing.T) {
	var (
		stringReadCloser    io.ReadCloser
		errStringReadCloser io.ReadCloser
		folderResource      glide.FoldersResource
	)
	stringReadCloser = tests_utils.ParseStructToIoReadCloser(&glide.Folder{
		Id:    "FOLDER ID",
		Kind:  "AGENT",
		Title: "THE MOST IMPORTANT FOLDER",
	})
	err := core.ErrorObject{
		Message: "ERROR SENDING EMAIL",
		Object:  "ERROR OBJECT EMAIL",
	}
	errStringReadCloser = tests_utils.ParseStructToIoReadCloser(&err)
	folderTrxId := "10"
	url := fmt.Sprintf("https://api.glide.com/transactions/%s/folders", folderTrxId)
	ttests := []tests_utils.GlideExternalApiTestCase[glide.FolderList]{
		{
			Name: "Should send email",
			Arrange: func(client glide.Client) {
				folderResource = glide.GetFoldersResource(client)
			},
			Act: func(client glide.Client) (*glide.FolderList, error) {
				return folderResource.List(folderTrxId)
			},
			ExpectedRequest:        tests_utils.MakeRequest(http.MethodGet, url, nil),
			MockResponse:           &http.Response{StatusCode: http.StatusOK, Body: stringReadCloser},
			ErrorInsteadOfResponse: nil,
			Assert:                 nil,
		},
		{
			Name: "Should not send email, some error happen",
			Arrange: func(client glide.Client) {
				folderResource = glide.GetFoldersResource(client)

			},
			Act: func(client glide.Client) (*glide.FolderList, error) {
				return folderResource.List(folderTrxId)
			},
			ExpectedRequest:        tests_utils.MakeRequest(http.MethodGet, url, nil),
			MockResponse:           &http.Response{StatusCode: http.StatusBadRequest, Body: errStringReadCloser},
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.FolderList, err error) {
				assert.NotNil(t, err)
				e := err.(*core.ApiErrorImpl)
				assert.Equal(t, e.StatusCode, http.StatusBadRequest)
				assert.Nil(t, response)
			},
		},
	}

	tests_utils.RunGlideExternalApiTestCases(t, ttests)
}
