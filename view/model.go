package view

type ViewModel struct {
	Children     map[string][]*ViewModel
	Vars         map[string]interface{}
	TemplateName string
}

type Vars map[string]interface{}

func NewViewModel(templateName string, vars map[string]interface{}) *ViewModel {
	if vars == nil {
		vars = make(map[string]interface{})
	}
	return &ViewModel{
		Children:     make(map[string][]*ViewModel),
		Vars:         vars,
		TemplateName: templateName,
	}
}

func (model *ViewModel) AddChild(child *ViewModel, captureTo string) *ViewModel {
	children, ok := model.Children[captureTo]
	if !ok {
		children = make([]*ViewModel, 0, 10)
	}

	model.Children[captureTo] = append(children, child)

	return model
}
