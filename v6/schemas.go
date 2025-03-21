package glide

import (
	"bytes"
	"io"
	"net/http"
	"strings"

	"github.com/retitle/go-sdk/v6/core"
)

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
	FirstName       string         `json:"first_name"`
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
	FirstName       string         `json:"first_name"`
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
	Id       string `json:"id"`
	FileName string `json:"file_name"`
	Object   string `json:"object,omitempty"`
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
	Files     []http.File         `json:"files,omitempty"`
	FilesMeta []*DocumentFileMeta `json:"files_meta,omitempty"`
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

type Envelope struct {
	Id        string `json:"id,omitempty"`
	CreatedAt int    `json:"created_at,omitempty"`
	Creator   string `json:"creator,omitempty"`
	Status    string `json:"status,omitempty"`
	Title     string `json:"title,omitempty"`
	Uuid      string `json:"uuid,omitempty"`
	Object    string `json:"object,omitempty"`
}

func (m Envelope) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
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

type Folder struct {
	Id                   string                   `json:"id,omitempty"`
	Can                  map[string]*bool         `json:"can,omitempty"`
	Kind                 string                   `json:"kind,omitempty"`
	LastModified         int                      `json:"last_modified,omitempty"`
	OrderIndex           int                      `json:"order_index,omitempty"`
	PropertyInfo         *PropertyInfo            `json:"property_info,omitempty"`
	Title                string                   `json:"title,omitempty"`
	TransactionDocuments *TransactionDocumentList `json:"transaction_documents,omitempty"`
	TransactionPackage   *TransactionPackage      `json:"transaction_package,omitempty"`
	Object               string                   `json:"object,omitempty"`
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
	Title string `json:"title,omitempty"`
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

type GlideForm struct {
	GlideFormSeriesId int      `json:"glide_form_series_id"`
	Tags              []string `json:"tags,omitempty"`
	Title             string   `json:"title"`
	Object            string   `json:"object,omitempty"`
}

func (m GlideForm) IsRef() bool {
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
	Data   []*Party `json:"data,omitempty"`
	Object string   `json:"object,omitempty"`
}

func (m OfferPartiesResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type OffersResponse struct {
	Data   string `json:"data,omitempty"`
	Object string `json:"object,omitempty"`
}

func (m OffersResponse) IsRef() bool {
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
	TransactionId string `json:"transaction_id,omitempty"`
	Object        string `json:"object,omitempty"`
}

func (m PartyCreatesResponse) IsRef() bool {
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
	TransactionId string `json:"transaction_id,omitempty"`
	Object        string `json:"object,omitempty"`
}

func (m PartyPatchesResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type PartyRecipient struct {
	Id      string   `json:"id,omitempty"`
	Contact *Contact `json:"contact,omitempty"`
	Roles   []string `json:"roles,omitempty"`
	Vers    string   `json:"vers,omitempty"`
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
}

type ReorderFoldersResponse struct {
	TransactionId string `json:"transaction_id,omitempty"`
	Object        string `json:"object,omitempty"`
}

func (m ReorderFoldersResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
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

type SignatureRequest struct {
	Id               string                       `json:"id,omitempty"`
	CompletedAt      int                          `json:"completed_at,omitempty"`
	CreatedAt        int                          `json:"created_at,omitempty"`
	Documents        []*SignatureRequestDocument  `json:"documents,omitempty"`
	EnvelopeId       string                       `json:"envelope_id,omitempty"`
	FlowId           string                       `json:"flow_id,omitempty"`
	IsArchived       *bool                        `json:"is_archived,omitempty"`
	Provider         string                       `json:"provider,omitempty"`
	Recipients       []*SignatureRequestRecipient `json:"recipients,omitempty"`
	RevisionFlowId   string                       `json:"revision_flow_id,omitempty"`
	SentAt           int                          `json:"sent_at,omitempty"`
	SignNowUrl       string                       `json:"sign_now_url,omitempty"`
	Status           string                       `json:"status,omitempty"`
	Title            string                       `json:"title,omitempty"`
	TransactionId    string                       `json:"transaction_id,omitempty"`
	TransactionTitle string                       `json:"transaction_title,omitempty"`
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

type SignatureRequestDocument struct {
	Id     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Seq    int    `json:"seq,omitempty"`
	Object string `json:"object,omitempty"`
}

func (m SignatureRequestDocument) IsRef() bool {
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

type SignatureRequestFlowRequest struct {
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

type SignatureRequestRecipient struct {
	Id            string `json:"id,omitempty"`
	Email         string `json:"email,omitempty"`
	Order         int    `json:"order,omitempty"`
	RecipientRole string `json:"recipient_role,omitempty"`
	Status        string `json:"status,omitempty"`
	UserName      string `json:"user_name,omitempty"`
	Object        string `json:"object,omitempty"`
}

func (m SignatureRequestRecipient) IsRef() bool {
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
	FolderId string `json:"folder_id,omitempty"`
	MimeType string `json:"mime_type,omitempty"`
	Title    string `json:"title,omitempty"`
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

type TransactionFormImport struct {
	FormId string `json:"form_id"`
	Title  string `json:"title,omitempty"`
}

type TransactionFormImports struct {
	FolderId string                   `json:"folder_id,omitempty"`
	Imports  []*TransactionFormImport `json:"imports"`
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

type UserContactRecipient struct {
	Contact           *Contact       `json:"contact,omitempty"`
	Roles             []string       `json:"roles,omitempty"`
	UserContactId     string         `json:"user_contact_id,omitempty"`
	UserContactSource *ContactSource `json:"user_contact_source,omitempty"`
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
