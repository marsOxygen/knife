package util

type Config struct {
	Basic  BasicConfig
	GitEnv []GitEnv
}
type BasicConfig struct {
	CodeDir string
}
type GitEnv struct {
	Pattern string
	Email   string
	Name    string
}

type RepoCache struct {
	Repos []RepoStruct
}
type RepoStruct struct {
	RepoPath  string
	LocalPath string
}
