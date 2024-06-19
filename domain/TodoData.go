package domain

import "errors"

type TodoStatus int

const (
	Planing TodoStatus = iota
	Doing
	Done
	Droped
)

type TodoElement struct {
	ID     int
	Title  string
	Status TodoStatus
}

type TodoData []TodoElement

// Factory
func GetNewTodoData() TodoData {
	return make(TodoData, 0)
}

func GetNewTodoElement(id int, title string, status TodoStatus) TodoElement {
	return TodoElement{ID: id, Title: title, Status: status}
}

func (data *TodoData) GetElementById(id int) (int, TodoElement) {
	for index, element := range *data {
		if element.ID == id {
			return index, element
		}
	}

	return -1, TodoElement{}
}

//
func (data *TodoData) AddTodoElement(title string, status TodoStatus) (TodoElement, error) {
	if len(title) < 3 {
		return TodoElement{}, errors.New("lengths need to be > 3")
	}
	el := GetNewTodoElement(len(*data), title, status)
	*data = append(*data, el)
	return el, nil
}

func (data *TodoData) DeleteTodoElement(id int) (TodoElement, error) {
	index, element := data.GetElementById(id)
	if index == -1 {
		return TodoElement{}, errors.New("nothing to delete")
	}

	*data = append((*data)[:index], (*data)[index+1:]...)
	return element, nil
}

func (element TodoElement) GetColor() string {
	if element.Status == Planing {
		return "orange"
	} else if element.Status == Doing {
		return "yellow"
	} else if element.Status == Done {
		return "green"
	} else if element.Status == Droped {
		return "black"
	} else {
		return ""
	}
}
