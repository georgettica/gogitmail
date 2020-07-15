package interfaces

// GitRemoteUser holds the gitlab and the github
type GitRemoteUser interface {
	GetID() string
}
