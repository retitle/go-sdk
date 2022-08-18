package glide

import (
	"fmt"
)

type DocumentsResource interface {
	DocumentSplit(documentSplitSchema DocumentSplitSchema, opts ...requestOption) (*DocumentSplitResponse, error)
	SignatureDetection(signatureDetectionSchema SignatureDetectionSchema, opts ...requestOption) (*SignatureDetectionResponse, error)
}

type documentsResourceImpl struct {
	client Client
}

func getDocumentsResource(client Client) DocumentsResource {
	return documentsResourceImpl{
		client: client,
	}
}

func (r documentsResourceImpl) DocumentSplit(documentSplitSchema DocumentSplitSchema, opts ...requestOption) (*DocumentSplitResponse, error) {
	res := DocumentSplitResponse{}
	if err := r.client.post(&res, true, fmt.Sprintf("/documents/document_split"), documentSplitSchema, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r documentsResourceImpl) SignatureDetection(signatureDetectionSchema SignatureDetectionSchema, opts ...requestOption) (*SignatureDetectionResponse, error) {
	res := SignatureDetectionResponse{}
	if err := r.client.post(&res, true, fmt.Sprintf("/documents/signature_detection"), signatureDetectionSchema, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}
