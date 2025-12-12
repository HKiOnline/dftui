package dfdb

const (
	// FileSystemProvider is the provider type for filesystem-based storage
	FileSystemProvider = "filesystem"
)

// New creates a new Provider based on the configuration.
// If the provider type is unknown, it defaults to filesystem provider.
func New(config ProviderConfiguration) (Provider, error) {
	switch config.Provider {
	case FileSystemProvider:
		return NewFsProvider(config.Filesystem.Directory)
	default:
		return NewFsProvider(config.Filesystem.Directory)
	}
}
