package coda

//NewTask - a helper to create a new task
func NewTask() Task {
	return Task{
		MetaData: make(map[string]interface{}),
	}
}

//SetMeta - set a public metadata record
func (s Task) SetMeta(name string, value interface{}) Task {
	s.MetaData[name] = value
	return s
}

//GetMeta - get a public metadata record
func (s Task) GetMeta(name string) interface{} {
	return s.MetaData[name]
}

// Save -- Safe way to update a task
func (s Task) Save(ts TaskSaver) {
	ts.SaveTask(s)
}
