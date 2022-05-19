package glide

import (
	"net/http"
	"strings"
)

type Address struct {
	City    string `json:"city"`
	State   string `json:"state"`
	Street  string `json:"street"`
	Unit    string `json:"unit"`
	ZipCode string `json:"zip_code"`
	Object  string `json:"object"`
}

func (m Address) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type Agent struct {
	CompanyLicenseNumber string `json:"company_license_number"`
	CompanyName          string `json:"company_name"`
	CompanyPhoneNumber   string `json:"company_phone_number"`
	LicenseNumber        string `json:"license_number"`
	LicenseState         string `json:"license_state"`
	NrdsNumber           string `json:"nrds_number"`
	Object               string `json:"object"`
}

func (m Agent) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type AgentRequest struct {
	CompanyLicenseNumber string `json:"company_license_number"`
	CompanyName          string `json:"company_name"`
	CompanyPhoneNumber   string `json:"company_phone_number"`
	LicenseNumber        string `json:"license_number"`
	LicenseState         string `json:"license_state"`
	NrdsNumber           string `json:"nrds_number"`
}

type Contact struct {
	Id              string  `json:"id"`
	Address         Address `json:"address"`
	Agent           Agent   `json:"agent"`
	AvatarUrl       string  `json:"avatar_url"`
	BrandLogoUrl    string  `json:"brand_logo_url"`
	CellPhone       string  `json:"cell_phone"`
	Email           string  `json:"email"`
	EntityName      string  `json:"entity_name"`
	EntityType      string  `json:"entity_type"`
	FirstName       string  `json:"first_name"`
	LastName        string  `json:"last_name"`
	PersonalWebsite string  `json:"personal_website"`
	Title           string  `json:"title"`
	Object          string  `json:"object"`
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

func (m ContactList) NextPageParams() *PageParams {
	if !m.HasMore {
		return nil
	}

	pageSize := len(m.Data)
	return &PageParams{
		StartingAfter: m.Data[pageSize-1].Id,
		Limit:         pageSize,
	}
}

type ContactCreate struct {
	Contact ContactRequest `json:"contact"`
}

type ContactCreateResponse struct {
	Contact Contact `json:"contact"`
	Object  string  `json:"object"`
}

func (m ContactCreateResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type ContactRequest struct {
	Address         map[string]interface{} `json:"address"`
	Agent           AgentRequest           `json:"agent"`
	AvatarUrl       string                 `json:"avatar_url"`
	BrandLogoUrl    string                 `json:"brand_logo_url"`
	CellPhone       string                 `json:"cell_phone"`
	Email           string                 `json:"email"`
	EntityName      string                 `json:"entity_name"`
	EntityType      string                 `json:"entity_type"`
	FirstName       string                 `json:"first_name"`
	LastName        string                 `json:"last_name"`
	PersonalWebsite string                 `json:"personal_website"`
	Title           string                 `json:"title"`
}

type ContactSource struct {
	Origin string `json:"origin"`
	Object string `json:"object"`
}

func (m ContactSource) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type ContactSourceRequest struct {
	Origin string `json:"origin"`
}

type ContactUpdate struct {
	Contact ContactRequest `json:"contact"`
	Roles   []string       `json:"roles"`
}

type ContactUpdateResponse struct {
	Contact Contact `json:"contact"`
	Id      string  `json:"id_"`
	Object  string  `json:"object"`
}

func (m ContactUpdateResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type CreateResponse struct {
	TransactionId string `json:"transaction_id"`
	Object        string `json:"object"`
}

func (m CreateResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type DocumentSplitAsyncResponse struct {
	ReqId       string                             `json:"req_id"`
	Suggestions map[string]DocumentSplitSuggestion `json:"suggestions"`
	Object      string                             `json:"object"`
}

func (m DocumentSplitAsyncResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type DocumentSplitResponse struct {
	ReqId  string                     `json:"req_id"`
	Result DocumentSplitAsyncResponse `json:"result"`
	Object string                     `json:"object"`
}

func (m DocumentSplitResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type DocumentSplitSchema struct {
	Files   []http.File      `json:"files"`
	ReState string           `json:"re_state"`
	ReqId   string           `json:"req_id"`
	Uploads []DocumentUpload `json:"uploads"`
}

type DocumentSplitSuggestion struct {
	EndPage      int    `json:"end_page"`
	Filename     string `json:"filename"`
	FormId       string `json:"form_id"`
	FormSeriesId string `json:"form_series_id"`
	StartPage    int    `json:"start_page"`
	Object       string `json:"object"`
}

func (m DocumentSplitSuggestion) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type DocumentUpload struct {
	Title string `json:"title"`
}

type DocumentZone struct {
	Id               string                 `json:"id"`
	FormId           string                 `json:"form_id"`
	Kind             string                 `json:"kind"`
	Name             string                 `json:"name"`
	OriginalLocation []DocumentZoneLocation `json:"original_location"`
	Page             int                    `json:"page"`
	Vertices         []DocumentZoneVertex   `json:"vertices"`
	Object           string                 `json:"object"`
}

func (m DocumentZone) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type DocumentZoneLocation struct {
	XMax   float64 `json:"x_max"`
	XMin   float64 `json:"x_min"`
	YMax   float64 `json:"y_max"`
	YMin   float64 `json:"y_min"`
	Object string  `json:"object"`
}

func (m DocumentZoneLocation) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type DocumentZoneVertex struct {
	X      int    `json:"x"`
	Y      int    `json:"y"`
	Object string `json:"object"`
}

func (m DocumentZoneVertex) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type Field struct {
	Timestamp int                    `json:"timestamp"`
	Value     map[string]interface{} `json:"value"`
	Object    string                 `json:"object"`
}

func (m Field) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type FieldOutOfDateDetail struct {
	ControlTimestamp int    `json:"control_timestamp"`
	Timestamp        int    `json:"timestamp"`
	Object           string `json:"object"`
}

func (m FieldOutOfDateDetail) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type FieldResponse struct {
	Timestamp int                    `json:"timestamp"`
	Value     map[string]interface{} `json:"value"`
	Object    string                 `json:"object"`
}

func (m FieldResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type FieldResponseWarnings struct {
	OutOfDateFields map[string]FieldOutOfDateDetail `json:"out_of_date_fields"`
	Object          string                          `json:"object"`
}

func (m FieldResponseWarnings) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type FieldWrite struct {
	ControlTimestamp int                    `json:"control_timestamp"`
	Value            map[string]interface{} `json:"value"`
}

type FieldWriteDict struct {
	ControlPolicy string                 `json:"control_policy"`
	Fields        TransactionFieldsWrite `json:"fields"`
}

type FieldsResponse struct {
	Result        FieldsResponseResult `json:"result"`
	TransactionId string               `json:"transaction_id"`
	Object        string               `json:"object"`
}

func (m FieldsResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type FieldsResponseResult struct {
	Fields   TransactionFields     `json:"fields"`
	Warnings FieldResponseWarnings `json:"warnings"`
	Object   string                `json:"object"`
}

func (m FieldsResponseResult) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type Folder struct {
	Id                   string                  `json:"id"`
	Kind                 string                  `json:"kind"`
	LastModified         int                     `json:"last_modified"`
	Title                string                  `json:"title"`
	TransactionDocuments TransactionDocumentList `json:"transaction_documents"`
	Object               string                  `json:"object"`
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

func (m FolderList) NextPageParams() *PageParams {
	if !m.HasMore {
		return nil
	}

	pageSize := len(m.Data)
	return &PageParams{
		StartingAfter: m.Data[pageSize-1].Id,
		Limit:         pageSize,
	}
}

type FolderCreate struct {
	Title string `json:"title"`
}

type FolderCreates struct {
	Creates []FolderCreate `json:"creates"`
}

type FolderCreatesResponse struct {
	Result        FolderCreatesResponseResult `json:"result"`
	TransactionId string                      `json:"transaction_id"`
	Object        string                      `json:"object"`
}

func (m FolderCreatesResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type FolderCreatesResponseResult struct {
	FolderIds []string `json:"folder_ids"`
	Object    string   `json:"object"`
}

func (m FolderCreatesResponseResult) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type FolderRename struct {
	FolderId string `json:"folder_id"`
	Title    string `json:"title"`
}

type FolderRenames struct {
	Renames []FolderRename `json:"renames"`
}

type FolderRenamesResponse struct {
	TransactionId string `json:"transaction_id"`
	Object        string `json:"object"`
}

func (m FolderRenamesResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type FormImportsResponse struct {
	TransactionId string `json:"transaction_id"`
	Object        string `json:"object"`
}

func (m FormImportsResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type ItemDeletes struct {
	Ids []string `json:"ids"`
}

type ItemDeletesResponse struct {
	TransactionId string `json:"transaction_id"`
	Object        string `json:"object"`
}

func (m ItemDeletesResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type LinkListingInfo struct {
	MlsKind   string `json:"mls_kind"`
	MlsNumber string `json:"mls_number"`
}

type LinkListingInfoResponse struct {
	TransactionId string `json:"transaction_id"`
	Object        string `json:"object"`
}

func (m LinkListingInfoResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type Listing struct {
	Id                      string   `json:"id"`
	Address                 Location `json:"address"`
	Bath                    float64  `json:"bath"`
	BathFull                float64  `json:"bath_full"`
	BathHalf                float64  `json:"bath_half"`
	BathOneQuarter          float64  `json:"bath_one_quarter"`
	BathThreeQuarter        float64  `json:"bath_three_quarter"`
	Bed                     float64  `json:"bed"`
	CloseDate               string   `json:"close_date"`
	ClosePrice              float64  `json:"close_price"`
	Dom                     float64  `json:"dom"`
	ListingDate             string   `json:"listing_date"`
	ListingPrice            float64  `json:"listing_price"`
	ListingType             string   `json:"listing_type"`
	MediaUrls               []string `json:"media_urls"`
	MlsKind                 string   `json:"mls_kind"`
	MlsNumber               string   `json:"mls_number"`
	MlsStatus               string   `json:"mls_status"`
	OriginalListPrice       float64  `json:"original_list_price"`
	PropertyType            string   `json:"property_type"`
	StatusDate              string   `json:"status_date"`
	UsedInActiveTransaction bool     `json:"used_in_active_transaction"`
	YearBuilt               string   `json:"year_built"`
	Object                  string   `json:"object"`
}

func (m Listing) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type ListingList struct {
	Data       []Listing `json:"data"`
	ListObject string    `json:"list_object"`
	Object     string    `json:"object"`
	HasMore    bool      `json:"has_more"`
}

func (m ListingList) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

func (m ListingList) NextPageParams() *PageParams {
	if !m.HasMore {
		return nil
	}

	pageSize := len(m.Data)
	return &PageParams{
		StartingAfter: m.Data[pageSize-1].Id,
		Limit:         pageSize,
	}
}

type Location struct {
	AgentAddress  string `json:"agent_address"`
	City          string `json:"city"`
	County        string `json:"county"`
	PrettyAddress string `json:"pretty_address"`
	State         string `json:"state"`
	Street        string `json:"street"`
	StreetNumber  string `json:"street_number"`
	StreetType    string `json:"street_type"`
	UnitNumber    string `json:"unit_number"`
	UnitType      string `json:"unit_type"`
	ZipCode       string `json:"zip_code"`
	Object        string `json:"object"`
}

func (m Location) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type Notification struct {
	Bcc              []string               `json:"bcc"`
	Cc               []string               `json:"cc"`
	Context          map[string]interface{} `json:"context"`
	IncludeSignature bool                   `json:"include_signature"`
	Recipients       []string               `json:"recipients"`
	SeparateEmails   bool                   `json:"separate_emails"`
	Template         string                 `json:"template"`
}

type NotificationResponse struct {
	Results []string `json:"results"`
	Object  string   `json:"object"`
}

func (m NotificationResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type Party struct {
	Id                string        `json:"id"`
	Contact           Contact       `json:"contact"`
	Roles             []string      `json:"roles"`
	Transaction       Transaction   `json:"transaction"`
	UserContactId     string        `json:"user_contact_id"`
	UserContactSource ContactSource `json:"user_contact_source"`
	UserId            string        `json:"user_id"`
	Object            string        `json:"object"`
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

func (m PartyList) NextPageParams() *PageParams {
	if !m.HasMore {
		return nil
	}

	pageSize := len(m.Data)
	return &PageParams{
		StartingAfter: m.Data[pageSize-1].Id,
		Limit:         pageSize,
	}
}

type PartyCreate struct {
	Body                string               `json:"body"`
	Contact             ContactRequest       `json:"contact"`
	Invite              bool                 `json:"invite"`
	InviteRestrictions  []string             `json:"invite_restrictions"`
	Roles               []string             `json:"roles"`
	Subject             string               `json:"subject"`
	SuppressInviteEmail bool                 `json:"suppress_invite_email"`
	UserContactId       string               `json:"user_contact_id"`
	UserContactSource   ContactSourceRequest `json:"user_contact_source"`
}

type PartyCreates struct {
	Creates []PartyCreate `json:"creates"`
}

type PartyCreatesResponse struct {
	TransactionId string `json:"transaction_id"`
	Object        string `json:"object"`
}

func (m PartyCreatesResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type PartyInvite struct {
	Body                string   `json:"body"`
	InviteRestrictions  []string `json:"invite_restrictions"`
	PartyId             string   `json:"party_id"`
	Subject             string   `json:"subject"`
	SuppressInviteEmail bool     `json:"suppress_invite_email"`
}

type PartyInvites struct {
	Invites []PartyInvite `json:"invites"`
}

type PartyInvitesResponse struct {
	TransactionId string `json:"transaction_id"`
	Object        string `json:"object"`
}

func (m PartyInvitesResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type PartyPatch struct {
	Contact ContactRequest `json:"contact"`
	PartyId string         `json:"party_id"`
	Roles   []string       `json:"roles"`
}

type PartyPatches struct {
	Patches []PartyPatch `json:"patches"`
}

type PartyPatchesResponse struct {
	TransactionId string `json:"transaction_id"`
	Object        string `json:"object"`
}

func (m PartyPatchesResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type PartyRemove struct {
	PartyId string `json:"party_id"`
}

type PartyRemoves struct {
	Removes []PartyRemove `json:"removes"`
}

type PartyRemovesResponse struct {
	TransactionId string `json:"transaction_id"`
	Object        string `json:"object"`
}

func (m PartyRemovesResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type PartyRoles struct {
	Data   []string `json:"data"`
	Object string   `json:"object"`
}

func (m PartyRoles) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type SignatureDetectionAnalysisResult struct {
	DocumentZone DocumentZone `json:"document_zone"`
	Score        float64      `json:"score"`
	Object       string       `json:"object"`
}

func (m SignatureDetectionAnalysisResult) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type SignatureDetectionAsyncResponse struct {
	ReqId      string                                      `json:"req_id"`
	Signatures map[string]SignatureDetectionAnalysisResult `json:"signatures"`
	Object     string                                      `json:"object"`
}

func (m SignatureDetectionAsyncResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type SignatureDetectionResponse struct {
	ReqId  string                          `json:"req_id"`
	Result SignatureDetectionAsyncResponse `json:"result"`
	Object string                          `json:"object"`
}

func (m SignatureDetectionResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type SignatureDetectionSchema struct {
	Files   []http.File      `json:"files"`
	Uploads []DocumentUpload `json:"uploads"`
}

type Transaction struct {
	Id                   string                  `json:"id"`
	Address              Address                 `json:"address"`
	Archived             bool                    `json:"archived"`
	Fields               TransactionFields       `json:"fields"`
	Folders              FolderList              `json:"folders"`
	IngestDocumentsEmail string                  `json:"ingest_documents_email"`
	IsLease              bool                    `json:"is_lease"`
	Parties              PartyList               `json:"parties"`
	ReState              string                  `json:"re_state"`
	Side                 string                  `json:"side"`
	Stage                string                  `json:"stage"`
	Title                string                  `json:"title"`
	TransactionDocuments TransactionDocumentList `json:"transaction_documents"`
	Object               string                  `json:"object"`
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

func (m TransactionList) NextPageParams() *PageParams {
	if !m.HasMore {
		return nil
	}

	pageSize := len(m.Data)
	return &PageParams{
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
	Archived bool `json:"archived"`
}

type TransactionCreate struct {
	AdditionalParties []PartyCreate `json:"additional_parties"`
	Address           Address       `json:"address"`
	CreatorRoles      []string      `json:"creator_roles"`
	IsLease           bool          `json:"is_lease"`
	ReState           string        `json:"re_state"`
	Stage             string        `json:"stage"`
	Title             string        `json:"title"`
}

type TransactionDocument struct {
	Id           string      `json:"id"`
	Folder       Folder      `json:"folder"`
	FolderKind   string      `json:"folder_kind"`
	LastModified int         `json:"last_modified"`
	Order        int         `json:"order"`
	Title        string      `json:"title"`
	Transaction  Transaction `json:"transaction"`
	Url          string      `json:"url"`
	Object       string      `json:"object"`
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

func (m TransactionDocumentList) NextPageParams() *PageParams {
	if !m.HasMore {
		return nil
	}

	pageSize := len(m.Data)
	return &PageParams{
		StartingAfter: m.Data[pageSize-1].Id,
		Limit:         pageSize,
	}
}

type TransactionDocumentAssignment struct {
	FolderId              string `json:"folder_id"`
	Order                 int    `json:"order"`
	TransactionDocumentId string `json:"transaction_document_id"`
}

type TransactionDocumentAssignments struct {
	Assignments []TransactionDocumentAssignment `json:"assignments"`
}

type TransactionDocumentAssignmentsResponse struct {
	TransactionId string `json:"transaction_id"`
	Object        string `json:"object"`
}

func (m TransactionDocumentAssignmentsResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type TransactionDocumentRename struct {
	Title                 string `json:"title"`
	TransactionDocumentId string `json:"transaction_document_id"`
}

type TransactionDocumentRenames struct {
	Renames []TransactionDocumentRename `json:"renames"`
}

type TransactionDocumentRenamesResponse struct {
	TransactionId string `json:"transaction_id"`
	Object        string `json:"object"`
}

func (m TransactionDocumentRenamesResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type TransactionDocumentRestoresResponse struct {
	TransactionId string `json:"transaction_id"`
	Object        string `json:"object"`
}

func (m TransactionDocumentRestoresResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type TransactionDocumentTrashes struct {
	TransactionDocumentIds []string `json:"transaction_document_ids"`
}

type TransactionDocumentTrashesResponse struct {
	TransactionId string `json:"transaction_id"`
	Object        string `json:"object"`
}

func (m TransactionDocumentTrashesResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type TransactionDocumentUpload struct {
	FolderId string `json:"folder_id"`
	Title    string `json:"title"`
}

type TransactionDocumentUploads struct {
	Files   []http.File                 `json:"files"`
	Uploads []TransactionDocumentUpload `json:"uploads"`
}

type TransactionDocumentsRestore struct {
	FolderId              string `json:"folder_id"`
	TransactionDocumentId string `json:"transaction_document_id"`
}

type TransactionDocumentsRestores struct {
	Restores []TransactionDocumentsRestore `json:"restores"`
}

type TransactionFormImport struct {
	FormId string `json:"form_id"`
	Title  string `json:"title"`
}

type TransactionFormImports struct {
	FolderId string                  `json:"folder_id"`
	Imports  []TransactionFormImport `json:"imports"`
}

type TransactionMeta struct {
	IsLease bool   `json:"is_lease"`
	Title   string `json:"title"`
}

type TransactionMetaUpdate struct {
	Data TransactionMeta `json:"data"`
}

type UpdateArchivalStatusResponse struct {
	TransactionId string `json:"transaction_id"`
	Object        string `json:"object"`
}

func (m UpdateArchivalStatusResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type UpdateTransactionMetaResponse struct {
	TransactionId string `json:"transaction_id"`
	Object        string `json:"object"`
}

func (m UpdateTransactionMetaResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type UploadsResponse struct {
	TransactionId string `json:"transaction_id"`
	Object        string `json:"object"`
}

func (m UploadsResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type User struct {
	Id           string  `json:"id"`
	AgentAddress Address `json:"agent_address"`
	Contact      Contact `json:"contact"`
	Object       string  `json:"object"`
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

func (m UserList) NextPageParams() *PageParams {
	if !m.HasMore {
		return nil
	}

	pageSize := len(m.Data)
	return &PageParams{
		StartingAfter: m.Data[pageSize-1].Id,
		Limit:         pageSize,
	}
}

type UserBillingInfo struct {
	StripeCustomerId string `json:"stripe_customer_id"`
	Object           string `json:"object"`
}

func (m UserBillingInfo) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type UserManagementSchema struct {
	Email           string `json:"email"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	LinkedSubjectId string `json:"linked_subject_id"`
	UsState         string `json:"us_state"`
}
