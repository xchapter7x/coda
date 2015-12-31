package coda

func NewTask() Task {
	return Task{}
}

//SetPublicMeta - set a public metadata record
func (s Task) SetMeta(name string, value interface{}) Task {
	if s.MetaData == nil {
		s.MetaData = make(map[string]interface{})
	}
	s.MetaData[name] = value
	return s
}

//GetPublicMeta - get a public metadata record
func (s Task) GetMeta(name string) interface{} {
	return s.MetaData[name]
}

// Save -- Safe way to update a task
func (s Task) Save(ts TaskSaver) {
	ts.SaveTask(s)
}
