package fakes

import (
	"time"

	"github.com/pivotal-pez/cfmgo"
	"github.com/xchapter7x/coda"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//FakeTaskManager - this is a fake representation of the task manager
type FakeTaskManager struct {
	ExpireEmitter      chan int64
	ReturnedTask       coda.Task
	ReturnedErr        error
	DefaultTaskExpires int64
}

//ScheduleTask ---
func (s *FakeTaskManager) ScheduleTask(t coda.Task, expireTime time.Time) coda.Task {
	return t
}

//SaveTask --
func (s *FakeTaskManager) SaveTask(t coda.Task) (coda.Task, error) {
	if s.ExpireEmitter != nil {
		s.ExpireEmitter <- t.Expires
	}
	return t, s.ReturnedErr
}

//FindAndStallTaskForCaller --
func (s *FakeTaskManager) FindAndStallTaskForCaller(callerName string) (t coda.Task, err error) {
	return s.ReturnedTask, s.ReturnedErr
}

//FindTask --
func (s *FakeTaskManager) FindTask(id string) (t coda.Task, err error) {
	return s.ReturnedTask, s.ReturnedErr
}

//NewTask --
func (s *FakeTaskManager) NewTask(callerName string, profile coda.ProfileType, status string) (t coda.Task) {
	t = coda.NewTask()
	t.CallerName = callerName
	t.Profile = profile
	t.Status = status
	t.ID = bson.NewObjectId()
	t.Timestamp = time.Now()
	t.MetaData = make(map[string]interface{})
	t.Expires = s.DefaultTaskExpires
	return
}

//FakeTask -
type FakeTask struct {
	ID              bson.ObjectId          `bson:"_id"`
	Timestamp       time.Time              `bson:"timestamp"`
	Status          string                 `bson:"status"`
	MetaData        map[string]interface{} `bson:"metadata"`
	PrivateMetaData map[string]interface{} `bson:"private_metadata"`
}

//NewFakeCollection ====
func NewFakeCollection(updated int) *FakeCollection {
	fakeCol := new(FakeCollection)

	if updated == -1 {
		fakeCol.FakeChangeInfo = nil
	} else {
		fakeCol.FakeChangeInfo = &mgo.ChangeInfo{
			Updated: updated,
		}
	}
	return fakeCol
}

//FakeCollection -
type FakeCollection struct {
	cfmgo.Collection
	ControlTask             coda.Task
	ErrControl              error
	FakeChangeInfo          *mgo.ChangeInfo
	ErrFindAndModify        error
	FakeResultFindAndModify interface{}
	AssignResult            func(result, setValue interface{})
}

//Close -
func (s *FakeCollection) Close() {

}

//Wake -
func (s *FakeCollection) Wake() {

}

//FindAndModify -
func (s *FakeCollection) FindAndModify(selector interface{}, update interface{}, result interface{}) (info *mgo.ChangeInfo, err error) {
	if s.AssignResult != nil {
		s.AssignResult(result, s.FakeResultFindAndModify)
	}
	return s.FakeChangeInfo, s.ErrFindAndModify
}

//UpsertID -
func (s *FakeCollection) UpsertID(id interface{}, result interface{}) (changInfo *mgo.ChangeInfo, err error) {
	return
}

//FindOne -
func (s *FakeCollection) FindOne(id string, result interface{}) (err error) {
	err = s.ErrControl
	result = s.ControlTask
	return
}
