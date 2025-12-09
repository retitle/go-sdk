package glide

import (
	"fmt"

	"github.com/retitle/go-sdk/v6/core"
)

type DocumentsResource interface {
	AnalyzeSynchronous(analyzeSchema AnalyzeSchema, files []core.File, opts ...core.RequestOption) (*DocumentAnalysisAsyncResponse, error)
	DocumentSplit(documentSplitSchema DocumentSplitSchema, files []core.File, opts ...core.RequestOption) (*DocumentSplitResponse, error)
	DownloadDetachedZip(opts ...core.RequestOption) (*BinaryResponse, error)
	DownloadZip(opts ...core.RequestOption) (*BinaryResponse, error)
	Duplicate(documentDuplicateSchema DocumentDuplicateSchema, opts ...core.RequestOption) (*DocumentDuplicateResponse, error)
	PspdfkitDetails(opts ...core.RequestOption) (*DocumentPspdfkitDetailsResponse, error)
	SignatureDetection(signatureDetectionSchema SignatureDetectionSchema, files []core.File, opts ...core.RequestOption) (*SignatureDetectionResponse, error)
	UploadFile(documentUploadSchema DocumentUploadSchema, files []core.File, opts ...core.RequestOption) (*DocumentUploadResponse, error)
	Download(id string, opts ...core.RequestOption) (*BinaryResponse, error)
	DownloadDetached(id string, opts ...core.RequestOption) (*BinaryResponse, error)
	GetApplicableTemplates(id string, opts ...core.RequestOption) (*ApplicableTemplatesResponse, error)
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

func (r documentsResourceImpl) DownloadDetachedZip(opts ...core.RequestOption) (*BinaryResponse, error) {
	res := BinaryResponse{}
	if err := r.client.GetStream(&res, true, fmt.Sprintf("/documents/download_detached_zip"), opts...); err != nil {
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

func (r documentsResourceImpl) Duplicate(documentDuplicateSchema DocumentDuplicateSchema, opts ...core.RequestOption) (*DocumentDuplicateResponse, error) {
	res := DocumentDuplicateResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/documents/duplicate"), documentDuplicateSchema, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r documentsResourceImpl) PspdfkitDetails(opts ...core.RequestOption) (*DocumentPspdfkitDetailsResponse, error) {
	res := DocumentPspdfkitDetailsResponse{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/documents/pspdfkit_details"), opts...); err != nil {
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

func (r documentsResourceImpl) UploadFile(documentUploadSchema DocumentUploadSchema, files []core.File, opts ...core.RequestOption) (*DocumentUploadResponse, error) {
	res := DocumentUploadResponse{}
	if err := r.client.PostWithFiles(&res, true, fmt.Sprintf("/documents/upload_file"), documentUploadSchema, files, opts...); err != nil {
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

func (r documentsResourceImpl) DownloadDetached(id string, opts ...core.RequestOption) (*BinaryResponse, error) {
	res := BinaryResponse{}
	if err := r.client.GetStream(&res, true, fmt.Sprintf("/documents/%s/download_detached", id), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r documentsResourceImpl) GetApplicableTemplates(id string, opts ...core.RequestOption) (*ApplicableTemplatesResponse, error) {
	res := ApplicableTemplatesResponse{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/documents/templates/%s", id), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}
