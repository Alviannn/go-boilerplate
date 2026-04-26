package services

import "github.com/samber/do/v2"

type helper struct {
	scope *do.Scope
}

func NewHelper(i do.Injector) (Helper, error) {
	svc, err := do.InvokeStruct[*helper](i)
	if err != nil {
		return nil, err
	}

	svc.scope = i.Scope("helper")
	return svc, nil
}

func (s *helper) GetBase() *Base {
	return do.MustInvoke[*Base](s.scope)
}
