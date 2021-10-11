package drive

//? Future usage

type Resource struct {
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

type User struct {
	Kind         string
	DisplayName  string
	PhotoLink    string
	Me           string
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
