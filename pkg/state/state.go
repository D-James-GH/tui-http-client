package state

import "github.com/d-james-gh/tui-http-client/pkg/request"

type State struct {
	Url    string
	Method string
	Result string
}

func (s *State) SendRequest() {
	if res, err := request.SendRequest(s.Method, s.Url); err == nil {
		s.Result = res
	}
}
