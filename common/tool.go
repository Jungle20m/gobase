package common

type IAppContext interface {
}

type appContext struct {
}

func NewAppContext() *appContext {
	return &appContext{}
}
