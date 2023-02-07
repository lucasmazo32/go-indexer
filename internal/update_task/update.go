package update_task

import (
	"log"
	"sync"

	"indexer.com/indexer/internal/map_from_file"
	"indexer.com/indexer/internal/zinc_send"
)

type UpdateTask struct {
	executeFunc func() error
	wg          *sync.WaitGroup
}

func NewUpdateTask(name string, wg *sync.WaitGroup) *UpdateTask {
	return &UpdateTask{
		executeFunc: func() error {
			fileMap := map_from_file.MapFromFile(name)
			e := zinc_send.SendToZinc(*fileMap)
			return e
		},
		wg: wg,
	}
}

func (u *UpdateTask) Execute() error {
	err := u.executeFunc()
	if u.wg != nil {
		defer u.wg.Done()
	}
	return err
}

func (u *UpdateTask) OnFailure(e error) {
	log.Println(e)
}
