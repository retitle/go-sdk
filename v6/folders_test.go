package glide_test

import (
	"fmt"
	"net/http"
	"testing"

	glide "github.com/retitle/go-sdk/v6"
	"github.com/retitle/go-sdk/v6/core"
	"github.com/retitle/go-sdk/v6/tests_utils"
	"github.com/stretchr/testify/assert"
)

func TestFoldersGetDetail(t *testing.T) {
	var (
		folderResource glide.FoldersResource
	)
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
			ExpectedRequest: tests_utils.MakeGetRequest(url),
			MockResponse: tests_utils.Make200Response(&glide.Folder{
				Id:    "FOLDER ID",
				Kind:  "AGENT",
				Title: "THE MOST IMPORTANT FOLDER",
			}),
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
			ExpectedRequest: tests_utils.MakeGetRequest(url),
			MockResponse: tests_utils.MakeResponseWithStatus(http.StatusBadRequest, &core.ErrorObject{
				Message: "ERROR SENDING EMAIL",
				Object:  "ERROR OBJECT EMAIL",
			}),
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
		folderResource glide.FoldersResource
	)

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
			ExpectedRequest: tests_utils.MakeGetRequest(url),
			MockResponse: tests_utils.Make200Response(&glide.Folder{
				Id:    "FOLDER ID",
				Kind:  "AGENT",
				Title: "THE MOST IMPORTANT FOLDER",
			}),
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
			ExpectedRequest: tests_utils.MakeGetRequest(url),
			MockResponse: tests_utils.MakeResponseWithStatus(http.StatusBadRequest, &core.ErrorObject{
				Message: "ERROR SENDING EMAIL",
				Object:  "ERROR OBJECT EMAIL",
			}),
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
		folderResource glide.FoldersResource
	)

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
			ExpectedRequest: tests_utils.MakeGetRequest(url),
			MockResponse: tests_utils.Make200Response(&glide.Folder{
				Id:    "FOLDER ID",
				Kind:  "AGENT",
				Title: "THE MOST IMPORTANT FOLDER",
			}),
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
			ExpectedRequest: tests_utils.MakeGetRequest(url),
			MockResponse: tests_utils.MakeResponseWithStatus(http.StatusBadRequest, &core.ErrorObject{
				Message: "ERROR SENDING EMAIL",
				Object:  "ERROR OBJECT EMAIL",
			}),
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
