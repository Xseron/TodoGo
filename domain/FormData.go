package domain

type FormData struct {
	Values map[string]string
	Errors map[string]string
}

// Factory
func NewFormData() *FormData {
	fd := &FormData{
		Values: make(map[string]string),
		Errors: make(map[string]string),
	}
	return fd
}

// Interactor
func (formData *FormData) hasErr(name string) (string, bool) {
	errorInfo, ok := formData.Errors[name]
	return errorInfo, ok
}

func (formData *FormData) hasVal(name string) (string, bool) {
	errorInfo, ok := formData.Values[name]
	return errorInfo, ok
}
