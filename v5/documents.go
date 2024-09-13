package glide

import (
	"fmt"

	"github.com/retitle/go-sdk/v5/core"
)

type DocumentsResource interface {
	AnalyzeSynchronous(analyzeSchema AnalyzeSchema, files []core.File, opts ...core.RequestOption) (*DocumentAnalysisAsyncResponse, error)
	DocumentSplit(documentSplitSchema DocumentSplitSchema, files []core.File, opts ...core.RequestOption) (*DocumentSplitResponse, error)
	DownloadZip(opts ...core.RequestOption) (*BinaryResponse, error)
	SignatureDetection(signatureDetectionSchema SignatureDetectionSchema, files []core.File, opts ...core.RequestOption) (*SignatureDetectionResponse, error)
	Download(id string, opts ...core.RequestOption) (*BinaryResponse, error)
}

type documentsResourceImpl struct {
	client Client
}

func GetDocumentsResource(client Client) DocumentsResource {
	return documentsResourceImpl{
		client: client,
	}
}

func (r documentsResourceImpl) AnalyzeSynchronous(analyzeSchema AnalyzeSchema, files []core.File, opts ...core.RequestOption) (*DocumentAnalysisAsyncResponse, error) {
	res := DocumentAnalysisAsyncResponse{}
	if err := r.client.PostWithFiles(&res, true, fmt.Sprintf("/documents/analyze_synchronous"), analyzeSchema, files, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r documentsResourceImpl) DocumentSplit(documentSplitSchema DocumentSplitSchema, files []core.File, opts ...core.RequestOption) (*DocumentSplitResponse, error) {
	res := DocumentSplitResponse{}
	if err := r.client.PostWithFiles(&res, true, fmt.Sprintf("/documents/document_split"), documentSplitSchema, files, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r documentsResourceImpl) DownloadZip(opts ...core.RequestOption) (*BinaryResponse, error) {
	res := BinaryResponse{}
	if err := r.client.GetStream(&res, true, fmt.Sprintf("/documents/download_zip"), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r documentsResourceImpl) SignatureDetection(signatureDetectionSchema SignatureDetectionSchema, files []core.File, opts ...core.RequestOption) (*SignatureDetectionResponse, error) {
	res := SignatureDetectionResponse{}
	if err := r.client.PostWithFiles(&res, true, fmt.Sprintf("/documents/signature_detection"), signatureDetectionSchema, files, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r documentsResourceImpl) Download(id string, opts ...core.RequestOption) (*BinaryResponse, error) {
	res := BinaryResponse{}
	if err := r.client.GetStream(&res, true, fmt.Sprintf("/documents/%s/download", id), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}
