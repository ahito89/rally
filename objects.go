package rally

import "time"

// PersistableObject definition 
type PersistableObject struct {
	CreationDate *time.Time
	ObjectID     int64
	ObjectUUID   string
	VersionID    string
}

// DomainObject definition
type DomainObject struct {
    PersistableObject
    
	Subscription
	RefURL string `json:"_ref"`
}

// WorkspaceDomainObject definition
type WorkspaceDomainObject struct {
	DomainObject
    
	Workspace *Workspace
}

// Workspace definition
type Workspace struct {
    DomainObject
    
	Children    []Project
	Description string
	Name        string
	Notes       string
	Owner       User
	Projects    []Project	
	SchemaVersion string
	State         string
	Style         string	
}

// User definition
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
	Role                   string
	SubscriptionAdmin      bool
	SubscriptionID         int
	SubscriptionPermission string
	WorkspacePermission string
}

// Project definition 
type Project struct {
    DomainObject
    
    Description string
    Iterations []Iteration
    Name string
    Notes string
    Owner User    
    SchemaVersion string
    State string
}

// Subscription definition
type Subscription struct {
	APIKeysEnabled          bool
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

// Ref definition
type Ref struct {
	Count         int
	Ref           string `json:"_ref"`
	Type          string `json:"_type"`
	RefObjectName string `json:"_refObjectName"`
	RefObjectUUID string `json:"_refObjectUUID"`
}

// HierarchicalRequirement definition
type HierarchicalRequirement struct {
	Requirement
    
	AcceptedDate  *time.Time
	Blocked       bool
	BlockedReason string
	Blocker       *Blocker
	CreationDate  *time.Time
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

// TestCase definition
type TestCase struct {
    Artifact
    
    DefectStatus string
    DragAndDropRank string
    LastBuild string
    LastRun string
    LastVerdict string
    Method string
    Objective string
    Package string
    PostConditions string
    PreConditions string
    Priority string
    Recycled bool
    Risk string
    Type string
    ValidationExpectedResult string
    ValidationInput string
}

// Task definition
type Task struct {
    Artifact
    
    Actuals float64
    Blocked bool
    BlockedReason string
    DragAndDropRank string
    Estimate float64
    Recycled bool
    State string
    TaskIndex int
    TimeSpent float64
    ToDo float64
}

// Release definition
type Release struct {
}

// Iteration definition
type Iteration struct {
	WorkspaceDomainObject
    
    EndDate *time.Time
    Name string
    Notes string
    PlanEstimate float64
    PlannedVelocity float64
    StartDate *time.Time
    State string
    TaskActualTotal float64
    TaskEstimateTotal float64
    TaskRemainingTotal float64
    Theme string
}

// Artifact definition
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

// PortfolioItem definition
type PortfolioItem struct {
	Artifact
}

// PortfolioItemFeature definition
type PortfolioItemFeature struct {
	PortfolioItem
    
	UserStories *Ref
}

// Defect definition
type Defect struct {
	SchedulableArtifact
}

// Blocker definition
type Blocker struct {
}

// Requirement definition
type Requirement struct {
	SchedulableArtifact
    
	Attachments *Ref
}

// SchedulableArtifact definition
type SchedulableArtifact struct {
	Artifact
}

// Attachment definition
type Attachment struct {
    WorkspaceDomainObject
    
    ContentType string
    Description string
    Name string
    Size int
}

// TestCaseResult definition
type TestCaseResult struct {
    WorkspaceDomainObject
    
    Build string
    Date *time.Time
    Duration float64
    Notes string
    Verdict string
}

// TestCaseStep definition
type TestCaseStep struct {
    WorkspaceDomainObject
    
    ExpectedResult string
    Input string
    StepIndex int
}