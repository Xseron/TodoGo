package domain

type UsersDataType map[string]*TodoData

func GetNewUserData() UsersDataType {
	return make(UsersDataType)
}

func (usersData UsersDataType) AddUser(addr string) *TodoData {
	newTodoData := GetNewTodoData()
	usersData[addr] = &newTodoData
	return &newTodoData
}

func (usersData UsersDataType) GetUser(addr string) *TodoData {
	data, ok := usersData[addr]
	if !ok {
		usersData.GetUser(addr)
		data = usersData[addr]
	}
	return data
}
