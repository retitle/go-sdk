package glide

import "strings"

type Address struct {
	City    string `json:"city"`
	State   string `json:"state"`
	Street  string `json:"street"`
	Unit    string `json:"unit"`
	ZipCode string `json:"zip_code"`
}

type Agent struct {
	CompanyLicenseNumber string `json:"company_license_number"`
	CompanyName          string `json:"company_name"`
	CompanyPhoneNumber   string `json:"company_phone_number"`
	LicenseNumber        string `json:"license_number"`
	LicenseState         string `json:"license_state"`
	NrdsNumber           string `json:"nrds_number"`
}

type Contact struct {
	Agent      Agent  `json:"agent"`
	CellPhone  string `json:"cell_phone"`
	Email      string `json:"email"`
	EntityName string `json:"entity_name"`
	EntityType string `json:"entity_type"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Title      string `json:"title"`
}

type CreateResponse struct {
	TransactionId string `json:"transaction_id"`
	Object        string `json:"object"`
}

func (m CreateResponse) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type FieldWriteDict struct {
	Fields TransactionFieldsWrite `json:"fields"`
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
	Fields TransactionFields `json:"fields"`
	Object string            `json:"object"`
}

func (m FieldsResponseResult) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}

type Folder struct {
	Id                   string                  `json:"id"`
	Kind                 string                  `json:"kind"`
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

type Party struct {
	Id          string      `json:"id"`
	Contact     Contact     `json:"contact"`
	Roles       []string    `json:"roles"`
	Transaction Transaction `json:"transaction"`
	Object      string      `json:"object"`
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
	Body                string   `json:"body"`
	Contact             Contact  `json:"contact"`
	Invite              bool     `json:"invite"`
	InviteRestrictions  []string `json:"invite_restrictions"`
	Roles               []string `json:"roles"`
	Subject             string   `json:"subject"`
	SuppressInviteEmail bool     `json:"suppress_invite_email"`
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

type PartyPatch struct {
	Contact Contact  `json:"contact"`
	PartyId string   `json:"party_id"`
	Roles   []string `json:"roles"`
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

type Transaction struct {
	Id                   string                  `json:"id"`
	Address              Address                 `json:"address"`
	Archived             bool                    `json:"archived"`
	Fields               TransactionFields       `json:"fields"`
	Folders              FolderList              `json:"folders"`
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
	Id          string      `json:"id"`
	Folder      Folder      `json:"folder"`
	FolderKind  string      `json:"folder_kind"`
	Order       int         `json:"order"`
	Title       string      `json:"title"`
	Transaction Transaction `json:"transaction"`
	Url         string      `json:"url"`
	Object      string      `json:"object"`
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
	FolderId string `json:"folder_id"`
	Order    int    `json:"order"`
	TdId     string `json:"td_id"`
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
	Id      string  `json:"id"`
	Contact Contact `json:"contact"`
	Object  string  `json:"object"`
}

func (m User) IsRef() bool {
	return strings.HasPrefix(m.Object, "/ref/")
}
