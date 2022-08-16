package glide

import (
	"fmt"

	"github.com/retitle/go-sdk/core"
)

type DocumentsResource interface {
	DocumentSplit(DocumentSplitSchema DocumentSplitSchema, opts ...core.RequestOption) (*DocumentSplitResponse, error)
	SignatureDetection(SignatureDetectionSchema SignatureDetectionSchema, opts ...core.RequestOption) (*SignatureDetectionResponse, error)
}

type documentsResourceImpl struct {
	client Client
}

func GetDocumentsResource(client Client) DocumentsResource {
	return documentsResourceImpl{
		client: client,
	}
}

func (r documentsResourceImpl) DocumentSplit(DocumentSplitSchema DocumentSplitSchema, opts ...core.RequestOption) (*DocumentSplitResponse, error) {
	res := DocumentSplitResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/documents/document_split"), DocumentSplitSchema, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r documentsResourceImpl) SignatureDetection(SignatureDetectionSchema SignatureDetectionSchema, opts ...core.RequestOption) (*SignatureDetectionResponse, error) {
	res := SignatureDetectionResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/documents/signature_detection"), SignatureDetectionSchema, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}
