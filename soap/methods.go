package soap

// SdmSoap is a struct that implements Client for CA Service Desk SOAP API
type SdmSoap struct {
	SdmSoapInterface
}

// Namespace was auto-generated from WSDL.
var Namespace = "http://www.ca.com/UnicenterServicePlus/ServiceDesk"

// SoapWebService was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type SoapWebService interface {
	// AddAssetLog was auto-generated from WSDL.
	AddAssetLog(AddAssetLog *AddAssetLog) (*AddAssetLogResponse, error)

	// AddBookmark was auto-generated from WSDL.
	AddBookmark(AddBookmark *AddBookmark) (*AddBookmarkResponse, error)

	// AddComment was auto-generated from WSDL.
	AddComment(AddComment *AddComment) (*AddCommentResponse, error)

	// AddMemberToGroup was auto-generated from WSDL.
	AddMemberToGroup(AddMemberToGroup *AddMemberToGroup) (*AddMemberToGroupResponse, error)

	// AttachChangeToRequest was auto-generated from WSDL.
	AttachChangeToRequest(AttachChangeToRequest *AttachChangeToRequest) (*AttachChangeToRequestResponse, error)

	// AttachURLLink was auto-generated from WSDL.
	AttachURLLink(AttachURLLink *AttachURLLink) (*AttachURLLinkResponse, error)

	// AttachURLLinkToTicket was auto-generated from WSDL.
	AttachURLLinkToTicket(AttachURLLinkToTicket *AttachURLLinkToTicket) (*AttachURLLinkToTicketResponse, error)

	// AttmntFolderLinkCount was auto-generated from WSDL.
	AttmntFolderLinkCount(AttmntFolderLinkCount *AttmntFolderLinkCount) (*AttmntFolderLinkCountResponse, error)

	// CallServerMethod was auto-generated from WSDL.
	CallServerMethod(CallServerMethod *CallServerMethod) (*CallServerMethodResponse, error)

	// ChangeStatus was auto-generated from WSDL.
	ChangeStatus(ChangeStatus *ChangeStatus) (*ChangeStatusResponse, error)

	// ClearNotification was auto-generated from WSDL.
	ClearNotification(ClearNotification *ClearNotification) (*ClearNotificationResponse, error)

	// CloseTicket was auto-generated from WSDL.
	CloseTicket(CloseTicket *CloseTicket) (*CloseTicketResponse, error)

	// CreateActivityLog was auto-generated from WSDL.
	CreateActivityLog(CreateActivityLog *CreateActivityLog) (*CreateActivityLogResponse, error)

	// CreateAsset was auto-generated from WSDL.
	CreateAsset(CreateAsset *CreateAsset) (*CreateAssetResponse, error)

	// CreateAssetParentChildRelationship was auto-generated from WSDL.
	CreateAssetParentChildRelationship(CreateAssetParentChildRelationship *CreateAssetParentChildRelationship) (*CreateAssetParentChildRelationshipResponse, error)

	// CreateAttachment was auto-generated from WSDL.
	CreateAttachment(CreateAttachment *CreateAttachment) (*CreateAttachmentResponse, error)

	// CreateAttmnt was auto-generated from WSDL.
	CreateAttmnt(CreateAttmnt *CreateAttmnt) (*CreateAttmntResponse, error)

	// CreateChangeOrder was auto-generated from WSDL.
	CreateChangeOrder(CreateChangeOrder *CreateChangeOrder) (*CreateChangeOrderResponse, error)

	// CreateDocument was auto-generated from WSDL.
	CreateDocument(CreateDocument *CreateDocument) (*CreateDocumentResponse, error)

	// CreateFolder was auto-generated from WSDL.
	CreateFolder(CreateFolder *CreateFolder) (*CreateFolderResponse, error)

	// CreateIssue was auto-generated from WSDL.
	CreateIssue(CreateIssue *CreateIssue) (*CreateIssueResponse, error)

	// CreateLrelRelationships was auto-generated from WSDL.
	CreateLrelRelationships(CreateLrelRelationships *CreateLrelRelationships) (*CreateLrelRelationshipsResponse, error)

	// CreateObject was auto-generated from WSDL.
	CreateObject(CreateObject *CreateObject) (*CreateObjectResponse, error)

	// CreateQuickTicket was auto-generated from WSDL.
	CreateQuickTicket(CreateQuickTicket *CreateQuickTicket) (*CreateQuickTicketResponse, error)

	// CreateRequest was auto-generated from WSDL.
	CreateRequest(CreateRequest *CreateRequest) (*CreateRequestResponse, error)

	// CreateTicket was auto-generated from WSDL.
	CreateTicket(CreateTicket *CreateTicket) (*CreateTicketResponse, error)

	// CreateWorkFlowTask was auto-generated from WSDL.
	CreateWorkFlowTask(CreateWorkFlowTask *CreateWorkFlowTask) (*CreateWorkFlowTaskResponse, error)

	// DeleteBookmark was auto-generated from WSDL.
	DeleteBookmark(DeleteBookmark *DeleteBookmark) (*DeleteBookmarkResponse, error)

	// DeleteComment was auto-generated from WSDL.
	DeleteComment(DeleteComment *DeleteComment) (*DeleteCommentResponse, error)

	// DeleteDocument was auto-generated from WSDL.
	DeleteDocument(DeleteDocument *DeleteDocument) (*DeleteDocumentResponse, error)

	// DeleteWorkFlowTask was auto-generated from WSDL.
	DeleteWorkFlowTask(DeleteWorkFlowTask *DeleteWorkFlowTask) (*DeleteWorkFlowTaskResponse, error)

	// DetachChangeFromRequest was auto-generated from WSDL.
	DetachChangeFromRequest(DetachChangeFromRequest *DetachChangeFromRequest) (*DetachChangeFromRequestResponse, error)

	// DoQuery was auto-generated from WSDL.
	DoQuery(DoQuery *DoQuery) (*DoQueryResponse, error)

	// DoSelect was auto-generated from WSDL.
	DoSelect(DoSelect *DoSelect) (*DoSelectResponse, error)

	// DoSelectKD was auto-generated from WSDL.
	DoSelectKD(DoSelectKD *DoSelectKD) (*DoSelectKDResponse, error)

	// Escalate was auto-generated from WSDL.
	Escalate(Escalate *Escalate) (*EscalateResponse, error)

	// Faq was auto-generated from WSDL.
	Faq(Faq *Faq) (*FaqResponse, error)

	// FindContacts was auto-generated from WSDL.
	FindContacts(FindContacts *FindContacts) (*FindContactsResponse, error)

	// FreeListHandles was auto-generated from WSDL.
	FreeListHandles(FreeListHandles *FreeListHandles) (*FreeListHandlesResponse, error)

	// GetAccessTypeForContact was auto-generated from WSDL.
	GetAccessTypeForContact(GetAccessTypeForContact *GetAccessTypeForContact) (*GetAccessTypeForContactResponse, error)

	// GetArtifact was auto-generated from WSDL.
	GetArtifact(GetArtifact *GetArtifact) (*GetArtifactResponse, error)

	// GetAssetExtensionInformation was auto-generated from WSDL.
	GetAssetExtensionInformation(GetAssetExtensionInformation *GetAssetExtensionInformation) (*GetAssetExtensionInformationResponse, error)

	// GetAttmntInfo was auto-generated from WSDL.
	GetAttmntInfo(GetAttmntInfo *GetAttmntInfo) (*GetAttmntInfoResponse, error)

	// GetAttmntList was auto-generated from WSDL.
	GetAttmntList(GetAttmntList *GetAttmntList) (*GetAttmntListResponse, error)

	// GetAttmntListPerKD was auto-generated from WSDL.
	GetAttmntListPerKD(GetAttmntListPerKD *GetAttmntListPerKD) (*GetAttmntListPerKDResponse, error)

	// GetBookmarks was auto-generated from WSDL.
	GetBookmarks(GetBookmarks *GetBookmarks) (*GetBookmarksResponse, error)

	// GetBopsid was auto-generated from WSDL.
	GetBopsid(GetBopsid *GetBopsid) (*GetBopsidResponse, error)

	// GetCategory was auto-generated from WSDL.
	GetCategory(GetCategory *GetCategory) (*GetCategoryResponse, error)

	// GetComments was auto-generated from WSDL.
	GetComments(GetComments *GetComments) (*GetCommentsResponse, error)

	// GetConfigurationMode was auto-generated from WSDL.
	GetConfigurationMode(GetConfigurationMode *GetConfigurationMode) (*GetConfigurationModeResponse, error)

	// GetContact was auto-generated from WSDL.
	GetContact(GetContact *GetContact) (*GetContactResponse, error)

	// GetDecisionTrees was auto-generated from WSDL.
	GetDecisionTrees(GetDecisionTrees *GetDecisionTrees) (*GetDecisionTreesResponse, error)

	// GetDependentAttrControls was auto-generated from WSDL.
	GetDependentAttrControls(GetDependentAttrControls *GetDependentAttrControls) (*GetDependentAttrControlsResponse, error)

	// GetDocument was auto-generated from WSDL.
	GetDocument(GetDocument *GetDocument) (*GetDocumentResponse, error)

	// GetDocumentTypes was auto-generated from WSDL.
	GetDocumentTypes(GetDocumentTypes *GetDocumentTypes) (*GetDocumentTypesResponse, error)

	// GetDocumentsByIDs was auto-generated from WSDL.
	GetDocumentsByIDs(GetDocumentsByIDs *GetDocumentsByIDs) (*GetDocumentsByIDsResponse, error)

	// GetFolderInfo was auto-generated from WSDL.
	GetFolderInfo(GetFolderInfo *GetFolderInfo) (*GetFolderInfoResponse, error)

	// GetFolderList was auto-generated from WSDL.
	GetFolderList(GetFolderList *GetFolderList) (*GetFolderListResponse, error)

	// GetGroupMemberListValues was auto-generated from WSDL.
	GetGroupMemberListValues(GetGroupMemberListValues *GetGroupMemberListValues) (*GetGroupMemberListValuesResponse, error)

	// GetHandleForUserid was auto-generated from WSDL.
	GetHandleForUserid(GetHandleForUserid *GetHandleForUserid) (*GetHandleForUseridResponse, error)

	// GetKDListPerAttmnt was auto-generated from WSDL.
	GetKDListPerAttmnt(GetKDListPerAttmnt *GetKDListPerAttmnt) (*GetKDListPerAttmntResponse, error)

	// GetListValues was auto-generated from WSDL.
	GetListValues(GetListValues *GetListValues) (*GetListValuesResponse, error)

	// GetLrelLength was auto-generated from WSDL.
	GetLrelLength(GetLrelLength *GetLrelLength) (*GetLrelLengthResponse, error)

	// GetLrelValues was auto-generated from WSDL.
	GetLrelValues(GetLrelValues *GetLrelValues) (*GetLrelValuesResponse, error)

	// GetNotificationsForContact was auto-generated from WSDL.
	GetNotificationsForContact(GetNotificationsForContact *GetNotificationsForContact) (*GetNotificationsForContactResponse, error)

	// GetObjectTypeInformation was auto-generated from WSDL.
	GetObjectTypeInformation(GetObjectTypeInformation *GetObjectTypeInformation) (*GetObjectTypeInformationResponse, error)

	// GetObjectValues was auto-generated from WSDL.
	GetObjectValues(GetObjectValues *GetObjectValues) (*GetObjectValuesResponse, error)

	// GetPendingChangeTaskListForContact was auto-generated from WSDL.
	GetPendingChangeTaskListForContact(GetPendingChangeTaskListForContact *GetPendingChangeTaskListForContact) (*GetPendingChangeTaskListForContactResponse, error)

	// GetPendingIssueTaskListForContact was auto-generated from WSDL.
	GetPendingIssueTaskListForContact(GetPendingIssueTaskListForContact *GetPendingIssueTaskListForContact) (*GetPendingIssueTaskListForContactResponse, error)

	// GetPermissionGroups was auto-generated from WSDL.
	GetPermissionGroups(GetPermissionGroups *GetPermissionGroups) (*GetPermissionGroupsResponse, error)

	// GetPolicyInfo was auto-generated from WSDL.
	GetPolicyInfo(GetPolicyInfo *GetPolicyInfo) (*GetPolicyInfoResponse, error)

	// GetPriorities was auto-generated from WSDL.
	GetPriorities(GetPriorities *GetPriorities) (*GetPrioritiesResponse, error)

	// GetPropertyInfoForCategory was auto-generated from WSDL.
	GetPropertyInfoForCategory(GetPropertyInfoForCategory *GetPropertyInfoForCategory) (*GetPropertyInfoForCategoryResponse, error)

	// GetQuestionsAsked was auto-generated from WSDL.
	GetQuestionsAsked(GetQuestionsAsked *GetQuestionsAsked) (*GetQuestionsAskedResponse, error)

	// GetRelatedList was auto-generated from WSDL.
	GetRelatedList(GetRelatedList *GetRelatedList) (*GetRelatedListResponse, error)

	// GetRelatedListValues was auto-generated from WSDL.
	GetRelatedListValues(GetRelatedListValues *GetRelatedListValues) (*GetRelatedListValuesResponse, error)

	// GetRepositoryInfo was auto-generated from WSDL.
	GetRepositoryInfo(GetRepositoryInfo *GetRepositoryInfo) (*GetRepositoryInfoResponse, error)

	// GetStatuses was auto-generated from WSDL.
	GetStatuses(GetStatuses *GetStatuses) (*GetStatusesResponse, error)

	// GetTaskListValues was auto-generated from WSDL.
	GetTaskListValues(GetTaskListValues *GetTaskListValues) (*GetTaskListValuesResponse, error)

	// GetTemplateList was auto-generated from WSDL.
	GetTemplateList(GetTemplateList *GetTemplateList) (*GetTemplateListResponse, error)

	// GetValidTaskTransitions was auto-generated from WSDL.
	GetValidTaskTransitions(GetValidTaskTransitions *GetValidTaskTransitions) (*GetValidTaskTransitionsResponse, error)

	// GetValidTransitions was auto-generated from WSDL.
	GetValidTransitions(GetValidTransitions *GetValidTransitions) (*GetValidTransitionsResponse, error)

	// GetWorkFlowTemplates was auto-generated from WSDL.
	GetWorkFlowTemplates(GetWorkFlowTemplates *GetWorkFlowTemplates) (*GetWorkFlowTemplatesResponse, error)

	// GetWorkflowTemplateList was auto-generated from WSDL.
	GetWorkflowTemplateList(GetWorkflowTemplateList *GetWorkflowTemplateList) (*GetWorkflowTemplateListResponse, error)

	// Impersonate was auto-generated from WSDL.
	Impersonate(Impersonate *Impersonate) (*ImpersonateResponse, error)

	// IsAttmntLinkedKD was auto-generated from WSDL.
	IsAttmntLinkedKD(IsAttmntLinkedKD *IsAttmntLinkedKD) (*IsAttmntLinkedKDResponse, error)

	// LogComment was auto-generated from WSDL.
	LogComment(LogComment *LogComment) (*LogCommentResponse, error)

	// Login was auto-generated from WSDL.
	Login(Login *Login) (*LoginResponse, error)

	// LoginService was auto-generated from WSDL.
	LoginService(LoginService *LoginService) (*LoginServiceResponse, error)

	// LoginServiceManaged was auto-generated from WSDL.
	LoginServiceManaged(LoginServiceManaged *LoginServiceManaged) (*LoginServiceManagedResponse, error)

	// LoginWithArtifact was auto-generated from WSDL.
	LoginWithArtifact(LoginWithArtifact *LoginWithArtifact) (*LoginWithArtifactResponse, error)

	// Logout was auto-generated from WSDL.
	Logout(Logout *Logout) (*LogoutResponse, error)

	// ModifyDocument was auto-generated from WSDL.
	ModifyDocument(ModifyDocument *ModifyDocument) (*ModifyDocumentResponse, error)

	// NotifyContacts was auto-generated from WSDL.
	NotifyContacts(NotifyContacts *NotifyContacts) (*NotifyContactsResponse, error)

	// RateDocument was auto-generated from WSDL.
	RateDocument(RateDocument *RateDocument) (*RateDocumentResponse, error)

	// RemoveAttachment was auto-generated from WSDL.
	RemoveAttachment(RemoveAttachment *RemoveAttachment) (*RemoveAttachmentResponse, error)

	// RemoveLrelRelationships was auto-generated from WSDL.
	RemoveLrelRelationships(RemoveLrelRelationships *RemoveLrelRelationships) (*RemoveLrelRelationshipsResponse, error)

	// RemoveMemberFromGroup was auto-generated from WSDL.
	RemoveMemberFromGroup(RemoveMemberFromGroup *RemoveMemberFromGroup) (*RemoveMemberFromGroupResponse, error)

	// Search was auto-generated from WSDL.
	Search(Search *Search) (*SearchResponse, error)

	// ServerStatus was auto-generated from WSDL.
	ServerStatus(ServerStatus *ServerStatus) (*ServerStatusResponse, error)

	// Transfer was auto-generated from WSDL.
	Transfer(Transfer *Transfer) (*TransferResponse, error)

	// UpdateObject was auto-generated from WSDL.
	UpdateObject(UpdateObject *UpdateObject) (*UpdateObjectResponse, error)

	// UpdateRating was auto-generated from WSDL.
	UpdateRating(UpdateRating *UpdateRating) (*UpdateRatingResponse, error)
}

// ArrayOfInt was auto-generated from WSDL.
type ArrayOfInt struct {
	Integer []*int `xml:"integer,omitempty" json:"integer,omitempty" yaml:"integer,omitempty"`
}

// ArrayOfString was auto-generated from WSDL.
type ArrayOfString struct {
	String []*string `xml:"string,omitempty" json:"string,omitempty" yaml:"string,omitempty"`
}

// ListResult was auto-generated from WSDL.
type ListResult struct {
	ListHandle *int `xml:"listHandle,omitempty" json:"listHandle,omitempty" yaml:"listHandle,omitempty"`
	ListLength *int `xml:"listLength,omitempty" json:"listLength,omitempty" yaml:"listLength,omitempty"`
}

// AddAssetLog was auto-generated from WSDL.
type AddAssetLog struct {
	Sid           *int    `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	AssetHandle   *string `xml:"assetHandle,omitempty" json:"assetHandle,omitempty" yaml:"assetHandle,omitempty"`
	ContactHandle *string `xml:"contactHandle,omitempty" json:"contactHandle,omitempty" yaml:"contactHandle,omitempty"`
	LogText       *string `xml:"logText,omitempty" json:"logText,omitempty" yaml:"logText,omitempty"`
}

// AddAssetLogResponse was auto-generated from WSDL.
type AddAssetLogResponse struct {
}

// AddBookmark was auto-generated from WSDL.
type AddBookmark struct {
	Sid       *int    `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	ContactId *string `xml:"contactId,omitempty" json:"contactId,omitempty" yaml:"contactId,omitempty"`
	DocId     *int    `xml:"docId,omitempty" json:"docId,omitempty" yaml:"docId,omitempty"`
}

// AddBookmarkResponse was auto-generated from WSDL.
type AddBookmarkResponse struct {
	AddBookmarkReturn *string `xml:"addBookmarkReturn,omitempty" json:"addBookmarkReturn,omitempty" yaml:"addBookmarkReturn,omitempty"`
}

// AddComment was auto-generated from WSDL.
type AddComment struct {
	Sid       *int    `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	Comment   *string `xml:"comment,omitempty" json:"comment,omitempty" yaml:"comment,omitempty"`
	DocId     *int    `xml:"docId,omitempty" json:"docId,omitempty" yaml:"docId,omitempty"`
	Email     *string `xml:"email,omitempty" json:"email,omitempty" yaml:"email,omitempty"`
	Username  *string `xml:"username,omitempty" json:"username,omitempty" yaml:"username,omitempty"`
	ContactId *string `xml:"contactId,omitempty" json:"contactId,omitempty" yaml:"contactId,omitempty"`
}

// AddCommentResponse was auto-generated from WSDL.
type AddCommentResponse struct {
	AddCommentReturn *string `xml:"addCommentReturn,omitempty" json:"addCommentReturn,omitempty" yaml:"addCommentReturn,omitempty"`
}

// AddMemberToGroup was auto-generated from WSDL.
type AddMemberToGroup struct {
	Sid           *int    `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	ContactHandle *string `xml:"contactHandle,omitempty" json:"contactHandle,omitempty" yaml:"contactHandle,omitempty"`
	GroupHandle   *string `xml:"groupHandle,omitempty" json:"groupHandle,omitempty" yaml:"groupHandle,omitempty"`
}

// AddMemberToGroupResponse was auto-generated from WSDL.
type AddMemberToGroupResponse struct {
}

// AttachChangeToRequest was auto-generated from WSDL.
type AttachChangeToRequest struct {
	Sid            *int           `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	Creator        *string        `xml:"creator,omitempty" json:"creator,omitempty" yaml:"creator,omitempty"`
	RequestHandle  *string        `xml:"requestHandle,omitempty" json:"requestHandle,omitempty" yaml:"requestHandle,omitempty"`
	ChangeHandle   *string        `xml:"changeHandle,omitempty" json:"changeHandle,omitempty" yaml:"changeHandle,omitempty"`
	ChangeAttrVals *ArrayOfString `xml:"changeAttrVals,omitempty" json:"changeAttrVals,omitempty" yaml:"changeAttrVals,omitempty"`
	Description    *string        `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
}

// AttachChangeToRequestResponse was auto-generated from WSDL.
type AttachChangeToRequestResponse struct {
	AttachChangeToRequestReturn *string `xml:"attachChangeToRequestReturn,omitempty" json:"attachChangeToRequestReturn,omitempty" yaml:"attachChangeToRequestReturn,omitempty"`
}

// AttachURLLink was auto-generated from WSDL.
type AttachURLLink struct {
	Sid         *int    `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	DocId       *int    `xml:"docId,omitempty" json:"docId,omitempty" yaml:"docId,omitempty"`
	Url         *string `xml:"url,omitempty" json:"url,omitempty" yaml:"url,omitempty"`
	AttmntName  *string `xml:"attmntName,omitempty" json:"attmntName,omitempty" yaml:"attmntName,omitempty"`
	Description *string `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
}

// AttachURLLinkResponse was auto-generated from WSDL.
type AttachURLLinkResponse struct {
	AttachURLLinkReturn *int `xml:"attachURLLinkReturn,omitempty" json:"attachURLLinkReturn,omitempty" yaml:"attachURLLinkReturn,omitempty"`
}

// AttachURLLinkToTicket was auto-generated from WSDL.
type AttachURLLinkToTicket struct {
	Sid          *int    `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	TicketHandle *string `xml:"ticketHandle,omitempty" json:"ticketHandle,omitempty" yaml:"ticketHandle,omitempty"`
	Url          *string `xml:"url,omitempty" json:"url,omitempty" yaml:"url,omitempty"`
	AttmntName   *string `xml:"attmntName,omitempty" json:"attmntName,omitempty" yaml:"attmntName,omitempty"`
	Description  *string `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
}

// AttachURLLinkToTicketResponse was auto-generated from WSDL.
type AttachURLLinkToTicketResponse struct {
	AttachURLLinkToTicketReturn *string `xml:"attachURLLinkToTicketReturn,omitempty" json:"attachURLLinkToTicketReturn,omitempty" yaml:"attachURLLinkToTicketReturn,omitempty"`
}

// AttmntFolderLinkCount was auto-generated from WSDL.
type AttmntFolderLinkCount struct {
	Sid      *int `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	FolderId *int `xml:"folderId,omitempty" json:"folderId,omitempty" yaml:"folderId,omitempty"`
}

// AttmntFolderLinkCountResponse was auto-generated from WSDL.
type AttmntFolderLinkCountResponse struct {
	AttmntFolderLinkCountReturn *int `xml:"attmntFolderLinkCountReturn,omitempty" json:"attmntFolderLinkCountReturn,omitempty" yaml:"attmntFolderLinkCountReturn,omitempty"`
}

// CallServerMethod was auto-generated from WSDL.
type CallServerMethod struct {
	Sid         *int           `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	MethodName  *string        `xml:"methodName,omitempty" json:"methodName,omitempty" yaml:"methodName,omitempty"`
	FactoryName *string        `xml:"factoryName,omitempty" json:"factoryName,omitempty" yaml:"factoryName,omitempty"`
	FormatList  *string        `xml:"formatList,omitempty" json:"formatList,omitempty" yaml:"formatList,omitempty"`
	Parameters  *ArrayOfString `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// CallServerMethodResponse was auto-generated from WSDL.
type CallServerMethodResponse struct {
	CallServerMethodReturn *string `xml:"callServerMethodReturn,omitempty" json:"callServerMethodReturn,omitempty" yaml:"callServerMethodReturn,omitempty"`
}

// ChangeStatus was auto-generated from WSDL.
type ChangeStatus struct {
	Sid             *int    `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	Creator         *string `xml:"creator,omitempty" json:"creator,omitempty" yaml:"creator,omitempty"`
	ObjectHandle    *string `xml:"objectHandle,omitempty" json:"objectHandle,omitempty" yaml:"objectHandle,omitempty"`
	Description     *string `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	NewStatusHandle *string `xml:"newStatusHandle,omitempty" json:"newStatusHandle,omitempty" yaml:"newStatusHandle,omitempty"`
}

// ChangeStatusResponse was auto-generated from WSDL.
type ChangeStatusResponse struct {
	ChangeStatusReturn *string `xml:"changeStatusReturn,omitempty" json:"changeStatusReturn,omitempty" yaml:"changeStatusReturn,omitempty"`
}

// ClearNotification was auto-generated from WSDL.
type ClearNotification struct {
	Sid      *int    `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	LrObject *string `xml:"lrObject,omitempty" json:"lrObject,omitempty" yaml:"lrObject,omitempty"`
	ClearBy  *string `xml:"clearBy,omitempty" json:"clearBy,omitempty" yaml:"clearBy,omitempty"`
}

// ClearNotificationResponse was auto-generated from WSDL.
type ClearNotificationResponse struct {
	ClearNotificationReturn *int `xml:"clearNotificationReturn,omitempty" json:"clearNotificationReturn,omitempty" yaml:"clearNotificationReturn,omitempty"`
}

// CloseTicket was auto-generated from WSDL.
type CloseTicket struct {
	Sid          *int    `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	Description  *string `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	TicketHandle *string `xml:"ticketHandle,omitempty" json:"ticketHandle,omitempty" yaml:"ticketHandle,omitempty"`
}

// CloseTicketResponse was auto-generated from WSDL.
type CloseTicketResponse struct {
	CloseTicketReturn *string `xml:"closeTicketReturn,omitempty" json:"closeTicketReturn,omitempty" yaml:"closeTicketReturn,omitempty"`
}

// CreateActivityLog was auto-generated from WSDL.
type CreateActivityLog struct {
	Sid          *int    `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	Creator      *string `xml:"creator,omitempty" json:"creator,omitempty" yaml:"creator,omitempty"`
	ObjectHandle *string `xml:"objectHandle,omitempty" json:"objectHandle,omitempty" yaml:"objectHandle,omitempty"`
	Description  *string `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	LogType      *string `xml:"logType,omitempty" json:"logType,omitempty" yaml:"logType,omitempty"`
	TimeSpent    *int    `xml:"timeSpent,omitempty" json:"timeSpent,omitempty" yaml:"timeSpent,omitempty"`
	Internal     *bool   `xml:"internal,omitempty" json:"internal,omitempty" yaml:"internal,omitempty"`
}

// CreateActivityLogResponse was auto-generated from WSDL.
type CreateActivityLogResponse struct {
	CreateActivityLogReturn *string `xml:"createActivityLogReturn,omitempty" json:"createActivityLogReturn,omitempty" yaml:"createActivityLogReturn,omitempty"`
}

// CreateAsset was auto-generated from WSDL.
type CreateAsset struct {
	Sid                *int           `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	AttrVals           *ArrayOfString `xml:"attrVals,omitempty" json:"attrVals,omitempty" yaml:"attrVals,omitempty"`
	Attributes         *ArrayOfString `xml:"attributes,omitempty" json:"attributes,omitempty" yaml:"attributes,omitempty"`
	CreateAssetResult  *string        `xml:"createAssetResult,omitempty" json:"createAssetResult,omitempty" yaml:"createAssetResult,omitempty"`
	NewAssetHandle     *string        `xml:"newAssetHandle,omitempty" json:"newAssetHandle,omitempty" yaml:"newAssetHandle,omitempty"`
	NewExtensionHandle *string        `xml:"newExtensionHandle,omitempty" json:"newExtensionHandle,omitempty" yaml:"newExtensionHandle,omitempty"`
	NewExtensionName   *string        `xml:"newExtensionName,omitempty" json:"newExtensionName,omitempty" yaml:"newExtensionName,omitempty"`
}

// CreateAssetParentChildRelationship was auto-generated from WSDL.
type CreateAssetParentChildRelationship struct {
	Sid          *int    `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	ParentHandle *string `xml:"parentHandle,omitempty" json:"parentHandle,omitempty" yaml:"parentHandle,omitempty"`
	ChildHandle  *string `xml:"childHandle,omitempty" json:"childHandle,omitempty" yaml:"childHandle,omitempty"`
}

// CreateAssetParentChildRelationshipResponse was auto-generated
// from WSDL.
type CreateAssetParentChildRelationshipResponse struct {
	CreateAssetParentChildRelationshipReturn *string `xml:"createAssetParentChildRelationshipReturn,omitempty" json:"createAssetParentChildRelationshipReturn,omitempty" yaml:"createAssetParentChildRelationshipReturn,omitempty"`
}

// CreateAssetResponse was auto-generated from WSDL.
type CreateAssetResponse struct {
	CreateAssetResult  *string `xml:"createAssetResult,omitempty" json:"createAssetResult,omitempty" yaml:"createAssetResult,omitempty"`
	NewAssetHandle     *string `xml:"newAssetHandle,omitempty" json:"newAssetHandle,omitempty" yaml:"newAssetHandle,omitempty"`
	NewExtensionHandle *string `xml:"newExtensionHandle,omitempty" json:"newExtensionHandle,omitempty" yaml:"newExtensionHandle,omitempty"`
	NewExtensionName   *string `xml:"newExtensionName,omitempty" json:"newExtensionName,omitempty" yaml:"newExtensionName,omitempty"`
}

// CreateAttachment was auto-generated from WSDL.
type CreateAttachment struct {
	Sid              *int    `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	RepositoryHandle *string `xml:"repositoryHandle,omitempty" json:"repositoryHandle,omitempty" yaml:"repositoryHandle,omitempty"`
	ObjectHandle     *string `xml:"objectHandle,omitempty" json:"objectHandle,omitempty" yaml:"objectHandle,omitempty"`
	Description      *string `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	FileName         *string `xml:"fileName,omitempty" json:"fileName,omitempty" yaml:"fileName,omitempty"`
}

// CreateAttachmentResponse was auto-generated from WSDL.
type CreateAttachmentResponse struct {
	CreateAttachmentReturn *string `xml:"createAttachmentReturn,omitempty" json:"createAttachmentReturn,omitempty" yaml:"createAttachmentReturn,omitempty"`
}

// CreateAttmnt was auto-generated from WSDL.
type CreateAttmnt struct {
	Sid              *int    `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	RepositoryHandle *string `xml:"repositoryHandle,omitempty" json:"repositoryHandle,omitempty" yaml:"repositoryHandle,omitempty"`
	FolderId         *int    `xml:"folderId,omitempty" json:"folderId,omitempty" yaml:"folderId,omitempty"`
	ObjectHandle     *int    `xml:"objectHandle,omitempty" json:"objectHandle,omitempty" yaml:"objectHandle,omitempty"`
	Description      *string `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	FileName         *string `xml:"fileName,omitempty" json:"fileName,omitempty" yaml:"fileName,omitempty"`
}

// CreateAttmntResponse was auto-generated from WSDL.
type CreateAttmntResponse struct {
	CreateAttmntReturn *string `xml:"createAttmntReturn,omitempty" json:"createAttmntReturn,omitempty" yaml:"createAttmntReturn,omitempty"`
}

// CreateChangeOrder was auto-generated from WSDL.
type CreateChangeOrder struct {
	Sid             *int           `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	CreatorHandle   *string        `xml:"creatorHandle,omitempty" json:"creatorHandle,omitempty" yaml:"creatorHandle,omitempty"`
	AttrVals        *ArrayOfString `xml:"attrVals,omitempty" json:"attrVals,omitempty" yaml:"attrVals,omitempty"`
	PropertyValues  *ArrayOfString `xml:"propertyValues,omitempty" json:"propertyValues,omitempty" yaml:"propertyValues,omitempty"`
	Template        *string        `xml:"template,omitempty" json:"template,omitempty" yaml:"template,omitempty"`
	Attributes      *ArrayOfString `xml:"attributes,omitempty" json:"attributes,omitempty" yaml:"attributes,omitempty"`
	NewChangeHandle *string        `xml:"newChangeHandle,omitempty" json:"newChangeHandle,omitempty" yaml:"newChangeHandle,omitempty"`
	NewChangeNumber *string        `xml:"newChangeNumber,omitempty" json:"newChangeNumber,omitempty" yaml:"newChangeNumber,omitempty"`
}

// CreateChangeOrderResponse was auto-generated from WSDL.
type CreateChangeOrderResponse struct {
	CreateChangeOrderReturn *string `xml:"createChangeOrderReturn,omitempty" json:"createChangeOrderReturn,omitempty" yaml:"createChangeOrderReturn,omitempty"`
	NewChangeHandle         *string `xml:"newChangeHandle,omitempty" json:"newChangeHandle,omitempty" yaml:"newChangeHandle,omitempty"`
	NewChangeNumber         *string `xml:"newChangeNumber,omitempty" json:"newChangeNumber,omitempty" yaml:"newChangeNumber,omitempty"`
}

// CreateDocument was auto-generated from WSDL.
type CreateDocument struct {
	Sid          *int           `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	KdAttributes *ArrayOfString `xml:"kdAttributes,omitempty" json:"kdAttributes,omitempty" yaml:"kdAttributes,omitempty"`
}

// CreateDocumentResponse was auto-generated from WSDL.
type CreateDocumentResponse struct {
	CreateDocumentReturn *string `xml:"createDocumentReturn,omitempty" json:"createDocumentReturn,omitempty" yaml:"createDocumentReturn,omitempty"`
}

// CreateFolder was auto-generated from WSDL.
type CreateFolder struct {
	Sid            *int    `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	ParentFolderId *int    `xml:"parentFolderId,omitempty" json:"parentFolderId,omitempty" yaml:"parentFolderId,omitempty"`
	RepId          *int    `xml:"repId,omitempty" json:"repId,omitempty" yaml:"repId,omitempty"`
	FolderType     *int    `xml:"folderType,omitempty" json:"folderType,omitempty" yaml:"folderType,omitempty"`
	Description    *string `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	FolderName     *string `xml:"folderName,omitempty" json:"folderName,omitempty" yaml:"folderName,omitempty"`
}

// CreateFolderResponse was auto-generated from WSDL.
type CreateFolderResponse struct {
	CreateFolderReturn *string `xml:"createFolderReturn,omitempty" json:"createFolderReturn,omitempty" yaml:"createFolderReturn,omitempty"`
}

// CreateIssue was auto-generated from WSDL.
type CreateIssue struct {
	Sid            *int           `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	CreatorHandle  *string        `xml:"creatorHandle,omitempty" json:"creatorHandle,omitempty" yaml:"creatorHandle,omitempty"`
	AttrVals       *ArrayOfString `xml:"attrVals,omitempty" json:"attrVals,omitempty" yaml:"attrVals,omitempty"`
	PropertyValues *ArrayOfString `xml:"propertyValues,omitempty" json:"propertyValues,omitempty" yaml:"propertyValues,omitempty"`
	Template       *string        `xml:"template,omitempty" json:"template,omitempty" yaml:"template,omitempty"`
	Attributes     *ArrayOfString `xml:"attributes,omitempty" json:"attributes,omitempty" yaml:"attributes,omitempty"`
	NewIssueHandle *string        `xml:"newIssueHandle,omitempty" json:"newIssueHandle,omitempty" yaml:"newIssueHandle,omitempty"`
	NewIssueNumber *string        `xml:"newIssueNumber,omitempty" json:"newIssueNumber,omitempty" yaml:"newIssueNumber,omitempty"`
}

// CreateIssueResponse was auto-generated from WSDL.
type CreateIssueResponse struct {
	CreateIssueReturn *string `xml:"createIssueReturn,omitempty" json:"createIssueReturn,omitempty" yaml:"createIssueReturn,omitempty"`
	NewIssueHandle    *string `xml:"newIssueHandle,omitempty" json:"newIssueHandle,omitempty" yaml:"newIssueHandle,omitempty"`
	NewIssueNumber    *string `xml:"newIssueNumber,omitempty" json:"newIssueNumber,omitempty" yaml:"newIssueNumber,omitempty"`
}

// CreateLrelRelationships was auto-generated from WSDL.
type CreateLrelRelationships struct {
	Sid              *int           `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	ContextObject    *string        `xml:"contextObject,omitempty" json:"contextObject,omitempty" yaml:"contextObject,omitempty"`
	LrelName         *string        `xml:"lrelName,omitempty" json:"lrelName,omitempty" yaml:"lrelName,omitempty"`
	AddObjectHandles *ArrayOfString `xml:"addObjectHandles,omitempty" json:"addObjectHandles,omitempty" yaml:"addObjectHandles,omitempty"`
}

// CreateLrelRelationshipsResponse was auto-generated from WSDL.
type CreateLrelRelationshipsResponse struct {
}

// CreateObject was auto-generated from WSDL.
type CreateObject struct {
	Sid                *int           `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	ObjectType         *string        `xml:"objectType,omitempty" json:"objectType,omitempty" yaml:"objectType,omitempty"`
	AttrVals           *ArrayOfString `xml:"attrVals,omitempty" json:"attrVals,omitempty" yaml:"attrVals,omitempty"`
	Attributes         *ArrayOfString `xml:"attributes,omitempty" json:"attributes,omitempty" yaml:"attributes,omitempty"`
	CreateObjectResult *string        `xml:"createObjectResult,omitempty" json:"createObjectResult,omitempty" yaml:"createObjectResult,omitempty"`
	NewHandle          *string        `xml:"newHandle,omitempty" json:"newHandle,omitempty" yaml:"newHandle,omitempty"`
}

// CreateObjectResponse was auto-generated from WSDL.
type CreateObjectResponse struct {
	CreateObjectResult *string `xml:"createObjectResult,omitempty" json:"createObjectResult,omitempty" yaml:"createObjectResult,omitempty"`
	NewHandle          *string `xml:"newHandle,omitempty" json:"newHandle,omitempty" yaml:"newHandle,omitempty"`
}

// CreateQuickTicket was auto-generated from WSDL.
type CreateQuickTicket struct {
	Sid             *int    `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	CustomerHandle  *string `xml:"customerHandle,omitempty" json:"customerHandle,omitempty" yaml:"customerHandle,omitempty"`
	Description     *string `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	NewTicketHandle *string `xml:"newTicketHandle,omitempty" json:"newTicketHandle,omitempty" yaml:"newTicketHandle,omitempty"`
	NewTicketNumber *string `xml:"newTicketNumber,omitempty" json:"newTicketNumber,omitempty" yaml:"newTicketNumber,omitempty"`
}

// CreateQuickTicketResponse was auto-generated from WSDL.
type CreateQuickTicketResponse struct {
	CreateQuickTicketReturn *string `xml:"createQuickTicketReturn,omitempty" json:"createQuickTicketReturn,omitempty" yaml:"createQuickTicketReturn,omitempty"`
	NewTicketHandle         *string `xml:"newTicketHandle,omitempty" json:"newTicketHandle,omitempty" yaml:"newTicketHandle,omitempty"`
	NewTicketNumber         *string `xml:"newTicketNumber,omitempty" json:"newTicketNumber,omitempty" yaml:"newTicketNumber,omitempty"`
}

// CreateRequest was auto-generated from WSDL.
type CreateRequest struct {
	Sid              *int           `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	CreatorHandle    *string        `xml:"creatorHandle,omitempty" json:"creatorHandle,omitempty" yaml:"creatorHandle,omitempty"`
	AttrVals         *ArrayOfString `xml:"attrVals,omitempty" json:"attrVals,omitempty" yaml:"attrVals,omitempty"`
	PropertyValues   *ArrayOfString `xml:"propertyValues,omitempty" json:"propertyValues,omitempty" yaml:"propertyValues,omitempty"`
	Template         *string        `xml:"template,omitempty" json:"template,omitempty" yaml:"template,omitempty"`
	Attributes       *ArrayOfString `xml:"attributes,omitempty" json:"attributes,omitempty" yaml:"attributes,omitempty"`
	NewRequestHandle *string        `xml:"newRequestHandle,omitempty" json:"newRequestHandle,omitempty" yaml:"newRequestHandle,omitempty"`
	NewRequestNumber *string        `xml:"newRequestNumber,omitempty" json:"newRequestNumber,omitempty" yaml:"newRequestNumber,omitempty"`
}

// CreateRequestResponse was auto-generated from WSDL.
type CreateRequestResponse struct {
	CreateRequestReturn *string `xml:"createRequestReturn,omitempty" json:"createRequestReturn,omitempty" yaml:"createRequestReturn,omitempty"`
	NewRequestHandle    *string `xml:"newRequestHandle,omitempty" json:"newRequestHandle,omitempty" yaml:"newRequestHandle,omitempty"`
	NewRequestNumber    *string `xml:"newRequestNumber,omitempty" json:"newRequestNumber,omitempty" yaml:"newRequestNumber,omitempty"`
}

// CreateTicket was auto-generated from WSDL.
type CreateTicket struct {
	Sid                   *int    `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	Description           *string `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	Problem_type          *string `xml:"problem_type,omitempty" json:"problem_type,omitempty" yaml:"problem_type,omitempty"`
	Userid                *string `xml:"userid,omitempty" json:"userid,omitempty" yaml:"userid,omitempty"`
	Asset                 *string `xml:"asset,omitempty" json:"asset,omitempty" yaml:"asset,omitempty"`
	Duplication_id        *string `xml:"duplication_id,omitempty" json:"duplication_id,omitempty" yaml:"duplication_id,omitempty"`
	NewTicketHandle       *string `xml:"newTicketHandle,omitempty" json:"newTicketHandle,omitempty" yaml:"newTicketHandle,omitempty"`
	NewTicketNumber       *string `xml:"newTicketNumber,omitempty" json:"newTicketNumber,omitempty" yaml:"newTicketNumber,omitempty"`
	ReturnUserData        *string `xml:"returnUserData,omitempty" json:"returnUserData,omitempty" yaml:"returnUserData,omitempty"`
	ReturnApplicationData *string `xml:"returnApplicationData,omitempty" json:"returnApplicationData,omitempty" yaml:"returnApplicationData,omitempty"`
}

// CreateTicketResponse was auto-generated from WSDL.
type CreateTicketResponse struct {
	CreateTicketReturn    *string `xml:"createTicketReturn,omitempty" json:"createTicketReturn,omitempty" yaml:"createTicketReturn,omitempty"`
	NewTicketHandle       *string `xml:"newTicketHandle,omitempty" json:"newTicketHandle,omitempty" yaml:"newTicketHandle,omitempty"`
	NewTicketNumber       *string `xml:"newTicketNumber,omitempty" json:"newTicketNumber,omitempty" yaml:"newTicketNumber,omitempty"`
	ReturnUserData        *string `xml:"returnUserData,omitempty" json:"returnUserData,omitempty" yaml:"returnUserData,omitempty"`
	ReturnApplicationData *string `xml:"returnApplicationData,omitempty" json:"returnApplicationData,omitempty" yaml:"returnApplicationData,omitempty"`
}

// CreateWorkFlowTask was auto-generated from WSDL.
type CreateWorkFlowTask struct {
	Sid                      *int           `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	AttrVals                 *ArrayOfString `xml:"attrVals,omitempty" json:"attrVals,omitempty" yaml:"attrVals,omitempty"`
	ObjectHandle             *string        `xml:"objectHandle,omitempty" json:"objectHandle,omitempty" yaml:"objectHandle,omitempty"`
	CreatorHandle            *string        `xml:"creatorHandle,omitempty" json:"creatorHandle,omitempty" yaml:"creatorHandle,omitempty"`
	SelectedWorkFlow         *string        `xml:"selectedWorkFlow,omitempty" json:"selectedWorkFlow,omitempty" yaml:"selectedWorkFlow,omitempty"`
	TaskType                 *string        `xml:"taskType,omitempty" json:"taskType,omitempty" yaml:"taskType,omitempty"`
	Attributes               *ArrayOfString `xml:"attributes,omitempty" json:"attributes,omitempty" yaml:"attributes,omitempty"`
	CreateWorkFlowTaskResult *string        `xml:"createWorkFlowTaskResult,omitempty" json:"createWorkFlowTaskResult,omitempty" yaml:"createWorkFlowTaskResult,omitempty"`
	NewHandle                *string        `xml:"newHandle,omitempty" json:"newHandle,omitempty" yaml:"newHandle,omitempty"`
}

// CreateWorkFlowTaskResponse was auto-generated from WSDL.
type CreateWorkFlowTaskResponse struct {
	CreateWorkFlowTaskResult *string `xml:"createWorkFlowTaskResult,omitempty" json:"createWorkFlowTaskResult,omitempty" yaml:"createWorkFlowTaskResult,omitempty"`
	NewHandle                *string `xml:"newHandle,omitempty" json:"newHandle,omitempty" yaml:"newHandle,omitempty"`
}

// DeleteBookmark was auto-generated from WSDL.
type DeleteBookmark struct {
	Sid       *int    `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	ContactId *string `xml:"contactId,omitempty" json:"contactId,omitempty" yaml:"contactId,omitempty"`
	DocId     *int    `xml:"docId,omitempty" json:"docId,omitempty" yaml:"docId,omitempty"`
}

// DeleteBookmarkResponse was auto-generated from WSDL.
type DeleteBookmarkResponse struct {
	DeleteBookmarkReturn *int `xml:"deleteBookmarkReturn,omitempty" json:"deleteBookmarkReturn,omitempty" yaml:"deleteBookmarkReturn,omitempty"`
}

// DeleteComment was auto-generated from WSDL.
type DeleteComment struct {
	Sid       *int `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	CommentId *int `xml:"commentId,omitempty" json:"commentId,omitempty" yaml:"commentId,omitempty"`
}

// DeleteCommentResponse was auto-generated from WSDL.
type DeleteCommentResponse struct {
	DeleteCommentReturn *int `xml:"deleteCommentReturn,omitempty" json:"deleteCommentReturn,omitempty" yaml:"deleteCommentReturn,omitempty"`
}

// DeleteDocument was auto-generated from WSDL.
type DeleteDocument struct {
	Sid   *int `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	DocId *int `xml:"docId,omitempty" json:"docId,omitempty" yaml:"docId,omitempty"`
}

// DeleteDocumentResponse was auto-generated from WSDL.
type DeleteDocumentResponse struct {
	DeleteDocumentReturn *int `xml:"deleteDocumentReturn,omitempty" json:"deleteDocumentReturn,omitempty" yaml:"deleteDocumentReturn,omitempty"`
}

// DeleteWorkFlowTask was auto-generated from WSDL.
type DeleteWorkFlowTask struct {
	Sid            *int    `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	WorkFlowHandle *string `xml:"workFlowHandle,omitempty" json:"workFlowHandle,omitempty" yaml:"workFlowHandle,omitempty"`
	ObjectHandle   *string `xml:"objectHandle,omitempty" json:"objectHandle,omitempty" yaml:"objectHandle,omitempty"`
}

// DeleteWorkFlowTaskResponse was auto-generated from WSDL.
type DeleteWorkFlowTaskResponse struct {
}

// DetachChangeFromRequest was auto-generated from WSDL.
type DetachChangeFromRequest struct {
	Sid           *int    `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	Creator       *string `xml:"creator,omitempty" json:"creator,omitempty" yaml:"creator,omitempty"`
	RequestHandle *string `xml:"requestHandle,omitempty" json:"requestHandle,omitempty" yaml:"requestHandle,omitempty"`
	Description   *string `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
}

// DetachChangeFromRequestResponse was auto-generated from WSDL.
type DetachChangeFromRequestResponse struct {
	DetachChangeFromRequestReturn *string `xml:"detachChangeFromRequestReturn,omitempty" json:"detachChangeFromRequestReturn,omitempty" yaml:"detachChangeFromRequestReturn,omitempty"`
}

// DoQuery was auto-generated from WSDL.
type DoQuery struct {
	Sid         *int    `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	ObjectType  *string `xml:"objectType,omitempty" json:"objectType,omitempty" yaml:"objectType,omitempty"`
	WhereClause *string `xml:"whereClause,omitempty" json:"whereClause,omitempty" yaml:"whereClause,omitempty"`
}

// DoQueryResponse was auto-generated from WSDL.
type DoQueryResponse struct {
	DoQueryReturn *ListResult `xml:"doQueryReturn,omitempty" json:"doQueryReturn,omitempty" yaml:"doQueryReturn,omitempty"`
}

// DoSelect was auto-generated from WSDL.
type DoSelect struct {
	Sid         *int           `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	ObjectType  *string        `xml:"objectType,omitempty" json:"objectType,omitempty" yaml:"objectType,omitempty"`
	WhereClause *string        `xml:"whereClause,omitempty" json:"whereClause,omitempty" yaml:"whereClause,omitempty"`
	MaxRows     *int           `xml:"maxRows,omitempty" json:"maxRows,omitempty" yaml:"maxRows,omitempty"`
	Attributes  *ArrayOfString `xml:"attributes,omitempty" json:"attributes,omitempty" yaml:"attributes,omitempty"`
}

// DoSelectKD was auto-generated from WSDL.
type DoSelectKD struct {
	Sid         *int           `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	WhereClause *string        `xml:"whereClause,omitempty" json:"whereClause,omitempty" yaml:"whereClause,omitempty"`
	SortBy      *string        `xml:"sortBy,omitempty" json:"sortBy,omitempty" yaml:"sortBy,omitempty"`
	Desc        *bool          `xml:"desc,omitempty" json:"desc,omitempty" yaml:"desc,omitempty"`
	MaxRows     *int           `xml:"maxRows,omitempty" json:"maxRows,omitempty" yaml:"maxRows,omitempty"`
	Attributes  *ArrayOfString `xml:"attributes,omitempty" json:"attributes,omitempty" yaml:"attributes,omitempty"`
	Skip        *int           `xml:"skip,omitempty" json:"skip,omitempty" yaml:"skip,omitempty"`
}

// DoSelectKDResponse was auto-generated from WSDL.
type DoSelectKDResponse struct {
	DoSelectKDReturn *string `xml:"doSelectKDReturn,omitempty" json:"doSelectKDReturn,omitempty" yaml:"doSelectKDReturn,omitempty"`
}

// DoSelectResponse was auto-generated from WSDL.
type DoSelectResponse struct {
	DoSelectReturn *string `xml:"doSelectReturn,omitempty" json:"doSelectReturn,omitempty" yaml:"doSelectReturn,omitempty"`
}

// Escalate was auto-generated from WSDL.
type Escalate struct {
	Sid                   *int    `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	Creator               *string `xml:"creator,omitempty" json:"creator,omitempty" yaml:"creator,omitempty"`
	ObjectHandle          *string `xml:"objectHandle,omitempty" json:"objectHandle,omitempty" yaml:"objectHandle,omitempty"`
	Description           *string `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	SetAssignee           *bool   `xml:"setAssignee,omitempty" json:"setAssignee,omitempty" yaml:"setAssignee,omitempty"`
	NewAssigneeHandle     *string `xml:"newAssigneeHandle,omitempty" json:"newAssigneeHandle,omitempty" yaml:"newAssigneeHandle,omitempty"`
	SetGroup              *bool   `xml:"setGroup,omitempty" json:"setGroup,omitempty" yaml:"setGroup,omitempty"`
	NewGroupHandle        *string `xml:"newGroupHandle,omitempty" json:"newGroupHandle,omitempty" yaml:"newGroupHandle,omitempty"`
	SetOrganization       *bool   `xml:"setOrganization,omitempty" json:"setOrganization,omitempty" yaml:"setOrganization,omitempty"`
	NewOrganizationHandle *string `xml:"newOrganizationHandle,omitempty" json:"newOrganizationHandle,omitempty" yaml:"newOrganizationHandle,omitempty"`
	SetPriority           *bool   `xml:"setPriority,omitempty" json:"setPriority,omitempty" yaml:"setPriority,omitempty"`
	NewPriorityHandle     *string `xml:"newPriorityHandle,omitempty" json:"newPriorityHandle,omitempty" yaml:"newPriorityHandle,omitempty"`
}

// EscalateResponse was auto-generated from WSDL.
type EscalateResponse struct {
	EscalateReturn *string `xml:"escalateReturn,omitempty" json:"escalateReturn,omitempty" yaml:"escalateReturn,omitempty"`
}

// Faq was auto-generated from WSDL.
type Faq struct {
	Sid          *int    `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	CategoryIds  *string `xml:"categoryIds,omitempty" json:"categoryIds,omitempty" yaml:"categoryIds,omitempty"`
	ResultSize   *int    `xml:"resultSize,omitempty" json:"resultSize,omitempty" yaml:"resultSize,omitempty"`
	PropertyList *string `xml:"propertyList,omitempty" json:"propertyList,omitempty" yaml:"propertyList,omitempty"`
	SortBy       *string `xml:"sortBy,omitempty" json:"sortBy,omitempty" yaml:"sortBy,omitempty"`
	Descending   *bool   `xml:"descending,omitempty" json:"descending,omitempty" yaml:"descending,omitempty"`
	WhereClause  *string `xml:"whereClause,omitempty" json:"whereClause,omitempty" yaml:"whereClause,omitempty"`
	MaxDocIDs    *int    `xml:"maxDocIDs,omitempty" json:"maxDocIDs,omitempty" yaml:"maxDocIDs,omitempty"`
}

// FaqResponse was auto-generated from WSDL.
type FaqResponse struct {
	FaqReturn *string `xml:"faqReturn,omitempty" json:"faqReturn,omitempty" yaml:"faqReturn,omitempty"`
}

// FindContacts was auto-generated from WSDL.
type FindContacts struct {
	Sid          *int    `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	UserName     *string `xml:"userName,omitempty" json:"userName,omitempty" yaml:"userName,omitempty"`
	LastName     *string `xml:"lastName,omitempty" json:"lastName,omitempty" yaml:"lastName,omitempty"`
	FirstName    *string `xml:"firstName,omitempty" json:"firstName,omitempty" yaml:"firstName,omitempty"`
	Email        *string `xml:"email,omitempty" json:"email,omitempty" yaml:"email,omitempty"`
	AccessType   *string `xml:"accessType,omitempty" json:"accessType,omitempty" yaml:"accessType,omitempty"`
	InactiveFlag *int    `xml:"inactiveFlag,omitempty" json:"inactiveFlag,omitempty" yaml:"inactiveFlag,omitempty"`
}

// FindContactsResponse was auto-generated from WSDL.
type FindContactsResponse struct {
	FindContactsReturn *string `xml:"findContactsReturn,omitempty" json:"findContactsReturn,omitempty" yaml:"findContactsReturn,omitempty"`
}

// FreeListHandles was auto-generated from WSDL.
type FreeListHandles struct {
	Sid     *int        `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	Handles *ArrayOfInt `xml:"handles,omitempty" json:"handles,omitempty" yaml:"handles,omitempty"`
}

// FreeListHandlesResponse was auto-generated from WSDL.
type FreeListHandlesResponse struct {
}

// GetAccessTypeForContact was auto-generated from WSDL.
type GetAccessTypeForContact struct {
	Sid           *int    `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	ContactHandle *string `xml:"contactHandle,omitempty" json:"contactHandle,omitempty" yaml:"contactHandle,omitempty"`
}

// GetAccessTypeForContactResponse was auto-generated from WSDL.
type GetAccessTypeForContactResponse struct {
	GetAccessTypeForContactReturn *string `xml:"getAccessTypeForContactReturn,omitempty" json:"getAccessTypeForContactReturn,omitempty" yaml:"getAccessTypeForContactReturn,omitempty"`
}

// GetArtifact was auto-generated from WSDL.
type GetArtifact struct {
	Sid     *int    `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	Contact *string `xml:"contact,omitempty" json:"contact,omitempty" yaml:"contact,omitempty"`
	Passwd  *string `xml:"passwd,omitempty" json:"passwd,omitempty" yaml:"passwd,omitempty"`
}

// GetArtifactResponse was auto-generated from WSDL.
type GetArtifactResponse struct {
	GetArtifactReturn *string `xml:"getArtifactReturn,omitempty" json:"getArtifactReturn,omitempty" yaml:"getArtifactReturn,omitempty"`
}

// GetAssetExtensionInformation was auto-generated from WSDL.
type GetAssetExtensionInformation struct {
	Sid                   *int           `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	AssetHandle           *string        `xml:"assetHandle,omitempty" json:"assetHandle,omitempty" yaml:"assetHandle,omitempty"`
	Attributes            *ArrayOfString `xml:"attributes,omitempty" json:"attributes,omitempty" yaml:"attributes,omitempty"`
	GetAssetExtInfoResult *string        `xml:"getAssetExtInfoResult,omitempty" json:"getAssetExtInfoResult,omitempty" yaml:"getAssetExtInfoResult,omitempty"`
	ExtensionHandle       *string        `xml:"extensionHandle,omitempty" json:"extensionHandle,omitempty" yaml:"extensionHandle,omitempty"`
	ExtensionName         *string        `xml:"extensionName,omitempty" json:"extensionName,omitempty" yaml:"extensionName,omitempty"`
}

// GetAssetExtensionInformationResponse was auto-generated from
// WSDL.
type GetAssetExtensionInformationResponse struct {
	GetAssetExtInfoResult *string `xml:"getAssetExtInfoResult,omitempty" json:"getAssetExtInfoResult,omitempty" yaml:"getAssetExtInfoResult,omitempty"`
	ExtensionHandle       *string `xml:"extensionHandle,omitempty" json:"extensionHandle,omitempty" yaml:"extensionHandle,omitempty"`
	ExtensionName         *string `xml:"extensionName,omitempty" json:"extensionName,omitempty" yaml:"extensionName,omitempty"`
}

// GetAttmntInfo was auto-generated from WSDL.
type GetAttmntInfo struct {
	Sid      *int `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	AttmntId *int `xml:"attmntId,omitempty" json:"attmntId,omitempty" yaml:"attmntId,omitempty"`
}

// GetAttmntInfoResponse was auto-generated from WSDL.
type GetAttmntInfoResponse struct {
	GetAttmntInfoReturn *string `xml:"getAttmntInfoReturn,omitempty" json:"getAttmntInfoReturn,omitempty" yaml:"getAttmntInfoReturn,omitempty"`
}

// GetAttmntList was auto-generated from WSDL.
type GetAttmntList struct {
	Sid      *int `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	FolderId *int `xml:"folderId,omitempty" json:"folderId,omitempty" yaml:"folderId,omitempty"`
	RepId    *int `xml:"repId,omitempty" json:"repId,omitempty" yaml:"repId,omitempty"`
}

// GetAttmntListPerKD was auto-generated from WSDL.
type GetAttmntListPerKD struct {
	Sid   *int `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	DocId *int `xml:"docId,omitempty" json:"docId,omitempty" yaml:"docId,omitempty"`
}

// GetAttmntListPerKDResponse was auto-generated from WSDL.
type GetAttmntListPerKDResponse struct {
	GetAttmntListPerKDReturn *string `xml:"getAttmntListPerKDReturn,omitempty" json:"getAttmntListPerKDReturn,omitempty" yaml:"getAttmntListPerKDReturn,omitempty"`
}

// GetAttmntListResponse was auto-generated from WSDL.
type GetAttmntListResponse struct {
	GetAttmntListReturn *string `xml:"getAttmntListReturn,omitempty" json:"getAttmntListReturn,omitempty" yaml:"getAttmntListReturn,omitempty"`
}

// GetBookmarks was auto-generated from WSDL.
type GetBookmarks struct {
	Sid       *int    `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	ContactId *string `xml:"contactId,omitempty" json:"contactId,omitempty" yaml:"contactId,omitempty"`
}

// GetBookmarksResponse was auto-generated from WSDL.
type GetBookmarksResponse struct {
	GetBookmarksReturn *string `xml:"getBookmarksReturn,omitempty" json:"getBookmarksReturn,omitempty" yaml:"getBookmarksReturn,omitempty"`
}

// GetBopsid was auto-generated from WSDL.
type GetBopsid struct {
	Sid     *int    `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	Contact *string `xml:"contact,omitempty" json:"contact,omitempty" yaml:"contact,omitempty"`
}

// GetBopsidResponse was auto-generated from WSDL.
type GetBopsidResponse struct {
	GetBopsidReturn *string `xml:"getBopsidReturn,omitempty" json:"getBopsidReturn,omitempty" yaml:"getBopsidReturn,omitempty"`
}

// GetCategory was auto-generated from WSDL.
type GetCategory struct {
	Sid              *int  `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	CatId            *int  `xml:"catId,omitempty" json:"catId,omitempty" yaml:"catId,omitempty"`
	GetCategoryPaths *bool `xml:"getCategoryPaths,omitempty" json:"getCategoryPaths,omitempty" yaml:"getCategoryPaths,omitempty"`
}

// GetCategoryResponse was auto-generated from WSDL.
type GetCategoryResponse struct {
	GetCategoryReturn *string `xml:"getCategoryReturn,omitempty" json:"getCategoryReturn,omitempty" yaml:"getCategoryReturn,omitempty"`
}

// GetComments was auto-generated from WSDL.
type GetComments struct {
	Sid    *int    `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	DocIds *string `xml:"docIds,omitempty" json:"docIds,omitempty" yaml:"docIds,omitempty"`
}

// GetCommentsResponse was auto-generated from WSDL.
type GetCommentsResponse struct {
	GetCommentsReturn *string `xml:"getCommentsReturn,omitempty" json:"getCommentsReturn,omitempty" yaml:"getCommentsReturn,omitempty"`
}

// GetConfigurationMode was auto-generated from WSDL.
type GetConfigurationMode struct {
	Sid *int `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
}

// GetConfigurationModeResponse was auto-generated from WSDL.
type GetConfigurationModeResponse struct {
	GetConfigurationModeReturn *string `xml:"getConfigurationModeReturn,omitempty" json:"getConfigurationModeReturn,omitempty" yaml:"getConfigurationModeReturn,omitempty"`
}

// GetContact was auto-generated from WSDL.
type GetContact struct {
	Sid       *int    `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	ContactId *string `xml:"contactId,omitempty" json:"contactId,omitempty" yaml:"contactId,omitempty"`
}

// GetContactResponse was auto-generated from WSDL.
type GetContactResponse struct {
	GetContactReturn *string `xml:"getContactReturn,omitempty" json:"getContactReturn,omitempty" yaml:"getContactReturn,omitempty"`
}

// GetDecisionTrees was auto-generated from WSDL.
type GetDecisionTrees struct {
	Sid          *int    `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	PropertyList *string `xml:"propertyList,omitempty" json:"propertyList,omitempty" yaml:"propertyList,omitempty"`
	SortBy       *string `xml:"sortBy,omitempty" json:"sortBy,omitempty" yaml:"sortBy,omitempty"`
	Descending   *bool   `xml:"descending,omitempty" json:"descending,omitempty" yaml:"descending,omitempty"`
}

// GetDecisionTreesResponse was auto-generated from WSDL.
type GetDecisionTreesResponse struct {
	GetDecisionTreesReturn *string `xml:"getDecisionTreesReturn,omitempty" json:"getDecisionTreesReturn,omitempty" yaml:"getDecisionTreesReturn,omitempty"`
}

// GetDependentAttrControls was auto-generated from WSDL.
type GetDependentAttrControls struct {
	Sid      *int           `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	Handle   *string        `xml:"handle,omitempty" json:"handle,omitempty" yaml:"handle,omitempty"`
	AttrVals *ArrayOfString `xml:"attrVals,omitempty" json:"attrVals,omitempty" yaml:"attrVals,omitempty"`
}

// GetDependentAttrControlsResponse was auto-generated from WSDL.
type GetDependentAttrControlsResponse struct {
	GetDependentAttrControlsReturn *string `xml:"getDependentAttrControlsReturn,omitempty" json:"getDependentAttrControlsReturn,omitempty" yaml:"getDependentAttrControlsReturn,omitempty"`
}

// GetDocument was auto-generated from WSDL.
type GetDocument struct {
	Sid          *int    `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	DocId        *int    `xml:"docId,omitempty" json:"docId,omitempty" yaml:"docId,omitempty"`
	PropertyList *string `xml:"propertyList,omitempty" json:"propertyList,omitempty" yaml:"propertyList,omitempty"`
	RelatedDoc   *bool   `xml:"relatedDoc,omitempty" json:"relatedDoc,omitempty" yaml:"relatedDoc,omitempty"`
	GetAttmnt    *bool   `xml:"getAttmnt,omitempty" json:"getAttmnt,omitempty" yaml:"getAttmnt,omitempty"`
	GetHistory   *bool   `xml:"getHistory,omitempty" json:"getHistory,omitempty" yaml:"getHistory,omitempty"`
	GetComment   *bool   `xml:"getComment,omitempty" json:"getComment,omitempty" yaml:"getComment,omitempty"`
	GetNotiList  *bool   `xml:"getNotiList,omitempty" json:"getNotiList,omitempty" yaml:"getNotiList,omitempty"`
}

// GetDocumentResponse was auto-generated from WSDL.
type GetDocumentResponse struct {
	GetDocumentReturn *string `xml:"getDocumentReturn,omitempty" json:"getDocumentReturn,omitempty" yaml:"getDocumentReturn,omitempty"`
}

// GetDocumentTypes was auto-generated from WSDL.
type GetDocumentTypes struct {
	Sid *int `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
}

// GetDocumentTypesResponse was auto-generated from WSDL.
type GetDocumentTypesResponse struct {
	GetDocumentTypesReturn *string `xml:"getDocumentTypesReturn,omitempty" json:"getDocumentTypesReturn,omitempty" yaml:"getDocumentTypesReturn,omitempty"`
}

// GetDocumentsByIDs was auto-generated from WSDL.
type GetDocumentsByIDs struct {
	Sid          *int    `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	DocIds       *string `xml:"docIds,omitempty" json:"docIds,omitempty" yaml:"docIds,omitempty"`
	PropertyList *string `xml:"propertyList,omitempty" json:"propertyList,omitempty" yaml:"propertyList,omitempty"`
	SortBy       *string `xml:"sortBy,omitempty" json:"sortBy,omitempty" yaml:"sortBy,omitempty"`
	Descending   *bool   `xml:"descending,omitempty" json:"descending,omitempty" yaml:"descending,omitempty"`
}

// GetDocumentsByIDsResponse was auto-generated from WSDL.
type GetDocumentsByIDsResponse struct {
	GetDocumentsByIDsReturn *string `xml:"getDocumentsByIDsReturn,omitempty" json:"getDocumentsByIDsReturn,omitempty" yaml:"getDocumentsByIDsReturn,omitempty"`
}

// GetFolderInfo was auto-generated from WSDL.
type GetFolderInfo struct {
	Sid      *int `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	FolderId *int `xml:"folderId,omitempty" json:"folderId,omitempty" yaml:"folderId,omitempty"`
}

// GetFolderInfoResponse was auto-generated from WSDL.
type GetFolderInfoResponse struct {
	GetFolderInfoReturn *string `xml:"getFolderInfoReturn,omitempty" json:"getFolderInfoReturn,omitempty" yaml:"getFolderInfoReturn,omitempty"`
}

// GetFolderList was auto-generated from WSDL.
type GetFolderList struct {
	Sid            *int `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	ParentFolderId *int `xml:"parentFolderId,omitempty" json:"parentFolderId,omitempty" yaml:"parentFolderId,omitempty"`
	RepId          *int `xml:"repId,omitempty" json:"repId,omitempty" yaml:"repId,omitempty"`
}

// GetFolderListResponse was auto-generated from WSDL.
type GetFolderListResponse struct {
	GetFolderListReturn *string `xml:"getFolderListReturn,omitempty" json:"getFolderListReturn,omitempty" yaml:"getFolderListReturn,omitempty"`
}

// GetGroupMemberListValues was auto-generated from WSDL.
type GetGroupMemberListValues struct {
	Sid         *int           `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	WhereClause *string        `xml:"whereClause,omitempty" json:"whereClause,omitempty" yaml:"whereClause,omitempty"`
	NumToFetch  *int           `xml:"numToFetch,omitempty" json:"numToFetch,omitempty" yaml:"numToFetch,omitempty"`
	Attributes  *ArrayOfString `xml:"attributes,omitempty" json:"attributes,omitempty" yaml:"attributes,omitempty"`
}

// GetGroupMemberListValuesResponse was auto-generated from WSDL.
type GetGroupMemberListValuesResponse struct {
	GetGroupMemberListValuesReturn *string `xml:"getGroupMemberListValuesReturn,omitempty" json:"getGroupMemberListValuesReturn,omitempty" yaml:"getGroupMemberListValuesReturn,omitempty"`
}

// GetHandleForUserid was auto-generated from WSDL.
type GetHandleForUserid struct {
	Sid    *int    `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	UserID *string `xml:"userID,omitempty" json:"userID,omitempty" yaml:"userID,omitempty"`
}

// GetHandleForUseridResponse was auto-generated from WSDL.
type GetHandleForUseridResponse struct {
	GetHandleForUseridReturn *string `xml:"getHandleForUseridReturn,omitempty" json:"getHandleForUseridReturn,omitempty" yaml:"getHandleForUseridReturn,omitempty"`
}

// GetKDListPerAttmnt was auto-generated from WSDL.
type GetKDListPerAttmnt struct {
	Sid      *int `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	AttmntId *int `xml:"attmntId,omitempty" json:"attmntId,omitempty" yaml:"attmntId,omitempty"`
}

// GetKDListPerAttmntResponse was auto-generated from WSDL.
type GetKDListPerAttmntResponse struct {
	GetKDListPerAttmntReturn *string `xml:"getKDListPerAttmntReturn,omitempty" json:"getKDListPerAttmntReturn,omitempty" yaml:"getKDListPerAttmntReturn,omitempty"`
}

// GetListValues was auto-generated from WSDL.
type GetListValues struct {
	Sid            *int           `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	ListHandle     *int           `xml:"listHandle,omitempty" json:"listHandle,omitempty" yaml:"listHandle,omitempty"`
	StartIndex     *int           `xml:"startIndex,omitempty" json:"startIndex,omitempty" yaml:"startIndex,omitempty"`
	EndIndex       *int           `xml:"endIndex,omitempty" json:"endIndex,omitempty" yaml:"endIndex,omitempty"`
	AttributeNames *ArrayOfString `xml:"attributeNames,omitempty" json:"attributeNames,omitempty" yaml:"attributeNames,omitempty"`
}

// GetListValuesResponse was auto-generated from WSDL.
type GetListValuesResponse struct {
	GetListValuesReturn *string `xml:"getListValuesReturn,omitempty" json:"getListValuesReturn,omitempty" yaml:"getListValuesReturn,omitempty"`
}

// GetLrelLength was auto-generated from WSDL.
type GetLrelLength struct {
	Sid           *int    `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	ContextObject *string `xml:"contextObject,omitempty" json:"contextObject,omitempty" yaml:"contextObject,omitempty"`
	LrelName      *string `xml:"lrelName,omitempty" json:"lrelName,omitempty" yaml:"lrelName,omitempty"`
}

// GetLrelLengthResponse was auto-generated from WSDL.
type GetLrelLengthResponse struct {
	GetLrelLengthReturn *int `xml:"getLrelLengthReturn,omitempty" json:"getLrelLengthReturn,omitempty" yaml:"getLrelLengthReturn,omitempty"`
}

// GetLrelValues was auto-generated from WSDL.
type GetLrelValues struct {
	Sid           *int           `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	ContextObject *string        `xml:"contextObject,omitempty" json:"contextObject,omitempty" yaml:"contextObject,omitempty"`
	LrelName      *string        `xml:"lrelName,omitempty" json:"lrelName,omitempty" yaml:"lrelName,omitempty"`
	StartIndex    *int           `xml:"startIndex,omitempty" json:"startIndex,omitempty" yaml:"startIndex,omitempty"`
	EndIndex      *int           `xml:"endIndex,omitempty" json:"endIndex,omitempty" yaml:"endIndex,omitempty"`
	Attributes    *ArrayOfString `xml:"attributes,omitempty" json:"attributes,omitempty" yaml:"attributes,omitempty"`
}

// GetLrelValuesResponse was auto-generated from WSDL.
type GetLrelValuesResponse struct {
	GetLrelValuesReturn *string `xml:"getLrelValuesReturn,omitempty" json:"getLrelValuesReturn,omitempty" yaml:"getLrelValuesReturn,omitempty"`
}

// GetNotificationsForContact was auto-generated from WSDL.
type GetNotificationsForContact struct {
	Sid           *int    `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	ContactHandle *string `xml:"contactHandle,omitempty" json:"contactHandle,omitempty" yaml:"contactHandle,omitempty"`
	QueryStatus   *int    `xml:"queryStatus,omitempty" json:"queryStatus,omitempty" yaml:"queryStatus,omitempty"`
}

// GetNotificationsForContactResponse was auto-generated from WSDL.
type GetNotificationsForContactResponse struct {
	GetNotificationsForContactReturn *ListResult `xml:"getNotificationsForContactReturn,omitempty" json:"getNotificationsForContactReturn,omitempty" yaml:"getNotificationsForContactReturn,omitempty"`
}

// GetObjectTypeInformation was auto-generated from WSDL.
type GetObjectTypeInformation struct {
	Sid     *int    `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	Factory *string `xml:"factory,omitempty" json:"factory,omitempty" yaml:"factory,omitempty"`
}

// GetObjectTypeInformationResponse was auto-generated from WSDL.
type GetObjectTypeInformationResponse struct {
	GetObjectTypeInformationReturn *string `xml:"getObjectTypeInformationReturn,omitempty" json:"getObjectTypeInformationReturn,omitempty" yaml:"getObjectTypeInformationReturn,omitempty"`
}

// GetObjectValues was auto-generated from WSDL.
type GetObjectValues struct {
	Sid          *int           `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	ObjectHandle *string        `xml:"objectHandle,omitempty" json:"objectHandle,omitempty" yaml:"objectHandle,omitempty"`
	Attributes   *ArrayOfString `xml:"attributes,omitempty" json:"attributes,omitempty" yaml:"attributes,omitempty"`
}

// GetObjectValuesResponse was auto-generated from WSDL.
type GetObjectValuesResponse struct {
	GetObjectValuesReturn *string `xml:"getObjectValuesReturn,omitempty" json:"getObjectValuesReturn,omitempty" yaml:"getObjectValuesReturn,omitempty"`
}

// GetPendingChangeTaskListForContact was auto-generated from WSDL.
type GetPendingChangeTaskListForContact struct {
	Sid           *int    `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	ContactHandle *string `xml:"contactHandle,omitempty" json:"contactHandle,omitempty" yaml:"contactHandle,omitempty"`
}

// GetPendingChangeTaskListForContactResponse was auto-generated
// from WSDL.
type GetPendingChangeTaskListForContactResponse struct {
	GetPendingChangeTaskListForContactReturn *ListResult `xml:"getPendingChangeTaskListForContactReturn,omitempty" json:"getPendingChangeTaskListForContactReturn,omitempty" yaml:"getPendingChangeTaskListForContactReturn,omitempty"`
}

// GetPendingIssueTaskListForContact was auto-generated from WSDL.
type GetPendingIssueTaskListForContact struct {
	Sid           *int    `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	ContactHandle *string `xml:"contactHandle,omitempty" json:"contactHandle,omitempty" yaml:"contactHandle,omitempty"`
}

// GetPendingIssueTaskListForContactResponse was auto-generated
// from WSDL.
type GetPendingIssueTaskListForContactResponse struct {
	GetPendingIssueTaskListForContactReturn *ListResult `xml:"getPendingIssueTaskListForContactReturn,omitempty" json:"getPendingIssueTaskListForContactReturn,omitempty" yaml:"getPendingIssueTaskListForContactReturn,omitempty"`
}

// GetPermissionGroups was auto-generated from WSDL.
type GetPermissionGroups struct {
	Sid     *int `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	GroupId *int `xml:"groupId,omitempty" json:"groupId,omitempty" yaml:"groupId,omitempty"`
}

// GetPermissionGroupsResponse was auto-generated from WSDL.
type GetPermissionGroupsResponse struct {
	GetPermissionGroupsReturn *string `xml:"getPermissionGroupsReturn,omitempty" json:"getPermissionGroupsReturn,omitempty" yaml:"getPermissionGroupsReturn,omitempty"`
}

// GetPolicyInfo was auto-generated from WSDL.
type GetPolicyInfo struct {
	Sid *int `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
}

// GetPolicyInfoResponse was auto-generated from WSDL.
type GetPolicyInfoResponse struct {
	GetPolicyInfoReturn *string `xml:"getPolicyInfoReturn,omitempty" json:"getPolicyInfoReturn,omitempty" yaml:"getPolicyInfoReturn,omitempty"`
}

// GetPriorities was auto-generated from WSDL.
type GetPriorities struct {
	Sid *int `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
}

// GetPrioritiesResponse was auto-generated from WSDL.
type GetPrioritiesResponse struct {
	GetPrioritiesReturn *string `xml:"getPrioritiesReturn,omitempty" json:"getPrioritiesReturn,omitempty" yaml:"getPrioritiesReturn,omitempty"`
}

// GetPropertyInfoForCategory was auto-generated from WSDL.
type GetPropertyInfoForCategory struct {
	Sid            *int           `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	CategoryHandle *string        `xml:"categoryHandle,omitempty" json:"categoryHandle,omitempty" yaml:"categoryHandle,omitempty"`
	Attributes     *ArrayOfString `xml:"attributes,omitempty" json:"attributes,omitempty" yaml:"attributes,omitempty"`
}

// GetPropertyInfoForCategoryResponse was auto-generated from WSDL.
type GetPropertyInfoForCategoryResponse struct {
	GetPropertyInfoForCategoryReturn *string `xml:"getPropertyInfoForCategoryReturn,omitempty" json:"getPropertyInfoForCategoryReturn,omitempty" yaml:"getPropertyInfoForCategoryReturn,omitempty"`
}

// GetQuestionsAsked was auto-generated from WSDL.
type GetQuestionsAsked struct {
	Sid        *int  `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	ResultSize *int  `xml:"resultSize,omitempty" json:"resultSize,omitempty" yaml:"resultSize,omitempty"`
	Descending *bool `xml:"descending,omitempty" json:"descending,omitempty" yaml:"descending,omitempty"`
}

// GetQuestionsAskedResponse was auto-generated from WSDL.
type GetQuestionsAskedResponse struct {
	GetQuestionsAskedReturn *string `xml:"getQuestionsAskedReturn,omitempty" json:"getQuestionsAskedReturn,omitempty" yaml:"getQuestionsAskedReturn,omitempty"`
}

// GetRelatedList was auto-generated from WSDL.
type GetRelatedList struct {
	Sid          *int    `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	ObjectHandle *string `xml:"objectHandle,omitempty" json:"objectHandle,omitempty" yaml:"objectHandle,omitempty"`
	ListName     *string `xml:"listName,omitempty" json:"listName,omitempty" yaml:"listName,omitempty"`
}

// GetRelatedListResponse was auto-generated from WSDL.
type GetRelatedListResponse struct {
	GetRelatedListReturn *ListResult `xml:"getRelatedListReturn,omitempty" json:"getRelatedListReturn,omitempty" yaml:"getRelatedListReturn,omitempty"`
}

// GetRelatedListValues was auto-generated from WSDL.
type GetRelatedListValues struct {
	Sid                        *int           `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	ObjectHandle               *string        `xml:"objectHandle,omitempty" json:"objectHandle,omitempty" yaml:"objectHandle,omitempty"`
	ListName                   *string        `xml:"listName,omitempty" json:"listName,omitempty" yaml:"listName,omitempty"`
	NumToFetch                 *int           `xml:"numToFetch,omitempty" json:"numToFetch,omitempty" yaml:"numToFetch,omitempty"`
	Attributes                 *ArrayOfString `xml:"attributes,omitempty" json:"attributes,omitempty" yaml:"attributes,omitempty"`
	GetRelatedListValuesResult *string        `xml:"getRelatedListValuesResult,omitempty" json:"getRelatedListValuesResult,omitempty" yaml:"getRelatedListValuesResult,omitempty"`
	NumRowsFound               *int           `xml:"numRowsFound,omitempty" json:"numRowsFound,omitempty" yaml:"numRowsFound,omitempty"`
}

// GetRelatedListValuesResponse was auto-generated from WSDL.
type GetRelatedListValuesResponse struct {
	GetRelatedListValuesResult *string `xml:"getRelatedListValuesResult,omitempty" json:"getRelatedListValuesResult,omitempty" yaml:"getRelatedListValuesResult,omitempty"`
	NumRowsFound               *int    `xml:"numRowsFound,omitempty" json:"numRowsFound,omitempty" yaml:"numRowsFound,omitempty"`
}

// GetRepositoryInfo was auto-generated from WSDL.
type GetRepositoryInfo struct {
	Sid          *int `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	RepositoryId *int `xml:"repositoryId,omitempty" json:"repositoryId,omitempty" yaml:"repositoryId,omitempty"`
}

// GetRepositoryInfoResponse was auto-generated from WSDL.
type GetRepositoryInfoResponse struct {
	GetRepositoryInfoReturn *string `xml:"getRepositoryInfoReturn,omitempty" json:"getRepositoryInfoReturn,omitempty" yaml:"getRepositoryInfoReturn,omitempty"`
}

// GetStatuses was auto-generated from WSDL.
type GetStatuses struct {
	Sid *int `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
}

// GetStatusesResponse was auto-generated from WSDL.
type GetStatusesResponse struct {
	GetStatusesReturn *string `xml:"getStatusesReturn,omitempty" json:"getStatusesReturn,omitempty" yaml:"getStatusesReturn,omitempty"`
}

// GetTaskListValues was auto-generated from WSDL.
type GetTaskListValues struct {
	Sid          *int           `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	ObjectHandle *string        `xml:"objectHandle,omitempty" json:"objectHandle,omitempty" yaml:"objectHandle,omitempty"`
	Attributes   *ArrayOfString `xml:"attributes,omitempty" json:"attributes,omitempty" yaml:"attributes,omitempty"`
}

// GetTaskListValuesResponse was auto-generated from WSDL.
type GetTaskListValuesResponse struct {
	GetTaskListValuesReturn *string `xml:"getTaskListValuesReturn,omitempty" json:"getTaskListValuesReturn,omitempty" yaml:"getTaskListValuesReturn,omitempty"`
}

// GetTemplateList was auto-generated from WSDL.
type GetTemplateList struct {
	Sid *int `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
}

// GetTemplateListResponse was auto-generated from WSDL.
type GetTemplateListResponse struct {
	GetTemplateListReturn *string `xml:"getTemplateListReturn,omitempty" json:"getTemplateListReturn,omitempty" yaml:"getTemplateListReturn,omitempty"`
}

// GetValidTaskTransitions was auto-generated from WSDL.
type GetValidTaskTransitions struct {
	Sid        *int           `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	TaskHandle *string        `xml:"taskHandle,omitempty" json:"taskHandle,omitempty" yaml:"taskHandle,omitempty"`
	Attributes *ArrayOfString `xml:"attributes,omitempty" json:"attributes,omitempty" yaml:"attributes,omitempty"`
}

// GetValidTaskTransitionsResponse was auto-generated from WSDL.
type GetValidTaskTransitionsResponse struct {
	GetValidTaskTransitionsReturn *string `xml:"getValidTaskTransitionsReturn,omitempty" json:"getValidTaskTransitionsReturn,omitempty" yaml:"getValidTaskTransitionsReturn,omitempty"`
}

// GetValidTransitions was auto-generated from WSDL.
type GetValidTransitions struct {
	Sid           *int    `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	Handle        *string `xml:"handle,omitempty" json:"handle,omitempty" yaml:"handle,omitempty"`
	TicketFactory *string `xml:"ticketFactory,omitempty" json:"ticketFactory,omitempty" yaml:"ticketFactory,omitempty"`
}

// GetValidTransitionsResponse was auto-generated from WSDL.
type GetValidTransitionsResponse struct {
	GetValidTransitionsReturn *string `xml:"getValidTransitionsReturn,omitempty" json:"getValidTransitionsReturn,omitempty" yaml:"getValidTransitionsReturn,omitempty"`
}

// GetWorkFlowTemplates was auto-generated from WSDL.
type GetWorkFlowTemplates struct {
	Sid          *int           `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	ObjectHandle *string        `xml:"objectHandle,omitempty" json:"objectHandle,omitempty" yaml:"objectHandle,omitempty"`
	Attributes   *ArrayOfString `xml:"attributes,omitempty" json:"attributes,omitempty" yaml:"attributes,omitempty"`
}

// GetWorkFlowTemplatesResponse was auto-generated from WSDL.
type GetWorkFlowTemplatesResponse struct {
	GetWorkFlowTemplatesReturn *string `xml:"getWorkFlowTemplatesReturn,omitempty" json:"getWorkFlowTemplatesReturn,omitempty" yaml:"getWorkFlowTemplatesReturn,omitempty"`
}

// GetWorkflowTemplateList was auto-generated from WSDL.
type GetWorkflowTemplateList struct {
	Sid *int `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
}

// GetWorkflowTemplateListResponse was auto-generated from WSDL.
type GetWorkflowTemplateListResponse struct {
	GetWorkflowTemplateListReturn *string `xml:"getWorkflowTemplateListReturn,omitempty" json:"getWorkflowTemplateListReturn,omitempty" yaml:"getWorkflowTemplateListReturn,omitempty"`
}

// Impersonate was auto-generated from WSDL.
type Impersonate struct {
	Sid    *int    `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	Userid *string `xml:"userid,omitempty" json:"userid,omitempty" yaml:"userid,omitempty"`
}

// ImpersonateResponse was auto-generated from WSDL.
type ImpersonateResponse struct {
	ImpersonateReturn *int `xml:"impersonateReturn,omitempty" json:"impersonateReturn,omitempty" yaml:"impersonateReturn,omitempty"`
}

// IsAttmntLinkedKD was auto-generated from WSDL.
type IsAttmntLinkedKD struct {
	Sid      *int `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	AttmntId *int `xml:"attmntId,omitempty" json:"attmntId,omitempty" yaml:"attmntId,omitempty"`
}

// IsAttmntLinkedKDResponse was auto-generated from WSDL.
type IsAttmntLinkedKDResponse struct {
	IsAttmntLinkedKDReturn *int `xml:"isAttmntLinkedKDReturn,omitempty" json:"isAttmntLinkedKDReturn,omitempty" yaml:"isAttmntLinkedKDReturn,omitempty"`
}

// LogComment was auto-generated from WSDL.
type LogComment struct {
	Sid          *int    `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	TicketHandle *string `xml:"ticketHandle,omitempty" json:"ticketHandle,omitempty" yaml:"ticketHandle,omitempty"`
	Comment      *string `xml:"comment,omitempty" json:"comment,omitempty" yaml:"comment,omitempty"`
	InternalFlag *int    `xml:"internalFlag,omitempty" json:"internalFlag,omitempty" yaml:"internalFlag,omitempty"`
}

// LogCommentResponse was auto-generated from WSDL.
type LogCommentResponse struct {
}

// Login was auto-generated from WSDL.
type Login struct {
	Username *string `xml:"username,omitempty" json:"username,omitempty" yaml:"username,omitempty"`
	Password *string `xml:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty"`
}

// LoginResponse was auto-generated from WSDL.
type LoginResponse struct {
	LoginReturn *int `xml:"loginReturn,omitempty" json:"loginReturn,omitempty" yaml:"loginReturn,omitempty"`
}

// LoginService was auto-generated from WSDL.
type LoginService struct {
	Username *string `xml:"username,omitempty" json:"username,omitempty" yaml:"username,omitempty"`
	Password *string `xml:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty"`
	Policy   *string `xml:"policy,omitempty" json:"policy,omitempty" yaml:"policy,omitempty"`
}

// LoginServiceManaged was auto-generated from WSDL.
type LoginServiceManaged struct {
	Policy           *string `xml:"policy,omitempty" json:"policy,omitempty" yaml:"policy,omitempty"`
	Encrypted_policy *string `xml:"encrypted_policy,omitempty" json:"encrypted_policy,omitempty" yaml:"encrypted_policy,omitempty"`
}

// LoginServiceManagedResponse was auto-generated from WSDL.
type LoginServiceManagedResponse struct {
	LoginServiceManagedReturn *string `xml:"loginServiceManagedReturn,omitempty" json:"loginServiceManagedReturn,omitempty" yaml:"loginServiceManagedReturn,omitempty"`
}

// LoginServiceResponse was auto-generated from WSDL.
type LoginServiceResponse struct {
	LoginServiceReturn *int `xml:"loginServiceReturn,omitempty" json:"loginServiceReturn,omitempty" yaml:"loginServiceReturn,omitempty"`
}

// LoginWithArtifact was auto-generated from WSDL.
type LoginWithArtifact struct {
	Userid   *string `xml:"userid,omitempty" json:"userid,omitempty" yaml:"userid,omitempty"`
	Artifact *string `xml:"artifact,omitempty" json:"artifact,omitempty" yaml:"artifact,omitempty"`
}

// LoginWithArtifactResponse was auto-generated from WSDL.
type LoginWithArtifactResponse struct {
	LoginWithArtifactReturn *int `xml:"loginWithArtifactReturn,omitempty" json:"loginWithArtifactReturn,omitempty" yaml:"loginWithArtifactReturn,omitempty"`
}

// Logout was auto-generated from WSDL.
type Logout struct {
	Sid *int `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
}

// LogoutResponse was auto-generated from WSDL.
type LogoutResponse struct {
}

// ModifyDocument was auto-generated from WSDL.
type ModifyDocument struct {
	Sid          *int           `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	DocId        *int           `xml:"docId,omitempty" json:"docId,omitempty" yaml:"docId,omitempty"`
	KdAttributes *ArrayOfString `xml:"kdAttributes,omitempty" json:"kdAttributes,omitempty" yaml:"kdAttributes,omitempty"`
}

// ModifyDocumentResponse was auto-generated from WSDL.
type ModifyDocumentResponse struct {
	ModifyDocumentReturn *string `xml:"modifyDocumentReturn,omitempty" json:"modifyDocumentReturn,omitempty" yaml:"modifyDocumentReturn,omitempty"`
}

// NotifyContacts was auto-generated from WSDL.
type NotifyContacts struct {
	Sid           *int           `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	Creator       *string        `xml:"creator,omitempty" json:"creator,omitempty" yaml:"creator,omitempty"`
	ContextObject *string        `xml:"contextObject,omitempty" json:"contextObject,omitempty" yaml:"contextObject,omitempty"`
	MessageTitle  *string        `xml:"messageTitle,omitempty" json:"messageTitle,omitempty" yaml:"messageTitle,omitempty"`
	MessageBody   *string        `xml:"messageBody,omitempty" json:"messageBody,omitempty" yaml:"messageBody,omitempty"`
	NotifyLevel   *int           `xml:"notifyLevel,omitempty" json:"notifyLevel,omitempty" yaml:"notifyLevel,omitempty"`
	Notifyees     *ArrayOfString `xml:"notifyees,omitempty" json:"notifyees,omitempty" yaml:"notifyees,omitempty"`
	Internal      *bool          `xml:"internal,omitempty" json:"internal,omitempty" yaml:"internal,omitempty"`
}

// NotifyContactsResponse was auto-generated from WSDL.
type NotifyContactsResponse struct {
	NotifyContactsReturn *string `xml:"notifyContactsReturn,omitempty" json:"notifyContactsReturn,omitempty" yaml:"notifyContactsReturn,omitempty"`
}

// RateDocument was auto-generated from WSDL.
type RateDocument struct {
	Sid              *int    `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	DocId            *int    `xml:"docId,omitempty" json:"docId,omitempty" yaml:"docId,omitempty"`
	Rating           *int    `xml:"rating,omitempty" json:"rating,omitempty" yaml:"rating,omitempty"`
	Multiplier       *int    `xml:"multiplier,omitempty" json:"multiplier,omitempty" yaml:"multiplier,omitempty"`
	TicketPerId      *string `xml:"ticketPerId,omitempty" json:"ticketPerId,omitempty" yaml:"ticketPerId,omitempty"`
	OnTicketAccept   *bool   `xml:"onTicketAccept,omitempty" json:"onTicketAccept,omitempty" yaml:"onTicketAccept,omitempty"`
	SolveUserProblem *bool   `xml:"solveUserProblem,omitempty" json:"solveUserProblem,omitempty" yaml:"solveUserProblem,omitempty"`
	IsDefault        *bool   `xml:"isDefault,omitempty" json:"isDefault,omitempty" yaml:"isDefault,omitempty"`
}

// RateDocumentResponse was auto-generated from WSDL.
type RateDocumentResponse struct {
	RateDocumentReturn *string `xml:"rateDocumentReturn,omitempty" json:"rateDocumentReturn,omitempty" yaml:"rateDocumentReturn,omitempty"`
}

// RemoveAttachment was auto-generated from WSDL.
type RemoveAttachment struct {
	Sid       *int    `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	AttHandle *string `xml:"attHandle,omitempty" json:"attHandle,omitempty" yaml:"attHandle,omitempty"`
}

// RemoveAttachmentResponse was auto-generated from WSDL.
type RemoveAttachmentResponse struct {
	RemoveAttachmentReturn *int `xml:"removeAttachmentReturn,omitempty" json:"removeAttachmentReturn,omitempty" yaml:"removeAttachmentReturn,omitempty"`
}

// RemoveLrelRelationships was auto-generated from WSDL.
type RemoveLrelRelationships struct {
	Sid                 *int           `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	ContextObject       *string        `xml:"contextObject,omitempty" json:"contextObject,omitempty" yaml:"contextObject,omitempty"`
	LrelName            *string        `xml:"lrelName,omitempty" json:"lrelName,omitempty" yaml:"lrelName,omitempty"`
	RemoveObjectHandles *ArrayOfString `xml:"removeObjectHandles,omitempty" json:"removeObjectHandles,omitempty" yaml:"removeObjectHandles,omitempty"`
}

// RemoveLrelRelationshipsResponse was auto-generated from WSDL.
type RemoveLrelRelationshipsResponse struct {
}

// RemoveMemberFromGroup was auto-generated from WSDL.
type RemoveMemberFromGroup struct {
	Sid           *int    `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	ContactHandle *string `xml:"contactHandle,omitempty" json:"contactHandle,omitempty" yaml:"contactHandle,omitempty"`
	GroupHandle   *string `xml:"groupHandle,omitempty" json:"groupHandle,omitempty" yaml:"groupHandle,omitempty"`
}

// RemoveMemberFromGroupResponse was auto-generated from WSDL.
type RemoveMemberFromGroupResponse struct {
}

// Search was auto-generated from WSDL.
type Search struct {
	Sid               *int    `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	Problem           *string `xml:"problem,omitempty" json:"problem,omitempty" yaml:"problem,omitempty"`
	ResultSize        *int    `xml:"resultSize,omitempty" json:"resultSize,omitempty" yaml:"resultSize,omitempty"`
	Properties        *string `xml:"properties,omitempty" json:"properties,omitempty" yaml:"properties,omitempty"`
	SortBy            *string `xml:"sortBy,omitempty" json:"sortBy,omitempty" yaml:"sortBy,omitempty"`
	Descending        *bool   `xml:"descending,omitempty" json:"descending,omitempty" yaml:"descending,omitempty"`
	RelatedCategories *bool   `xml:"relatedCategories,omitempty" json:"relatedCategories,omitempty" yaml:"relatedCategories,omitempty"`
	SearchType        *int    `xml:"searchType,omitempty" json:"searchType,omitempty" yaml:"searchType,omitempty"`
	MatchType         *int    `xml:"matchType,omitempty" json:"matchType,omitempty" yaml:"matchType,omitempty"`
	SearchField       *int    `xml:"searchField,omitempty" json:"searchField,omitempty" yaml:"searchField,omitempty"`
	CategoryPath      *string `xml:"categoryPath,omitempty" json:"categoryPath,omitempty" yaml:"categoryPath,omitempty"`
	WhereClause       *string `xml:"whereClause,omitempty" json:"whereClause,omitempty" yaml:"whereClause,omitempty"`
	MaxDocIDs         *int    `xml:"maxDocIDs,omitempty" json:"maxDocIDs,omitempty" yaml:"maxDocIDs,omitempty"`
}

// SearchResponse was auto-generated from WSDL.
type SearchResponse struct {
	SearchReturn *string `xml:"searchReturn,omitempty" json:"searchReturn,omitempty" yaml:"searchReturn,omitempty"`
}

// ServerStatus was auto-generated from WSDL.
type ServerStatus struct {
	Sid *int `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
}

// ServerStatusResponse was auto-generated from WSDL.
type ServerStatusResponse struct {
	ServerStatusReturn *int `xml:"serverStatusReturn,omitempty" json:"serverStatusReturn,omitempty" yaml:"serverStatusReturn,omitempty"`
}

// Transfer was auto-generated from WSDL.
type Transfer struct {
	Sid                   *int    `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	Creator               *string `xml:"creator,omitempty" json:"creator,omitempty" yaml:"creator,omitempty"`
	ObjectHandle          *string `xml:"objectHandle,omitempty" json:"objectHandle,omitempty" yaml:"objectHandle,omitempty"`
	Description           *string `xml:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty"`
	SetAssignee           *bool   `xml:"setAssignee,omitempty" json:"setAssignee,omitempty" yaml:"setAssignee,omitempty"`
	NewAssigneeHandle     *string `xml:"newAssigneeHandle,omitempty" json:"newAssigneeHandle,omitempty" yaml:"newAssigneeHandle,omitempty"`
	SetGroup              *bool   `xml:"setGroup,omitempty" json:"setGroup,omitempty" yaml:"setGroup,omitempty"`
	NewGroupHandle        *string `xml:"newGroupHandle,omitempty" json:"newGroupHandle,omitempty" yaml:"newGroupHandle,omitempty"`
	SetOrganization       *bool   `xml:"setOrganization,omitempty" json:"setOrganization,omitempty" yaml:"setOrganization,omitempty"`
	NewOrganizationHandle *string `xml:"newOrganizationHandle,omitempty" json:"newOrganizationHandle,omitempty" yaml:"newOrganizationHandle,omitempty"`
}

// TransferResponse was auto-generated from WSDL.
type TransferResponse struct {
	TransferReturn *string `xml:"transferReturn,omitempty" json:"transferReturn,omitempty" yaml:"transferReturn,omitempty"`
}

// UpdateObject was auto-generated from WSDL.
type UpdateObject struct {
	Sid          *int           `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	ObjectHandle *string        `xml:"objectHandle,omitempty" json:"objectHandle,omitempty" yaml:"objectHandle,omitempty"`
	AttrVals     *ArrayOfString `xml:"attrVals,omitempty" json:"attrVals,omitempty" yaml:"attrVals,omitempty"`
	Attributes   *ArrayOfString `xml:"attributes,omitempty" json:"attributes,omitempty" yaml:"attributes,omitempty"`
}

// UpdateObjectResponse was auto-generated from WSDL.
type UpdateObjectResponse struct {
	UpdateObjectReturn *string `xml:"updateObjectReturn,omitempty" json:"updateObjectReturn,omitempty" yaml:"updateObjectReturn,omitempty"`
}

// UpdateRating was auto-generated from WSDL.
type UpdateRating struct {
	Sid  *int `xml:"sid,omitempty" json:"sid,omitempty" yaml:"sid,omitempty"`
	BuId *int `xml:"buId,omitempty" json:"buId,omitempty" yaml:"buId,omitempty"`
	Rate *int `xml:"rate,omitempty" json:"rate,omitempty" yaml:"rate,omitempty"`
}

// UpdateRatingResponse was auto-generated from WSDL.
type UpdateRatingResponse struct {
	UpdateRatingReturn *string `xml:"updateRatingReturn,omitempty" json:"updateRatingReturn,omitempty" yaml:"updateRatingReturn,omitempty"`
}

// Operation wrapper for AddAssetLog.
// OperationAddAssetLogRequest was auto-generated from WSDL.
type OperationAddAssetLogRequest struct {
	AddAssetLog *AddAssetLog `xml:"addAssetLog,omitempty" json:"addAssetLog,omitempty" yaml:"addAssetLog,omitempty"`
}

// Operation wrapper for AddAssetLog.
// OperationAddAssetLogResponse was auto-generated from WSDL.
type OperationAddAssetLogResponse struct {
	AddAssetLogResponse *AddAssetLogResponse `xml:"addAssetLogResponse,omitempty" json:"addAssetLogResponse,omitempty" yaml:"addAssetLogResponse,omitempty"`
}

// Operation wrapper for AddBookmark.
// OperationAddBookmarkRequest was auto-generated from WSDL.
type OperationAddBookmarkRequest struct {
	AddBookmark *AddBookmark `xml:"addBookmark,omitempty" json:"addBookmark,omitempty" yaml:"addBookmark,omitempty"`
}

// Operation wrapper for AddBookmark.
// OperationAddBookmarkResponse was auto-generated from WSDL.
type OperationAddBookmarkResponse struct {
	AddBookmarkResponse *AddBookmarkResponse `xml:"addBookmarkResponse,omitempty" json:"addBookmarkResponse,omitempty" yaml:"addBookmarkResponse,omitempty"`
}

// Operation wrapper for AddComment.
// OperationAddCommentRequest was auto-generated from WSDL.
type OperationAddCommentRequest struct {
	AddComment *AddComment `xml:"addComment,omitempty" json:"addComment,omitempty" yaml:"addComment,omitempty"`
}

// Operation wrapper for AddComment.
// OperationAddCommentResponse was auto-generated from WSDL.
type OperationAddCommentResponse struct {
	AddCommentResponse *AddCommentResponse `xml:"addCommentResponse,omitempty" json:"addCommentResponse,omitempty" yaml:"addCommentResponse,omitempty"`
}

// Operation wrapper for AddMemberToGroup.
// OperationAddMemberToGroupRequest was auto-generated from WSDL.
type OperationAddMemberToGroupRequest struct {
	AddMemberToGroup *AddMemberToGroup `xml:"addMemberToGroup,omitempty" json:"addMemberToGroup,omitempty" yaml:"addMemberToGroup,omitempty"`
}

// Operation wrapper for AddMemberToGroup.
// OperationAddMemberToGroupResponse was auto-generated from WSDL.
type OperationAddMemberToGroupResponse struct {
	AddMemberToGroupResponse *AddMemberToGroupResponse `xml:"addMemberToGroupResponse,omitempty" json:"addMemberToGroupResponse,omitempty" yaml:"addMemberToGroupResponse,omitempty"`
}

// Operation wrapper for AttachChangeToRequest.
// OperationAttachChangeToRequestRequest was auto-generated from
// WSDL.
type OperationAttachChangeToRequestRequest struct {
	AttachChangeToRequest *AttachChangeToRequest `xml:"attachChangeToRequest,omitempty" json:"attachChangeToRequest,omitempty" yaml:"attachChangeToRequest,omitempty"`
}

// Operation wrapper for AttachChangeToRequest.
// OperationAttachChangeToRequestResponse was auto-generated from
// WSDL.
type OperationAttachChangeToRequestResponse struct {
	AttachChangeToRequestResponse *AttachChangeToRequestResponse `xml:"attachChangeToRequestResponse,omitempty" json:"attachChangeToRequestResponse,omitempty" yaml:"attachChangeToRequestResponse,omitempty"`
}

// Operation wrapper for AttachURLLink.
// OperationAttachURLLinkRequest was auto-generated from WSDL.
type OperationAttachURLLinkRequest struct {
	AttachURLLink *AttachURLLink `xml:"attachURLLink,omitempty" json:"attachURLLink,omitempty" yaml:"attachURLLink,omitempty"`
}

// Operation wrapper for AttachURLLink.
// OperationAttachURLLinkResponse was auto-generated from WSDL.
type OperationAttachURLLinkResponse struct {
	AttachURLLinkResponse *AttachURLLinkResponse `xml:"attachURLLinkResponse,omitempty" json:"attachURLLinkResponse,omitempty" yaml:"attachURLLinkResponse,omitempty"`
}

// Operation wrapper for AttachURLLinkToTicket.
// OperationAttachURLLinkToTicketRequest was auto-generated from
// WSDL.
type OperationAttachURLLinkToTicketRequest struct {
	AttachURLLinkToTicket *AttachURLLinkToTicket `xml:"attachURLLinkToTicket,omitempty" json:"attachURLLinkToTicket,omitempty" yaml:"attachURLLinkToTicket,omitempty"`
}

// Operation wrapper for AttachURLLinkToTicket.
// OperationAttachURLLinkToTicketResponse was auto-generated from
// WSDL.
type OperationAttachURLLinkToTicketResponse struct {
	AttachURLLinkToTicketResponse *AttachURLLinkToTicketResponse `xml:"attachURLLinkToTicketResponse,omitempty" json:"attachURLLinkToTicketResponse,omitempty" yaml:"attachURLLinkToTicketResponse,omitempty"`
}

// Operation wrapper for AttmntFolderLinkCount.
// OperationAttmntFolderLinkCountRequest was auto-generated from
// WSDL.
type OperationAttmntFolderLinkCountRequest struct {
	AttmntFolderLinkCount *AttmntFolderLinkCount `xml:"attmntFolderLinkCount,omitempty" json:"attmntFolderLinkCount,omitempty" yaml:"attmntFolderLinkCount,omitempty"`
}

// Operation wrapper for AttmntFolderLinkCount.
// OperationAttmntFolderLinkCountResponse was auto-generated from
// WSDL.
type OperationAttmntFolderLinkCountResponse struct {
	AttmntFolderLinkCountResponse *AttmntFolderLinkCountResponse `xml:"attmntFolderLinkCountResponse,omitempty" json:"attmntFolderLinkCountResponse,omitempty" yaml:"attmntFolderLinkCountResponse,omitempty"`
}

// Operation wrapper for CallServerMethod.
// OperationCallServerMethodRequest was auto-generated from WSDL.
type OperationCallServerMethodRequest struct {
	CallServerMethod *CallServerMethod `xml:"callServerMethod,omitempty" json:"callServerMethod,omitempty" yaml:"callServerMethod,omitempty"`
}

// Operation wrapper for CallServerMethod.
// OperationCallServerMethodResponse was auto-generated from WSDL.
type OperationCallServerMethodResponse struct {
	CallServerMethodResponse *CallServerMethodResponse `xml:"callServerMethodResponse,omitempty" json:"callServerMethodResponse,omitempty" yaml:"callServerMethodResponse,omitempty"`
}

// Operation wrapper for ChangeStatus.
// OperationChangeStatusRequest was auto-generated from WSDL.
type OperationChangeStatusRequest struct {
	ChangeStatus *ChangeStatus `xml:"changeStatus,omitempty" json:"changeStatus,omitempty" yaml:"changeStatus,omitempty"`
}

// Operation wrapper for ChangeStatus.
// OperationChangeStatusResponse was auto-generated from WSDL.
type OperationChangeStatusResponse struct {
	ChangeStatusResponse *ChangeStatusResponse `xml:"changeStatusResponse,omitempty" json:"changeStatusResponse,omitempty" yaml:"changeStatusResponse,omitempty"`
}

// Operation wrapper for ClearNotification.
// OperationClearNotificationRequest was auto-generated from WSDL.
type OperationClearNotificationRequest struct {
	ClearNotification *ClearNotification `xml:"clearNotification,omitempty" json:"clearNotification,omitempty" yaml:"clearNotification,omitempty"`
}

// Operation wrapper for ClearNotification.
// OperationClearNotificationResponse was auto-generated from WSDL.
type OperationClearNotificationResponse struct {
	ClearNotificationResponse *ClearNotificationResponse `xml:"clearNotificationResponse,omitempty" json:"clearNotificationResponse,omitempty" yaml:"clearNotificationResponse,omitempty"`
}

// Operation wrapper for CloseTicket.
// OperationCloseTicketRequest was auto-generated from WSDL.
type OperationCloseTicketRequest struct {
	CloseTicket *CloseTicket `xml:"closeTicket,omitempty" json:"closeTicket,omitempty" yaml:"closeTicket,omitempty"`
}

// Operation wrapper for CloseTicket.
// OperationCloseTicketResponse was auto-generated from WSDL.
type OperationCloseTicketResponse struct {
	CloseTicketResponse *CloseTicketResponse `xml:"closeTicketResponse,omitempty" json:"closeTicketResponse,omitempty" yaml:"closeTicketResponse,omitempty"`
}

// Operation wrapper for CreateActivityLog.
// OperationCreateActivityLogRequest was auto-generated from WSDL.
type OperationCreateActivityLogRequest struct {
	CreateActivityLog *CreateActivityLog `xml:"createActivityLog,omitempty" json:"createActivityLog,omitempty" yaml:"createActivityLog,omitempty"`
}

// Operation wrapper for CreateActivityLog.
// OperationCreateActivityLogResponse was auto-generated from WSDL.
type OperationCreateActivityLogResponse struct {
	CreateActivityLogResponse *CreateActivityLogResponse `xml:"createActivityLogResponse,omitempty" json:"createActivityLogResponse,omitempty" yaml:"createActivityLogResponse,omitempty"`
}

// Operation wrapper for CreateAsset.
// OperationCreateAssetRequest was auto-generated from WSDL.
type OperationCreateAssetRequest struct {
	CreateAsset *CreateAsset `xml:"createAsset,omitempty" json:"createAsset,omitempty" yaml:"createAsset,omitempty"`
}

// Operation wrapper for CreateAsset.
// OperationCreateAssetResponse was auto-generated from WSDL.
type OperationCreateAssetResponse struct {
	CreateAssetResponse *CreateAssetResponse `xml:"createAssetResponse,omitempty" json:"createAssetResponse,omitempty" yaml:"createAssetResponse,omitempty"`
}

// Operation wrapper for CreateAssetParentChildRelationship.
// OperationCreateAssetParentChildRelationshipRequest was auto-generated
// from WSDL.
type OperationCreateAssetParentChildRelationshipRequest struct {
	CreateAssetParentChildRelationship *CreateAssetParentChildRelationship `xml:"createAssetParentChildRelationship,omitempty" json:"createAssetParentChildRelationship,omitempty" yaml:"createAssetParentChildRelationship,omitempty"`
}

// Operation wrapper for CreateAssetParentChildRelationship.
// OperationCreateAssetParentChildRelationshipResponse was auto-generated
// from WSDL.
type OperationCreateAssetParentChildRelationshipResponse struct {
	CreateAssetParentChildRelationshipResponse *CreateAssetParentChildRelationshipResponse `xml:"createAssetParentChildRelationshipResponse,omitempty" json:"createAssetParentChildRelationshipResponse,omitempty" yaml:"createAssetParentChildRelationshipResponse,omitempty"`
}

// Operation wrapper for CreateAttachment.
// OperationCreateAttachmentRequest was auto-generated from WSDL.
type OperationCreateAttachmentRequest struct {
	CreateAttachment *CreateAttachment `xml:"createAttachment,omitempty" json:"createAttachment,omitempty" yaml:"createAttachment,omitempty"`
}

// Operation wrapper for CreateAttachment.
// OperationCreateAttachmentResponse was auto-generated from WSDL.
type OperationCreateAttachmentResponse struct {
	CreateAttachmentResponse *CreateAttachmentResponse `xml:"createAttachmentResponse,omitempty" json:"createAttachmentResponse,omitempty" yaml:"createAttachmentResponse,omitempty"`
}

// Operation wrapper for CreateAttmnt.
// OperationCreateAttmntRequest was auto-generated from WSDL.
type OperationCreateAttmntRequest struct {
	CreateAttmnt *CreateAttmnt `xml:"createAttmnt,omitempty" json:"createAttmnt,omitempty" yaml:"createAttmnt,omitempty"`
}

// Operation wrapper for CreateAttmnt.
// OperationCreateAttmntResponse was auto-generated from WSDL.
type OperationCreateAttmntResponse struct {
	CreateAttmntResponse *CreateAttmntResponse `xml:"createAttmntResponse,omitempty" json:"createAttmntResponse,omitempty" yaml:"createAttmntResponse,omitempty"`
}

// Operation wrapper for CreateChangeOrder.
// OperationCreateChangeOrderRequest was auto-generated from WSDL.
type OperationCreateChangeOrderRequest struct {
	CreateChangeOrder *CreateChangeOrder `xml:"createChangeOrder,omitempty" json:"createChangeOrder,omitempty" yaml:"createChangeOrder,omitempty"`
}

// Operation wrapper for CreateChangeOrder.
// OperationCreateChangeOrderResponse was auto-generated from WSDL.
type OperationCreateChangeOrderResponse struct {
	CreateChangeOrderResponse *CreateChangeOrderResponse `xml:"createChangeOrderResponse,omitempty" json:"createChangeOrderResponse,omitempty" yaml:"createChangeOrderResponse,omitempty"`
}

// Operation wrapper for CreateDocument.
// OperationCreateDocumentRequest was auto-generated from WSDL.
type OperationCreateDocumentRequest struct {
	CreateDocument *CreateDocument `xml:"createDocument,omitempty" json:"createDocument,omitempty" yaml:"createDocument,omitempty"`
}

// Operation wrapper for CreateDocument.
// OperationCreateDocumentResponse was auto-generated from WSDL.
type OperationCreateDocumentResponse struct {
	CreateDocumentResponse *CreateDocumentResponse `xml:"createDocumentResponse,omitempty" json:"createDocumentResponse,omitempty" yaml:"createDocumentResponse,omitempty"`
}

// Operation wrapper for CreateFolder.
// OperationCreateFolderRequest was auto-generated from WSDL.
type OperationCreateFolderRequest struct {
	CreateFolder *CreateFolder `xml:"createFolder,omitempty" json:"createFolder,omitempty" yaml:"createFolder,omitempty"`
}

// Operation wrapper for CreateFolder.
// OperationCreateFolderResponse was auto-generated from WSDL.
type OperationCreateFolderResponse struct {
	CreateFolderResponse *CreateFolderResponse `xml:"createFolderResponse,omitempty" json:"createFolderResponse,omitempty" yaml:"createFolderResponse,omitempty"`
}

// Operation wrapper for CreateIssue.
// OperationCreateIssueRequest was auto-generated from WSDL.
type OperationCreateIssueRequest struct {
	CreateIssue *CreateIssue `xml:"createIssue,omitempty" json:"createIssue,omitempty" yaml:"createIssue,omitempty"`
}

// Operation wrapper for CreateIssue.
// OperationCreateIssueResponse was auto-generated from WSDL.
type OperationCreateIssueResponse struct {
	CreateIssueResponse *CreateIssueResponse `xml:"createIssueResponse,omitempty" json:"createIssueResponse,omitempty" yaml:"createIssueResponse,omitempty"`
}

// Operation wrapper for CreateLrelRelationships.
// OperationCreateLrelRelationshipsRequest was auto-generated from
// WSDL.
type OperationCreateLrelRelationshipsRequest struct {
	CreateLrelRelationships *CreateLrelRelationships `xml:"createLrelRelationships,omitempty" json:"createLrelRelationships,omitempty" yaml:"createLrelRelationships,omitempty"`
}

// Operation wrapper for CreateLrelRelationships.
// OperationCreateLrelRelationshipsResponse was auto-generated
// from WSDL.
type OperationCreateLrelRelationshipsResponse struct {
	CreateLrelRelationshipsResponse *CreateLrelRelationshipsResponse `xml:"createLrelRelationshipsResponse,omitempty" json:"createLrelRelationshipsResponse,omitempty" yaml:"createLrelRelationshipsResponse,omitempty"`
}

// Operation wrapper for CreateObject.
// OperationCreateObjectRequest was auto-generated from WSDL.
type OperationCreateObjectRequest struct {
	CreateObject *CreateObject `xml:"createObject,omitempty" json:"createObject,omitempty" yaml:"createObject,omitempty"`
}

// Operation wrapper for CreateObject.
// OperationCreateObjectResponse was auto-generated from WSDL.
type OperationCreateObjectResponse struct {
	CreateObjectResponse *CreateObjectResponse `xml:"createObjectResponse,omitempty" json:"createObjectResponse,omitempty" yaml:"createObjectResponse,omitempty"`
}

// Operation wrapper for CreateQuickTicket.
// OperationCreateQuickTicketRequest was auto-generated from WSDL.
type OperationCreateQuickTicketRequest struct {
	CreateQuickTicket *CreateQuickTicket `xml:"createQuickTicket,omitempty" json:"createQuickTicket,omitempty" yaml:"createQuickTicket,omitempty"`
}

// Operation wrapper for CreateQuickTicket.
// OperationCreateQuickTicketResponse was auto-generated from WSDL.
type OperationCreateQuickTicketResponse struct {
	CreateQuickTicketResponse *CreateQuickTicketResponse `xml:"createQuickTicketResponse,omitempty" json:"createQuickTicketResponse,omitempty" yaml:"createQuickTicketResponse,omitempty"`
}

// Operation wrapper for CreateRequest.
// OperationCreateRequestRequest was auto-generated from WSDL.
type OperationCreateRequestRequest struct {
	CreateRequest *CreateRequest `xml:"createRequest,omitempty" json:"createRequest,omitempty" yaml:"createRequest,omitempty"`
}

// Operation wrapper for CreateRequest.
// OperationCreateRequestResponse was auto-generated from WSDL.
type OperationCreateRequestResponse struct {
	CreateRequestResponse *CreateRequestResponse `xml:"createRequestResponse,omitempty" json:"createRequestResponse,omitempty" yaml:"createRequestResponse,omitempty"`
}

// Operation wrapper for CreateTicket.
// OperationCreateTicketRequest was auto-generated from WSDL.
type OperationCreateTicketRequest struct {
	CreateTicket *CreateTicket `xml:"createTicket,omitempty" json:"createTicket,omitempty" yaml:"createTicket,omitempty"`
}

// Operation wrapper for CreateTicket.
// OperationCreateTicketResponse was auto-generated from WSDL.
type OperationCreateTicketResponse struct {
	CreateTicketResponse *CreateTicketResponse `xml:"createTicketResponse,omitempty" json:"createTicketResponse,omitempty" yaml:"createTicketResponse,omitempty"`
}

// Operation wrapper for CreateWorkFlowTask.
// OperationCreateWorkFlowTaskRequest was auto-generated from WSDL.
type OperationCreateWorkFlowTaskRequest struct {
	CreateWorkFlowTask *CreateWorkFlowTask `xml:"createWorkFlowTask,omitempty" json:"createWorkFlowTask,omitempty" yaml:"createWorkFlowTask,omitempty"`
}

// Operation wrapper for CreateWorkFlowTask.
// OperationCreateWorkFlowTaskResponse was auto-generated from
// WSDL.
type OperationCreateWorkFlowTaskResponse struct {
	CreateWorkFlowTaskResponse *CreateWorkFlowTaskResponse `xml:"createWorkFlowTaskResponse,omitempty" json:"createWorkFlowTaskResponse,omitempty" yaml:"createWorkFlowTaskResponse,omitempty"`
}

// Operation wrapper for DeleteBookmark.
// OperationDeleteBookmarkRequest was auto-generated from WSDL.
type OperationDeleteBookmarkRequest struct {
	DeleteBookmark *DeleteBookmark `xml:"deleteBookmark,omitempty" json:"deleteBookmark,omitempty" yaml:"deleteBookmark,omitempty"`
}

// Operation wrapper for DeleteBookmark.
// OperationDeleteBookmarkResponse was auto-generated from WSDL.
type OperationDeleteBookmarkResponse struct {
	DeleteBookmarkResponse *DeleteBookmarkResponse `xml:"deleteBookmarkResponse,omitempty" json:"deleteBookmarkResponse,omitempty" yaml:"deleteBookmarkResponse,omitempty"`
}

// Operation wrapper for DeleteComment.
// OperationDeleteCommentRequest was auto-generated from WSDL.
type OperationDeleteCommentRequest struct {
	DeleteComment *DeleteComment `xml:"deleteComment,omitempty" json:"deleteComment,omitempty" yaml:"deleteComment,omitempty"`
}

// Operation wrapper for DeleteComment.
// OperationDeleteCommentResponse was auto-generated from WSDL.
type OperationDeleteCommentResponse struct {
	DeleteCommentResponse *DeleteCommentResponse `xml:"deleteCommentResponse,omitempty" json:"deleteCommentResponse,omitempty" yaml:"deleteCommentResponse,omitempty"`
}

// Operation wrapper for DeleteDocument.
// OperationDeleteDocumentRequest was auto-generated from WSDL.
type OperationDeleteDocumentRequest struct {
	DeleteDocument *DeleteDocument `xml:"deleteDocument,omitempty" json:"deleteDocument,omitempty" yaml:"deleteDocument,omitempty"`
}

// Operation wrapper for DeleteDocument.
// OperationDeleteDocumentResponse was auto-generated from WSDL.
type OperationDeleteDocumentResponse struct {
	DeleteDocumentResponse *DeleteDocumentResponse `xml:"deleteDocumentResponse,omitempty" json:"deleteDocumentResponse,omitempty" yaml:"deleteDocumentResponse,omitempty"`
}

// Operation wrapper for DeleteWorkFlowTask.
// OperationDeleteWorkFlowTaskRequest was auto-generated from WSDL.
type OperationDeleteWorkFlowTaskRequest struct {
	DeleteWorkFlowTask *DeleteWorkFlowTask `xml:"deleteWorkFlowTask,omitempty" json:"deleteWorkFlowTask,omitempty" yaml:"deleteWorkFlowTask,omitempty"`
}

// Operation wrapper for DeleteWorkFlowTask.
// OperationDeleteWorkFlowTaskResponse was auto-generated from
// WSDL.
type OperationDeleteWorkFlowTaskResponse struct {
	DeleteWorkFlowTaskResponse *DeleteWorkFlowTaskResponse `xml:"deleteWorkFlowTaskResponse,omitempty" json:"deleteWorkFlowTaskResponse,omitempty" yaml:"deleteWorkFlowTaskResponse,omitempty"`
}

// Operation wrapper for DetachChangeFromRequest.
// OperationDetachChangeFromRequestRequest was auto-generated from
// WSDL.
type OperationDetachChangeFromRequestRequest struct {
	DetachChangeFromRequest *DetachChangeFromRequest `xml:"detachChangeFromRequest,omitempty" json:"detachChangeFromRequest,omitempty" yaml:"detachChangeFromRequest,omitempty"`
}

// Operation wrapper for DetachChangeFromRequest.
// OperationDetachChangeFromRequestResponse was auto-generated
// from WSDL.
type OperationDetachChangeFromRequestResponse struct {
	DetachChangeFromRequestResponse *DetachChangeFromRequestResponse `xml:"detachChangeFromRequestResponse,omitempty" json:"detachChangeFromRequestResponse,omitempty" yaml:"detachChangeFromRequestResponse,omitempty"`
}

// Operation wrapper for DoQuery.
// OperationDoQueryRequest was auto-generated from WSDL.
type OperationDoQueryRequest struct {
	DoQuery *DoQuery `xml:"doQuery,omitempty" json:"doQuery,omitempty" yaml:"doQuery,omitempty"`
}

// Operation wrapper for DoQuery.
// OperationDoQueryResponse was auto-generated from WSDL.
type OperationDoQueryResponse struct {
	DoQueryResponse *DoQueryResponse `xml:"doQueryResponse,omitempty" json:"doQueryResponse,omitempty" yaml:"doQueryResponse,omitempty"`
}

// Operation wrapper for DoSelect.
// OperationDoSelectRequest was auto-generated from WSDL.
type OperationDoSelectRequest struct {
	DoSelect *DoSelect `xml:"doSelect,omitempty" json:"doSelect,omitempty" yaml:"doSelect,omitempty"`
}

// Operation wrapper for DoSelect.
// OperationDoSelectResponse was auto-generated from WSDL.
type OperationDoSelectResponse struct {
	DoSelectResponse *DoSelectResponse `xml:"doSelectResponse,omitempty" json:"doSelectResponse,omitempty" yaml:"doSelectResponse,omitempty"`
}

// Operation wrapper for DoSelectKD.
// OperationDoSelectKDRequest was auto-generated from WSDL.
type OperationDoSelectKDRequest struct {
	DoSelectKD *DoSelectKD `xml:"doSelectKD,omitempty" json:"doSelectKD,omitempty" yaml:"doSelectKD,omitempty"`
}

// Operation wrapper for DoSelectKD.
// OperationDoSelectKDResponse was auto-generated from WSDL.
type OperationDoSelectKDResponse struct {
	DoSelectKDResponse *DoSelectKDResponse `xml:"doSelectKDResponse,omitempty" json:"doSelectKDResponse,omitempty" yaml:"doSelectKDResponse,omitempty"`
}

// Operation wrapper for Escalate.
// OperationEscalateRequest was auto-generated from WSDL.
type OperationEscalateRequest struct {
	Escalate *Escalate `xml:"escalate,omitempty" json:"escalate,omitempty" yaml:"escalate,omitempty"`
}

// Operation wrapper for Escalate.
// OperationEscalateResponse was auto-generated from WSDL.
type OperationEscalateResponse struct {
	EscalateResponse *EscalateResponse `xml:"escalateResponse,omitempty" json:"escalateResponse,omitempty" yaml:"escalateResponse,omitempty"`
}

// Operation wrapper for Faq.
// OperationFaqRequest was auto-generated from WSDL.
type OperationFaqRequest struct {
	Faq *Faq `xml:"faq,omitempty" json:"faq,omitempty" yaml:"faq,omitempty"`
}

// Operation wrapper for Faq.
// OperationFaqResponse was auto-generated from WSDL.
type OperationFaqResponse struct {
	FaqResponse *FaqResponse `xml:"faqResponse,omitempty" json:"faqResponse,omitempty" yaml:"faqResponse,omitempty"`
}

// Operation wrapper for FindContacts.
// OperationFindContactsRequest was auto-generated from WSDL.
type OperationFindContactsRequest struct {
	FindContacts *FindContacts `xml:"findContacts,omitempty" json:"findContacts,omitempty" yaml:"findContacts,omitempty"`
}

// Operation wrapper for FindContacts.
// OperationFindContactsResponse was auto-generated from WSDL.
type OperationFindContactsResponse struct {
	FindContactsResponse *FindContactsResponse `xml:"findContactsResponse,omitempty" json:"findContactsResponse,omitempty" yaml:"findContactsResponse,omitempty"`
}

// Operation wrapper for FreeListHandles.
// OperationFreeListHandlesRequest was auto-generated from WSDL.
type OperationFreeListHandlesRequest struct {
	FreeListHandles *FreeListHandles `xml:"freeListHandles,omitempty" json:"freeListHandles,omitempty" yaml:"freeListHandles,omitempty"`
}

// Operation wrapper for FreeListHandles.
// OperationFreeListHandlesResponse was auto-generated from WSDL.
type OperationFreeListHandlesResponse struct {
	FreeListHandlesResponse *FreeListHandlesResponse `xml:"freeListHandlesResponse,omitempty" json:"freeListHandlesResponse,omitempty" yaml:"freeListHandlesResponse,omitempty"`
}

// Operation wrapper for GetAccessTypeForContact.
// OperationGetAccessTypeForContactRequest was auto-generated from
// WSDL.
type OperationGetAccessTypeForContactRequest struct {
	GetAccessTypeForContact *GetAccessTypeForContact `xml:"getAccessTypeForContact,omitempty" json:"getAccessTypeForContact,omitempty" yaml:"getAccessTypeForContact,omitempty"`
}

// Operation wrapper for GetAccessTypeForContact.
// OperationGetAccessTypeForContactResponse was auto-generated
// from WSDL.
type OperationGetAccessTypeForContactResponse struct {
	GetAccessTypeForContactResponse *GetAccessTypeForContactResponse `xml:"getAccessTypeForContactResponse,omitempty" json:"getAccessTypeForContactResponse,omitempty" yaml:"getAccessTypeForContactResponse,omitempty"`
}

// Operation wrapper for GetArtifact.
// OperationGetArtifactRequest was auto-generated from WSDL.
type OperationGetArtifactRequest struct {
	GetArtifact *GetArtifact `xml:"getArtifact,omitempty" json:"getArtifact,omitempty" yaml:"getArtifact,omitempty"`
}

// Operation wrapper for GetArtifact.
// OperationGetArtifactResponse was auto-generated from WSDL.
type OperationGetArtifactResponse struct {
	GetArtifactResponse *GetArtifactResponse `xml:"getArtifactResponse,omitempty" json:"getArtifactResponse,omitempty" yaml:"getArtifactResponse,omitempty"`
}

// Operation wrapper for GetAssetExtensionInformation.
// OperationGetAssetExtensionInformationRequest was auto-generated
// from WSDL.
type OperationGetAssetExtensionInformationRequest struct {
	GetAssetExtensionInformation *GetAssetExtensionInformation `xml:"getAssetExtensionInformation,omitempty" json:"getAssetExtensionInformation,omitempty" yaml:"getAssetExtensionInformation,omitempty"`
}

// Operation wrapper for GetAssetExtensionInformation.
// OperationGetAssetExtensionInformationResponse was auto-generated
// from WSDL.
type OperationGetAssetExtensionInformationResponse struct {
	GetAssetExtensionInformationResponse *GetAssetExtensionInformationResponse `xml:"getAssetExtensionInformationResponse,omitempty" json:"getAssetExtensionInformationResponse,omitempty" yaml:"getAssetExtensionInformationResponse,omitempty"`
}

// Operation wrapper for GetAttmntInfo.
// OperationGetAttmntInfoRequest was auto-generated from WSDL.
type OperationGetAttmntInfoRequest struct {
	GetAttmntInfo *GetAttmntInfo `xml:"getAttmntInfo,omitempty" json:"getAttmntInfo,omitempty" yaml:"getAttmntInfo,omitempty"`
}

// Operation wrapper for GetAttmntInfo.
// OperationGetAttmntInfoResponse was auto-generated from WSDL.
type OperationGetAttmntInfoResponse struct {
	GetAttmntInfoResponse *GetAttmntInfoResponse `xml:"getAttmntInfoResponse,omitempty" json:"getAttmntInfoResponse,omitempty" yaml:"getAttmntInfoResponse,omitempty"`
}

// Operation wrapper for GetAttmntList.
// OperationGetAttmntListRequest was auto-generated from WSDL.
type OperationGetAttmntListRequest struct {
	GetAttmntList *GetAttmntList `xml:"getAttmntList,omitempty" json:"getAttmntList,omitempty" yaml:"getAttmntList,omitempty"`
}

// Operation wrapper for GetAttmntList.
// OperationGetAttmntListResponse was auto-generated from WSDL.
type OperationGetAttmntListResponse struct {
	GetAttmntListResponse *GetAttmntListResponse `xml:"getAttmntListResponse,omitempty" json:"getAttmntListResponse,omitempty" yaml:"getAttmntListResponse,omitempty"`
}

// Operation wrapper for GetAttmntListPerKD.
// OperationGetAttmntListPerKDRequest was auto-generated from WSDL.
type OperationGetAttmntListPerKDRequest struct {
	GetAttmntListPerKD *GetAttmntListPerKD `xml:"getAttmntListPerKD,omitempty" json:"getAttmntListPerKD,omitempty" yaml:"getAttmntListPerKD,omitempty"`
}

// Operation wrapper for GetAttmntListPerKD.
// OperationGetAttmntListPerKDResponse was auto-generated from
// WSDL.
type OperationGetAttmntListPerKDResponse struct {
	GetAttmntListPerKDResponse *GetAttmntListPerKDResponse `xml:"getAttmntListPerKDResponse,omitempty" json:"getAttmntListPerKDResponse,omitempty" yaml:"getAttmntListPerKDResponse,omitempty"`
}

// Operation wrapper for GetBookmarks.
// OperationGetBookmarksRequest was auto-generated from WSDL.
type OperationGetBookmarksRequest struct {
	GetBookmarks *GetBookmarks `xml:"getBookmarks,omitempty" json:"getBookmarks,omitempty" yaml:"getBookmarks,omitempty"`
}

// Operation wrapper for GetBookmarks.
// OperationGetBookmarksResponse was auto-generated from WSDL.
type OperationGetBookmarksResponse struct {
	GetBookmarksResponse *GetBookmarksResponse `xml:"getBookmarksResponse,omitempty" json:"getBookmarksResponse,omitempty" yaml:"getBookmarksResponse,omitempty"`
}

// Operation wrapper for GetBopsid.
// OperationGetBopsidRequest was auto-generated from WSDL.
type OperationGetBopsidRequest struct {
	GetBopsid *GetBopsid `xml:"getBopsid,omitempty" json:"getBopsid,omitempty" yaml:"getBopsid,omitempty"`
}

// Operation wrapper for GetBopsid.
// OperationGetBopsidResponse was auto-generated from WSDL.
type OperationGetBopsidResponse struct {
	GetBopsidResponse *GetBopsidResponse `xml:"getBopsidResponse,omitempty" json:"getBopsidResponse,omitempty" yaml:"getBopsidResponse,omitempty"`
}

// Operation wrapper for GetCategory.
// OperationGetCategoryRequest was auto-generated from WSDL.
type OperationGetCategoryRequest struct {
	GetCategory *GetCategory `xml:"getCategory,omitempty" json:"getCategory,omitempty" yaml:"getCategory,omitempty"`
}

// Operation wrapper for GetCategory.
// OperationGetCategoryResponse was auto-generated from WSDL.
type OperationGetCategoryResponse struct {
	GetCategoryResponse *GetCategoryResponse `xml:"getCategoryResponse,omitempty" json:"getCategoryResponse,omitempty" yaml:"getCategoryResponse,omitempty"`
}

// Operation wrapper for GetComments.
// OperationGetCommentsRequest was auto-generated from WSDL.
type OperationGetCommentsRequest struct {
	GetComments *GetComments `xml:"getComments,omitempty" json:"getComments,omitempty" yaml:"getComments,omitempty"`
}

// Operation wrapper for GetComments.
// OperationGetCommentsResponse was auto-generated from WSDL.
type OperationGetCommentsResponse struct {
	GetCommentsResponse *GetCommentsResponse `xml:"getCommentsResponse,omitempty" json:"getCommentsResponse,omitempty" yaml:"getCommentsResponse,omitempty"`
}

// Operation wrapper for GetConfigurationMode.
// OperationGetConfigurationModeRequest was auto-generated from
// WSDL.
type OperationGetConfigurationModeRequest struct {
	GetConfigurationMode *GetConfigurationMode `xml:"getConfigurationMode,omitempty" json:"getConfigurationMode,omitempty" yaml:"getConfigurationMode,omitempty"`
}

// Operation wrapper for GetConfigurationMode.
// OperationGetConfigurationModeResponse was auto-generated from
// WSDL.
type OperationGetConfigurationModeResponse struct {
	GetConfigurationModeResponse *GetConfigurationModeResponse `xml:"getConfigurationModeResponse,omitempty" json:"getConfigurationModeResponse,omitempty" yaml:"getConfigurationModeResponse,omitempty"`
}

// Operation wrapper for GetContact.
// OperationGetContactRequest was auto-generated from WSDL.
type OperationGetContactRequest struct {
	GetContact *GetContact `xml:"getContact,omitempty" json:"getContact,omitempty" yaml:"getContact,omitempty"`
}

// Operation wrapper for GetContact.
// OperationGetContactResponse was auto-generated from WSDL.
type OperationGetContactResponse struct {
	GetContactResponse *GetContactResponse `xml:"getContactResponse,omitempty" json:"getContactResponse,omitempty" yaml:"getContactResponse,omitempty"`
}

// Operation wrapper for GetDecisionTrees.
// OperationGetDecisionTreesRequest was auto-generated from WSDL.
type OperationGetDecisionTreesRequest struct {
	GetDecisionTrees *GetDecisionTrees `xml:"getDecisionTrees,omitempty" json:"getDecisionTrees,omitempty" yaml:"getDecisionTrees,omitempty"`
}

// Operation wrapper for GetDecisionTrees.
// OperationGetDecisionTreesResponse was auto-generated from WSDL.
type OperationGetDecisionTreesResponse struct {
	GetDecisionTreesResponse *GetDecisionTreesResponse `xml:"getDecisionTreesResponse,omitempty" json:"getDecisionTreesResponse,omitempty" yaml:"getDecisionTreesResponse,omitempty"`
}

// Operation wrapper for GetDependentAttrControls.
// OperationGetDependentAttrControlsRequest was auto-generated
// from WSDL.
type OperationGetDependentAttrControlsRequest struct {
	GetDependentAttrControls *GetDependentAttrControls `xml:"getDependentAttrControls,omitempty" json:"getDependentAttrControls,omitempty" yaml:"getDependentAttrControls,omitempty"`
}

// Operation wrapper for GetDependentAttrControls.
// OperationGetDependentAttrControlsResponse was auto-generated
// from WSDL.
type OperationGetDependentAttrControlsResponse struct {
	GetDependentAttrControlsResponse *GetDependentAttrControlsResponse `xml:"getDependentAttrControlsResponse,omitempty" json:"getDependentAttrControlsResponse,omitempty" yaml:"getDependentAttrControlsResponse,omitempty"`
}

// Operation wrapper for GetDocument.
// OperationGetDocumentRequest was auto-generated from WSDL.
type OperationGetDocumentRequest struct {
	GetDocument *GetDocument `xml:"getDocument,omitempty" json:"getDocument,omitempty" yaml:"getDocument,omitempty"`
}

// Operation wrapper for GetDocument.
// OperationGetDocumentResponse was auto-generated from WSDL.
type OperationGetDocumentResponse struct {
	GetDocumentResponse *GetDocumentResponse `xml:"getDocumentResponse,omitempty" json:"getDocumentResponse,omitempty" yaml:"getDocumentResponse,omitempty"`
}

// Operation wrapper for GetDocumentTypes.
// OperationGetDocumentTypesRequest was auto-generated from WSDL.
type OperationGetDocumentTypesRequest struct {
	GetDocumentTypes *GetDocumentTypes `xml:"getDocumentTypes,omitempty" json:"getDocumentTypes,omitempty" yaml:"getDocumentTypes,omitempty"`
}

// Operation wrapper for GetDocumentTypes.
// OperationGetDocumentTypesResponse was auto-generated from WSDL.
type OperationGetDocumentTypesResponse struct {
	GetDocumentTypesResponse *GetDocumentTypesResponse `xml:"getDocumentTypesResponse,omitempty" json:"getDocumentTypesResponse,omitempty" yaml:"getDocumentTypesResponse,omitempty"`
}

// Operation wrapper for GetDocumentsByIDs.
// OperationGetDocumentsByIDsRequest was auto-generated from WSDL.
type OperationGetDocumentsByIDsRequest struct {
	GetDocumentsByIDs *GetDocumentsByIDs `xml:"getDocumentsByIDs,omitempty" json:"getDocumentsByIDs,omitempty" yaml:"getDocumentsByIDs,omitempty"`
}

// Operation wrapper for GetDocumentsByIDs.
// OperationGetDocumentsByIDsResponse was auto-generated from WSDL.
type OperationGetDocumentsByIDsResponse struct {
	GetDocumentsByIDsResponse *GetDocumentsByIDsResponse `xml:"getDocumentsByIDsResponse,omitempty" json:"getDocumentsByIDsResponse,omitempty" yaml:"getDocumentsByIDsResponse,omitempty"`
}

// Operation wrapper for GetFolderInfo.
// OperationGetFolderInfoRequest was auto-generated from WSDL.
type OperationGetFolderInfoRequest struct {
	GetFolderInfo *GetFolderInfo `xml:"getFolderInfo,omitempty" json:"getFolderInfo,omitempty" yaml:"getFolderInfo,omitempty"`
}

// Operation wrapper for GetFolderInfo.
// OperationGetFolderInfoResponse was auto-generated from WSDL.
type OperationGetFolderInfoResponse struct {
	GetFolderInfoResponse *GetFolderInfoResponse `xml:"getFolderInfoResponse,omitempty" json:"getFolderInfoResponse,omitempty" yaml:"getFolderInfoResponse,omitempty"`
}

// Operation wrapper for GetFolderList.
// OperationGetFolderListRequest was auto-generated from WSDL.
type OperationGetFolderListRequest struct {
	GetFolderList *GetFolderList `xml:"getFolderList,omitempty" json:"getFolderList,omitempty" yaml:"getFolderList,omitempty"`
}

// Operation wrapper for GetFolderList.
// OperationGetFolderListResponse was auto-generated from WSDL.
type OperationGetFolderListResponse struct {
	GetFolderListResponse *GetFolderListResponse `xml:"getFolderListResponse,omitempty" json:"getFolderListResponse,omitempty" yaml:"getFolderListResponse,omitempty"`
}

// Operation wrapper for GetGroupMemberListValues.
// OperationGetGroupMemberListValuesRequest was auto-generated
// from WSDL.
type OperationGetGroupMemberListValuesRequest struct {
	GetGroupMemberListValues *GetGroupMemberListValues `xml:"getGroupMemberListValues,omitempty" json:"getGroupMemberListValues,omitempty" yaml:"getGroupMemberListValues,omitempty"`
}

// Operation wrapper for GetGroupMemberListValues.
// OperationGetGroupMemberListValuesResponse was auto-generated
// from WSDL.
type OperationGetGroupMemberListValuesResponse struct {
	GetGroupMemberListValuesResponse *GetGroupMemberListValuesResponse `xml:"getGroupMemberListValuesResponse,omitempty" json:"getGroupMemberListValuesResponse,omitempty" yaml:"getGroupMemberListValuesResponse,omitempty"`
}

// Operation wrapper for GetHandleForUserid.
// OperationGetHandleForUseridRequest was auto-generated from WSDL.
type OperationGetHandleForUseridRequest struct {
	GetHandleForUserid *GetHandleForUserid `xml:"getHandleForUserid,omitempty" json:"getHandleForUserid,omitempty" yaml:"getHandleForUserid,omitempty"`
}

// Operation wrapper for GetHandleForUserid.
// OperationGetHandleForUseridResponse was auto-generated from
// WSDL.
type OperationGetHandleForUseridResponse struct {
	GetHandleForUseridResponse *GetHandleForUseridResponse `xml:"getHandleForUseridResponse,omitempty" json:"getHandleForUseridResponse,omitempty" yaml:"getHandleForUseridResponse,omitempty"`
}

// Operation wrapper for GetKDListPerAttmnt.
// OperationGetKDListPerAttmntRequest was auto-generated from WSDL.
type OperationGetKDListPerAttmntRequest struct {
	GetKDListPerAttmnt *GetKDListPerAttmnt `xml:"getKDListPerAttmnt,omitempty" json:"getKDListPerAttmnt,omitempty" yaml:"getKDListPerAttmnt,omitempty"`
}

// Operation wrapper for GetKDListPerAttmnt.
// OperationGetKDListPerAttmntResponse was auto-generated from
// WSDL.
type OperationGetKDListPerAttmntResponse struct {
	GetKDListPerAttmntResponse *GetKDListPerAttmntResponse `xml:"getKDListPerAttmntResponse,omitempty" json:"getKDListPerAttmntResponse,omitempty" yaml:"getKDListPerAttmntResponse,omitempty"`
}

// Operation wrapper for GetListValues.
// OperationGetListValuesRequest was auto-generated from WSDL.
type OperationGetListValuesRequest struct {
	GetListValues *GetListValues `xml:"getListValues,omitempty" json:"getListValues,omitempty" yaml:"getListValues,omitempty"`
}

// Operation wrapper for GetListValues.
// OperationGetListValuesResponse was auto-generated from WSDL.
type OperationGetListValuesResponse struct {
	GetListValuesResponse *GetListValuesResponse `xml:"getListValuesResponse,omitempty" json:"getListValuesResponse,omitempty" yaml:"getListValuesResponse,omitempty"`
}

// Operation wrapper for GetLrelLength.
// OperationGetLrelLengthRequest was auto-generated from WSDL.
type OperationGetLrelLengthRequest struct {
	GetLrelLength *GetLrelLength `xml:"getLrelLength,omitempty" json:"getLrelLength,omitempty" yaml:"getLrelLength,omitempty"`
}

// Operation wrapper for GetLrelLength.
// OperationGetLrelLengthResponse was auto-generated from WSDL.
type OperationGetLrelLengthResponse struct {
	GetLrelLengthResponse *GetLrelLengthResponse `xml:"getLrelLengthResponse,omitempty" json:"getLrelLengthResponse,omitempty" yaml:"getLrelLengthResponse,omitempty"`
}

// Operation wrapper for GetLrelValues.
// OperationGetLrelValuesRequest was auto-generated from WSDL.
type OperationGetLrelValuesRequest struct {
	GetLrelValues *GetLrelValues `xml:"getLrelValues,omitempty" json:"getLrelValues,omitempty" yaml:"getLrelValues,omitempty"`
}

// Operation wrapper for GetLrelValues.
// OperationGetLrelValuesResponse was auto-generated from WSDL.
type OperationGetLrelValuesResponse struct {
	GetLrelValuesResponse *GetLrelValuesResponse `xml:"getLrelValuesResponse,omitempty" json:"getLrelValuesResponse,omitempty" yaml:"getLrelValuesResponse,omitempty"`
}

// Operation wrapper for GetNotificationsForContact.
// OperationGetNotificationsForContactRequest was auto-generated
// from WSDL.
type OperationGetNotificationsForContactRequest struct {
	GetNotificationsForContact *GetNotificationsForContact `xml:"getNotificationsForContact,omitempty" json:"getNotificationsForContact,omitempty" yaml:"getNotificationsForContact,omitempty"`
}

// Operation wrapper for GetNotificationsForContact.
// OperationGetNotificationsForContactResponse was auto-generated
// from WSDL.
type OperationGetNotificationsForContactResponse struct {
	GetNotificationsForContactResponse *GetNotificationsForContactResponse `xml:"getNotificationsForContactResponse,omitempty" json:"getNotificationsForContactResponse,omitempty" yaml:"getNotificationsForContactResponse,omitempty"`
}

// Operation wrapper for GetObjectTypeInformation.
// OperationGetObjectTypeInformationRequest was auto-generated
// from WSDL.
type OperationGetObjectTypeInformationRequest struct {
	GetObjectTypeInformation *GetObjectTypeInformation `xml:"getObjectTypeInformation,omitempty" json:"getObjectTypeInformation,omitempty" yaml:"getObjectTypeInformation,omitempty"`
}

// Operation wrapper for GetObjectTypeInformation.
// OperationGetObjectTypeInformationResponse was auto-generated
// from WSDL.
type OperationGetObjectTypeInformationResponse struct {
	GetObjectTypeInformationResponse *GetObjectTypeInformationResponse `xml:"getObjectTypeInformationResponse,omitempty" json:"getObjectTypeInformationResponse,omitempty" yaml:"getObjectTypeInformationResponse,omitempty"`
}

// Operation wrapper for GetObjectValues.
// OperationGetObjectValuesRequest was auto-generated from WSDL.
type OperationGetObjectValuesRequest struct {
	GetObjectValues *GetObjectValues `xml:"getObjectValues,omitempty" json:"getObjectValues,omitempty" yaml:"getObjectValues,omitempty"`
}

// Operation wrapper for GetObjectValues.
// OperationGetObjectValuesResponse was auto-generated from WSDL.
type OperationGetObjectValuesResponse struct {
	GetObjectValuesResponse *GetObjectValuesResponse `xml:"getObjectValuesResponse,omitempty" json:"getObjectValuesResponse,omitempty" yaml:"getObjectValuesResponse,omitempty"`
}

// Operation wrapper for GetPendingChangeTaskListForContact.
// OperationGetPendingChangeTaskListForContactRequest was auto-generated
// from WSDL.
type OperationGetPendingChangeTaskListForContactRequest struct {
	GetPendingChangeTaskListForContact *GetPendingChangeTaskListForContact `xml:"getPendingChangeTaskListForContact,omitempty" json:"getPendingChangeTaskListForContact,omitempty" yaml:"getPendingChangeTaskListForContact,omitempty"`
}

// Operation wrapper for GetPendingChangeTaskListForContact.
// OperationGetPendingChangeTaskListForContactResponse was auto-generated
// from WSDL.
type OperationGetPendingChangeTaskListForContactResponse struct {
	GetPendingChangeTaskListForContactResponse *GetPendingChangeTaskListForContactResponse `xml:"getPendingChangeTaskListForContactResponse,omitempty" json:"getPendingChangeTaskListForContactResponse,omitempty" yaml:"getPendingChangeTaskListForContactResponse,omitempty"`
}

// Operation wrapper for GetPendingIssueTaskListForContact.
// OperationGetPendingIssueTaskListForContactRequest was auto-generated
// from WSDL.
type OperationGetPendingIssueTaskListForContactRequest struct {
	GetPendingIssueTaskListForContact *GetPendingIssueTaskListForContact `xml:"getPendingIssueTaskListForContact,omitempty" json:"getPendingIssueTaskListForContact,omitempty" yaml:"getPendingIssueTaskListForContact,omitempty"`
}

// Operation wrapper for GetPendingIssueTaskListForContact.
// OperationGetPendingIssueTaskListForContactResponse was auto-generated
// from WSDL.
type OperationGetPendingIssueTaskListForContactResponse struct {
	GetPendingIssueTaskListForContactResponse *GetPendingIssueTaskListForContactResponse `xml:"getPendingIssueTaskListForContactResponse,omitempty" json:"getPendingIssueTaskListForContactResponse,omitempty" yaml:"getPendingIssueTaskListForContactResponse,omitempty"`
}

// Operation wrapper for GetPermissionGroups.
// OperationGetPermissionGroupsRequest was auto-generated from
// WSDL.
type OperationGetPermissionGroupsRequest struct {
	GetPermissionGroups *GetPermissionGroups `xml:"getPermissionGroups,omitempty" json:"getPermissionGroups,omitempty" yaml:"getPermissionGroups,omitempty"`
}

// Operation wrapper for GetPermissionGroups.
// OperationGetPermissionGroupsResponse was auto-generated from
// WSDL.
type OperationGetPermissionGroupsResponse struct {
	GetPermissionGroupsResponse *GetPermissionGroupsResponse `xml:"getPermissionGroupsResponse,omitempty" json:"getPermissionGroupsResponse,omitempty" yaml:"getPermissionGroupsResponse,omitempty"`
}

// Operation wrapper for GetPolicyInfo.
// OperationGetPolicyInfoRequest was auto-generated from WSDL.
type OperationGetPolicyInfoRequest struct {
	GetPolicyInfo *GetPolicyInfo `xml:"getPolicyInfo,omitempty" json:"getPolicyInfo,omitempty" yaml:"getPolicyInfo,omitempty"`
}

// Operation wrapper for GetPolicyInfo.
// OperationGetPolicyInfoResponse was auto-generated from WSDL.
type OperationGetPolicyInfoResponse struct {
	GetPolicyInfoResponse *GetPolicyInfoResponse `xml:"getPolicyInfoResponse,omitempty" json:"getPolicyInfoResponse,omitempty" yaml:"getPolicyInfoResponse,omitempty"`
}

// Operation wrapper for GetPriorities.
// OperationGetPrioritiesRequest was auto-generated from WSDL.
type OperationGetPrioritiesRequest struct {
	GetPriorities *GetPriorities `xml:"getPriorities,omitempty" json:"getPriorities,omitempty" yaml:"getPriorities,omitempty"`
}

// Operation wrapper for GetPriorities.
// OperationGetPrioritiesResponse was auto-generated from WSDL.
type OperationGetPrioritiesResponse struct {
	GetPrioritiesResponse *GetPrioritiesResponse `xml:"getPrioritiesResponse,omitempty" json:"getPrioritiesResponse,omitempty" yaml:"getPrioritiesResponse,omitempty"`
}

// Operation wrapper for GetPropertyInfoForCategory.
// OperationGetPropertyInfoForCategoryRequest was auto-generated
// from WSDL.
type OperationGetPropertyInfoForCategoryRequest struct {
	GetPropertyInfoForCategory *GetPropertyInfoForCategory `xml:"getPropertyInfoForCategory,omitempty" json:"getPropertyInfoForCategory,omitempty" yaml:"getPropertyInfoForCategory,omitempty"`
}

// Operation wrapper for GetPropertyInfoForCategory.
// OperationGetPropertyInfoForCategoryResponse was auto-generated
// from WSDL.
type OperationGetPropertyInfoForCategoryResponse struct {
	GetPropertyInfoForCategoryResponse *GetPropertyInfoForCategoryResponse `xml:"getPropertyInfoForCategoryResponse,omitempty" json:"getPropertyInfoForCategoryResponse,omitempty" yaml:"getPropertyInfoForCategoryResponse,omitempty"`
}

// Operation wrapper for GetQuestionsAsked.
// OperationGetQuestionsAskedRequest was auto-generated from WSDL.
type OperationGetQuestionsAskedRequest struct {
	GetQuestionsAsked *GetQuestionsAsked `xml:"getQuestionsAsked,omitempty" json:"getQuestionsAsked,omitempty" yaml:"getQuestionsAsked,omitempty"`
}

// Operation wrapper for GetQuestionsAsked.
// OperationGetQuestionsAskedResponse was auto-generated from WSDL.
type OperationGetQuestionsAskedResponse struct {
	GetQuestionsAskedResponse *GetQuestionsAskedResponse `xml:"getQuestionsAskedResponse,omitempty" json:"getQuestionsAskedResponse,omitempty" yaml:"getQuestionsAskedResponse,omitempty"`
}

// Operation wrapper for GetRelatedList.
// OperationGetRelatedListRequest was auto-generated from WSDL.
type OperationGetRelatedListRequest struct {
	GetRelatedList *GetRelatedList `xml:"getRelatedList,omitempty" json:"getRelatedList,omitempty" yaml:"getRelatedList,omitempty"`
}

// Operation wrapper for GetRelatedList.
// OperationGetRelatedListResponse was auto-generated from WSDL.
type OperationGetRelatedListResponse struct {
	GetRelatedListResponse *GetRelatedListResponse `xml:"getRelatedListResponse,omitempty" json:"getRelatedListResponse,omitempty" yaml:"getRelatedListResponse,omitempty"`
}

// Operation wrapper for GetRelatedListValues.
// OperationGetRelatedListValuesRequest was auto-generated from
// WSDL.
type OperationGetRelatedListValuesRequest struct {
	GetRelatedListValues *GetRelatedListValues `xml:"getRelatedListValues,omitempty" json:"getRelatedListValues,omitempty" yaml:"getRelatedListValues,omitempty"`
}

// Operation wrapper for GetRelatedListValues.
// OperationGetRelatedListValuesResponse was auto-generated from
// WSDL.
type OperationGetRelatedListValuesResponse struct {
	GetRelatedListValuesResponse *GetRelatedListValuesResponse `xml:"getRelatedListValuesResponse,omitempty" json:"getRelatedListValuesResponse,omitempty" yaml:"getRelatedListValuesResponse,omitempty"`
}

// Operation wrapper for GetRepositoryInfo.
// OperationGetRepositoryInfoRequest was auto-generated from WSDL.
type OperationGetRepositoryInfoRequest struct {
	GetRepositoryInfo *GetRepositoryInfo `xml:"getRepositoryInfo,omitempty" json:"getRepositoryInfo,omitempty" yaml:"getRepositoryInfo,omitempty"`
}

// Operation wrapper for GetRepositoryInfo.
// OperationGetRepositoryInfoResponse was auto-generated from WSDL.
type OperationGetRepositoryInfoResponse struct {
	GetRepositoryInfoResponse *GetRepositoryInfoResponse `xml:"getRepositoryInfoResponse,omitempty" json:"getRepositoryInfoResponse,omitempty" yaml:"getRepositoryInfoResponse,omitempty"`
}

// Operation wrapper for GetStatuses.
// OperationGetStatusesRequest was auto-generated from WSDL.
type OperationGetStatusesRequest struct {
	GetStatuses *GetStatuses `xml:"getStatuses,omitempty" json:"getStatuses,omitempty" yaml:"getStatuses,omitempty"`
}

// Operation wrapper for GetStatuses.
// OperationGetStatusesResponse was auto-generated from WSDL.
type OperationGetStatusesResponse struct {
	GetStatusesResponse *GetStatusesResponse `xml:"getStatusesResponse,omitempty" json:"getStatusesResponse,omitempty" yaml:"getStatusesResponse,omitempty"`
}

// Operation wrapper for GetTaskListValues.
// OperationGetTaskListValuesRequest was auto-generated from WSDL.
type OperationGetTaskListValuesRequest struct {
	GetTaskListValues *GetTaskListValues `xml:"getTaskListValues,omitempty" json:"getTaskListValues,omitempty" yaml:"getTaskListValues,omitempty"`
}

// Operation wrapper for GetTaskListValues.
// OperationGetTaskListValuesResponse was auto-generated from WSDL.
type OperationGetTaskListValuesResponse struct {
	GetTaskListValuesResponse *GetTaskListValuesResponse `xml:"getTaskListValuesResponse,omitempty" json:"getTaskListValuesResponse,omitempty" yaml:"getTaskListValuesResponse,omitempty"`
}

// Operation wrapper for GetTemplateList.
// OperationGetTemplateListRequest was auto-generated from WSDL.
type OperationGetTemplateListRequest struct {
	GetTemplateList *GetTemplateList `xml:"getTemplateList,omitempty" json:"getTemplateList,omitempty" yaml:"getTemplateList,omitempty"`
}

// Operation wrapper for GetTemplateList.
// OperationGetTemplateListResponse was auto-generated from WSDL.
type OperationGetTemplateListResponse struct {
	GetTemplateListResponse *GetTemplateListResponse `xml:"getTemplateListResponse,omitempty" json:"getTemplateListResponse,omitempty" yaml:"getTemplateListResponse,omitempty"`
}

// Operation wrapper for GetValidTaskTransitions.
// OperationGetValidTaskTransitionsRequest was auto-generated from
// WSDL.
type OperationGetValidTaskTransitionsRequest struct {
	GetValidTaskTransitions *GetValidTaskTransitions `xml:"getValidTaskTransitions,omitempty" json:"getValidTaskTransitions,omitempty" yaml:"getValidTaskTransitions,omitempty"`
}

// Operation wrapper for GetValidTaskTransitions.
// OperationGetValidTaskTransitionsResponse was auto-generated
// from WSDL.
type OperationGetValidTaskTransitionsResponse struct {
	GetValidTaskTransitionsResponse *GetValidTaskTransitionsResponse `xml:"getValidTaskTransitionsResponse,omitempty" json:"getValidTaskTransitionsResponse,omitempty" yaml:"getValidTaskTransitionsResponse,omitempty"`
}

// Operation wrapper for GetValidTransitions.
// OperationGetValidTransitionsRequest was auto-generated from
// WSDL.
type OperationGetValidTransitionsRequest struct {
	GetValidTransitions *GetValidTransitions `xml:"getValidTransitions,omitempty" json:"getValidTransitions,omitempty" yaml:"getValidTransitions,omitempty"`
}

// Operation wrapper for GetValidTransitions.
// OperationGetValidTransitionsResponse was auto-generated from
// WSDL.
type OperationGetValidTransitionsResponse struct {
	GetValidTransitionsResponse *GetValidTransitionsResponse `xml:"getValidTransitionsResponse,omitempty" json:"getValidTransitionsResponse,omitempty" yaml:"getValidTransitionsResponse,omitempty"`
}

// Operation wrapper for GetWorkFlowTemplates.
// OperationGetWorkFlowTemplatesRequest was auto-generated from
// WSDL.
type OperationGetWorkFlowTemplatesRequest struct {
	GetWorkFlowTemplates *GetWorkFlowTemplates `xml:"getWorkFlowTemplates,omitempty" json:"getWorkFlowTemplates,omitempty" yaml:"getWorkFlowTemplates,omitempty"`
}

// Operation wrapper for GetWorkFlowTemplates.
// OperationGetWorkFlowTemplatesResponse was auto-generated from
// WSDL.
type OperationGetWorkFlowTemplatesResponse struct {
	GetWorkFlowTemplatesResponse *GetWorkFlowTemplatesResponse `xml:"getWorkFlowTemplatesResponse,omitempty" json:"getWorkFlowTemplatesResponse,omitempty" yaml:"getWorkFlowTemplatesResponse,omitempty"`
}

// Operation wrapper for GetWorkflowTemplateList.
// OperationGetWorkflowTemplateListRequest was auto-generated from
// WSDL.
type OperationGetWorkflowTemplateListRequest struct {
	GetWorkflowTemplateList *GetWorkflowTemplateList `xml:"getWorkflowTemplateList,omitempty" json:"getWorkflowTemplateList,omitempty" yaml:"getWorkflowTemplateList,omitempty"`
}

// Operation wrapper for GetWorkflowTemplateList.
// OperationGetWorkflowTemplateListResponse was auto-generated
// from WSDL.
type OperationGetWorkflowTemplateListResponse struct {
	GetWorkflowTemplateListResponse *GetWorkflowTemplateListResponse `xml:"getWorkflowTemplateListResponse,omitempty" json:"getWorkflowTemplateListResponse,omitempty" yaml:"getWorkflowTemplateListResponse,omitempty"`
}

// Operation wrapper for Impersonate.
// OperationImpersonateRequest was auto-generated from WSDL.
type OperationImpersonateRequest struct {
	Impersonate *Impersonate `xml:"impersonate,omitempty" json:"impersonate,omitempty" yaml:"impersonate,omitempty"`
}

// Operation wrapper for Impersonate.
// OperationImpersonateResponse was auto-generated from WSDL.
type OperationImpersonateResponse struct {
	ImpersonateResponse *ImpersonateResponse `xml:"impersonateResponse,omitempty" json:"impersonateResponse,omitempty" yaml:"impersonateResponse,omitempty"`
}

// Operation wrapper for IsAttmntLinkedKD.
// OperationIsAttmntLinkedKDRequest was auto-generated from WSDL.
type OperationIsAttmntLinkedKDRequest struct {
	IsAttmntLinkedKD *IsAttmntLinkedKD `xml:"isAttmntLinkedKD,omitempty" json:"isAttmntLinkedKD,omitempty" yaml:"isAttmntLinkedKD,omitempty"`
}

// Operation wrapper for IsAttmntLinkedKD.
// OperationIsAttmntLinkedKDResponse was auto-generated from WSDL.
type OperationIsAttmntLinkedKDResponse struct {
	IsAttmntLinkedKDResponse *IsAttmntLinkedKDResponse `xml:"isAttmntLinkedKDResponse,omitempty" json:"isAttmntLinkedKDResponse,omitempty" yaml:"isAttmntLinkedKDResponse,omitempty"`
}

// Operation wrapper for LogComment.
// OperationLogCommentRequest was auto-generated from WSDL.
type OperationLogCommentRequest struct {
	LogComment *LogComment `xml:"logComment,omitempty" json:"logComment,omitempty" yaml:"logComment,omitempty"`
}

// Operation wrapper for LogComment.
// OperationLogCommentResponse was auto-generated from WSDL.
type OperationLogCommentResponse struct {
	LogCommentResponse *LogCommentResponse `xml:"logCommentResponse,omitempty" json:"logCommentResponse,omitempty" yaml:"logCommentResponse,omitempty"`
}

// Operation wrapper for Login.
// OperationLoginRequest was auto-generated from WSDL.
type OperationLoginRequest struct {
	Login *Login `xml:"login,omitempty" json:"login,omitempty" yaml:"login,omitempty"`
}

// Operation wrapper for Login.
// OperationLoginResponse was auto-generated from WSDL.
type OperationLoginResponse struct {
	LoginResponse *LoginResponse `xml:"loginResponse,omitempty" json:"loginResponse,omitempty" yaml:"loginResponse,omitempty"`
}

// Operation wrapper for LoginService.
// OperationLoginServiceRequest was auto-generated from WSDL.
type OperationLoginServiceRequest struct {
	LoginService *LoginService `xml:"loginService,omitempty" json:"loginService,omitempty" yaml:"loginService,omitempty"`
}

// Operation wrapper for LoginService.
// OperationLoginServiceResponse was auto-generated from WSDL.
type OperationLoginServiceResponse struct {
	LoginServiceResponse *LoginServiceResponse `xml:"loginServiceResponse,omitempty" json:"loginServiceResponse,omitempty" yaml:"loginServiceResponse,omitempty"`
}

// Operation wrapper for LoginServiceManaged.
// OperationLoginServiceManagedRequest was auto-generated from
// WSDL.
type OperationLoginServiceManagedRequest struct {
	LoginServiceManaged *LoginServiceManaged `xml:"loginServiceManaged,omitempty" json:"loginServiceManaged,omitempty" yaml:"loginServiceManaged,omitempty"`
}

// Operation wrapper for LoginServiceManaged.
// OperationLoginServiceManagedResponse was auto-generated from
// WSDL.
type OperationLoginServiceManagedResponse struct {
	LoginServiceManagedResponse *LoginServiceManagedResponse `xml:"loginServiceManagedResponse,omitempty" json:"loginServiceManagedResponse,omitempty" yaml:"loginServiceManagedResponse,omitempty"`
}

// Operation wrapper for LoginWithArtifact.
// OperationLoginWithArtifactRequest was auto-generated from WSDL.
type OperationLoginWithArtifactRequest struct {
	LoginWithArtifact *LoginWithArtifact `xml:"loginWithArtifact,omitempty" json:"loginWithArtifact,omitempty" yaml:"loginWithArtifact,omitempty"`
}

// Operation wrapper for LoginWithArtifact.
// OperationLoginWithArtifactResponse was auto-generated from WSDL.
type OperationLoginWithArtifactResponse struct {
	LoginWithArtifactResponse *LoginWithArtifactResponse `xml:"loginWithArtifactResponse,omitempty" json:"loginWithArtifactResponse,omitempty" yaml:"loginWithArtifactResponse,omitempty"`
}

// Operation wrapper for Logout.
// OperationLogoutRequest was auto-generated from WSDL.
type OperationLogoutRequest struct {
	Logout *Logout `xml:"logout,omitempty" json:"logout,omitempty" yaml:"logout,omitempty"`
}

// Operation wrapper for Logout.
// OperationLogoutResponse was auto-generated from WSDL.
type OperationLogoutResponse struct {
	LogoutResponse *LogoutResponse `xml:"logoutResponse,omitempty" json:"logoutResponse,omitempty" yaml:"logoutResponse,omitempty"`
}

// Operation wrapper for ModifyDocument.
// OperationModifyDocumentRequest was auto-generated from WSDL.
type OperationModifyDocumentRequest struct {
	ModifyDocument *ModifyDocument `xml:"modifyDocument,omitempty" json:"modifyDocument,omitempty" yaml:"modifyDocument,omitempty"`
}

// Operation wrapper for ModifyDocument.
// OperationModifyDocumentResponse was auto-generated from WSDL.
type OperationModifyDocumentResponse struct {
	ModifyDocumentResponse *ModifyDocumentResponse `xml:"modifyDocumentResponse,omitempty" json:"modifyDocumentResponse,omitempty" yaml:"modifyDocumentResponse,omitempty"`
}

// Operation wrapper for NotifyContacts.
// OperationNotifyContactsRequest was auto-generated from WSDL.
type OperationNotifyContactsRequest struct {
	NotifyContacts *NotifyContacts `xml:"notifyContacts,omitempty" json:"notifyContacts,omitempty" yaml:"notifyContacts,omitempty"`
}

// Operation wrapper for NotifyContacts.
// OperationNotifyContactsResponse was auto-generated from WSDL.
type OperationNotifyContactsResponse struct {
	NotifyContactsResponse *NotifyContactsResponse `xml:"notifyContactsResponse,omitempty" json:"notifyContactsResponse,omitempty" yaml:"notifyContactsResponse,omitempty"`
}

// Operation wrapper for RateDocument.
// OperationRateDocumentRequest was auto-generated from WSDL.
type OperationRateDocumentRequest struct {
	RateDocument *RateDocument `xml:"rateDocument,omitempty" json:"rateDocument,omitempty" yaml:"rateDocument,omitempty"`
}

// Operation wrapper for RateDocument.
// OperationRateDocumentResponse was auto-generated from WSDL.
type OperationRateDocumentResponse struct {
	RateDocumentResponse *RateDocumentResponse `xml:"rateDocumentResponse,omitempty" json:"rateDocumentResponse,omitempty" yaml:"rateDocumentResponse,omitempty"`
}

// Operation wrapper for RemoveAttachment.
// OperationRemoveAttachmentRequest was auto-generated from WSDL.
type OperationRemoveAttachmentRequest struct {
	RemoveAttachment *RemoveAttachment `xml:"removeAttachment,omitempty" json:"removeAttachment,omitempty" yaml:"removeAttachment,omitempty"`
}

// Operation wrapper for RemoveAttachment.
// OperationRemoveAttachmentResponse was auto-generated from WSDL.
type OperationRemoveAttachmentResponse struct {
	RemoveAttachmentResponse *RemoveAttachmentResponse `xml:"removeAttachmentResponse,omitempty" json:"removeAttachmentResponse,omitempty" yaml:"removeAttachmentResponse,omitempty"`
}

// Operation wrapper for RemoveLrelRelationships.
// OperationRemoveLrelRelationshipsRequest was auto-generated from
// WSDL.
type OperationRemoveLrelRelationshipsRequest struct {
	RemoveLrelRelationships *RemoveLrelRelationships `xml:"removeLrelRelationships,omitempty" json:"removeLrelRelationships,omitempty" yaml:"removeLrelRelationships,omitempty"`
}

// Operation wrapper for RemoveLrelRelationships.
// OperationRemoveLrelRelationshipsResponse was auto-generated
// from WSDL.
type OperationRemoveLrelRelationshipsResponse struct {
	RemoveLrelRelationshipsResponse *RemoveLrelRelationshipsResponse `xml:"removeLrelRelationshipsResponse,omitempty" json:"removeLrelRelationshipsResponse,omitempty" yaml:"removeLrelRelationshipsResponse,omitempty"`
}

// Operation wrapper for RemoveMemberFromGroup.
// OperationRemoveMemberFromGroupRequest was auto-generated from
// WSDL.
type OperationRemoveMemberFromGroupRequest struct {
	RemoveMemberFromGroup *RemoveMemberFromGroup `xml:"removeMemberFromGroup,omitempty" json:"removeMemberFromGroup,omitempty" yaml:"removeMemberFromGroup,omitempty"`
}

// Operation wrapper for RemoveMemberFromGroup.
// OperationRemoveMemberFromGroupResponse was auto-generated from
// WSDL.
type OperationRemoveMemberFromGroupResponse struct {
	RemoveMemberFromGroupResponse *RemoveMemberFromGroupResponse `xml:"removeMemberFromGroupResponse,omitempty" json:"removeMemberFromGroupResponse,omitempty" yaml:"removeMemberFromGroupResponse,omitempty"`
}

// Operation wrapper for Search.
// OperationSearchRequest was auto-generated from WSDL.
type OperationSearchRequest struct {
	Search *Search `xml:"search,omitempty" json:"search,omitempty" yaml:"search,omitempty"`
}

// Operation wrapper for Search.
// OperationSearchResponse was auto-generated from WSDL.
type OperationSearchResponse struct {
	SearchResponse *SearchResponse `xml:"searchResponse,omitempty" json:"searchResponse,omitempty" yaml:"searchResponse,omitempty"`
}

// Operation wrapper for ServerStatus.
// OperationServerStatusRequest was auto-generated from WSDL.
type OperationServerStatusRequest struct {
	ServerStatus *ServerStatus `xml:"serverStatus,omitempty" json:"serverStatus,omitempty" yaml:"serverStatus,omitempty"`
}

// Operation wrapper for ServerStatus.
// OperationServerStatusResponse was auto-generated from WSDL.
type OperationServerStatusResponse struct {
	ServerStatusResponse *ServerStatusResponse `xml:"serverStatusResponse,omitempty" json:"serverStatusResponse,omitempty" yaml:"serverStatusResponse,omitempty"`
}

// Operation wrapper for Transfer.
// OperationTransferRequest was auto-generated from WSDL.
type OperationTransferRequest struct {
	Transfer *Transfer `xml:"transfer,omitempty" json:"transfer,omitempty" yaml:"transfer,omitempty"`
}

// Operation wrapper for Transfer.
// OperationTransferResponse was auto-generated from WSDL.
type OperationTransferResponse struct {
	TransferResponse *TransferResponse `xml:"transferResponse,omitempty" json:"transferResponse,omitempty" yaml:"transferResponse,omitempty"`
}

// Operation wrapper for UpdateObject.
// OperationUpdateObjectRequest was auto-generated from WSDL.
type OperationUpdateObjectRequest struct {
	UpdateObject *UpdateObject `xml:"updateObject,omitempty" json:"updateObject,omitempty" yaml:"updateObject,omitempty"`
}

// Operation wrapper for UpdateObject.
// OperationUpdateObjectResponse was auto-generated from WSDL.
type OperationUpdateObjectResponse struct {
	UpdateObjectResponse *UpdateObjectResponse `xml:"updateObjectResponse,omitempty" json:"updateObjectResponse,omitempty" yaml:"updateObjectResponse,omitempty"`
}

// Operation wrapper for UpdateRating.
// OperationUpdateRatingRequest was auto-generated from WSDL.
type OperationUpdateRatingRequest struct {
	UpdateRating *UpdateRating `xml:"updateRating,omitempty" json:"updateRating,omitempty" yaml:"updateRating,omitempty"`
}

// Operation wrapper for UpdateRating.
// OperationUpdateRatingResponse was auto-generated from WSDL.
type OperationUpdateRatingResponse struct {
	UpdateRatingResponse *UpdateRatingResponse `xml:"updateRatingResponse,omitempty" json:"updateRatingResponse,omitempty" yaml:"updateRatingResponse,omitempty"`
}

// LoginCall is unified login method which reduces need to implement
// preparation for each method manually.
//
// Func name caused by conflict w/ SOAP legacy Login method.
func (p *SdmSoap) LoginCall(login LoginIntfc) (*int, error) {
	return login.Login(p)
}

// AddAssetLog was auto-generated from WSDL.
func (p *SdmSoap) AddAssetLog(AddAssetLog *AddAssetLog) (*AddAssetLogResponse, error) {
	α := struct {
		OperationAddAssetLogRequest `xml:"impl:addAssetLog"`
	}{
		OperationAddAssetLogRequest{
			AddAssetLog,
		},
	}

	γ := struct {
		OperationAddAssetLogResponse `xml:"addAssetLogResponse"`
	}{}
	if err := p.RoundTripWithAction("AddAssetLog", α, &γ); err != nil {
		return nil, err
	}
	return γ.AddAssetLogResponse, nil
}

// AddBookmark was auto-generated from WSDL.
func (p *SdmSoap) AddBookmark(AddBookmark *AddBookmark) (*AddBookmarkResponse, error) {
	α := struct {
		OperationAddBookmarkRequest `xml:"impl:addBookmark"`
	}{
		OperationAddBookmarkRequest{
			AddBookmark,
		},
	}

	γ := struct {
		OperationAddBookmarkResponse `xml:"addBookmarkResponse"`
	}{}
	if err := p.RoundTripWithAction("AddBookmark", α, &γ); err != nil {
		return nil, err
	}
	return γ.AddBookmarkResponse, nil
}

// AddComment was auto-generated from WSDL.
func (p *SdmSoap) AddComment(AddComment *AddComment) (*AddCommentResponse, error) {
	α := struct {
		OperationAddCommentRequest `xml:"impl:addComment"`
	}{
		OperationAddCommentRequest{
			AddComment,
		},
	}

	γ := struct {
		OperationAddCommentResponse `xml:"addCommentResponse"`
	}{}
	if err := p.RoundTripWithAction("AddComment", α, &γ); err != nil {
		return nil, err
	}
	return γ.AddCommentResponse, nil
}

// AddMemberToGroup was auto-generated from WSDL.
func (p *SdmSoap) AddMemberToGroup(AddMemberToGroup *AddMemberToGroup) (*AddMemberToGroupResponse, error) {
	α := struct {
		OperationAddMemberToGroupRequest `xml:"impl:addMemberToGroup"`
	}{
		OperationAddMemberToGroupRequest{
			AddMemberToGroup,
		},
	}

	γ := struct {
		OperationAddMemberToGroupResponse `xml:"addMemberToGroupResponse"`
	}{}
	if err := p.RoundTripWithAction("AddMemberToGroup", α, &γ); err != nil {
		return nil, err
	}
	return γ.AddMemberToGroupResponse, nil
}

// AttachChangeToRequest was auto-generated from WSDL.
func (p *SdmSoap) AttachChangeToRequest(AttachChangeToRequest *AttachChangeToRequest) (*AttachChangeToRequestResponse, error) {
	α := struct {
		OperationAttachChangeToRequestRequest `xml:"impl:attachChangeToRequest"`
	}{
		OperationAttachChangeToRequestRequest{
			AttachChangeToRequest,
		},
	}

	γ := struct {
		OperationAttachChangeToRequestResponse `xml:"attachChangeToRequestResponse"`
	}{}
	if err := p.RoundTripWithAction("AttachChangeToRequest", α, &γ); err != nil {
		return nil, err
	}
	return γ.AttachChangeToRequestResponse, nil
}

// AttachURLLink was auto-generated from WSDL.
func (p *SdmSoap) AttachURLLink(AttachURLLink *AttachURLLink) (*AttachURLLinkResponse, error) {
	α := struct {
		OperationAttachURLLinkRequest `xml:"impl:attachURLLink"`
	}{
		OperationAttachURLLinkRequest{
			AttachURLLink,
		},
	}

	γ := struct {
		OperationAttachURLLinkResponse `xml:"attachURLLinkResponse"`
	}{}
	if err := p.RoundTripWithAction("AttachURLLink", α, &γ); err != nil {
		return nil, err
	}
	return γ.AttachURLLinkResponse, nil
}

// AttachURLLinkToTicket was auto-generated from WSDL.
func (p *SdmSoap) AttachURLLinkToTicket(AttachURLLinkToTicket *AttachURLLinkToTicket) (*AttachURLLinkToTicketResponse, error) {
	α := struct {
		OperationAttachURLLinkToTicketRequest `xml:"impl:attachURLLinkToTicket"`
	}{
		OperationAttachURLLinkToTicketRequest{
			AttachURLLinkToTicket,
		},
	}

	γ := struct {
		OperationAttachURLLinkToTicketResponse `xml:"attachURLLinkToTicketResponse"`
	}{}
	if err := p.RoundTripWithAction("AttachURLLinkToTicket", α, &γ); err != nil {
		return nil, err
	}
	return γ.AttachURLLinkToTicketResponse, nil
}

// AttmntFolderLinkCount was auto-generated from WSDL.
func (p *SdmSoap) AttmntFolderLinkCount(AttmntFolderLinkCount *AttmntFolderLinkCount) (*AttmntFolderLinkCountResponse, error) {
	α := struct {
		OperationAttmntFolderLinkCountRequest `xml:"impl:attmntFolderLinkCount"`
	}{
		OperationAttmntFolderLinkCountRequest{
			AttmntFolderLinkCount,
		},
	}

	γ := struct {
		OperationAttmntFolderLinkCountResponse `xml:"attmntFolderLinkCountResponse"`
	}{}
	if err := p.RoundTripWithAction("AttmntFolderLinkCount", α, &γ); err != nil {
		return nil, err
	}
	return γ.AttmntFolderLinkCountResponse, nil
}

// CallServerMethod was auto-generated from WSDL.
func (p *SdmSoap) CallServerMethod(CallServerMethod *CallServerMethod) (*CallServerMethodResponse, error) {
	α := struct {
		OperationCallServerMethodRequest `xml:"impl:callServerMethod"`
	}{
		OperationCallServerMethodRequest{
			CallServerMethod,
		},
	}

	γ := struct {
		OperationCallServerMethodResponse `xml:"callServerMethodResponse"`
	}{}
	if err := p.RoundTripWithAction("CallServerMethod", α, &γ); err != nil {
		return nil, err
	}
	return γ.CallServerMethodResponse, nil
}

// ChangeStatus was auto-generated from WSDL.
func (p *SdmSoap) ChangeStatus(ChangeStatus *ChangeStatus) (*ChangeStatusResponse, error) {
	α := struct {
		OperationChangeStatusRequest `xml:"impl:changeStatus"`
	}{
		OperationChangeStatusRequest{
			ChangeStatus,
		},
	}

	γ := struct {
		OperationChangeStatusResponse `xml:"changeStatusResponse"`
	}{}
	if err := p.RoundTripWithAction("ChangeStatus", α, &γ); err != nil {
		return nil, err
	}
	return γ.ChangeStatusResponse, nil
}

// ClearNotification was auto-generated from WSDL.
func (p *SdmSoap) ClearNotification(ClearNotification *ClearNotification) (*ClearNotificationResponse, error) {
	α := struct {
		OperationClearNotificationRequest `xml:"impl:clearNotification"`
	}{
		OperationClearNotificationRequest{
			ClearNotification,
		},
	}

	γ := struct {
		OperationClearNotificationResponse `xml:"clearNotificationResponse"`
	}{}
	if err := p.RoundTripWithAction("ClearNotification", α, &γ); err != nil {
		return nil, err
	}
	return γ.ClearNotificationResponse, nil
}

// CloseTicket was auto-generated from WSDL.
func (p *SdmSoap) CloseTicket(CloseTicket *CloseTicket) (*CloseTicketResponse, error) {
	α := struct {
		OperationCloseTicketRequest `xml:"impl:closeTicket"`
	}{
		OperationCloseTicketRequest{
			CloseTicket,
		},
	}

	γ := struct {
		OperationCloseTicketResponse `xml:"closeTicketResponse"`
	}{}
	if err := p.RoundTripWithAction("CloseTicket", α, &γ); err != nil {
		return nil, err
	}
	return γ.CloseTicketResponse, nil
}

// CreateActivityLog was auto-generated from WSDL.
func (p *SdmSoap) CreateActivityLog(CreateActivityLog *CreateActivityLog) (*CreateActivityLogResponse, error) {
	α := struct {
		OperationCreateActivityLogRequest `xml:"impl:createActivityLog"`
	}{
		OperationCreateActivityLogRequest{
			CreateActivityLog,
		},
	}

	γ := struct {
		OperationCreateActivityLogResponse `xml:"createActivityLogResponse"`
	}{}
	if err := p.RoundTripWithAction("CreateActivityLog", α, &γ); err != nil {
		return nil, err
	}
	return γ.CreateActivityLogResponse, nil
}

// CreateAsset was auto-generated from WSDL.
func (p *SdmSoap) CreateAsset(CreateAsset *CreateAsset) (*CreateAssetResponse, error) {
	α := struct {
		OperationCreateAssetRequest `xml:"impl:createAsset"`
	}{
		OperationCreateAssetRequest{
			CreateAsset,
		},
	}

	γ := struct {
		OperationCreateAssetResponse `xml:"createAssetResponse"`
	}{}
	if err := p.RoundTripWithAction("CreateAsset", α, &γ); err != nil {
		return nil, err
	}
	return γ.CreateAssetResponse, nil
}

// CreateAssetParentChildRelationship was auto-generated from WSDL.
func (p *SdmSoap) CreateAssetParentChildRelationship(CreateAssetParentChildRelationship *CreateAssetParentChildRelationship) (*CreateAssetParentChildRelationshipResponse, error) {
	α := struct {
		OperationCreateAssetParentChildRelationshipRequest `xml:"impl:createAssetParentChildRelationship"`
	}{
		OperationCreateAssetParentChildRelationshipRequest{
			CreateAssetParentChildRelationship,
		},
	}

	γ := struct {
		OperationCreateAssetParentChildRelationshipResponse `xml:"createAssetParentChildRelationshipResponse"`
	}{}
	if err := p.RoundTripWithAction("CreateAssetParentChildRelationship", α, &γ); err != nil {
		return nil, err
	}
	return γ.CreateAssetParentChildRelationshipResponse, nil
}

// CreateAttachment was auto-generated from WSDL.
func (p *SdmSoap) CreateAttachment(CreateAttachment *CreateAttachment) (*CreateAttachmentResponse, error) {
	α := struct {
		OperationCreateAttachmentRequest `xml:"impl:createAttachment"`
	}{
		OperationCreateAttachmentRequest{
			CreateAttachment,
		},
	}

	γ := struct {
		OperationCreateAttachmentResponse `xml:"createAttachmentResponse"`
	}{}
	if err := p.RoundTripWithAction("CreateAttachment", α, &γ); err != nil {
		return nil, err
	}
	return γ.CreateAttachmentResponse, nil
}

// CreateAttmnt was auto-generated from WSDL.
func (p *SdmSoap) CreateAttmnt(CreateAttmnt *CreateAttmnt) (*CreateAttmntResponse, error) {
	α := struct {
		OperationCreateAttmntRequest `xml:"impl:createAttmnt"`
	}{
		OperationCreateAttmntRequest{
			CreateAttmnt,
		},
	}

	γ := struct {
		OperationCreateAttmntResponse `xml:"createAttmntResponse"`
	}{}
	if err := p.RoundTripWithAction("CreateAttmnt", α, &γ); err != nil {
		return nil, err
	}
	return γ.CreateAttmntResponse, nil
}

// CreateChangeOrder was auto-generated from WSDL.
func (p *SdmSoap) CreateChangeOrder(CreateChangeOrder *CreateChangeOrder) (*CreateChangeOrderResponse, error) {
	α := struct {
		OperationCreateChangeOrderRequest `xml:"impl:createChangeOrder"`
	}{
		OperationCreateChangeOrderRequest{
			CreateChangeOrder,
		},
	}

	γ := struct {
		OperationCreateChangeOrderResponse `xml:"createChangeOrderResponse"`
	}{}
	if err := p.RoundTripWithAction("CreateChangeOrder", α, &γ); err != nil {
		return nil, err
	}
	return γ.CreateChangeOrderResponse, nil
}

// CreateDocument was auto-generated from WSDL.
func (p *SdmSoap) CreateDocument(CreateDocument *CreateDocument) (*CreateDocumentResponse, error) {
	α := struct {
		OperationCreateDocumentRequest `xml:"impl:createDocument"`
	}{
		OperationCreateDocumentRequest{
			CreateDocument,
		},
	}

	γ := struct {
		OperationCreateDocumentResponse `xml:"createDocumentResponse"`
	}{}
	if err := p.RoundTripWithAction("CreateDocument", α, &γ); err != nil {
		return nil, err
	}
	return γ.CreateDocumentResponse, nil
}

// CreateFolder was auto-generated from WSDL.
func (p *SdmSoap) CreateFolder(CreateFolder *CreateFolder) (*CreateFolderResponse, error) {
	α := struct {
		OperationCreateFolderRequest `xml:"impl:createFolder"`
	}{
		OperationCreateFolderRequest{
			CreateFolder,
		},
	}

	γ := struct {
		OperationCreateFolderResponse `xml:"createFolderResponse"`
	}{}
	if err := p.RoundTripWithAction("CreateFolder", α, &γ); err != nil {
		return nil, err
	}
	return γ.CreateFolderResponse, nil
}

// CreateIssue was auto-generated from WSDL.
func (p *SdmSoap) CreateIssue(CreateIssue *CreateIssue) (*CreateIssueResponse, error) {
	α := struct {
		OperationCreateIssueRequest `xml:"impl:createIssue"`
	}{
		OperationCreateIssueRequest{
			CreateIssue,
		},
	}

	γ := struct {
		OperationCreateIssueResponse `xml:"createIssueResponse"`
	}{}
	if err := p.RoundTripWithAction("CreateIssue", α, &γ); err != nil {
		return nil, err
	}
	return γ.CreateIssueResponse, nil
}

// CreateLrelRelationships was auto-generated from WSDL.
func (p *SdmSoap) CreateLrelRelationships(CreateLrelRelationships *CreateLrelRelationships) (*CreateLrelRelationshipsResponse, error) {
	α := struct {
		OperationCreateLrelRelationshipsRequest `xml:"impl:createLrelRelationships"`
	}{
		OperationCreateLrelRelationshipsRequest{
			CreateLrelRelationships,
		},
	}

	γ := struct {
		OperationCreateLrelRelationshipsResponse `xml:"createLrelRelationshipsResponse"`
	}{}
	if err := p.RoundTripWithAction("CreateLrelRelationships", α, &γ); err != nil {
		return nil, err
	}
	return γ.CreateLrelRelationshipsResponse, nil
}

// CreateObject was auto-generated from WSDL.
func (p *SdmSoap) CreateObject(CreateObject *CreateObject) (*CreateObjectResponse, error) {
	α := struct {
		OperationCreateObjectRequest `xml:"impl:createObject"`
	}{
		OperationCreateObjectRequest{
			CreateObject,
		},
	}

	γ := struct {
		OperationCreateObjectResponse `xml:"createObjectResponse"`
	}{}
	if err := p.RoundTripWithAction("CreateObject", α, &γ); err != nil {
		return nil, err
	}
	return γ.CreateObjectResponse, nil
}

// CreateQuickTicket was auto-generated from WSDL.
func (p *SdmSoap) CreateQuickTicket(CreateQuickTicket *CreateQuickTicket) (*CreateQuickTicketResponse, error) {
	α := struct {
		OperationCreateQuickTicketRequest `xml:"impl:createQuickTicket"`
	}{
		OperationCreateQuickTicketRequest{
			CreateQuickTicket,
		},
	}

	γ := struct {
		OperationCreateQuickTicketResponse `xml:"createQuickTicketResponse"`
	}{}
	if err := p.RoundTripWithAction("CreateQuickTicket", α, &γ); err != nil {
		return nil, err
	}
	return γ.CreateQuickTicketResponse, nil
}

// CreateRequest was auto-generated from WSDL.
func (p *SdmSoap) CreateRequest(CreateRequest *CreateRequest) (*CreateRequestResponse, error) {
	α := struct {
		OperationCreateRequestRequest `xml:"impl:createRequest"`
	}{
		OperationCreateRequestRequest{
			CreateRequest,
		},
	}

	γ := struct {
		OperationCreateRequestResponse `xml:"createRequestResponse"`
	}{}
	if err := p.RoundTripWithAction("CreateRequest", α, &γ); err != nil {
		return nil, err
	}
	return γ.CreateRequestResponse, nil
}

// CreateTicket was auto-generated from WSDL.
func (p *SdmSoap) CreateTicket(CreateTicket *CreateTicket) (*CreateTicketResponse, error) {
	α := struct {
		OperationCreateTicketRequest `xml:"impl:createTicket"`
	}{
		OperationCreateTicketRequest{
			CreateTicket,
		},
	}

	γ := struct {
		OperationCreateTicketResponse `xml:"createTicketResponse"`
	}{}
	if err := p.RoundTripWithAction("CreateTicket", α, &γ); err != nil {
		return nil, err
	}
	return γ.CreateTicketResponse, nil
}

// CreateWorkFlowTask was auto-generated from WSDL.
func (p *SdmSoap) CreateWorkFlowTask(CreateWorkFlowTask *CreateWorkFlowTask) (*CreateWorkFlowTaskResponse, error) {
	α := struct {
		OperationCreateWorkFlowTaskRequest `xml:"impl:createWorkFlowTask"`
	}{
		OperationCreateWorkFlowTaskRequest{
			CreateWorkFlowTask,
		},
	}

	γ := struct {
		OperationCreateWorkFlowTaskResponse `xml:"createWorkFlowTaskResponse"`
	}{}
	if err := p.RoundTripWithAction("CreateWorkFlowTask", α, &γ); err != nil {
		return nil, err
	}
	return γ.CreateWorkFlowTaskResponse, nil
}

// DeleteBookmark was auto-generated from WSDL.
func (p *SdmSoap) DeleteBookmark(DeleteBookmark *DeleteBookmark) (*DeleteBookmarkResponse, error) {
	α := struct {
		OperationDeleteBookmarkRequest `xml:"impl:deleteBookmark"`
	}{
		OperationDeleteBookmarkRequest{
			DeleteBookmark,
		},
	}

	γ := struct {
		OperationDeleteBookmarkResponse `xml:"deleteBookmarkResponse"`
	}{}
	if err := p.RoundTripWithAction("DeleteBookmark", α, &γ); err != nil {
		return nil, err
	}
	return γ.DeleteBookmarkResponse, nil
}

// DeleteComment was auto-generated from WSDL.
func (p *SdmSoap) DeleteComment(DeleteComment *DeleteComment) (*DeleteCommentResponse, error) {
	α := struct {
		OperationDeleteCommentRequest `xml:"impl:deleteComment"`
	}{
		OperationDeleteCommentRequest{
			DeleteComment,
		},
	}

	γ := struct {
		OperationDeleteCommentResponse `xml:"deleteCommentResponse"`
	}{}
	if err := p.RoundTripWithAction("DeleteComment", α, &γ); err != nil {
		return nil, err
	}
	return γ.DeleteCommentResponse, nil
}

// DeleteDocument was auto-generated from WSDL.
func (p *SdmSoap) DeleteDocument(DeleteDocument *DeleteDocument) (*DeleteDocumentResponse, error) {
	α := struct {
		OperationDeleteDocumentRequest `xml:"impl:deleteDocument"`
	}{
		OperationDeleteDocumentRequest{
			DeleteDocument,
		},
	}

	γ := struct {
		OperationDeleteDocumentResponse `xml:"deleteDocumentResponse"`
	}{}
	if err := p.RoundTripWithAction("DeleteDocument", α, &γ); err != nil {
		return nil, err
	}
	return γ.DeleteDocumentResponse, nil
}

// DeleteWorkFlowTask was auto-generated from WSDL.
func (p *SdmSoap) DeleteWorkFlowTask(DeleteWorkFlowTask *DeleteWorkFlowTask) (*DeleteWorkFlowTaskResponse, error) {
	α := struct {
		OperationDeleteWorkFlowTaskRequest `xml:"impl:deleteWorkFlowTask"`
	}{
		OperationDeleteWorkFlowTaskRequest{
			DeleteWorkFlowTask,
		},
	}

	γ := struct {
		OperationDeleteWorkFlowTaskResponse `xml:"deleteWorkFlowTaskResponse"`
	}{}
	if err := p.RoundTripWithAction("DeleteWorkFlowTask", α, &γ); err != nil {
		return nil, err
	}
	return γ.DeleteWorkFlowTaskResponse, nil
}

// DetachChangeFromRequest was auto-generated from WSDL.
func (p *SdmSoap) DetachChangeFromRequest(DetachChangeFromRequest *DetachChangeFromRequest) (*DetachChangeFromRequestResponse, error) {
	α := struct {
		OperationDetachChangeFromRequestRequest `xml:"impl:detachChangeFromRequest"`
	}{
		OperationDetachChangeFromRequestRequest{
			DetachChangeFromRequest,
		},
	}

	γ := struct {
		OperationDetachChangeFromRequestResponse `xml:"detachChangeFromRequestResponse"`
	}{}
	if err := p.RoundTripWithAction("DetachChangeFromRequest", α, &γ); err != nil {
		return nil, err
	}
	return γ.DetachChangeFromRequestResponse, nil
}

// DoQuery was auto-generated from WSDL.
func (p *SdmSoap) DoQuery(DoQuery *DoQuery) (*DoQueryResponse, error) {
	α := struct {
		OperationDoQueryRequest `xml:"impl:doQuery"`
	}{
		OperationDoQueryRequest{
			DoQuery,
		},
	}

	γ := struct {
		OperationDoQueryResponse `xml:"doQueryResponse"`
	}{}
	if err := p.RoundTripWithAction("DoQuery", α, &γ); err != nil {
		return nil, err
	}
	return γ.DoQueryResponse, nil
}

// DoSelect was auto-generated from WSDL.
func (p *SdmSoap) DoSelect(DoSelect *DoSelect) (*DoSelectResponse, error) {
	α := struct {
		OperationDoSelectRequest `xml:"impl:doSelect"`
	}{
		OperationDoSelectRequest{
			DoSelect,
		},
	}

	γ := struct {
		OperationDoSelectResponse `xml:"doSelectResponse"`
	}{}
	if err := p.RoundTripWithAction("DoSelect", α, &γ); err != nil {
		return nil, err
	}
	return γ.DoSelectResponse, nil
}

// DoSelectKD was auto-generated from WSDL.
func (p *SdmSoap) DoSelectKD(DoSelectKD *DoSelectKD) (*DoSelectKDResponse, error) {
	α := struct {
		OperationDoSelectKDRequest `xml:"impl:doSelectKD"`
	}{
		OperationDoSelectKDRequest{
			DoSelectKD,
		},
	}

	γ := struct {
		OperationDoSelectKDResponse `xml:"doSelectKDResponse"`
	}{}
	if err := p.RoundTripWithAction("DoSelectKD", α, &γ); err != nil {
		return nil, err
	}
	return γ.DoSelectKDResponse, nil
}

// Escalate was auto-generated from WSDL.
func (p *SdmSoap) Escalate(Escalate *Escalate) (*EscalateResponse, error) {
	α := struct {
		OperationEscalateRequest `xml:"impl:escalate"`
	}{
		OperationEscalateRequest{
			Escalate,
		},
	}

	γ := struct {
		OperationEscalateResponse `xml:"escalateResponse"`
	}{}
	if err := p.RoundTripWithAction("Escalate", α, &γ); err != nil {
		return nil, err
	}
	return γ.EscalateResponse, nil
}

// Faq was auto-generated from WSDL.
func (p *SdmSoap) Faq(Faq *Faq) (*FaqResponse, error) {
	α := struct {
		OperationFaqRequest `xml:"impl:faq"`
	}{
		OperationFaqRequest{
			Faq,
		},
	}

	γ := struct {
		OperationFaqResponse `xml:"faqResponse"`
	}{}
	if err := p.RoundTripWithAction("Faq", α, &γ); err != nil {
		return nil, err
	}
	return γ.FaqResponse, nil
}

// FindContacts was auto-generated from WSDL.
func (p *SdmSoap) FindContacts(FindContacts *FindContacts) (*FindContactsResponse, error) {
	α := struct {
		OperationFindContactsRequest `xml:"impl:findContacts"`
	}{
		OperationFindContactsRequest{
			FindContacts,
		},
	}

	γ := struct {
		OperationFindContactsResponse `xml:"findContactsResponse"`
	}{}
	if err := p.RoundTripWithAction("FindContacts", α, &γ); err != nil {
		return nil, err
	}
	return γ.FindContactsResponse, nil
}

// FreeListHandles was auto-generated from WSDL.
func (p *SdmSoap) FreeListHandles(FreeListHandles *FreeListHandles) (*FreeListHandlesResponse, error) {
	α := struct {
		OperationFreeListHandlesRequest `xml:"impl:freeListHandles"`
	}{
		OperationFreeListHandlesRequest{
			FreeListHandles,
		},
	}

	γ := struct {
		OperationFreeListHandlesResponse `xml:"freeListHandlesResponse"`
	}{}
	if err := p.RoundTripWithAction("FreeListHandles", α, &γ); err != nil {
		return nil, err
	}
	return γ.FreeListHandlesResponse, nil
}

// GetAccessTypeForContact was auto-generated from WSDL.
func (p *SdmSoap) GetAccessTypeForContact(GetAccessTypeForContact *GetAccessTypeForContact) (*GetAccessTypeForContactResponse, error) {
	α := struct {
		OperationGetAccessTypeForContactRequest `xml:"impl:getAccessTypeForContact"`
	}{
		OperationGetAccessTypeForContactRequest{
			GetAccessTypeForContact,
		},
	}

	γ := struct {
		OperationGetAccessTypeForContactResponse `xml:"getAccessTypeForContactResponse"`
	}{}
	if err := p.RoundTripWithAction("GetAccessTypeForContact", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetAccessTypeForContactResponse, nil
}

// GetArtifact was auto-generated from WSDL.
func (p *SdmSoap) GetArtifact(GetArtifact *GetArtifact) (*GetArtifactResponse, error) {
	α := struct {
		OperationGetArtifactRequest `xml:"impl:getArtifact"`
	}{
		OperationGetArtifactRequest{
			GetArtifact,
		},
	}

	γ := struct {
		OperationGetArtifactResponse `xml:"getArtifactResponse"`
	}{}
	if err := p.RoundTripWithAction("GetArtifact", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetArtifactResponse, nil
}

// GetAssetExtensionInformation was auto-generated from WSDL.
func (p *SdmSoap) GetAssetExtensionInformation(GetAssetExtensionInformation *GetAssetExtensionInformation) (*GetAssetExtensionInformationResponse, error) {
	α := struct {
		OperationGetAssetExtensionInformationRequest `xml:"impl:getAssetExtensionInformation"`
	}{
		OperationGetAssetExtensionInformationRequest{
			GetAssetExtensionInformation,
		},
	}

	γ := struct {
		OperationGetAssetExtensionInformationResponse `xml:"getAssetExtensionInformationResponse"`
	}{}
	if err := p.RoundTripWithAction("GetAssetExtensionInformation", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetAssetExtensionInformationResponse, nil
}

// GetAttmntInfo was auto-generated from WSDL.
func (p *SdmSoap) GetAttmntInfo(GetAttmntInfo *GetAttmntInfo) (*GetAttmntInfoResponse, error) {
	α := struct {
		OperationGetAttmntInfoRequest `xml:"impl:getAttmntInfo"`
	}{
		OperationGetAttmntInfoRequest{
			GetAttmntInfo,
		},
	}

	γ := struct {
		OperationGetAttmntInfoResponse `xml:"getAttmntInfoResponse"`
	}{}
	if err := p.RoundTripWithAction("GetAttmntInfo", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetAttmntInfoResponse, nil
}

// GetAttmntList was auto-generated from WSDL.
func (p *SdmSoap) GetAttmntList(GetAttmntList *GetAttmntList) (*GetAttmntListResponse, error) {
	α := struct {
		OperationGetAttmntListRequest `xml:"impl:getAttmntList"`
	}{
		OperationGetAttmntListRequest{
			GetAttmntList,
		},
	}

	γ := struct {
		OperationGetAttmntListResponse `xml:"getAttmntListResponse"`
	}{}
	if err := p.RoundTripWithAction("GetAttmntList", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetAttmntListResponse, nil
}

// GetAttmntListPerKD was auto-generated from WSDL.
func (p *SdmSoap) GetAttmntListPerKD(GetAttmntListPerKD *GetAttmntListPerKD) (*GetAttmntListPerKDResponse, error) {
	α := struct {
		OperationGetAttmntListPerKDRequest `xml:"impl:getAttmntListPerKD"`
	}{
		OperationGetAttmntListPerKDRequest{
			GetAttmntListPerKD,
		},
	}

	γ := struct {
		OperationGetAttmntListPerKDResponse `xml:"getAttmntListPerKDResponse"`
	}{}
	if err := p.RoundTripWithAction("GetAttmntListPerKD", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetAttmntListPerKDResponse, nil
}

// GetBookmarks was auto-generated from WSDL.
func (p *SdmSoap) GetBookmarks(GetBookmarks *GetBookmarks) (*GetBookmarksResponse, error) {
	α := struct {
		OperationGetBookmarksRequest `xml:"impl:getBookmarks"`
	}{
		OperationGetBookmarksRequest{
			GetBookmarks,
		},
	}

	γ := struct {
		OperationGetBookmarksResponse `xml:"getBookmarksResponse"`
	}{}
	if err := p.RoundTripWithAction("GetBookmarks", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetBookmarksResponse, nil
}

// GetBopsid was auto-generated from WSDL.
func (p *SdmSoap) GetBopsid(GetBopsid *GetBopsid) (*GetBopsidResponse, error) {
	α := struct {
		OperationGetBopsidRequest `xml:"impl:getBopsid"`
	}{
		OperationGetBopsidRequest{
			GetBopsid,
		},
	}

	γ := struct {
		OperationGetBopsidResponse `xml:"getBopsidResponse"`
	}{}
	if err := p.RoundTripWithAction("GetBopsid", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetBopsidResponse, nil
}

// GetCategory was auto-generated from WSDL.
func (p *SdmSoap) GetCategory(GetCategory *GetCategory) (*GetCategoryResponse, error) {
	α := struct {
		OperationGetCategoryRequest `xml:"impl:getCategory"`
	}{
		OperationGetCategoryRequest{
			GetCategory,
		},
	}

	γ := struct {
		OperationGetCategoryResponse `xml:"getCategoryResponse"`
	}{}
	if err := p.RoundTripWithAction("GetCategory", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetCategoryResponse, nil
}

// GetComments was auto-generated from WSDL.
func (p *SdmSoap) GetComments(GetComments *GetComments) (*GetCommentsResponse, error) {
	α := struct {
		OperationGetCommentsRequest `xml:"impl:getComments"`
	}{
		OperationGetCommentsRequest{
			GetComments,
		},
	}

	γ := struct {
		OperationGetCommentsResponse `xml:"getCommentsResponse"`
	}{}
	if err := p.RoundTripWithAction("GetComments", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetCommentsResponse, nil
}

// GetConfigurationMode was auto-generated from WSDL.
func (p *SdmSoap) GetConfigurationMode(GetConfigurationMode *GetConfigurationMode) (*GetConfigurationModeResponse, error) {
	α := struct {
		OperationGetConfigurationModeRequest `xml:"impl:getConfigurationMode"`
	}{
		OperationGetConfigurationModeRequest{
			GetConfigurationMode,
		},
	}

	γ := struct {
		OperationGetConfigurationModeResponse `xml:"getConfigurationModeResponse"`
	}{}
	if err := p.RoundTripWithAction("GetConfigurationMode", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetConfigurationModeResponse, nil
}

// GetContact was auto-generated from WSDL.
func (p *SdmSoap) GetContact(GetContact *GetContact) (*GetContactResponse, error) {
	α := struct {
		OperationGetContactRequest `xml:"impl:getContact"`
	}{
		OperationGetContactRequest{
			GetContact,
		},
	}

	γ := struct {
		OperationGetContactResponse `xml:"getContactResponse"`
	}{}
	if err := p.RoundTripWithAction("GetContact", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetContactResponse, nil
}

// GetDecisionTrees was auto-generated from WSDL.
func (p *SdmSoap) GetDecisionTrees(GetDecisionTrees *GetDecisionTrees) (*GetDecisionTreesResponse, error) {
	α := struct {
		OperationGetDecisionTreesRequest `xml:"impl:getDecisionTrees"`
	}{
		OperationGetDecisionTreesRequest{
			GetDecisionTrees,
		},
	}

	γ := struct {
		OperationGetDecisionTreesResponse `xml:"getDecisionTreesResponse"`
	}{}
	if err := p.RoundTripWithAction("GetDecisionTrees", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetDecisionTreesResponse, nil
}

// GetDependentAttrControls was auto-generated from WSDL.
func (p *SdmSoap) GetDependentAttrControls(GetDependentAttrControls *GetDependentAttrControls) (*GetDependentAttrControlsResponse, error) {
	α := struct {
		OperationGetDependentAttrControlsRequest `xml:"impl:getDependentAttrControls"`
	}{
		OperationGetDependentAttrControlsRequest{
			GetDependentAttrControls,
		},
	}

	γ := struct {
		OperationGetDependentAttrControlsResponse `xml:"getDependentAttrControlsResponse"`
	}{}
	if err := p.RoundTripWithAction("GetDependentAttrControls", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetDependentAttrControlsResponse, nil
}

// GetDocument was auto-generated from WSDL.
func (p *SdmSoap) GetDocument(GetDocument *GetDocument) (*GetDocumentResponse, error) {
	α := struct {
		OperationGetDocumentRequest `xml:"impl:getDocument"`
	}{
		OperationGetDocumentRequest{
			GetDocument,
		},
	}

	γ := struct {
		OperationGetDocumentResponse `xml:"getDocumentResponse"`
	}{}
	if err := p.RoundTripWithAction("GetDocument", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetDocumentResponse, nil
}

// GetDocumentTypes was auto-generated from WSDL.
func (p *SdmSoap) GetDocumentTypes(GetDocumentTypes *GetDocumentTypes) (*GetDocumentTypesResponse, error) {
	α := struct {
		OperationGetDocumentTypesRequest `xml:"impl:getDocumentTypes"`
	}{
		OperationGetDocumentTypesRequest{
			GetDocumentTypes,
		},
	}

	γ := struct {
		OperationGetDocumentTypesResponse `xml:"getDocumentTypesResponse"`
	}{}
	if err := p.RoundTripWithAction("GetDocumentTypes", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetDocumentTypesResponse, nil
}

// GetDocumentsByIDs was auto-generated from WSDL.
func (p *SdmSoap) GetDocumentsByIDs(GetDocumentsByIDs *GetDocumentsByIDs) (*GetDocumentsByIDsResponse, error) {
	α := struct {
		OperationGetDocumentsByIDsRequest `xml:"impl:getDocumentsByIDs"`
	}{
		OperationGetDocumentsByIDsRequest{
			GetDocumentsByIDs,
		},
	}

	γ := struct {
		OperationGetDocumentsByIDsResponse `xml:"getDocumentsByIDsResponse"`
	}{}
	if err := p.RoundTripWithAction("GetDocumentsByIDs", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetDocumentsByIDsResponse, nil
}

// GetFolderInfo was auto-generated from WSDL.
func (p *SdmSoap) GetFolderInfo(GetFolderInfo *GetFolderInfo) (*GetFolderInfoResponse, error) {
	α := struct {
		OperationGetFolderInfoRequest `xml:"impl:getFolderInfo"`
	}{
		OperationGetFolderInfoRequest{
			GetFolderInfo,
		},
	}

	γ := struct {
		OperationGetFolderInfoResponse `xml:"getFolderInfoResponse"`
	}{}
	if err := p.RoundTripWithAction("GetFolderInfo", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetFolderInfoResponse, nil
}

// GetFolderList was auto-generated from WSDL.
func (p *SdmSoap) GetFolderList(GetFolderList *GetFolderList) (*GetFolderListResponse, error) {
	α := struct {
		OperationGetFolderListRequest `xml:"impl:getFolderList"`
	}{
		OperationGetFolderListRequest{
			GetFolderList,
		},
	}

	γ := struct {
		OperationGetFolderListResponse `xml:"getFolderListResponse"`
	}{}
	if err := p.RoundTripWithAction("GetFolderList", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetFolderListResponse, nil
}

// GetGroupMemberListValues was auto-generated from WSDL.
func (p *SdmSoap) GetGroupMemberListValues(GetGroupMemberListValues *GetGroupMemberListValues) (*GetGroupMemberListValuesResponse, error) {
	α := struct {
		OperationGetGroupMemberListValuesRequest `xml:"impl:getGroupMemberListValues"`
	}{
		OperationGetGroupMemberListValuesRequest{
			GetGroupMemberListValues,
		},
	}

	γ := struct {
		OperationGetGroupMemberListValuesResponse `xml:"getGroupMemberListValuesResponse"`
	}{}
	if err := p.RoundTripWithAction("GetGroupMemberListValues", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetGroupMemberListValuesResponse, nil
}

// GetHandleForUserid was auto-generated from WSDL.
func (p *SdmSoap) GetHandleForUserid(GetHandleForUserid *GetHandleForUserid) (*GetHandleForUseridResponse, error) {
	α := struct {
		OperationGetHandleForUseridRequest `xml:"impl:getHandleForUserid"`
	}{
		OperationGetHandleForUseridRequest{
			GetHandleForUserid,
		},
	}

	γ := struct {
		OperationGetHandleForUseridResponse `xml:"getHandleForUseridResponse"`
	}{}
	if err := p.RoundTripWithAction("GetHandleForUserid", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetHandleForUseridResponse, nil
}

// GetKDListPerAttmnt was auto-generated from WSDL.
func (p *SdmSoap) GetKDListPerAttmnt(GetKDListPerAttmnt *GetKDListPerAttmnt) (*GetKDListPerAttmntResponse, error) {
	α := struct {
		OperationGetKDListPerAttmntRequest `xml:"impl:getKDListPerAttmnt"`
	}{
		OperationGetKDListPerAttmntRequest{
			GetKDListPerAttmnt,
		},
	}

	γ := struct {
		OperationGetKDListPerAttmntResponse `xml:"getKDListPerAttmntResponse"`
	}{}
	if err := p.RoundTripWithAction("GetKDListPerAttmnt", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetKDListPerAttmntResponse, nil
}

// GetListValues was auto-generated from WSDL.
func (p *SdmSoap) GetListValues(GetListValues *GetListValues) (*GetListValuesResponse, error) {
	α := struct {
		OperationGetListValuesRequest `xml:"impl:getListValues"`
	}{
		OperationGetListValuesRequest{
			GetListValues,
		},
	}

	γ := struct {
		OperationGetListValuesResponse `xml:"getListValuesResponse"`
	}{}
	if err := p.RoundTripWithAction("GetListValues", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetListValuesResponse, nil
}

// GetLrelLength was auto-generated from WSDL.
func (p *SdmSoap) GetLrelLength(GetLrelLength *GetLrelLength) (*GetLrelLengthResponse, error) {
	α := struct {
		OperationGetLrelLengthRequest `xml:"impl:getLrelLength"`
	}{
		OperationGetLrelLengthRequest{
			GetLrelLength,
		},
	}

	γ := struct {
		OperationGetLrelLengthResponse `xml:"getLrelLengthResponse"`
	}{}
	if err := p.RoundTripWithAction("GetLrelLength", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetLrelLengthResponse, nil
}

// GetLrelValues was auto-generated from WSDL.
func (p *SdmSoap) GetLrelValues(GetLrelValues *GetLrelValues) (*GetLrelValuesResponse, error) {
	α := struct {
		OperationGetLrelValuesRequest `xml:"impl:getLrelValues"`
	}{
		OperationGetLrelValuesRequest{
			GetLrelValues,
		},
	}

	γ := struct {
		OperationGetLrelValuesResponse `xml:"getLrelValuesResponse"`
	}{}
	if err := p.RoundTripWithAction("GetLrelValues", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetLrelValuesResponse, nil
}

// GetNotificationsForContact was auto-generated from WSDL.
func (p *SdmSoap) GetNotificationsForContact(GetNotificationsForContact *GetNotificationsForContact) (*GetNotificationsForContactResponse, error) {
	α := struct {
		OperationGetNotificationsForContactRequest `xml:"impl:getNotificationsForContact"`
	}{
		OperationGetNotificationsForContactRequest{
			GetNotificationsForContact,
		},
	}

	γ := struct {
		OperationGetNotificationsForContactResponse `xml:"getNotificationsForContactResponse"`
	}{}
	if err := p.RoundTripWithAction("GetNotificationsForContact", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetNotificationsForContactResponse, nil
}

// GetObjectTypeInformation was auto-generated from WSDL.
func (p *SdmSoap) GetObjectTypeInformation(GetObjectTypeInformation *GetObjectTypeInformation) (*GetObjectTypeInformationResponse, error) {
	α := struct {
		OperationGetObjectTypeInformationRequest `xml:"impl:getObjectTypeInformation"`
	}{
		OperationGetObjectTypeInformationRequest{
			GetObjectTypeInformation,
		},
	}

	γ := struct {
		OperationGetObjectTypeInformationResponse `xml:"getObjectTypeInformationResponse"`
	}{}
	if err := p.RoundTripWithAction("GetObjectTypeInformation", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetObjectTypeInformationResponse, nil
}

// GetObjectValues was auto-generated from WSDL.
func (p *SdmSoap) GetObjectValues(GetObjectValues *GetObjectValues) (*GetObjectValuesResponse, error) {
	α := struct {
		OperationGetObjectValuesRequest `xml:"impl:getObjectValues"`
	}{
		OperationGetObjectValuesRequest{
			GetObjectValues,
		},
	}

	γ := struct {
		OperationGetObjectValuesResponse `xml:"getObjectValuesResponse"`
	}{}
	if err := p.RoundTripWithAction("GetObjectValues", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetObjectValuesResponse, nil
}

// GetPendingChangeTaskListForContact was auto-generated from WSDL.
func (p *SdmSoap) GetPendingChangeTaskListForContact(GetPendingChangeTaskListForContact *GetPendingChangeTaskListForContact) (*GetPendingChangeTaskListForContactResponse, error) {
	α := struct {
		OperationGetPendingChangeTaskListForContactRequest `xml:"impl:getPendingChangeTaskListForContact"`
	}{
		OperationGetPendingChangeTaskListForContactRequest{
			GetPendingChangeTaskListForContact,
		},
	}

	γ := struct {
		OperationGetPendingChangeTaskListForContactResponse `xml:"getPendingChangeTaskListForContactResponse"`
	}{}
	if err := p.RoundTripWithAction("GetPendingChangeTaskListForContact", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetPendingChangeTaskListForContactResponse, nil
}

// GetPendingIssueTaskListForContact was auto-generated from WSDL.
func (p *SdmSoap) GetPendingIssueTaskListForContact(GetPendingIssueTaskListForContact *GetPendingIssueTaskListForContact) (*GetPendingIssueTaskListForContactResponse, error) {
	α := struct {
		OperationGetPendingIssueTaskListForContactRequest `xml:"impl:getPendingIssueTaskListForContact"`
	}{
		OperationGetPendingIssueTaskListForContactRequest{
			GetPendingIssueTaskListForContact,
		},
	}

	γ := struct {
		OperationGetPendingIssueTaskListForContactResponse `xml:"getPendingIssueTaskListForContactResponse"`
	}{}
	if err := p.RoundTripWithAction("GetPendingIssueTaskListForContact", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetPendingIssueTaskListForContactResponse, nil
}

// GetPermissionGroups was auto-generated from WSDL.
func (p *SdmSoap) GetPermissionGroups(GetPermissionGroups *GetPermissionGroups) (*GetPermissionGroupsResponse, error) {
	α := struct {
		OperationGetPermissionGroupsRequest `xml:"impl:getPermissionGroups"`
	}{
		OperationGetPermissionGroupsRequest{
			GetPermissionGroups,
		},
	}

	γ := struct {
		OperationGetPermissionGroupsResponse `xml:"getPermissionGroupsResponse"`
	}{}
	if err := p.RoundTripWithAction("GetPermissionGroups", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetPermissionGroupsResponse, nil
}

// GetPolicyInfo was auto-generated from WSDL.
func (p *SdmSoap) GetPolicyInfo(GetPolicyInfo *GetPolicyInfo) (*GetPolicyInfoResponse, error) {
	α := struct {
		OperationGetPolicyInfoRequest `xml:"impl:getPolicyInfo"`
	}{
		OperationGetPolicyInfoRequest{
			GetPolicyInfo,
		},
	}

	γ := struct {
		OperationGetPolicyInfoResponse `xml:"getPolicyInfoResponse"`
	}{}
	if err := p.RoundTripWithAction("GetPolicyInfo", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetPolicyInfoResponse, nil
}

// GetPriorities was auto-generated from WSDL.
func (p *SdmSoap) GetPriorities(GetPriorities *GetPriorities) (*GetPrioritiesResponse, error) {
	α := struct {
		OperationGetPrioritiesRequest `xml:"impl:getPriorities"`
	}{
		OperationGetPrioritiesRequest{
			GetPriorities,
		},
	}

	γ := struct {
		OperationGetPrioritiesResponse `xml:"getPrioritiesResponse"`
	}{}
	if err := p.RoundTripWithAction("GetPriorities", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetPrioritiesResponse, nil
}

// GetPropertyInfoForCategory was auto-generated from WSDL.
func (p *SdmSoap) GetPropertyInfoForCategory(GetPropertyInfoForCategory *GetPropertyInfoForCategory) (*GetPropertyInfoForCategoryResponse, error) {
	α := struct {
		OperationGetPropertyInfoForCategoryRequest `xml:"impl:getPropertyInfoForCategory"`
	}{
		OperationGetPropertyInfoForCategoryRequest{
			GetPropertyInfoForCategory,
		},
	}

	γ := struct {
		OperationGetPropertyInfoForCategoryResponse `xml:"getPropertyInfoForCategoryResponse"`
	}{}
	if err := p.RoundTripWithAction("GetPropertyInfoForCategory", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetPropertyInfoForCategoryResponse, nil
}

// GetQuestionsAsked was auto-generated from WSDL.
func (p *SdmSoap) GetQuestionsAsked(GetQuestionsAsked *GetQuestionsAsked) (*GetQuestionsAskedResponse, error) {
	α := struct {
		OperationGetQuestionsAskedRequest `xml:"impl:getQuestionsAsked"`
	}{
		OperationGetQuestionsAskedRequest{
			GetQuestionsAsked,
		},
	}

	γ := struct {
		OperationGetQuestionsAskedResponse `xml:"getQuestionsAskedResponse"`
	}{}
	if err := p.RoundTripWithAction("GetQuestionsAsked", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetQuestionsAskedResponse, nil
}

// GetRelatedList was auto-generated from WSDL.
func (p *SdmSoap) GetRelatedList(GetRelatedList *GetRelatedList) (*GetRelatedListResponse, error) {
	α := struct {
		OperationGetRelatedListRequest `xml:"impl:getRelatedList"`
	}{
		OperationGetRelatedListRequest{
			GetRelatedList,
		},
	}

	γ := struct {
		OperationGetRelatedListResponse `xml:"getRelatedListResponse"`
	}{}
	if err := p.RoundTripWithAction("GetRelatedList", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetRelatedListResponse, nil
}

// GetRelatedListValues was auto-generated from WSDL.
func (p *SdmSoap) GetRelatedListValues(GetRelatedListValues *GetRelatedListValues) (*GetRelatedListValuesResponse, error) {
	α := struct {
		OperationGetRelatedListValuesRequest `xml:"impl:getRelatedListValues"`
	}{
		OperationGetRelatedListValuesRequest{
			GetRelatedListValues,
		},
	}

	γ := struct {
		OperationGetRelatedListValuesResponse `xml:"getRelatedListValuesResponse"`
	}{}
	if err := p.RoundTripWithAction("GetRelatedListValues", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetRelatedListValuesResponse, nil
}

// GetRepositoryInfo was auto-generated from WSDL.
func (p *SdmSoap) GetRepositoryInfo(GetRepositoryInfo *GetRepositoryInfo) (*GetRepositoryInfoResponse, error) {
	α := struct {
		OperationGetRepositoryInfoRequest `xml:"impl:getRepositoryInfo"`
	}{
		OperationGetRepositoryInfoRequest{
			GetRepositoryInfo,
		},
	}

	γ := struct {
		OperationGetRepositoryInfoResponse `xml:"getRepositoryInfoResponse"`
	}{}
	if err := p.RoundTripWithAction("GetRepositoryInfo", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetRepositoryInfoResponse, nil
}

// GetStatuses was auto-generated from WSDL.
func (p *SdmSoap) GetStatuses(GetStatuses *GetStatuses) (*GetStatusesResponse, error) {
	α := struct {
		OperationGetStatusesRequest `xml:"impl:getStatuses"`
	}{
		OperationGetStatusesRequest{
			GetStatuses,
		},
	}

	γ := struct {
		OperationGetStatusesResponse `xml:"getStatusesResponse"`
	}{}
	if err := p.RoundTripWithAction("GetStatuses", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetStatusesResponse, nil
}

// GetTaskListValues was auto-generated from WSDL.
func (p *SdmSoap) GetTaskListValues(GetTaskListValues *GetTaskListValues) (*GetTaskListValuesResponse, error) {
	α := struct {
		OperationGetTaskListValuesRequest `xml:"impl:getTaskListValues"`
	}{
		OperationGetTaskListValuesRequest{
			GetTaskListValues,
		},
	}

	γ := struct {
		OperationGetTaskListValuesResponse `xml:"getTaskListValuesResponse"`
	}{}
	if err := p.RoundTripWithAction("GetTaskListValues", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetTaskListValuesResponse, nil
}

// GetTemplateList was auto-generated from WSDL.
func (p *SdmSoap) GetTemplateList(GetTemplateList *GetTemplateList) (*GetTemplateListResponse, error) {
	α := struct {
		OperationGetTemplateListRequest `xml:"impl:getTemplateList"`
	}{
		OperationGetTemplateListRequest{
			GetTemplateList,
		},
	}

	γ := struct {
		OperationGetTemplateListResponse `xml:"getTemplateListResponse"`
	}{}
	if err := p.RoundTripWithAction("GetTemplateList", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetTemplateListResponse, nil
}

// GetValidTaskTransitions was auto-generated from WSDL.
func (p *SdmSoap) GetValidTaskTransitions(GetValidTaskTransitions *GetValidTaskTransitions) (*GetValidTaskTransitionsResponse, error) {
	α := struct {
		OperationGetValidTaskTransitionsRequest `xml:"impl:getValidTaskTransitions"`
	}{
		OperationGetValidTaskTransitionsRequest{
			GetValidTaskTransitions,
		},
	}

	γ := struct {
		OperationGetValidTaskTransitionsResponse `xml:"getValidTaskTransitionsResponse"`
	}{}
	if err := p.RoundTripWithAction("GetValidTaskTransitions", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetValidTaskTransitionsResponse, nil
}

// GetValidTransitions was auto-generated from WSDL.
func (p *SdmSoap) GetValidTransitions(GetValidTransitions *GetValidTransitions) (*GetValidTransitionsResponse, error) {
	α := struct {
		OperationGetValidTransitionsRequest `xml:"impl:getValidTransitions"`
	}{
		OperationGetValidTransitionsRequest{
			GetValidTransitions,
		},
	}

	γ := struct {
		OperationGetValidTransitionsResponse `xml:"getValidTransitionsResponse"`
	}{}
	if err := p.RoundTripWithAction("GetValidTransitions", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetValidTransitionsResponse, nil
}

// GetWorkFlowTemplates was auto-generated from WSDL.
func (p *SdmSoap) GetWorkFlowTemplates(GetWorkFlowTemplates *GetWorkFlowTemplates) (*GetWorkFlowTemplatesResponse, error) {
	α := struct {
		OperationGetWorkFlowTemplatesRequest `xml:"impl:getWorkFlowTemplates"`
	}{
		OperationGetWorkFlowTemplatesRequest{
			GetWorkFlowTemplates,
		},
	}

	γ := struct {
		OperationGetWorkFlowTemplatesResponse `xml:"getWorkFlowTemplatesResponse"`
	}{}
	if err := p.RoundTripWithAction("GetWorkFlowTemplates", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetWorkFlowTemplatesResponse, nil
}

// GetWorkflowTemplateList was auto-generated from WSDL.
func (p *SdmSoap) GetWorkflowTemplateList(GetWorkflowTemplateList *GetWorkflowTemplateList) (*GetWorkflowTemplateListResponse, error) {
	α := struct {
		OperationGetWorkflowTemplateListRequest `xml:"impl:getWorkflowTemplateList"`
	}{
		OperationGetWorkflowTemplateListRequest{
			GetWorkflowTemplateList,
		},
	}

	γ := struct {
		OperationGetWorkflowTemplateListResponse `xml:"getWorkflowTemplateListResponse"`
	}{}
	if err := p.RoundTripWithAction("GetWorkflowTemplateList", α, &γ); err != nil {
		return nil, err
	}
	return γ.GetWorkflowTemplateListResponse, nil
}

// Impersonate was auto-generated from WSDL.
func (p *SdmSoap) Impersonate(Impersonate *Impersonate) (*ImpersonateResponse, error) {
	α := struct {
		OperationImpersonateRequest `xml:"impl:impersonate"`
	}{
		OperationImpersonateRequest{
			Impersonate,
		},
	}

	γ := struct {
		OperationImpersonateResponse `xml:"impersonateResponse"`
	}{}
	if err := p.RoundTripWithAction("Impersonate", α, &γ); err != nil {
		return nil, err
	}
	return γ.ImpersonateResponse, nil
}

// IsAttmntLinkedKD was auto-generated from WSDL.
func (p *SdmSoap) IsAttmntLinkedKD(IsAttmntLinkedKD *IsAttmntLinkedKD) (*IsAttmntLinkedKDResponse, error) {
	α := struct {
		OperationIsAttmntLinkedKDRequest `xml:"impl:isAttmntLinkedKD"`
	}{
		OperationIsAttmntLinkedKDRequest{
			IsAttmntLinkedKD,
		},
	}

	γ := struct {
		OperationIsAttmntLinkedKDResponse `xml:"isAttmntLinkedKDResponse"`
	}{}
	if err := p.RoundTripWithAction("IsAttmntLinkedKD", α, &γ); err != nil {
		return nil, err
	}
	return γ.IsAttmntLinkedKDResponse, nil
}

// LogComment was auto-generated from WSDL.
func (p *SdmSoap) LogComment(LogComment *LogComment) (*LogCommentResponse, error) {
	α := struct {
		OperationLogCommentRequest `xml:"impl:logComment"`
	}{
		OperationLogCommentRequest{
			LogComment,
		},
	}

	γ := struct {
		OperationLogCommentResponse `xml:"logCommentResponse"`
	}{}
	if err := p.RoundTripWithAction("LogComment", α, &γ); err != nil {
		return nil, err
	}
	return γ.LogCommentResponse, nil
}

// Login was auto-generated from WSDL.
func (p *SdmSoap) Login(Login *Login) (*LoginResponse, error) {
	α := struct {
		OperationLoginRequest `xml:"impl:login"`
	}{
		OperationLoginRequest{
			Login,
		},
	}

	γ := struct {
		OperationLoginResponse `xml:"loginResponse"`
	}{}
	if err := p.RoundTripWithAction("Login", α, &γ); err != nil {
		return nil, err
	}
	return γ.LoginResponse, nil
}

// LoginService was auto-generated from WSDL.
func (p *SdmSoap) LoginService(LoginService *LoginService) (*LoginServiceResponse, error) {
	α := struct {
		OperationLoginServiceRequest `xml:"impl:loginService"`
	}{
		OperationLoginServiceRequest{
			LoginService,
		},
	}

	γ := struct {
		OperationLoginServiceResponse `xml:"loginServiceResponse"`
	}{}
	if err := p.RoundTripWithAction("LoginService", α, &γ); err != nil {
		return nil, err
	}
	return γ.LoginServiceResponse, nil
}

// LoginServiceManaged was auto-generated from WSDL.
func (p *SdmSoap) LoginServiceManaged(LoginServiceManaged *LoginServiceManaged) (*LoginServiceManagedResponse, error) {
	α := struct {
		OperationLoginServiceManagedRequest `xml:"impl:loginServiceManaged"`
	}{
		OperationLoginServiceManagedRequest{
			LoginServiceManaged,
		},
	}

	γ := struct {
		OperationLoginServiceManagedResponse `xml:"loginServiceManagedResponse"`
	}{}
	if err := p.RoundTripWithAction("LoginServiceManaged", α, &γ); err != nil {
		return nil, err
	}
	return γ.LoginServiceManagedResponse, nil
}

// LoginWithArtifact was auto-generated from WSDL.
func (p *SdmSoap) LoginWithArtifact(LoginWithArtifact *LoginWithArtifact) (*LoginWithArtifactResponse, error) {
	α := struct {
		OperationLoginWithArtifactRequest `xml:"impl:loginWithArtifact"`
	}{
		OperationLoginWithArtifactRequest{
			LoginWithArtifact,
		},
	}

	γ := struct {
		OperationLoginWithArtifactResponse `xml:"loginWithArtifactResponse"`
	}{}
	if err := p.RoundTripWithAction("LoginWithArtifact", α, &γ); err != nil {
		return nil, err
	}
	return γ.LoginWithArtifactResponse, nil
}

// Logout was auto-generated from WSDL.
func (p *SdmSoap) Logout(Logout *Logout) (*LogoutResponse, error) {
	α := struct {
		OperationLogoutRequest `xml:"impl:logout"`
	}{
		OperationLogoutRequest{
			Logout,
		},
	}

	γ := struct {
		OperationLogoutResponse `xml:"logoutResponse"`
	}{}
	if err := p.RoundTripWithAction("Logout", α, &γ); err != nil {
		return nil, err
	}
	return γ.LogoutResponse, nil
}

// ModifyDocument was auto-generated from WSDL.
func (p *SdmSoap) ModifyDocument(ModifyDocument *ModifyDocument) (*ModifyDocumentResponse, error) {
	α := struct {
		OperationModifyDocumentRequest `xml:"impl:modifyDocument"`
	}{
		OperationModifyDocumentRequest{
			ModifyDocument,
		},
	}

	γ := struct {
		OperationModifyDocumentResponse `xml:"modifyDocumentResponse"`
	}{}
	if err := p.RoundTripWithAction("ModifyDocument", α, &γ); err != nil {
		return nil, err
	}
	return γ.ModifyDocumentResponse, nil
}

// NotifyContacts was auto-generated from WSDL.
func (p *SdmSoap) NotifyContacts(NotifyContacts *NotifyContacts) (*NotifyContactsResponse, error) {
	α := struct {
		OperationNotifyContactsRequest `xml:"impl:notifyContacts"`
	}{
		OperationNotifyContactsRequest{
			NotifyContacts,
		},
	}

	γ := struct {
		OperationNotifyContactsResponse `xml:"notifyContactsResponse"`
	}{}
	if err := p.RoundTripWithAction("NotifyContacts", α, &γ); err != nil {
		return nil, err
	}
	return γ.NotifyContactsResponse, nil
}

// RateDocument was auto-generated from WSDL.
func (p *SdmSoap) RateDocument(RateDocument *RateDocument) (*RateDocumentResponse, error) {
	α := struct {
		OperationRateDocumentRequest `xml:"impl:rateDocument"`
	}{
		OperationRateDocumentRequest{
			RateDocument,
		},
	}

	γ := struct {
		OperationRateDocumentResponse `xml:"rateDocumentResponse"`
	}{}
	if err := p.RoundTripWithAction("RateDocument", α, &γ); err != nil {
		return nil, err
	}
	return γ.RateDocumentResponse, nil
}

// RemoveAttachment was auto-generated from WSDL.
func (p *SdmSoap) RemoveAttachment(RemoveAttachment *RemoveAttachment) (*RemoveAttachmentResponse, error) {
	α := struct {
		OperationRemoveAttachmentRequest `xml:"impl:removeAttachment"`
	}{
		OperationRemoveAttachmentRequest{
			RemoveAttachment,
		},
	}

	γ := struct {
		OperationRemoveAttachmentResponse `xml:"removeAttachmentResponse"`
	}{}
	if err := p.RoundTripWithAction("RemoveAttachment", α, &γ); err != nil {
		return nil, err
	}
	return γ.RemoveAttachmentResponse, nil
}

// RemoveLrelRelationships was auto-generated from WSDL.
func (p *SdmSoap) RemoveLrelRelationships(RemoveLrelRelationships *RemoveLrelRelationships) (*RemoveLrelRelationshipsResponse, error) {
	α := struct {
		OperationRemoveLrelRelationshipsRequest `xml:"impl:removeLrelRelationships"`
	}{
		OperationRemoveLrelRelationshipsRequest{
			RemoveLrelRelationships,
		},
	}

	γ := struct {
		OperationRemoveLrelRelationshipsResponse `xml:"removeLrelRelationshipsResponse"`
	}{}
	if err := p.RoundTripWithAction("RemoveLrelRelationships", α, &γ); err != nil {
		return nil, err
	}
	return γ.RemoveLrelRelationshipsResponse, nil
}

// RemoveMemberFromGroup was auto-generated from WSDL.
func (p *SdmSoap) RemoveMemberFromGroup(RemoveMemberFromGroup *RemoveMemberFromGroup) (*RemoveMemberFromGroupResponse, error) {
	α := struct {
		OperationRemoveMemberFromGroupRequest `xml:"impl:removeMemberFromGroup"`
	}{
		OperationRemoveMemberFromGroupRequest{
			RemoveMemberFromGroup,
		},
	}

	γ := struct {
		OperationRemoveMemberFromGroupResponse `xml:"removeMemberFromGroupResponse"`
	}{}
	if err := p.RoundTripWithAction("RemoveMemberFromGroup", α, &γ); err != nil {
		return nil, err
	}
	return γ.RemoveMemberFromGroupResponse, nil
}

// Search was auto-generated from WSDL.
func (p *SdmSoap) Search(Search *Search) (*SearchResponse, error) {
	α := struct {
		OperationSearchRequest `xml:"impl:search"`
	}{
		OperationSearchRequest{
			Search,
		},
	}

	γ := struct {
		OperationSearchResponse `xml:"searchResponse"`
	}{}
	if err := p.RoundTripWithAction("Search", α, &γ); err != nil {
		return nil, err
	}
	return γ.SearchResponse, nil
}

// ServerStatus was auto-generated from WSDL.
func (p *SdmSoap) ServerStatus(ServerStatus *ServerStatus) (*ServerStatusResponse, error) {
	α := struct {
		OperationServerStatusRequest `xml:"impl:serverStatus"`
	}{
		OperationServerStatusRequest{
			ServerStatus,
		},
	}

	γ := struct {
		OperationServerStatusResponse `xml:"serverStatusResponse"`
	}{}
	if err := p.RoundTripWithAction("ServerStatus", α, &γ); err != nil {
		return nil, err
	}
	return γ.ServerStatusResponse, nil
}

// Transfer was auto-generated from WSDL.
func (p *SdmSoap) Transfer(Transfer *Transfer) (*TransferResponse, error) {
	α := struct {
		OperationTransferRequest `xml:"impl:transfer"`
	}{
		OperationTransferRequest{
			Transfer,
		},
	}

	γ := struct {
		OperationTransferResponse `xml:"transferResponse"`
	}{}
	if err := p.RoundTripWithAction("Transfer", α, &γ); err != nil {
		return nil, err
	}
	return γ.TransferResponse, nil
}

// UpdateObject was auto-generated from WSDL.
func (p *SdmSoap) UpdateObject(UpdateObject *UpdateObject) (*UpdateObjectResponse, error) {
	α := struct {
		OperationUpdateObjectRequest `xml:"impl:updateObject"`
	}{
		OperationUpdateObjectRequest{
			UpdateObject,
		},
	}

	γ := struct {
		OperationUpdateObjectResponse `xml:"updateObjectResponse"`
	}{}
	if err := p.RoundTripWithAction("UpdateObject", α, &γ); err != nil {
		return nil, err
	}
	return γ.UpdateObjectResponse, nil
}

// UpdateRating was auto-generated from WSDL.
func (p *SdmSoap) UpdateRating(UpdateRating *UpdateRating) (*UpdateRatingResponse, error) {
	α := struct {
		OperationUpdateRatingRequest `xml:"impl:updateRating"`
	}{
		OperationUpdateRatingRequest{
			UpdateRating,
		},
	}

	γ := struct {
		OperationUpdateRatingResponse `xml:"updateRatingResponse"`
	}{}
	if err := p.RoundTripWithAction("UpdateRating", α, &γ); err != nil {
		return nil, err
	}
	return γ.UpdateRatingResponse, nil
}
