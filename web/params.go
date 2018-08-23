package web

const (
	PAGE_DEFAULT          = 0
	PER_PAGE_DEFAULT      = 60
	PER_PAGE_MAXIMUM      = 200
	LOGS_PER_PAGE_DEFAULT = 10000
	LOGS_PER_PAGE_MAXIMUM = 10000
)


type Params struct {
	UserId         string
	TokenId        string
	PostId         string
	Username       string
	Category       string
	Page           int
	PerPage        int
	LogsPerPage    int
}