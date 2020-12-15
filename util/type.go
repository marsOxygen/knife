package util

type Config struct {
	Basic    BasicConfig
	Match    []MatchConfig
	Lollipop map[string]string
}
type BasicConfig struct {
	CodeDir string
}
type MatchConfig struct {
	Pattern string
	Data    map[string]string
	Hook    map[string][]string
}

type RepoCache struct {
	Repos []RepoStruct
}
type RepoStruct struct {
	RepoPath  string
	LocalPath string
}
