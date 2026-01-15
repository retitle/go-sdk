package glide

import (
	"bytes"
	"io"
	"net/http"
	"strings"

	"github.com/retitle/go-sdk/v6/core"
)

type AccessControl struct {
	Id         string `json:"id,omitempty"`
	Deleted    int    `json:"deleted,omitempty"`
	ObjectId   string `json:"object_id,omitempty"`
	ObjectType string `json:"object_type,omitempty"`
	Policy     string `json:"policy,omitempty"`
	RoleId     string `json:"role_id,omitempty"`
	RoleType   string `json:"role_type,omitempty"`
}

type AccessPolicy struct {
	Acl []*AccessControl `json:"acl,omitempty"`
}

type Activity struct {
	Id        string           `json:"id,omitempty"`
	Context   *ActivityContext `json:"context,omitempty"`
	CreatedAt int              `json:"created_at,omitempty"`
	Kind      string           `json:"kind,omitempty"`
	Object    string           `json:"object,omitempty"`
}

func (m Activity) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type ActivityContext struct {
	Documents  []*ActivityContextItem `json:"documents,omitempty"`
	Party      *ActivityContextItem   `json:"party,omitempty"`
	Recipients []*ActivityContextItem `json:"recipients,omitempty"`
	User       *ActivityContextItem   `json:"user,omitempty"`
	Object     string                 `json:"object,omitempty"`
}

func (m ActivityContext) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type ActivityContextItem struct {
	Id     string `json:"id,omitempty"`
	Link   string `json:"link,omitempty"`
	Name   string `json:"name,omitempty"`
	Object string `json:"object,omitempty"`
}

func (m ActivityContextItem) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type ActivityListWithCursor struct {
	Cursor  string      `json:"cursor,omitempty"`
	Data    []*Activity `json:"data,omitempty"`
	HasMore *bool       `json:"has_more,omitempty"`
	Total   int         `json:"total,omitempty"`
	Object  string      `json:"object,omitempty"`
}

func (m ActivityListWithCursor) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type Address struct {
	City    string `json:"city"`
	County  string `json:"county,omitempty"`
	State   string `json:"state"`
	Street  string `json:"street"`
	Unit    string `json:"unit,omitempty"`
	ZipCode string `json:"zip_code"`
	Object  string `json:"object,omitempty"`
}

func (m Address) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type Agent struct {
	Address              *Address `json:"address,omitempty"`
	CompanyLicenseNumber string   `json:"company_license_number,omitempty"`
	CompanyName          string   `json:"company_name,omitempty"`
	CompanyPhoneNumber   string   `json:"company_phone_number,omitempty"`
	LicenseNumber        string   `json:"license_number,omitempty"`
	LicenseState         string   `json:"license_state,omitempty"`
	NrdsNumber           string   `json:"nrds_number,omitempty"`
	Object               string   `json:"object,omitempty"`
}

func (m Agent) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type AgentRequest struct {
	Address              *Address `json:"address,omitempty"`
	AddressId            string   `json:"address_id,omitempty"`
	CompanyLicenseNumber string   `json:"company_license_number,omitempty"`
	CompanyName          string   `json:"company_name,omitempty"`
	CompanyPhoneNumber   string   `json:"company_phone_number,omitempty"`
	LicenseNumber        string   `json:"license_number,omitempty"`
	LicenseState         string   `json:"license_state,omitempty"`
	NrdsNumber           string   `json:"nrds_number,omitempty"`
}

type AnalyzeSchema struct {
	Files     []http.File `json:"files,omitempty"`
	FilesMeta []*FileMeta `json:"files_meta,omitempty"`
}

type ApplyTemplatesResponse struct {
	IsDelayed     *bool                              `json:"is_delayed,omitempty"`
	Result        *TransactionAppliedTemplatesResult `json:"result,omitempty"`
	TransactionId string                             `json:"transaction_id,omitempty"`
	Object        string                             `json:"object,omitempty"`
}

func (m ApplyTemplatesResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type Annotations struct {
	Id             string  `json:"id,omitempty"`
	FieldId        string  `json:"field_id,omitempty"`
	Height         float64 `json:"height,omitempty"`
	Kind           string  `json:"kind,omitempty"`
	Left           float64 `json:"left,omitempty"`
	RecipientColor string  `json:"recipient_color,omitempty"`
	RecipientRole  string  `json:"recipient_role,omitempty"`
	Source         string  `json:"source,omitempty"`
	Top            float64 `json:"top,omitempty"`
	Width          float64 `json:"width,omitempty"`
}

type BinaryResponse struct {
	ContentDisposition string       `json:"content_disposition,omitempty"`
	ContentType        string       `json:"content_type,omitempty"`
	Data               bytes.Buffer `json:"data,omitempty"`
	Object             string       `json:"object,omitempty"`
}

func (m *BinaryResponse) SetData(dataSource io.Reader, metadata core.BinaryMetadata) error {
	m.ContentType = metadata.ContentType
	m.ContentDisposition = metadata.ContentDisposition
	_, err := io.Copy(&m.Data, dataSource)
	return err
}

func (m BinaryResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type ConditionalLinking struct {
	LinkId        string `json:"link_id,omitempty"`
	LinkNamespace string `json:"link_namespace,omitempty"`
	Rule          string `json:"rule,omitempty"`
	Terms         string `json:"terms,omitempty"`
}

type Contact struct {
	Id              string         `json:"id,omitempty"`
	Address         *Address       `json:"address,omitempty"`
	Agent           *Agent         `json:"agent,omitempty"`
	AvatarUrl       string         `json:"avatar_url,omitempty"`
	BrandLogoUrl    string         `json:"brand_logo_url,omitempty"`
	CellPhone       string         `json:"cell_phone,omitempty"`
	ContactSource   *ContactSource `json:"contact_source,omitempty"`
	Email           string         `json:"email,omitempty"`
	EntityName      string         `json:"entity_name,omitempty"`
	EntityType      string         `json:"entity_type,omitempty"`
	FaxPhone        string         `json:"fax_phone,omitempty"`
	FirstName       string         `json:"first_name,omitempty"`
	LastName        string         `json:"last_name,omitempty"`
	PersonalWebsite string         `json:"personal_website,omitempty"`
	TeamId          string         `json:"team_id,omitempty"`
	Title           string         `json:"title,omitempty"`
	Object          string         `json:"object,omitempty"`
}

func (m Contact) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type ContactList struct {
	Data       []Contact `json:"data"`
	ListObject string    `json:"list_object"`
	Object     string    `json:"object"`
	HasMore    bool      `json:"has_more"`
}

func (m ContactList) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

func (m ContactList) NextPageParams() core.PageParams {
	if !m.HasMore {
		return nil
	}

	pageSize := len(m.Data)
	return &core.PageParamsImpl{
		StartingAfter: m.Data[pageSize-1].Id,
		Limit:         pageSize,
	}
}

type Contact1 struct {
	Email     string `json:"email,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Object    string `json:"object,omitempty"`
}

func (m Contact1) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type ContactCreate struct {
	Contact *ContactRequest `json:"contact"`
}

type ContactCreateResponse struct {
	Contact *Contact `json:"contact,omitempty"`
	Object  string   `json:"object,omitempty"`
}

func (m ContactCreateResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type ContactRequest struct {
	Address         *Address       `json:"address,omitempty"`
	AddressId       string         `json:"address_id,omitempty"`
	Agent           *AgentRequest  `json:"agent,omitempty"`
	AvatarUrl       string         `json:"avatar_url,omitempty"`
	BrandLogoUrl    string         `json:"brand_logo_url,omitempty"`
	CellPhone       string         `json:"cell_phone,omitempty"`
	ContactSource   *ContactSource `json:"contact_source,omitempty"`
	Email           string         `json:"email,omitempty"`
	EntityName      string         `json:"entity_name,omitempty"`
	EntityType      string         `json:"entity_type,omitempty"`
	FaxPhone        string         `json:"fax_phone,omitempty"`
	FirstName       string         `json:"first_name,omitempty"`
	LastName        string         `json:"last_name,omitempty"`
	PersonalWebsite string         `json:"personal_website,omitempty"`
	Title           string         `json:"title,omitempty"`
}

type ContactSource struct {
	Id     string `json:"id,omitempty"`
	Origin string `json:"origin,omitempty"`
	Object string `json:"object,omitempty"`
}

func (m ContactSource) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type ContactSourceRequest struct {
	Id     string `json:"id,omitempty"`
	Origin string `json:"origin,omitempty"`
}

type ContactUpdate struct {
	Contact *ContactRequest `json:"contact,omitempty"`
	Roles   []string        `json:"roles,omitempty"`
}

type ContactUpdateResponse struct {
	Contact *Contact `json:"contact,omitempty"`
	Id      string   `json:"id_,omitempty"`
	Object  string   `json:"object,omitempty"`
}

func (m ContactUpdateResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type CreateResponse struct {
	TransactionId string `json:"transaction_id,omitempty"`
	Object        string `json:"object,omitempty"`
}

func (m CreateResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type DealTemplate struct {
	Tags              []string             `json:"tags,omitempty"`
	TemplateDocuments []*TemplateDocuments `json:"template_documents,omitempty"`
	TemplateId        string               `json:"template_id,omitempty"`
	TemplateTitle     string               `json:"template_title,omitempty"`
	Object            string               `json:"object,omitempty"`
}

func (m DealTemplate) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type DealTemplatesResponse struct {
	DealTemplates []*DealTemplate `json:"deal_templates,omitempty"`
	Object        string          `json:"object,omitempty"`
}

func (m DealTemplatesResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type DeletedParties struct {
	Data   []*DeletedParty `json:"data,omitempty"`
	Object string          `json:"object,omitempty"`
}

func (m DeletedParties) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type DeletedParty struct {
	Contact   *Contact `json:"contact,omitempty"`
	DeletedAt int      `json:"deleted_at,omitempty"`
	PartyId   string   `json:"party_id,omitempty"`
	Roles     []string `json:"roles,omitempty"`
	Object    string   `json:"object,omitempty"`
}

func (m DeletedParty) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type Document struct {
	Id                 string `json:"id"`
	FileName           string `json:"file_name"`
	Object             string `json:"object,omitempty"`
	PspdfkitDocumentId string `json:"pspdfkit_document_id,omitempty"`
}

func (m Document) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type DocumentAnalysisAsyncResponse struct {
	ResultsByFileReferenceId map[string]*DocumentAnalysisResult `json:"results_by_file_reference_id,omitempty"`
	Object                   string                             `json:"object,omitempty"`
}

func (m DocumentAnalysisAsyncResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type DocumentAnalysisResult struct {
	DatapointExtractionSucceeded *bool        `json:"datapoint_extraction_succeeded"`
	Error                        string       `json:"error,omitempty"`
	FormMatches                  []*FormMatch `json:"form_matches,omitempty"`
	FormMatchingSucceeded        *bool        `json:"form_matching_succeeded"`
	SignatureDetectionSucceeded  *bool        `json:"signature_detection_succeeded"`
	Status                       string       `json:"status,omitempty"`
	Object                       string       `json:"object,omitempty"`
}

func (m DocumentAnalysisResult) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type DocumentDuplicateResponse struct {
	Documents []*Document `json:"documents,omitempty"`
	Object    string      `json:"object,omitempty"`
}

func (m DocumentDuplicateResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type DocumentDuplicateSchema struct {
	CopyAnalysis     *bool             `json:"copy_analysis,omitempty"`
	CopyPspdfkitData *bool             `json:"copy_pspdfkit_data,omitempty"`
	CreatePspdfkit   *bool             `json:"create_pspdfkit,omitempty"`
	DocumentUuids    []string          `json:"document_uuids"`
	RecipientIdMap   map[string]string `json:"recipient_id_map,omitempty"`
}

type DocumentFileMeta struct {
	FileName string `json:"file_name"`
}

type DocumentMergeSchema struct {
	DeleteOriginalDocuments       *bool    `json:"delete_original_documents,omitempty"`
	IsAsync                       *bool    `json:"is_async,omitempty"`
	NewDocumentFolderId           string   `json:"new_document_folder_id"`
	NewDocumentTitle              string   `json:"new_document_title"`
	TransactionDocumentVersionIds []string `json:"transaction_document_version_ids,omitempty"`
}

type DocumentPspdfkitDetailsResponse struct {
	JwtToken           string `json:"jwt_token,omitempty"`
	PspdfkitDocumentId string `json:"pspdfkit_document_id,omitempty"`
	Object             string `json:"object,omitempty"`
}

func (m DocumentPspdfkitDetailsResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type DocumentSplitAsyncResponse struct {
	ReqId       string                              `json:"req_id,omitempty"`
	Suggestions map[string]*DocumentSplitSuggestion `json:"suggestions,omitempty"`
	Object      string                              `json:"object,omitempty"`
}

func (m DocumentSplitAsyncResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type DocumentSplitResponse struct {
	ReqId  string                      `json:"req_id,omitempty"`
	Result *DocumentSplitAsyncResponse `json:"result,omitempty"`
	Object string                      `json:"object,omitempty"`
}

func (m DocumentSplitResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type DocumentSplitSchema struct {
	Files   []http.File       `json:"files,omitempty"`
	ReState string            `json:"re_state,omitempty"`
	ReqId   string            `json:"req_id"`
	Uploads []*DocumentUpload `json:"uploads,omitempty"`
}

type DocumentSplitSuggestion struct {
	EndPage      int    `json:"end_page,omitempty"`
	Filename     string `json:"filename,omitempty"`
	FormId       string `json:"form_id,omitempty"`
	FormSeriesId string `json:"form_series_id,omitempty"`
	StartPage    int    `json:"start_page,omitempty"`
	Object       string `json:"object,omitempty"`
}

func (m DocumentSplitSuggestion) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type DocumentUpload struct {
	Title string `json:"title,omitempty"`
}

type DocumentUploadResponse struct {
	Documents []*Document    `json:"documents,omitempty"`
	Errors    []*UploadError `json:"errors,omitempty"`
	Object    string         `json:"object,omitempty"`
}

func (m DocumentUploadResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type DocumentUploadSchema struct {
	Files          []http.File         `json:"files,omitempty"`
	FilesMeta      []*DocumentFileMeta `json:"files_meta,omitempty"`
	CreatePspdfkit bool                `json:"create_pspdfkit,omitempty"`
}

type DocumentZone struct {
	Id               string                  `json:"id,omitempty"`
	FormId           string                  `json:"form_id,omitempty"`
	Kind             string                  `json:"kind,omitempty"`
	Name             string                  `json:"name,omitempty"`
	OriginalLocation []*DocumentZoneLocation `json:"original_location,omitempty"`
	Page             int                     `json:"page,omitempty"`
	Vertices         []*DocumentZoneVertex   `json:"vertices,omitempty"`
	Object           string                  `json:"object,omitempty"`
}

func (m DocumentZone) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type DocumentZoneLocation struct {
	XMax   float64 `json:"x_max,omitempty"`
	XMin   float64 `json:"x_min,omitempty"`
	YMax   float64 `json:"y_max,omitempty"`
	YMin   float64 `json:"y_min,omitempty"`
	Object string  `json:"object,omitempty"`
}

func (m DocumentZoneLocation) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type DocumentZoneVertex struct {
	X      int    `json:"x,omitempty"`
	Y      int    `json:"y,omitempty"`
	Object string `json:"object,omitempty"`
}

func (m DocumentZoneVertex) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type ESignAnnotation struct {
	Id           string       `json:"id,omitempty"`
	DefaultValue string       `json:"default_value,omitempty"`
	FieldId      string       `json:"field_id,omitempty"`
	FieldParams  *FieldParams `json:"field_params,omitempty"`
	FieldPart    int          `json:"field_part,omitempty"`
	Height       float64      `json:"height,omitempty"`
	IsRequired   *bool        `json:"is_required,omitempty"`
	Kind         string       `json:"kind,omitempty"`
	Left         float64      `json:"left,omitempty"`
	Optional     *bool        `json:"optional,omitempty"`
	OwnerId      string       `json:"owner_id,omitempty"`
	Page         int          `json:"page,omitempty"`
	Top          float64      `json:"top,omitempty"`
	Width        float64      `json:"width,omitempty"`
	Object       string       `json:"object,omitempty"`
}

func (m ESignAnnotation) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type ESignFillConfig struct {
	Annotations []*ESignAnnotation `json:"annotations,omitempty"`
	TdvId       string             `json:"tdv_id,omitempty"`
	Object      string             `json:"object,omitempty"`
}

func (m ESignFillConfig) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type Envelope struct {
	Id                string                          `json:"id,omitempty"`
	Activities        *EnvelopeActivityListWithCursor `json:"activities,omitempty"`
	CertificateDocId  string                          `json:"certificate_doc_id,omitempty"`
	CreatedAt         int                             `json:"created_at,omitempty"`
	Creator           string                          `json:"creator,omitempty"`
	EnvelopeDocument  *EnvelopeDocumentList           `json:"envelope_document,omitempty"`
	EnvelopeRecipient *EnvelopeRecipientList          `json:"envelope_recipient,omitempty"`
	SigningUrl        string                          `json:"signing_url,omitempty"`
	Status            string                          `json:"status,omitempty"`
	Step              *StepList                       `json:"step,omitempty"`
	Title             string                          `json:"title,omitempty"`
	Uuid              string                          `json:"uuid,omitempty"`
	Object            string                          `json:"object,omitempty"`
}

func (m Envelope) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type EnvelopeActivity struct {
	Id        string                   `json:"id,omitempty"`
	Context   *EnvelopeActivityContext `json:"context,omitempty"`
	CreatedAt int                      `json:"created_at,omitempty"`
	Kind      string                   `json:"kind,omitempty"`
	Object    string                   `json:"object,omitempty"`
}

func (m EnvelopeActivity) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type EnvelopeActivityContext struct {
	Recipients []*EnvelopeActivityContextItem `json:"recipients,omitempty"`
	User       *EnvelopeActivityContextItem   `json:"user,omitempty"`
	Object     string                         `json:"object,omitempty"`
}

func (m EnvelopeActivityContext) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type EnvelopeActivityContextItem struct {
	Id     string `json:"id,omitempty"`
	Link   string `json:"link,omitempty"`
	Name   string `json:"name,omitempty"`
	Object string `json:"object,omitempty"`
}

func (m EnvelopeActivityContextItem) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type EnvelopeActivityListWithCursor struct {
	Cursor  string              `json:"cursor,omitempty"`
	Data    []*EnvelopeActivity `json:"data,omitempty"`
	HasMore *bool               `json:"has_more,omitempty"`
	Total   int                 `json:"total,omitempty"`
	Object  string              `json:"object,omitempty"`
}

func (m EnvelopeActivityListWithCursor) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type EnvelopeCancelRevisionResponse struct {
	EnvelopeId string `json:"envelope_id,omitempty"`
	Object     string `json:"object,omitempty"`
}

func (m EnvelopeCancelRevisionResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type EnvelopeContact struct {
	Agent          *EnvelopeContactAgent `json:"agent,omitempty"`
	CellPhone      string                `json:"cell_phone,omitempty"`
	Email          string                `json:"email,omitempty"`
	EmailAddresses []string              `json:"email_addresses,omitempty"`
	EntityType     string                `json:"entity_type,omitempty"`
	FirstName      string                `json:"first_name,omitempty"`
	LastName       string                `json:"last_name,omitempty"`
	PhoneNumbers   []string              `json:"phone_numbers,omitempty"`
}

type EnvelopeContactAgent struct {
	CompanyName string `json:"company_name,omitempty"`
}

type EnvelopeCreateIntentSchema struct {
	AccessPolicy  *AccessPolicy              `json:"access_policy,omitempty"`
	CallbackPath  string                     `json:"callback_path,omitempty"`
	Documents     []*InitialEnvelopeDocument `json:"documents,omitempty"`
	EmailSubject  string                     `json:"email_subject,omitempty"`
	EmailMessage  string                     `json:"email_message,omitempty"`
	ExternalId    string                     `json:"external_id,omitempty"`
	Recipients    []*InitialRecipient        `json:"recipients,omitempty"`
	TransactionId string                     `json:"transaction_id,omitempty"`
}

type EnvelopeCreateResponse struct {
	Envelope   *EnvelopeResponse `json:"envelope,omitempty"`
	EnvelopeId string            `json:"envelope_id,omitempty"`
	Object     string            `json:"object,omitempty"`
}

func (m EnvelopeCreateResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type EnvelopeDocument struct {
	Id         string `json:"id,omitempty"`
	DocumentId string `json:"document_id,omitempty"`
	ExternalId string `json:"external_id,omitempty"`
	Filename   string `json:"filename,omitempty"`
	Seq        int    `json:"seq,omitempty"`
	Title      string `json:"title,omitempty"`
	Object     string `json:"object,omitempty"`
}

func (m EnvelopeDocument) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type EnvelopeDocumentList struct {
	Data       []EnvelopeDocument `json:"data"`
	ListObject string             `json:"list_object"`
	Object     string             `json:"object"`
	HasMore    bool               `json:"has_more"`
}

func (m EnvelopeDocumentList) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

func (m EnvelopeDocumentList) NextPageParams() core.PageParams {
	if !m.HasMore {
		return nil
	}

	pageSize := len(m.Data)
	return &core.PageParamsImpl{
		StartingAfter: m.Data[pageSize-1].Id,
		Limit:         pageSize,
	}
}

type EnvelopeField struct {
	Id                 string                `json:"id,omitempty"`
	ConditionalLinking []*ConditionalLinking `json:"conditional_linking,omitempty"`
	FillConditions     []*FillCondition      `json:"fill_conditions,omitempty"`
	Kind               string                `json:"kind,omitempty"`
	LinkId             string                `json:"link_id,omitempty"`
	LinkNamespace      string                `json:"link_namespace,omitempty"`
	OverflowPdfFormat  string                `json:"overflow_pdf_format,omitempty"`
	OwnerId            string                `json:"owner_id,omitempty"`
}

type EnvelopeRecipient struct {
	Id             string          `json:"id,omitempty"`
	Contact        *Contact1       `json:"contact,omitempty"`
	ExternalId     string          `json:"external_id,omitempty"`
	Index          int             `json:"index,omitempty"`
	InitialsImage  *SignatureImage `json:"initials_image,omitempty"`
	RecipientRole  string          `json:"recipient_role,omitempty"`
	SignatureImage *SignatureImage `json:"signature_image,omitempty"`
	SigningStatus  string          `json:"signing_status,omitempty"`
	Object         string          `json:"object,omitempty"`
}

func (m EnvelopeRecipient) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type EnvelopeRecipientList struct {
	Data       []EnvelopeRecipient `json:"data"`
	ListObject string              `json:"list_object"`
	Object     string              `json:"object"`
	HasMore    bool                `json:"has_more"`
}

func (m EnvelopeRecipientList) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

func (m EnvelopeRecipientList) NextPageParams() core.PageParams {
	if !m.HasMore {
		return nil
	}

	pageSize := len(m.Data)
	return &core.PageParamsImpl{
		StartingAfter: m.Data[pageSize-1].Id,
		Limit:         pageSize,
	}
}

type EnvelopeResendResponse struct {
	EnvelopeId string `json:"envelope_id,omitempty"`
	Object     string `json:"object,omitempty"`
}

func (m EnvelopeResendResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type EnvelopeResponse struct {
	EnvelopeDocuments  []*EnvelopeDocument  `json:"envelope_documents,omitempty"`
	EnvelopeRecipients []*EnvelopeRecipient `json:"envelope_recipients,omitempty"`
	LatestVersionId    string               `json:"latest_version_id,omitempty"`
	SigningUrl         string               `json:"signing_url,omitempty"`
	Uuid               string               `json:"uuid,omitempty"`
	Object             string               `json:"object,omitempty"`
}

func (m EnvelopeResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type EnvelopeSendRevisionResponse struct {
	Envelope   *EnvelopeResponse `json:"envelope,omitempty"`
	EnvelopeId string            `json:"envelope_id,omitempty"`
	Object     string            `json:"object,omitempty"`
}

func (m EnvelopeSendRevisionResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type EnvelopeSendRevisionSchema struct {
	Documents        []*InitialEnvelopeDocument `json:"documents,omitempty"`
	EmailMessage     string                     `json:"email_message,omitempty"`
	EmailSubject     string                     `json:"email_subject,omitempty"`
	ExternalId       string                     `json:"external_id,omitempty"`
	LockedSignerKeys []string                   `json:"locked_signer_keys,omitempty"`
	Recipients       []*InitialRecipient        `json:"recipients,omitempty"`
}

type EnvelopeStartRevisionResponse struct {
	EnvelopeId string `json:"envelope_id,omitempty"`
	IsDelayed  *bool  `json:"is_delayed,omitempty"`
	Object     string `json:"object,omitempty"`
}

func (m EnvelopeStartRevisionResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type EnvelopeVoidResponse struct {
	EnvelopeId string `json:"envelope_id,omitempty"`
	Object     string `json:"object,omitempty"`
}

func (m EnvelopeVoidResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type EnvelopeVoidSchema struct {
	Message string `json:"message,omitempty"`
}

type EsignTabConfig struct {
	FillConfigs []*ESignFillConfig `json:"fill_configs,omitempty"`
}

type ExtractedField struct {
	ExtractedData          string `json:"extracted_data,omitempty"`
	FormPage               int    `json:"form_page,omitempty"`
	GlideDataDictionaryKey string `json:"glide_data_dictionary_key"`
	Object                 string `json:"object,omitempty"`
}

func (m ExtractedField) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type Field struct {
	Timestamp int                    `json:"timestamp,omitempty"`
	Value     map[string]interface{} `json:"value,omitempty"`
	Object    string                 `json:"object,omitempty"`
}

func (m Field) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type FieldOutOfDateDetail struct {
	ControlTimestamp int    `json:"control_timestamp,omitempty"`
	Timestamp        int    `json:"timestamp,omitempty"`
	Object           string `json:"object,omitempty"`
}

func (m FieldOutOfDateDetail) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type FieldParams struct {
	Options []string `json:"options,omitempty"`
	Object  string   `json:"object,omitempty"`
}

func (m FieldParams) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type FieldResponse struct {
	Timestamp int                    `json:"timestamp,omitempty"`
	Value     map[string]interface{} `json:"value,omitempty"`
	Object    string                 `json:"object,omitempty"`
}

func (m FieldResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type FieldResponseWarnings struct {
	OutOfDateFields map[string]*FieldOutOfDateDetail `json:"out_of_date_fields,omitempty"`
	Object          string                           `json:"object,omitempty"`
}

func (m FieldResponseWarnings) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type FieldWrite struct {
	ControlTimestamp int                    `json:"control_timestamp,omitempty"`
	Value            map[string]interface{} `json:"value,omitempty"`
}

type FieldWriteDict struct {
	ControlPolicy string                 `json:"control_policy,omitempty"`
	Fields        TransactionFieldsWrite `json:"fields,omitempty"`
}

type FieldsResponse struct {
	Result        *FieldsResponseResult `json:"result,omitempty"`
	TransactionId string                `json:"transaction_id,omitempty"`
	Object        string                `json:"object,omitempty"`
}

func (m FieldsResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type FieldsResponseResult struct {
	Fields   TransactionFields      `json:"fields,omitempty"`
	Warnings *FieldResponseWarnings `json:"warnings,omitempty"`
	Object   string                 `json:"object,omitempty"`
}

func (m FieldsResponseResult) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type FileMeta struct {
	Data                       http.File `json:"data,omitempty"`
	FileName                   string    `json:"file_name"`
	FileReferenceId            string    `json:"file_reference_id"`
	IncludeDatapointExtraction *bool     `json:"include_datapoint_extraction,omitempty"`
	IncludeFormMatching        *bool     `json:"include_form_matching,omitempty"`
	IncludeSignatureDetection  *bool     `json:"include_signature_detection,omitempty"`
	MimeType                   string    `json:"mime_type,omitempty"`
	StateCode                  string    `json:"state_code"`
	Url                        string    `json:"url,omitempty"`
}

type FillCondition struct {
	Id      string `json:"id,omitempty"`
	Level   string `json:"level,omitempty"`
	Message string `json:"message,omitempty"`
	Rule    string `json:"rule,omitempty"`
	Terms   string `json:"terms,omitempty"`
}

type FillConfig struct {
	Annotations        []*Annotations      `json:"annotations,omitempty"`
	ReformForm         *Form               `json:"reform_form,omitempty"`
	ReformFormPrepared *ReformFormPrepared `json:"reform_form_prepared,omitempty"`
}

type Folder struct {
	Id                        string                     `json:"id,omitempty"`
	Can                       map[string]*bool           `json:"can,omitempty"`
	IntegratedServicesPartner *IntegratedServicesPartner `json:"integrated_services_partner,omitempty"`
	Kind                      string                     `json:"kind,omitempty"`
	LastModified              int                        `json:"last_modified,omitempty"`
	OrderIndex                int                        `json:"order_index,omitempty"`
	PropertyInfo              *PropertyInfo              `json:"property_info,omitempty"`
	Title                     string                     `json:"title,omitempty"`
	TransactionDocuments      *TransactionDocumentList   `json:"transaction_documents,omitempty"`
	TransactionPackage        *TransactionPackage        `json:"transaction_package,omitempty"`
	Object                    string                     `json:"object,omitempty"`
}

func (m Folder) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type FolderList struct {
	Data       []Folder `json:"data"`
	ListObject string   `json:"list_object"`
	Object     string   `json:"object"`
	HasMore    bool     `json:"has_more"`
}

func (m FolderList) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

func (m FolderList) NextPageParams() core.PageParams {
	if !m.HasMore {
		return nil
	}

	pageSize := len(m.Data)
	return &core.PageParamsImpl{
		StartingAfter: m.Data[pageSize-1].Id,
		Limit:         pageSize,
	}
}

type FolderCreate struct {
	IntegratedServicesPartner *FolderCreateIntegratedServicesPartner `json:"integrated_services_partner,omitempty"`
	Kind                      int                                    `json:"kind,omitempty"`
	Title                     string                                 `json:"title,omitempty"`
}

type FolderCreateIntegratedServicesPartner struct {
	DisplayName string `json:"display_name,omitempty"`
	EnumValue   int    `json:"enum_value,omitempty"`
}

type FolderCreates struct {
	Creates []*FolderCreate `json:"creates,omitempty"`
}

type FolderCreatesResponse struct {
	Result        *FolderCreatesResponseResult `json:"result,omitempty"`
	TransactionId string                       `json:"transaction_id,omitempty"`
	Object        string                       `json:"object,omitempty"`
}

func (m FolderCreatesResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type FolderCreatesResponseResult struct {
	FolderIds []string `json:"folder_ids,omitempty"`
	Object    string   `json:"object,omitempty"`
}

func (m FolderCreatesResponseResult) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type FolderRename struct {
	FolderId string `json:"folder_id,omitempty"`
	Title    string `json:"title,omitempty"`
}

type FolderRenames struct {
	Renames []*FolderRename `json:"renames,omitempty"`
}

type FolderRenamesResponse struct {
	TransactionId string `json:"transaction_id,omitempty"`
	Object        string `json:"object,omitempty"`
}

func (m FolderRenamesResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type Form struct {
	Fields []*EnvelopeField `json:"fields,omitempty"`
}

type FormImportsResponse struct {
	TransactionId string `json:"transaction_id,omitempty"`
	Object        string `json:"object,omitempty"`
}

func (m FormImportsResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type FormMatch struct {
	EndPage                 int                         `json:"end_page"`
	ExtractedFields         []*ExtractedField           `json:"extracted_fields,omitempty"`
	Form                    *GlideForm                  `json:"form,omitempty"`
	Score                   float64                     `json:"score,omitempty"`
	SignatureResultsByParty map[string]*SignatureResult `json:"signature_results_by_party,omitempty"`
	StartPage               int                         `json:"start_page"`
	Object                  string                      `json:"object,omitempty"`
}

func (m FormMatch) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type GetSignatureRequestFlowDocumentsResponse struct {
	DestinationFolderId          string                                         `json:"destination_folder_id,omitempty"`
	HideFieldMoveOriginalToTrash *bool                                          `json:"hide_field_move_original_to_trash,omitempty"`
	LockedTransactionDocumentIds []string                                       `json:"locked_transaction_document_ids,omitempty"`
	MakeDocumentsVisibleInCd     *bool                                          `json:"make_documents_visible_in_cd,omitempty"`
	MoveOriginalToTrash          *bool                                          `json:"move_original_to_trash,omitempty"`
	TransactionDocumentIds       []string                                       `json:"transaction_document_ids,omitempty"`
	TransactionDocuments         []*SignatureRequestExpandedTransactionDocument `json:"transaction_documents,omitempty"`
	Object                       string                                         `json:"object,omitempty"`
}

func (m GetSignatureRequestFlowDocumentsResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type GetSignatureRequestFlowRecipientsResponse struct {
	IsSigningOrderApplied *bool        `json:"is_signing_order_applied,omitempty"`
	LockedOrderSignerKeys []string     `json:"locked_order_signer_keys,omitempty"`
	LockedSignerKeys      []string     `json:"locked_signer_keys,omitempty"`
	Recipients            []*Recipient `json:"recipients,omitempty"`
	Object                string       `json:"object,omitempty"`
}

func (m GetSignatureRequestFlowRecipientsResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type GetSignatureRequestFlowReviewResponse struct {
	MakeDocumentsVisibleInCd *bool                                          `json:"make_documents_visible_in_cd,omitempty"`
	MessageBody              string                                         `json:"message_body,omitempty"`
	MessageSubject           string                                         `json:"message_subject,omitempty"`
	Recipients               []*Recipient                                   `json:"recipients,omitempty"`
	Subject                  string                                         `json:"subject,omitempty"`
	TransactionDocumentIds   []string                                       `json:"transaction_document_ids,omitempty"`
	TransactionDocuments     []*SignatureRequestExpandedTransactionDocument `json:"transaction_documents,omitempty"`
	Object                   string                                         `json:"object,omitempty"`
}

func (m GetSignatureRequestFlowReviewResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type GlideForm struct {
	GlideFormSeriesId int      `json:"glide_form_series_id"`
	Tags              []string `json:"tags,omitempty"`
	Title             string   `json:"title"`
	Object            string   `json:"object,omitempty"`
}

func (m GlideForm) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type Imports struct {
	Id    string `json:"id,omitempty"`
	Alias string `json:"alias,omitempty"`
}

type InitialEnvelopeDocument struct {
	Id          string                 `json:"id,omitempty"`
	Annotations []*SignatureAnnotation `json:"annotations,omitempty"`
	ExternalId  string                 `json:"external_id,omitempty"`
}

type InitialRecipient struct {
	Contact       *EnvelopeContact `json:"contact,omitempty"`
	ExternalId    string           `json:"external_id,omitempty"`
	Order         string           `json:"order,omitempty"`
	RecipientRole string           `json:"recipient_role,omitempty"`
	Source        *RecipientSource `json:"source,omitempty"`
}

type IntegratedServicesPartner struct {
	BannerDismissed *bool  `json:"banner_dismissed,omitempty"`
	DisplayName     string `json:"display_name,omitempty"`
	EnumValue       int    `json:"enum_value,omitempty"`
	Object          string `json:"object,omitempty"`
}

func (m IntegratedServicesPartner) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type ItemDeletes struct {
	Ids []string `json:"ids,omitempty"`
}

type ItemDeletesResponse struct {
	TransactionId string `json:"transaction_id,omitempty"`
	Object        string `json:"object,omitempty"`
}

func (m ItemDeletesResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type Member struct {
	Id     string `json:"id,omitempty"`
	Offer  *Offer `json:"offer,omitempty"`
	Object string `json:"object,omitempty"`
}

func (m Member) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type MergeDocumentsResponse struct {
	IsDelayed     *bool  `json:"is_delayed,omitempty"`
	TransactionId string `json:"transaction_id,omitempty"`
	Object        string `json:"object,omitempty"`
}

func (m MergeDocumentsResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type NewContactRecipient struct {
	Contact *Contact `json:"contact,omitempty"`
	Roles   []string `json:"roles,omitempty"`
	Object  string   `json:"object,omitempty"`
}

func (m NewContactRecipient) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type Notification struct {
	Bcc              []string               `json:"bcc,omitempty"`
	Cc               []string               `json:"cc,omitempty"`
	Context          map[string]interface{} `json:"context,omitempty"`
	IncludeSignature *bool                  `json:"include_signature,omitempty"`
	Recipients       []string               `json:"recipients,omitempty"`
	SeparateEmails   *bool                  `json:"separate_emails,omitempty"`
	Template         string                 `json:"template"`
}

type NotificationResponse struct {
	Results []string `json:"results"`
	Object  string   `json:"object,omitempty"`
}

func (m NotificationResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type Offer struct {
	Id              string   `json:"id,omitempty"`
	Archived        *bool    `json:"archived,omitempty"`
	Favorite        *bool    `json:"favorite,omitempty"`
	Notifications   []string `json:"notifications,omitempty"`
	Status          string   `json:"status,omitempty"`
	TransactionSide string   `json:"transaction_side,omitempty"`
	Object          string   `json:"object,omitempty"`
}

func (m Offer) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type OfferPartiesResponse struct {
	Data   []*OfferParty `json:"data,omitempty"`
	Object string        `json:"object,omitempty"`
}

func (m OfferPartiesResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type OfferParty struct {
	Id               string       `json:"id,omitempty"`
	ClientVisibility string       `json:"client_visibility,omitempty"`
	Contact          *Contact     `json:"contact,omitempty"`
	CreatedAt        int          `json:"created_at,omitempty"`
	Roles            []string     `json:"roles,omitempty"`
	Transaction      *Transaction `json:"transaction,omitempty"`
	UpdatedAt        int          `json:"updated_at,omitempty"`
	UserId           string       `json:"user_id,omitempty"`
	Object           string       `json:"object,omitempty"`
}

func (m OfferParty) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type OffersResponse struct {
	Data   string `json:"data,omitempty"`
	Object string `json:"object,omitempty"`
}

func (m OffersResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type Output struct {
	Id      string     `json:"id,omitempty"`
	Handler string     `json:"handler,omitempty"`
	Imports []*Imports `json:"imports,omitempty"`
	OutKind string     `json:"out_kind,omitempty"`
}

type PartnerFolderUploadsResponse struct {
	Result        *TransactionDocumentUploadResult `json:"result,omitempty"`
	TransactionId string                           `json:"transaction_id,omitempty"`
	Object        string                           `json:"object,omitempty"`
}

func (m PartnerFolderUploadsResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type Party struct {
	Id               string       `json:"id,omitempty"`
	ClientVisibility string       `json:"client_visibility,omitempty"`
	Contact          *Contact     `json:"contact,omitempty"`
	CreatedAt        int          `json:"created_at,omitempty"`
	Roles            []string     `json:"roles,omitempty"`
	Transaction      *Transaction `json:"transaction,omitempty"`
	UpdatedAt        int          `json:"updated_at,omitempty"`
	UserId           string       `json:"user_id,omitempty"`
	Object           string       `json:"object,omitempty"`
}

func (m Party) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type PartyList struct {
	Data       []Party `json:"data"`
	ListObject string  `json:"list_object"`
	Object     string  `json:"object"`
	HasMore    bool    `json:"has_more"`
}

func (m PartyList) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

func (m PartyList) NextPageParams() core.PageParams {
	if !m.HasMore {
		return nil
	}

	pageSize := len(m.Data)
	return &core.PageParamsImpl{
		StartingAfter: m.Data[pageSize-1].Id,
		Limit:         pageSize,
	}
}

type PartyCreate struct {
	Body                  string          `json:"body,omitempty"`
	ClientVisibility      string          `json:"client_visibility,omitempty"`
	Contact               *ContactRequest `json:"contact,omitempty"`
	Invite                *bool           `json:"invite,omitempty"`
	InviteRestrictions    []string        `json:"invite_restrictions,omitempty"`
	PromoteToPrimaryAgent *bool           `json:"promote_to_primary_agent,omitempty"`
	Roles                 []string        `json:"roles,omitempty"`
	Subject               string          `json:"subject,omitempty"`
	SuppressInviteEmail   *bool           `json:"suppress_invite_email,omitempty"`
	TeamId                string          `json:"team_id,omitempty"`
}

type PartyCreates struct {
	Creates []*PartyCreate `json:"creates"`
}

type PartyCreatesResponse struct {
	Result        *PartyCreatesResult `json:"result,omitempty"`
	TransactionId string              `json:"transaction_id,omitempty"`
	Object        string              `json:"object,omitempty"`
}

func (m PartyCreatesResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type PartyCreatesResult struct {
	Parties []*Party `json:"parties,omitempty"`
	Object  string   `json:"object,omitempty"`
}

func (m PartyCreatesResult) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type PartyInvite struct {
	Body                string   `json:"body,omitempty"`
	InviteRestrictions  []string `json:"invite_restrictions,omitempty"`
	PartyId             string   `json:"party_id"`
	Subject             string   `json:"subject,omitempty"`
	SuppressInviteEmail *bool    `json:"suppress_invite_email,omitempty"`
}

type PartyInvites struct {
	Invites []*PartyInvite `json:"invites"`
}

type PartyInvitesResponse struct {
	TransactionId string `json:"transaction_id,omitempty"`
	Object        string `json:"object,omitempty"`
}

func (m PartyInvitesResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type PartyPatch struct {
	ClientVisibility string          `json:"client_visibility,omitempty"`
	Contact          *ContactRequest `json:"contact,omitempty"`
	PartyId          string          `json:"party_id,omitempty"`
	Roles            []string        `json:"roles,omitempty"`
	TeamId           string          `json:"team_id,omitempty"`
}

type PartyPatches struct {
	Patches []*PartyPatch `json:"patches,omitempty"`
}

type PartyPatchesResponse struct {
	Result        *PartyPatchesResult `json:"result,omitempty"`
	TransactionId string              `json:"transaction_id,omitempty"`
	Object        string              `json:"object,omitempty"`
}

func (m PartyPatchesResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type PartyPatchesResult struct {
	Parties []*Party `json:"parties,omitempty"`
	Object  string   `json:"object,omitempty"`
}

func (m PartyPatchesResult) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type PartyRecipient struct {
	Id      string   `json:"id,omitempty"`
	Contact *Contact `json:"contact,omitempty"`
	Roles   []string `json:"roles,omitempty"`
	Vers    string   `json:"vers,omitempty"`
	Object  string   `json:"object,omitempty"`
}

func (m PartyRecipient) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type PartyRemove struct {
	PartyId string `json:"party_id,omitempty"`
}

type PartyRemoves struct {
	Removes []*PartyRemove `json:"removes,omitempty"`
}

type PartyRemovesResponse struct {
	TransactionId string `json:"transaction_id,omitempty"`
	Object        string `json:"object,omitempty"`
}

func (m PartyRemovesResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type PartyRoles struct {
	Data   []string `json:"data,omitempty"`
	Object string   `json:"object,omitempty"`
}

func (m PartyRoles) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type PartyUpdateContactDetails struct {
	ClientVisibility      string          `json:"client_visibility,omitempty"`
	Contact               *ContactRequest `json:"contact,omitempty"`
	PartyId               string          `json:"party_id,omitempty"`
	PromoteToPrimaryAgent *bool           `json:"promote_to_primary_agent,omitempty"`
	Roles                 []string        `json:"roles,omitempty"`
	TeamId                string          `json:"team_id,omitempty"`
}

type PartyUpdateContactDetailsResponse struct {
	TransactionId string `json:"transaction_id,omitempty"`
	Object        string `json:"object,omitempty"`
}

func (m PartyUpdateContactDetailsResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type PartyUpdateContactSource struct {
	ContactSource *ContactSource `json:"contact_source,omitempty"`
	PartyId       string         `json:"party_id"`
}

type PartyUpdateContactSourceResponse struct {
	TransactionId string `json:"transaction_id,omitempty"`
	Object        string `json:"object,omitempty"`
}

func (m PartyUpdateContactSourceResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type PropertyInfo struct {
	Id           string       `json:"id,omitempty"`
	Address      *Address     `json:"address,omitempty"`
	EmailAddress string       `json:"email_address,omitempty"`
	IsSecondary  *bool        `json:"is_secondary,omitempty"`
	PropertyType string       `json:"property_type,omitempty"`
	Transaction  *Transaction `json:"transaction,omitempty"`
	Object       string       `json:"object,omitempty"`
}

func (m PropertyInfo) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type PropertyInfoList struct {
	Data       []PropertyInfo `json:"data"`
	ListObject string         `json:"list_object"`
	Object     string         `json:"object"`
	HasMore    bool           `json:"has_more"`
}

func (m PropertyInfoList) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

func (m PropertyInfoList) NextPageParams() core.PageParams {
	if !m.HasMore {
		return nil
	}

	pageSize := len(m.Data)
	return &core.PageParamsImpl{
		StartingAfter: m.Data[pageSize-1].Id,
		Limit:         pageSize,
	}
}

type Recipient struct {
	Key           string                `json:"key,omitempty"`
	Kind          string                `json:"kind,omitempty"`
	NewContact    *NewContactRecipient  `json:"new_contact,omitempty"`
	Order         int                   `json:"order,omitempty"`
	Party         *PartyRecipient       `json:"party,omitempty"`
	RecipientRole string                `json:"recipient_role,omitempty"`
	UserContact   *UserContactRecipient `json:"user_contact,omitempty"`
	Object        string                `json:"object,omitempty"`
}

func (m Recipient) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type RecipientSource struct {
	Kind        string       `json:"kind,omitempty"`
	UserContact *UserContact `json:"user_contact,omitempty"`
}

type ReformFormPrepareField struct {
	Id                 string                `json:"id,omitempty"`
	ConditionalLinking []*ConditionalLinking `json:"conditional_linking,omitempty"`
	FillConditions     []*FillCondition      `json:"fill_conditions,omitempty"`
	Kind               string                `json:"kind,omitempty"`
	LinkId             string                `json:"link_id,omitempty"`
	LinkNamespace      string                `json:"link_namespace,omitempty"`
	OverflowPdfFormat  string                `json:"overflow_pdf_format,omitempty"`
	OwnerId            string                `json:"owner_id,omitempty"`
	Validations        []*Validation         `json:"validations,omitempty"`
}

type ReformFormPrepared struct {
	Fields  []*ReformFormPrepareField `json:"fields,omitempty"`
	Outputs []*Output                 `json:"outputs,omitempty"`
}

type ReorderFoldersResponse struct {
	TransactionId string `json:"transaction_id,omitempty"`
	Object        string `json:"object,omitempty"`
}

func (m ReorderFoldersResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type ReplacePrimaryAgent struct {
	PartyId string `json:"party_id"`
	Object  string `json:"object,omitempty"`
}

func (m ReplacePrimaryAgent) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type ReplacePrimaryAgentResponse struct {
	Result        *ReplacePrimaryAgent `json:"result,omitempty"`
	TransactionId string               `json:"transaction_id,omitempty"`
	Object        string               `json:"object,omitempty"`
}

func (m ReplacePrimaryAgentResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type Signature struct {
	Kind   string `json:"kind,omitempty"`
	Text   string `json:"text,omitempty"`
	Object string `json:"object,omitempty"`
}

func (m Signature) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type SignatureAnnotation struct {
	Id             string                     `json:"id,omitempty"`
	Color          string                     `json:"color,omitempty"`
	ExtraParams    *SignatureAnnotationParams `json:"extra_params,omitempty"`
	FieldId        string                     `json:"field_id,omitempty"`
	Height         float64                    `json:"height,omitempty"`
	Kind           string                     `json:"kind,omitempty"`
	Left           float64                    `json:"left,omitempty"`
	MultiPartIndex int                        `json:"multi_part_index,omitempty"`
	Optional       *bool                      `json:"optional,omitempty"`
	PageIndex      int                        `json:"page_index,omitempty"`
	ReadOnly       *bool                      `json:"read_only,omitempty"`
	RecipientId    string                     `json:"recipient_id,omitempty"`
	Top            float64                    `json:"top,omitempty"`
	Width          float64                    `json:"width,omitempty"`
}

type SignatureAnnotationParams struct {
	AllowInput   *bool                  `json:"allow_input,omitempty"`
	DefaultValue map[string]interface{} `json:"default_value,omitempty"`
	ExportValue  map[string]interface{} `json:"export_value,omitempty"`
	Options      []string               `json:"options,omitempty"`
}

type SignatureDetectionAnalysisResult struct {
	DocumentZone *DocumentZone `json:"document_zone,omitempty"`
	Score        float64       `json:"score,omitempty"`
	Object       string        `json:"object,omitempty"`
}

func (m SignatureDetectionAnalysisResult) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type SignatureDetectionAsyncResponse struct {
	ReqId      string                                       `json:"req_id,omitempty"`
	Signatures map[string]*SignatureDetectionAnalysisResult `json:"signatures,omitempty"`
	Object     string                                       `json:"object,omitempty"`
}

func (m SignatureDetectionAsyncResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type SignatureDetectionResponse struct {
	ReqId  string                           `json:"req_id,omitempty"`
	Result *SignatureDetectionAsyncResponse `json:"result,omitempty"`
	Object string                           `json:"object,omitempty"`
}

func (m SignatureDetectionResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type SignatureDetectionSchema struct {
	Files   []http.File       `json:"files,omitempty"`
	Uploads []*DocumentUpload `json:"uploads,omitempty"`
}

type SignatureImage struct {
	DocumentId string     `json:"document_id,omitempty"`
	Signature  *Signature `json:"signature,omitempty"`
	Object     string     `json:"object,omitempty"`
}

func (m SignatureImage) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type SignatureIntent struct {
	MaxLength int `json:"max_length,omitempty"`
}

type SignatureRequest struct {
	Id               string                       `json:"id,omitempty"`
	Activities       *ActivityListWithCursor      `json:"activities,omitempty"`
	CompletedAt      int                          `json:"completed_at,omitempty"`
	CreatedAt        int                          `json:"created_at,omitempty"`
	CurrentFlowPage  string                       `json:"current_flow_page,omitempty"`
	DealId           string                       `json:"deal_id,omitempty"`
	Documents        []*SignatureRequestDocument  `json:"documents,omitempty"`
	EnvelopeId       string                       `json:"envelope_id,omitempty"`
	FlowId           string                       `json:"flow_id,omitempty"`
	IsArchived       *bool                        `json:"is_archived,omitempty"`
	Message          string                       `json:"message,omitempty"`
	Provider         string                       `json:"provider,omitempty"`
	Recipients       []*SignatureRequestRecipient `json:"recipients,omitempty"`
	RevisionFlowId   string                       `json:"revision_flow_id,omitempty"`
	SentAt           int                          `json:"sent_at,omitempty"`
	SignNowUrl       string                       `json:"sign_now_url,omitempty"`
	Status           string                       `json:"status,omitempty"`
	Title            string                       `json:"title,omitempty"`
	TransactionId    string                       `json:"transaction_id,omitempty"`
	TransactionTitle string                       `json:"transaction_title,omitempty"`
	UpdatedAt        int                          `json:"updated_at,omitempty"`
	Object           string                       `json:"object,omitempty"`
}

func (m SignatureRequest) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type SignatureRequestList struct {
	Data       []SignatureRequest `json:"data"`
	ListObject string             `json:"list_object"`
	Object     string             `json:"object"`
	HasMore    bool               `json:"has_more"`
}

func (m SignatureRequestList) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

func (m SignatureRequestList) NextPageParams() core.PageParams {
	if !m.HasMore {
		return nil
	}

	pageSize := len(m.Data)
	return &core.PageParamsImpl{
		StartingAfter: m.Data[pageSize-1].Id,
		Limit:         pageSize,
	}
}

type SignatureRequestArchiveResponse struct {
	TransactionId string `json:"transaction_id,omitempty"`
	Object        string `json:"object,omitempty"`
}

func (m SignatureRequestArchiveResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type SignatureRequestDocument struct {
	Id                 string `json:"id,omitempty"`
	EnvelopeDocumentId string `json:"envelope_document_id,omitempty"`
	LatestVersionId    string `json:"latest_version_id,omitempty"`
	Name               string `json:"name,omitempty"`
	Seq                int    `json:"seq,omitempty"`
	Object             string `json:"object,omitempty"`
}

func (m SignatureRequestDocument) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type SignatureRequestDuplicateResponse struct {
	FlowId             string `json:"flow_id,omitempty"`
	SignatureRequestId string `json:"signature_request_id,omitempty"`
	Object             string `json:"object,omitempty"`
}

func (m SignatureRequestDuplicateResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type SignatureRequestExpandedTransactionDocument struct {
	Id                 string `json:"id,omitempty"`
	EnvelopeDocumentId string `json:"envelope_document_id,omitempty"`
	LastModified       int    `json:"last_modified,omitempty"`
	LatestVersionId    string `json:"latest_version_id,omitempty"`
	SignatureStatus    string `json:"signature_status,omitempty"`
	Title              string `json:"title,omitempty"`
	Object             string `json:"object,omitempty"`
}

func (m SignatureRequestExpandedTransactionDocument) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type SignatureRequestFlowDocumentsRequest struct {
	FlowId                  string   `json:"flow_id,omitempty"`
	MoveOriginalToTrash     *bool    `json:"move_original_to_trash,omitempty"`
	TransactionDocumentsIds []string `json:"transaction_documents_ids,omitempty"`
}

type SignatureRequestFlowDocumentsResponse struct {
	FlowId string `json:"flow_id,omitempty"`
	Object string `json:"object,omitempty"`
}

func (m SignatureRequestFlowDocumentsResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type SignatureRequestFlowRecipientsRequest struct {
	FlowId     string       `json:"flow_id"`
	Recipients []*Recipient `json:"recipients,omitempty"`
}

type SignatureRequestFlowRecipientsResponse struct {
	FlowId string `json:"flow_id,omitempty"`
	Object string `json:"object,omitempty"`
}

func (m SignatureRequestFlowRecipientsResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type SignatureRequestFlowRequest struct {
	ExternalTaskId          string       `json:"external_task_id,omitempty"`
	Recipients              []*Recipient `json:"recipients,omitempty"`
	TransactionDocumentsIds []string     `json:"transaction_documents_ids,omitempty"`
	TransactionId           string       `json:"transaction_id,omitempty"`
}

type SignatureRequestFlowResponse struct {
	FlowId string `json:"flow_id,omitempty"`
	Object string `json:"object,omitempty"`
}

func (m SignatureRequestFlowResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type SignatureRequestFlowReviewRequest struct {
	Body                     string `json:"body,omitempty"`
	FlowId                   string `json:"flow_id"`
	MakeDocumentsVisibleInCd *bool  `json:"make_documents_visible_in_cd,omitempty"`
	Subject                  string `json:"subject,omitempty"`
}

type SignatureRequestFlowReviewResponse struct {
	FlowId string `json:"flow_id,omitempty"`
	Object string `json:"object,omitempty"`
}

func (m SignatureRequestFlowReviewResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type SignatureRequestFlowSendRequest struct {
	FlowId string `json:"flow_id"`
}

type SignatureRequestFlowSendResponse struct {
	SignatureRequest *SignatureRequest `json:"signature_request,omitempty"`
	Object           string            `json:"object,omitempty"`
}

func (m SignatureRequestFlowSendResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type SignatureRequestFlowSendRevisionResponse struct {
	SignatureRequest *SignatureRequest `json:"signature_request,omitempty"`
	Object           string            `json:"object,omitempty"`
}

func (m SignatureRequestFlowSendRevisionResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type SignatureRequestRecipient struct {
	Id            string `json:"id,omitempty"`
	Email         string `json:"email,omitempty"`
	FirstName     string `json:"first_name,omitempty"`
	LastName      string `json:"last_name,omitempty"`
	Order         int    `json:"order,omitempty"`
	RecipientRole string `json:"recipient_role,omitempty"`
	Status        string `json:"status,omitempty"`
	UserName      string `json:"user_name,omitempty"`
	Object        string `json:"object,omitempty"`
}

func (m SignatureRequestRecipient) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type SignatureRequestReviseRequest struct {
	RevisionFlowId string `json:"revision_flow_id"`
}

type SignatureRequestSaveTabConfigRequest struct {
	Data   *EsignTabConfig `json:"data"`
	FlowId string          `json:"flow_id"`
}

type SignatureRequestSaveTabConfigResponse struct {
	FlowId string `json:"flow_id,omitempty"`
	Object string `json:"object,omitempty"`
}

func (m SignatureRequestSaveTabConfigResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type SignatureRequestStartRevisionResponse struct {
	FlowId string `json:"flow_id,omitempty"`
	Object string `json:"object,omitempty"`
}

func (m SignatureRequestStartRevisionResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type SignatureRequestTabConfigDetailResponse struct {
	FillConfigs             []*ESignFillConfig                             `json:"fill_configs,omitempty"`
	LockedSignerKeysByTdvId map[string][]string                            `json:"locked_signer_keys_by_tdv_id,omitempty"`
	Recipients              []*Recipient                                   `json:"recipients,omitempty"`
	TransactionDocuments    []*SignatureRequestExpandedTransactionDocument `json:"transaction_documents,omitempty"`
	Object                  string                                         `json:"object,omitempty"`
}

func (m SignatureRequestTabConfigDetailResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type SignatureRequestVoidRequest struct {
	VoidReason string `json:"void_reason,omitempty"`
}

type SignatureRequestVoidResponse struct {
	TransactionId string `json:"transaction_id,omitempty"`
	Object        string `json:"object,omitempty"`
}

func (m SignatureRequestVoidResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type SignatureResult struct {
	SignedSignatureFields int    `json:"signed_signature_fields"`
	TotalSignatureFields  int    `json:"total_signature_fields"`
	Object                string `json:"object,omitempty"`
}

func (m SignatureResult) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type SplitAnnotationTextRequest struct {
	SplitTexts []*TextSplitRequest `json:"split_texts"`
}

type SplitAnnotationTextResponse struct {
	Splits []*TextSplitResponse `json:"splits,omitempty"`
	Object string               `json:"object,omitempty"`
}

func (m SplitAnnotationTextResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type Step struct {
	Id       string `json:"id,omitempty"`
	ClosedAt int    `json:"closed_at,omitempty"`
	State    string `json:"state,omitempty"`
	Status   string `json:"status,omitempty"`
	StepKind string `json:"step_kind,omitempty"`
	Object   string `json:"object,omitempty"`
}

func (m Step) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type StepList struct {
	Data       []Step `json:"data"`
	ListObject string `json:"list_object"`
	Object     string `json:"object"`
	HasMore    bool   `json:"has_more"`
}

func (m StepList) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

func (m StepList) NextPageParams() core.PageParams {
	if !m.HasMore {
		return nil
	}

	pageSize := len(m.Data)
	return &core.PageParamsImpl{
		StartingAfter: m.Data[pageSize-1].Id,
		Limit:         pageSize,
	}
}

type Task struct {
	Id          string       `json:"id,omitempty"`
	BoardId     string       `json:"board_id,omitempty"`
	Name        string       `json:"name,omitempty"`
	OrderIndex  int          `json:"order_index,omitempty"`
	Status      string       `json:"status,omitempty"`
	TaskKind    string       `json:"task_kind,omitempty"`
	Title       string       `json:"title,omitempty"`
	Transaction *Transaction `json:"transaction,omitempty"`
	Type        string       `json:"type,omitempty"`
	Object      string       `json:"object,omitempty"`
}

func (m Task) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type TaskList struct {
	Data       []Task `json:"data"`
	ListObject string `json:"list_object"`
	Object     string `json:"object"`
	HasMore    bool   `json:"has_more"`
}

func (m TaskList) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

func (m TaskList) NextPageParams() core.PageParams {
	if !m.HasMore {
		return nil
	}

	pageSize := len(m.Data)
	return &core.PageParamsImpl{
		StartingAfter: m.Data[pageSize-1].Id,
		Limit:         pageSize,
	}
}

type TemplateDocuments struct {
	Id     string `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	Object string `json:"object,omitempty"`
}

func (m TemplateDocuments) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type TextSplitRequest struct {
	Annotations []*ESignAnnotation `json:"annotations,omitempty"`
	Text        string             `json:"text,omitempty"`
}

type TextSplitResponse struct {
	Remainder string   `json:"remainder,omitempty"`
	TextParts []string `json:"text_parts,omitempty"`
	Object    string   `json:"object,omitempty"`
}

func (m TextSplitResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type TimelineSignatureRequest struct {
	Id         string                       `json:"id,omitempty"`
	Recipients []*SignatureRequestRecipient `json:"recipients,omitempty"`
	SignNowUrl string                       `json:"sign_now_url,omitempty"`
	Object     string                       `json:"object,omitempty"`
}

func (m TimelineSignatureRequest) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type TimelineSignatureRequests struct {
	Data   []*TimelineSignatureRequest `json:"data,omitempty"`
	Object string                      `json:"object,omitempty"`
}

func (m TimelineSignatureRequests) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type Transaction struct {
	Id                           string                   `json:"id,omitempty"`
	Address                      *Address                 `json:"address,omitempty"`
	Archived                     *bool                    `json:"archived,omitempty"`
	DealId                       string                   `json:"deal_id,omitempty"`
	Fields                       TransactionFields        `json:"fields,omitempty"`
	Folders                      *FolderList              `json:"folders,omitempty"`
	IngestDocumentsEmail         string                   `json:"ingest_documents_email,omitempty"`
	IsLease                      *bool                    `json:"is_lease,omitempty"`
	IsReferral                   *bool                    `json:"is_referral,omitempty"`
	Parties                      *PartyList               `json:"parties,omitempty"`
	PropertiesInfo               *PropertyInfoList        `json:"properties_info,omitempty"`
	ReState                      string                   `json:"re_state,omitempty"`
	SecondaryAddressesIds        []string                 `json:"secondary_addresses_ids,omitempty"`
	Side                         string                   `json:"side,omitempty"`
	Stage                        string                   `json:"stage,omitempty"`
	Tasks                        *TaskList                `json:"tasks,omitempty"`
	Title                        string                   `json:"title,omitempty"`
	TransactionDocuments         *TransactionDocumentList `json:"transaction_documents,omitempty"`
	TransactionPackages          *TransactionPackageList  `json:"transaction_packages,omitempty"`
	TransactionSignatureRequests *SignatureRequestList    `json:"transaction_signature_requests,omitempty"`
	Object                       string                   `json:"object,omitempty"`
}

func (m Transaction) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type TransactionList struct {
	Data       []Transaction `json:"data"`
	ListObject string        `json:"list_object"`
	Object     string        `json:"object"`
	HasMore    bool          `json:"has_more"`
}

func (m TransactionList) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

func (m TransactionList) NextPageParams() core.PageParams {
	if !m.HasMore {
		return nil
	}

	pageSize := len(m.Data)
	return &core.PageParamsImpl{
		StartingAfter: m.Data[pageSize-1].Id,
		Limit:         pageSize,
	}
}

type TransactionFieldValue = interface{}

type TransactionFieldValues = map[string]TransactionFieldValue

type TransactionField struct {
	Value     TransactionFieldValue `json:"value"`
	Timestamp int                   `json:"timestamp"`
}

type TransactionFields = map[string]TransactionField

type TransactionFieldWrite struct {
	Value            TransactionFieldValue `json:"value"`
	ControlTimestamp int                   `json:"control_timestamp"`
}

type TransactionFieldsWrite = map[string]TransactionFieldWrite

func GetFieldWrite(value TransactionFieldValue, controlTimestamp int) TransactionFieldWrite {
	return TransactionFieldWrite{
		Value:            value,
		ControlTimestamp: controlTimestamp,
	}
}

func GetFieldWriteNoControl(value TransactionFieldValue) TransactionFieldWrite {
	return GetFieldWrite(value, 0)
}

func (t Transaction) GetFields(fieldIds ...string) TransactionFields {
	requestedFieldIds := map[string]bool{}
	for _, fk := range fieldIds {
		requestedFieldIds[fk] = true
	}
	res := TransactionFields{}
	for k, v := range t.Fields {
		if _, found := requestedFieldIds[k]; len(fieldIds) == 0 || found {
			res[k] = v
		}
	}
	return res
}

func (t Transaction) GetFieldsWrite(fieldValues TransactionFieldValues) TransactionFieldsWrite {
	res := TransactionFieldsWrite{}
	for k, v := range fieldValues {
		val, found := t.Fields[k]
		var controlTimestamp int
		if found {
			controlTimestamp = val.Timestamp
		} else {
			controlTimestamp = 0
		}
		res[k] = GetFieldWrite(v, controlTimestamp)
	}
	return res
}

func CombineFieldsWrites(fieldWrites ...TransactionFieldsWrite) TransactionFieldsWrite {
	res := TransactionFieldsWrite{}
	for _, fields := range fieldWrites {
		for k, v := range fields {
			res[k] = v
		}
	}
	return res
}

type TransactionAppliedTemplatesResult struct {
	CreatedDocuments []*TransactionDocument `json:"created_documents"`
	Object           string                 `json:"object,omitempty"`
}

func (m TransactionAppliedTemplatesResult) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type TransactionArchivalStatus struct {
	Archived *bool `json:"archived,omitempty"`
}

type TransactionByOrgSchema struct {
	Cursor  string   `json:"cursor,omitempty"`
	Data    []string `json:"data,omitempty"`
	HasMore *bool    `json:"has_more,omitempty"`
	Total   int      `json:"total,omitempty"`
	Object  string   `json:"object,omitempty"`
}

func (m TransactionByOrgSchema) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type TransactionCreate struct {
	AdditionalParties []*PartyCreate      `json:"additional_parties,omitempty"`
	Address           *Address            `json:"address,omitempty"`
	Creator           *TransactionCreator `json:"creator,omitempty"`
	CreatorRoles      []string            `json:"creator_roles,omitempty"`
	DealId            string              `json:"deal_id,omitempty"`
	IsLease           *bool               `json:"is_lease,omitempty"`
	IsReferral        *bool               `json:"is_referral,omitempty"`
	ReState           string              `json:"re_state,omitempty"`
	Stage             string              `json:"stage,omitempty"`
	TeamId            string              `json:"team_id,omitempty"`
	Title             string              `json:"title,omitempty"`
}

type TransactionCreator struct {
	UserContactId     string                `json:"user_contact_id,omitempty"`
	UserContactSource *ContactSourceRequest `json:"user_contact_source,omitempty"`
}

type TransactionDocument struct {
	Id                        string           `json:"id,omitempty"`
	Can                       map[string]*bool `json:"can,omitempty"`
	ClientDocumentProperty    string           `json:"client_document_property,omitempty"`
	ClientDocumentType        string           `json:"client_document_type,omitempty"`
	ClientVisibilityChangedAt int              `json:"client_visibility_changed_at,omitempty"`
	ClientVisibilityStatus    string           `json:"client_visibility_status,omitempty"`
	Folder                    *Folder          `json:"folder,omitempty"`
	FolderKind                string           `json:"folder_kind,omitempty"`
	FormId                    string           `json:"form_id,omitempty"`
	Kind                      string           `json:"kind,omitempty"`
	LastModified              int              `json:"last_modified,omitempty"`
	LatestVersionId           string           `json:"latest_version_id,omitempty"`
	Order                     int              `json:"order,omitempty"`
	Origin                    string           `json:"origin,omitempty"`
	PageCount                 int              `json:"page_count,omitempty"`
	SignatureStatus           string           `json:"signature_status,omitempty"`
	Title                     string           `json:"title,omitempty"`
	Transaction               *Transaction     `json:"transaction,omitempty"`
	UploadedBy                *UploadedBy      `json:"uploaded_by,omitempty"`
	Url                       string           `json:"url,omitempty"`
	Object                    string           `json:"object,omitempty"`
}

func (m TransactionDocument) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type TransactionDocumentList struct {
	Data       []TransactionDocument `json:"data"`
	ListObject string                `json:"list_object"`
	Object     string                `json:"object"`
	HasMore    bool                  `json:"has_more"`
}

func (m TransactionDocumentList) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

func (m TransactionDocumentList) NextPageParams() core.PageParams {
	if !m.HasMore {
		return nil
	}

	pageSize := len(m.Data)
	return &core.PageParamsImpl{
		StartingAfter: m.Data[pageSize-1].Id,
		Limit:         pageSize,
	}
}

type TransactionDocumentAssignment struct {
	FolderId              string `json:"folder_id,omitempty"`
	Order                 int    `json:"order,omitempty"`
	TransactionDocumentId string `json:"transaction_document_id,omitempty"`
}

type TransactionDocumentAssignments struct {
	Assignments []*TransactionDocumentAssignment `json:"assignments,omitempty"`
}

type TransactionDocumentAssignmentsResponse struct {
	TransactionId string `json:"transaction_id,omitempty"`
	Object        string `json:"object,omitempty"`
}

func (m TransactionDocumentAssignmentsResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type TransactionDocumentRename struct {
	Title                 string `json:"title,omitempty"`
	TransactionDocumentId string `json:"transaction_document_id,omitempty"`
}

type TransactionDocumentRenames struct {
	Renames []*TransactionDocumentRename `json:"renames,omitempty"`
}

type TransactionDocumentRenamesResponse struct {
	TransactionId string `json:"transaction_id,omitempty"`
	Object        string `json:"object,omitempty"`
}

func (m TransactionDocumentRenamesResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type TransactionDocumentReorderFolder struct {
	FolderId   string `json:"folder_id,omitempty"`
	OrderIndex int    `json:"order_index,omitempty"`
}

type TransactionDocumentReorderFolders struct {
	Folders []*TransactionDocumentReorderFolder `json:"folders,omitempty"`
}

type TransactionDocumentRestoresResponse struct {
	TransactionId string `json:"transaction_id,omitempty"`
	Object        string `json:"object,omitempty"`
}

func (m TransactionDocumentRestoresResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type TransactionDocumentTrashes struct {
	TransactionDocumentIds []string `json:"transaction_document_ids,omitempty"`
}

type TransactionDocumentTrashesResponse struct {
	TransactionId string `json:"transaction_id,omitempty"`
	Object        string `json:"object,omitempty"`
}

func (m TransactionDocumentTrashesResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type TransactionDocumentUpload struct {
	Folder                    *FolderCreate                          `json:"folder,omitempty"`
	FolderId                  string                                 `json:"folder_id,omitempty"`
	IntegratedServicesPartner *FolderCreateIntegratedServicesPartner `json:"integrated_services_partner,omitempty"`
	MimeType                  string                                 `json:"mime_type,omitempty"`
	Title                     string                                 `json:"title,omitempty"`
}

type TransactionDocumentUploadResult struct {
	TransactionDocuments []*TransactionDocument `json:"transaction_documents,omitempty"`
	Object               string                 `json:"object,omitempty"`
}

func (m TransactionDocumentUploadResult) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type TransactionDocumentUploads struct {
	Files   []http.File                  `json:"files,omitempty"`
	Uploads []*TransactionDocumentUpload `json:"uploads,omitempty"`
}

type TransactionDocumentsRestore struct {
	FolderId              string `json:"folder_id,omitempty"`
	TransactionDocumentId string `json:"transaction_document_id,omitempty"`
}

type TransactionDocumentsRestores struct {
	Restores []*TransactionDocumentsRestore `json:"restores,omitempty"`
}

type TransactionForm struct {
	Id     string `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	Object string `json:"object,omitempty"`
}

func (m TransactionForm) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type TransactionFormImport struct {
	FormId string `json:"form_id"`
	Title  string `json:"title,omitempty"`
}

type TransactionFormImports struct {
	FolderId string                   `json:"folder_id,omitempty"`
	Imports  []*TransactionFormImport `json:"imports"`
}

type TransactionFormLibrary struct {
	Id     string             `json:"id,omitempty"`
	Forms  []*TransactionForm `json:"forms,omitempty"`
	Title  string             `json:"title,omitempty"`
	Object string             `json:"object,omitempty"`
}

func (m TransactionFormLibrary) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type TransactionFormsResponse struct {
	Libraries []*TransactionFormLibrary `json:"libraries,omitempty"`
	Object    string                    `json:"object,omitempty"`
}

func (m TransactionFormsResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type TransactionMeta struct {
	IsLease *bool  `json:"is_lease,omitempty"`
	Title   string `json:"title,omitempty"`
}

type TransactionMetaUpdate struct {
	Data *TransactionMeta `json:"data,omitempty"`
}

type TransactionPackage struct {
	Id            string        `json:"id,omitempty"`
	Members       []*Member     `json:"members,omitempty"`
	PackageId     string        `json:"package_id,omitempty"`
	PackageKind   string        `json:"package_kind,omitempty"`
	PackageStatus string        `json:"package_status,omitempty"`
	PropertyInfo  *PropertyInfo `json:"property_info,omitempty"`
	Title         string        `json:"title,omitempty"`
	Transaction   *Transaction  `json:"transaction,omitempty"`
	Object        string        `json:"object,omitempty"`
}

func (m TransactionPackage) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type TransactionPackageList struct {
	Data       []TransactionPackage `json:"data"`
	ListObject string               `json:"list_object"`
	Object     string               `json:"object"`
	HasMore    bool                 `json:"has_more"`
}

func (m TransactionPackageList) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

func (m TransactionPackageList) NextPageParams() core.PageParams {
	if !m.HasMore {
		return nil
	}

	pageSize := len(m.Data)
	return &core.PageParamsImpl{
		StartingAfter: m.Data[pageSize-1].Id,
		Limit:         pageSize,
	}
}

type TransactionSelectedTemplate struct {
	TemplateId             string   `json:"template_id"`
	TransactionDocumentIds []string `json:"transaction_document_ids"`
}

type TransactionSelectedTemplates struct {
	SelectedTemplateDocuments []*TransactionSelectedTemplate `json:"selected_template_documents"`
	TemplateIds               []string                       `json:"template_ids"`
}

type UpdateArchivalStatusResponse struct {
	TransactionId string `json:"transaction_id,omitempty"`
	Object        string `json:"object,omitempty"`
}

func (m UpdateArchivalStatusResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type UpdateTransactionMetaResponse struct {
	TransactionId string `json:"transaction_id,omitempty"`
	Object        string `json:"object,omitempty"`
}

func (m UpdateTransactionMetaResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type UploadError struct {
	Error    string `json:"error"`
	Filename string `json:"filename"`
	Size     int    `json:"size,omitempty"`
	Object   string `json:"object,omitempty"`
}

func (m UploadError) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type UploadedBy struct {
	Name     string `json:"name,omitempty"`
	PersonId string `json:"person_id,omitempty"`
	Object   string `json:"object,omitempty"`
}

func (m UploadedBy) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type UploadsResponse struct {
	Result        *TransactionDocumentUploadResult `json:"result,omitempty"`
	TransactionId string                           `json:"transaction_id,omitempty"`
	Object        string                           `json:"object,omitempty"`
}

func (m UploadsResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type User struct {
	Id           string   `json:"id,omitempty"`
	AgentAddress *Address `json:"agent_address,omitempty"`
	Contact      *Contact `json:"contact,omitempty"`
	Uuid         string   `json:"uuid,omitempty"`
	Object       string   `json:"object,omitempty"`
}

func (m User) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type UserList struct {
	Data       []User `json:"data"`
	ListObject string `json:"list_object"`
	Object     string `json:"object"`
	HasMore    bool   `json:"has_more"`
}

func (m UserList) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

func (m UserList) NextPageParams() core.PageParams {
	if !m.HasMore {
		return nil
	}

	pageSize := len(m.Data)
	return &core.PageParamsImpl{
		StartingAfter: m.Data[pageSize-1].Id,
		Limit:         pageSize,
	}
}

type UserBillingInfo struct {
	StripeCustomerId string `json:"stripe_customer_id,omitempty"`
	Object           string `json:"object,omitempty"`
}

func (m UserBillingInfo) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type UserContact struct {
	UserContactId string `json:"user_contact_id,omitempty"`
}

type UserContactRecipient struct {
	Contact           *Contact       `json:"contact,omitempty"`
	Roles             []string       `json:"roles,omitempty"`
	UserContactId     string         `json:"user_contact_id,omitempty"`
	UserContactSource *ContactSource `json:"user_contact_source,omitempty"`
	Object            string         `json:"object,omitempty"`
}

func (m UserContactRecipient) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type UserManagementSchema struct {
	Email           string   `json:"email"`
	FirstName       string   `json:"first_name"`
	LastName        string   `json:"last_name"`
	LinkedSubjectId string   `json:"linked_subject_id"`
	MarketIds       []string `json:"market_ids,omitempty"`
	SubmarketIds    []string `json:"submarket_ids,omitempty"`
	UsState         string   `json:"us_state,omitempty"`
}

type Validation struct {
	Kind      string           `json:"kind,omitempty"`
	Signature *SignatureIntent `json:"signature,omitempty"`
}
