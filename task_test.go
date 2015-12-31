package coda_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/xchapter7x/coda"
)

var _ = Describe("Task", func() {
	Describe(".SetMeta()", func() {
		Context("when called on a task", func() {
			var (
				controlPublicMetaKey = "hithere"
				controlCaller        = "caller_nnnn"
				controlStatus        = "somestatus"
				task                 = NewTaskManager(nil).NewTask(controlCaller, TaskLongPollQueue, controlStatus)
			)
			BeforeEach(func() {
				task = task.SetMeta(controlPublicMetaKey, controlStatus)
			})
			It("then it should return a task with the added meta data", func() {
				Ω(task.MetaData).Should(Equal(map[string]interface{}{
					controlPublicMetaKey: controlStatus,
				}))
			})
		})
	})
})
