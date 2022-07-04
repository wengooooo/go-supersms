package go_supersms

type NumberDetail struct {
	Pid    string `json:"pid"`
	Phone  string `json:"phone"`
	Taskid int    `json:"taskid"`
	Cost   int    `json:"cost"`
}

type CodeDetail struct {
	Code   string `json:"code"`
	Phone  string `json:"phone"`
	Taskid int    `json:"taskid"`
}

type ReleaseDetail struct {
	Message string `json:"code"`
}
