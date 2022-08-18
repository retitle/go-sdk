package glide

import (
	"fmt"

	"github.com/retitle/go-sdk/v3/core"
)

type DocumentsResource interface {
	DocumentSplit(documentsplitschema DocumentSplitSchema, opts ...core.RequestOption) (*DocumentSplitResponse, error)
	SignatureDetection(signaturedetectionschema SignatureDetectionSchema, opts ...core.RequestOption) (*SignatureDetectionResponse, error)
}

type documentsResourceImpl struct {
	client Client
}

func GetDocumentsResource(client Client) DocumentsResource {
	return documentsResourceImpl{
		client: client,
	}
}

func (r documentsResourceImpl) DocumentSplit(documentsplitschema DocumentSplitSchema, opts ...core.RequestOption) (*DocumentSplitResponse, error) {
	res := DocumentSplitResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/documents/document_split"), documentsplitschema, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r documentsResourceImpl) SignatureDetection(signaturedetectionschema SignatureDetectionSchema, opts ...core.RequestOption) (*SignatureDetectionResponse, error) {
	res := SignatureDetectionResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/documents/signature_detection"), signaturedetectionschema, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}
