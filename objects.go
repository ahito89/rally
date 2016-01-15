package rally

import "time"

// https://github.com/RallyTools/RallyRestToolkitForPython/blob/master/pyral/entity.py

type PersistableObject struct {
	CreationDate *time.Time
	ObjectID     string
	ObjectUUID   string
	VersionId    string
}

type DomainObject struct {
	Subscription

	RefURL string `json:"_ref"`
}

type WorkspaceDomainObject struct {
	DomainObject
	Workspace *Workspace
}

type Workspace struct {
	Children    []Project
	Description string
	Name        string
	Notes       string
	Owner       User
	Projects    []Project
	//RevisionHistory
	SchemaVersion string
	State         string
	Style         string
	// and more...
}

type User struct {
	*DomainObject
	CostCenter string
	Department string
	Disabled   bool

	// Name and e-mail
	DisplayName      string
	FirstName        string
	MiddleName       string
	LastName         string
	ShortDisplayName string
	UserName         string
	EmailAddress     string

	LandingPage            string
	LastLoginDate          *time.Time
	LastPasswordUpdateDate *time.Time
	NetworkID              string
	OfficeLocation         string
	OnpremLdapUsername     string
	Phone                  string
	Planner                bool
	//RevisionHistory
	Role                   string
	SubscriptionAdmin      bool
	SubscriptionID         int
	SubscriptionPermission string
	//TeamMemberships
	//UserPermissions
	//UserProfile
	WorkspacePermission string
}

type Project struct {
}

type Subscription struct {
	ApiKeysEnabled          bool
	EmailEnabled            bool
	ExpirationDate          *time.Time
	MaximumCustomUserFields int
	MaximumProjects         int
	Modules                 string
	Name                    string
	PasswordExpirationDays  int
	PreviousPasswordCount   int
	ProjectHierarchyEnabled bool
	// and more
}

type Ref struct {
	Count         int
	Ref           string `json:"_ref"`
	Type          string `json:"_type"`
	RefObjectName string `json:"_refObjectName"`
	RefObjectUUID string `json:"_refObjectUUID"`
}

type HierarchicalRequirement struct {
	Requirement
	AcceptedDate  *time.Time
	Blocked       bool
	BlockedReason string
	Blocker       *Blocker
	CreationDate  *time.Time
	// All the C_ fields seem to be custom fields !! We'll add support for that later.
	// C_AcceptanceCriteria string
	// C_CVSSRating         float64
	// C_CVSSVector         string
	// C_CXDLink            string
	Changesets          *Ref
	Children            *Ref
	Defects             *Ref
	DefectStatus        string // make special values ?
	DirectChildrenCount int
	DragAndDropRank     string
	Feature             *PortfolioItemFeature
	HasParent           bool
	InProgressDate      *time.Time
	Iteration           *Ref
	Parent              *Ref
	PlanEstimate        float64
	PortfolioItem       *Ref
	Predecessors        *Ref
	Recycled            bool
	Release             *Ref
	TaskActualTotal     float64
	TaskEstimateTotal   float64
	TaskRemainingTotal  float64
	Tasks               *Ref
	TaskStatus          string
	TestCases           *Ref
	TestCaseStatus      string
}

type TestCase struct {
}

type Task struct {
}

type Release struct {
}

type Iteration struct {
	WorkspaceDomainObject
}

type Artifact struct {
	WorkspaceDomainObject
	Description    string
	DisplayColor   string
	Expedite       bool
	FormattedID    string
	LastUpdateDate *time.Time
	Name           string
	Notes          string
	Ready          bool
	Tags           *Ref
}

type PortfolioItem struct {
	Artifact
}

type PortfolioItemFeature struct {
	PortfolioItem
	UserStories *Ref
}

type Defect struct {
	SchedulableArtifact
}

type Blocker struct {
}

type Requirement struct {
	SchedulableArtifact
	Attachments *Ref
}

type SchedulableArtifact struct {
	Artifact
}

type Attachment struct {
}

// class Workspace       (DomainObject): pass
// class Blocker         (DomainObject): pass
// class UserPermission  (DomainObject): pass
// class WorkspacePermission   (UserPermission): pass
// class ProjectPermission     (UserPermission): pass

// class WorkspaceDomainObject(DomainObject):
//     """
//         This is an Abstract Base class, with a convenience method (details) that
//         formats the attrbutes and corresponding values into an easily viewable
//         mulitiline string representation.
//     """
//     COMMON_ATTRIBUTES = ['_type',
//                          'oid', 'ref', 'ObjectID', '_ref',
//                          '_CreatedAt', '_hydrated',
//                          'Name', 'Subscription', 'Workspace',
//                          'FormattedID'
//                         ]
// class WorkspaceConfiguration(WorkspaceDomainObject): pass
// class Type                  (WorkspaceDomainObject): pass
// class Program               (WorkspaceDomainObject): pass
// class Project               (WorkspaceDomainObject): pass
// class Release               (WorkspaceDomainObject): pass
// class Iteration             (WorkspaceDomainObject): pass  # query capable only
// class ArtifactNotification  (WorkspaceDomainObject): pass  # query capable only
// class AttributeDefinition   (WorkspaceDomainObject): pass  # query capable only
// class TypeDefinition        (WorkspaceDomainObject): pass  # query capable only
// class Attachment            (WorkspaceDomainObject): pass
// class AttachmentContent     (WorkspaceDomainObject): pass
// class Build                 (WorkspaceDomainObject): pass  # query capable only
// class BuildDefinition       (WorkspaceDomainObject): pass  # query capable only
// class BuildMetric           (WorkspaceDomainObject): pass  # query capable only
// class BuildMetricDefinition (WorkspaceDomainObject): pass  # query capable only
// class Change                (WorkspaceDomainObject): pass
// class Changeset             (WorkspaceDomainObject): pass
// class ConversationPost      (WorkspaceDomainObject): pass  # query capable only
// class Milestone             (WorkspaceDomainObject): pass
// class Preference            (WorkspaceDomainObject): pass
// class PreliminaryEstimate   (WorkspaceDomainObject): pass
// class SCMRepository         (WorkspaceDomainObject): pass
// class State                 (WorkspaceDomainObject): pass
// class TestCaseStep          (WorkspaceDomainObject): pass
// class TestCaseResult        (WorkspaceDomainObject): pass
// class TestFolder            (WorkspaceDomainObject): pass
// class Tag                   (WorkspaceDomainObject): pass
// class TimeEntryItem         (WorkspaceDomainObject): pass
// class TimeEntryValue        (WorkspaceDomainObject): pass
// class UserIterationCapacity (WorkspaceDomainObject): pass
// class RecycleBinEntry       (WorkspaceDomainObject): pass
// class RevisionHistory       (WorkspaceDomainObject): pass

// class Revision              (WorkspaceDomainObject):
//     INFO_ATTRS = ['RevisionNumber', 'Description', 'CreationDate', 'User']

// classFor = { 'Persistable'             : Persistable,
//              'DomainObject'            : DomainObject,
//              'WorkspaceDomainObject'   : WorkspaceDomainObject,
//              'Subscription'            : Subscription,
//              'User'                    : User,
//              'UserProfile'             : UserProfile,
//              'UserPermission'          : UserPermission,
//              'Workspace'               : Workspace,
//              'WorkspaceConfiguration'  : WorkspaceConfiguration,
//              'WorkspacePermission'     : WorkspacePermission,
//              'Type'                    : Type,
//              'TypeDefinition'          : TypeDefinition,
//              'AttributeDefinition'     : AttributeDefinition,
//              'Program'                 : Program,
//              'Project'                 : Project,
//              'ProjectPermission'       : ProjectPermission,
//              'Artifact'                : Artifact,
//              'ArtifactNotification'    : ArtifactNotification,
//              'Release'                 : Release,
//              'Iteration'               : Iteration,
//              'Requirement'             : Requirement,
//              'HierarchicalRequirement' : HierarchicalRequirement,
//              'UserStory'               : UserStory,
//              'Story'                   : Story,
//              'Task'                    : Task,
//              'Tag'                     : Tag,
//              'Preference'              : Preference,
//              'SCMRepository'           : SCMRepository,
//              'RevisionHistory'         : RevisionHistory,
//              'Revision'                : Revision,
//              'Attachment'              : Attachment,
//              'AttachmentContent'       : AttachmentContent,
//              'TestCase'                : TestCase,
//              'TestCaseStep'            : TestCaseStep,
//              'TestCaseResult'          : TestCaseResult,
//              'TestSet'                 : TestSet,
//              'TestFolder'              : TestFolder,
//              'TimeEntryItem'           : TimeEntryItem,
//              'TimeEntryValue'          : TimeEntryValue,
//              'Build'                   : Build,
//              'BuildDefinition'         : BuildDefinition,
//              'BuildMetric'             : BuildMetric,
//              'BuildMetricDefinition'   : BuildMetricDefinition,
//              'Defect'                  : Defect,
//              'DefectSuite'             : DefectSuite,
//              'Change'                  : Change,
//              'Changeset'               : Changeset,
//              'PortfolioItem'           : PortfolioItem,
//              'PortfolioItem_Strategy'  : PortfolioItem_Strategy,
//              'PortfolioItem_Initiative': PortfolioItem_Initiative,
//              'PortfolioItem_Theme'     : PortfolioItem_Theme,
//              'PortfolioItem_Feature'   : PortfolioItem_Feature,
//              'State'                   : State,
//              'PreliminaryEstimate'     : PreliminaryEstimate,
//              'WebLinkDefinition'       : WebLinkDefinition,
//              'Milestone'               : Milestone,
//              'ConversationPost'        : ConversationPost,
//              'Blocker'                 : Blocker,
//              'AllowedAttributeValue'   : AllowedAttributeValue,
//              'AllowedQueryOperator'    : AllowedQueryOperator,
//              'CustomField'             : CustomField,
//              'UserIterationCapacity'   : UserIterationCapacity,
//              'CumulativeFlowData'      : CumulativeFlowData,
//              'ReleaseCumulativeFlowData'   : ReleaseCumulativeFlowData,
//              'IterationCumulativeFlowData' : IterationCumulativeFlowData,
//              'RecycleBinEntry'         : RecycleBinEntry,
//              'SearchObject'            : SearchObject,
//            }
