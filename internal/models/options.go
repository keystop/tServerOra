package models

//Options interface for program options.
type Options interface {
	ServAddr() string
	RepoFileName() string
	DBConnString() string
}
