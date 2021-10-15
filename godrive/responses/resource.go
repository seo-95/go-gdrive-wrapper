package drive

//? Future usage

type FileList struct {
	Kind             string
	NextPageToken    string
	IncompleteSearch string
	Files            []File
}

type AboutResource struct {
	Kind                string
	User                User
	StorageQuota        Storage
	ImportFormats       map[string][]string
	ExportFormats       map[string][]string
	MaxImportSizes      map[string][]uint64
	MaxUploadSize       uint64
	AppInstalled        bool
	FolderColorPalette  []string
	TeamDriveThemes     []DriveTheme
	canCreateTeamDrives bool
	canCreateDrives     bool
}

type File struct {
	Kind                         string
	Id                           string
	Name                         string
	MimeType                     string
	Description                  string
	Starred                      bool
	Trashed                      bool
	ExplicitlyTrashed            bool
	TrashingUser                 User
	TrashedTime                  string //datetime,
	Parents                      []string
	Properties                   map[string][]string
	AppProperties                map[string][]string
	Spaces                       []string
	Version                      uint32 //long
	WebContentLink               string
	WebViewLink                  string
	IconLink                     string
	HasThumbnail                 bool
	ThumbnailLink                string
	ThumbnailVersion             uint64
	ViewedByMe                   bool
	ViewedByMeTime               string // datetime,
	CreatedTime                  string // datetime,
	ModifiedTime                 string // datetime,
	ModifiedByMeTime             string // datetime,
	ModifiedByMe                 bool
	SharedWithMeTime             string // datetime,
	SharingUser                  User
	Owners                       []User
	TeamDriveId                  string
	DriveId                      string
	lastModifyingUser            User
	Shared                       bool
	OwnedByMe                    bool
	Capabilities                 Capabilities
	ViewersCanCopyContent        bool
	CopyRequiresWriterPermission bool
	WritersCanShare              bool
	Permissions                  []Permission
	PermissionIds                []string
	HasAugmentedPermissions      bool
	FolderColorRgb               string
	OriginalFilename             string
	FullFileExtension            string
	FileExtension                string
	Md5Checksum                  string
	Size                         uint32 // long
	QuotaBytesUsed               uint32 // long
	HeadRevisionId               string
	ContentHints                 Hints
	ImageMediaMetadata           ImageMetadata
	Time                         string
	VideoMediaMetadata           VideoMetadata
	IsAppAuthorized              bool
	ExportLinks                  map[string]string
	ShortcutDetails              Shortcut
	ContentRestrictions          []ContentRestriction
	ResourceKey                  string
	LinkShareMetadata            LinkShareMetadata
}

type User struct {
	Kind         string
	DisplayName  string
	PhotoLink    string
	Me           bool
	PermissionId string
	EmailAddress string
}

type Storage struct {
	Limit             uint64
	Usage             uint64
	UsageInDrive      uint64
	UsageInDriveTrash uint64
}

type DriveTheme struct {
	Id                  string
	BackgroundImageLink string
	colorRgb            string
}

type Capabilities struct {
	CanAddChildren                        bool
	CanAddFolderFromAnotherDrive          bool
	CanAddMyDriveParent                   bool
	CanChangeCopyRequiresWriterPermission bool
	CanChangeSecurityUpdateEnabled        bool
	CanChangeViewersCanCopyContent        bool
	CanComment                            bool
	CanCopy                               bool
	CanDelete                             bool
	CanDeleteChildren                     bool
	CanDownload                           bool
	CanEdit                               bool
	CanListChildren                       bool
	CanModifyContent                      bool
	CanModifyContentRestriction           bool
	CanMoveChildrenOutOfTeamDrive         bool
	CanMoveChildrenOutOfDrive             bool
	CanMoveChildrenWithinTeamDrive        bool
	CanMoveChildrenWithinDrive            bool
	CanMoveItemIntoTeamDrive              bool
	CanMoveItemOutOfTeamDrive             bool
	CanMoveItemOutOfDrive                 bool
	CanMoveItemWithinTeamDrive            bool
	CanMoveItemWithinDrive                bool
	CanMoveTeamDriveItem                  bool
	CanReadRevisions                      bool
	CanReadTeamDrive                      bool
	CanReadDrive                          bool
	CanRemoveChildren                     bool
	CanRemoveMyDriveParent                bool
	CanRename                             bool
	CanShare                              bool
	CanTrash                              bool
	CanTrashChildren                      bool
	CanUntrash                            bool
}

type Permission struct {
	Kind                       string
	Id                         string
	Type                       string
	EmailAddress               string
	Domain                     string
	Role                       string
	View                       string
	AllowFileDiscovery         bool
	DisplayName                string
	PhotoLink                  string
	ExpirationTime             string // datetime,
	TeamDrivePermissionDetails []TeamDrivePermissionDetail
	PermissionDetails          []PermissionDetail
	Deleted                    bool
}

type Hints struct {
	Thumbnail     Thumbnail
	IndexableText string
}

type ImageMetadata struct {
	Width            uint16
	Height           uint16
	Rotation         uint16
	Location         Location
	CameraMake       string
	CameraModel      string
	ExposureTime     float32
	Aperture         float32
	FlashUsed        bool
	FocalLength      float32
	IsoSpeed         uint16
	MeteringMode     string
	Sensor           string
	ExposureMode     string
	ColorSpace       string
	WhiteBalance     string
	ExposureBias     float32
	MaxApertureValue float32
	SubjectDistance  uint16
	Lens             string
}

type Location struct {
	Latitude  float64
	Longitude float64
	Altitude  float64
}

type VideoMetadata struct {
	Width          uint16
	Height         uint16
	DurationMillis uint32 //long
}

type Shortcut struct {
	TargetId          string
	TargetMimeType    string
	TargetResourceKey string
}

type ContentRestriction struct {
	ReadOnly        bool
	Reason          string
	RestrictingUser User
	RestrictionTime string // datetime
	Type            string
}

type LinkShareMetadata struct {
	SecurityUpdateEligible bool
	SecurityUpdateEnabled  bool
}

type TeamDrivePermissionDetail struct {
	TeamDrivePermissionType string
	Role                    string
	InheritedFrom           string
	Inherited               bool
}

type PermissionDetail struct {
	PermissionType string
	Role           string
	InheritedFrom  string
	Inherited      bool
}

type Thumbnail struct {
	Image    byte
	MimeType string
}
