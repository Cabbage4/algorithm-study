package mvcc

import "sync"

type Record struct {
	nextIdLock sync.Mutex
	nextId     int

	nextTranslationIdLock sync.Mutex
	nextTranslationId     int

	dataMpLock sync.RWMutex
	dataMp     map[string]*Row
}

func (r *Record) getNextId() int {
	r.nextIdLock.Lock()
	defer r.nextIdLock.Unlock()

	ans := r.nextId
	r.nextId++
	return ans
}

func (r *Record) GetNextTranslationId() int {
	r.nextTranslationIdLock.Lock()
	defer r.nextTranslationIdLock.Unlock()

	ans := r.nextTranslationId
	r.nextTranslationId++
	return ans
}

func (r *Record) Add(translationId int, key string, data interface{}) {
	id := r.getNextId()

	row := &Row{
		mutex:         sync.RWMutex{},
		id:            id,
		translationId: translationId,
		key:           key,
		data:          data,
	}

	r.dataMpLock.Lock()
	defer r.dataMpLock.Unlock()
	r.dataMp[key] = row
}

func (r *Record) Update(translationId int, key string, data interface{}) {
	row, ok := r.dataMp[key]
	if !ok {
		return
	}

	row.mutex.Lock()
	defer row.mutex.Unlock()

	row.undoLog = &UndoLog{
		Row: &Row{
			mutex:         sync.RWMutex{},
			id:            row.id,
			translationId: row.translationId,
			undoLog:       row.undoLog,
			key:           row.key,
			data:          row.data,
		},
		nextUndoLog: row.undoLog,
	}
	row.data = data
	row.translationId = translationId
}

func (r *Record) Get(translationId int, keys ...string) []interface{} {
	if len(keys) == 0 {
		return nil
	}

	var ans []interface{}
	for _, key := range keys {
		r.dataMpLock.RLock()
		row, ok := r.dataMp[key]
		r.dataMpLock.RUnlock()

		if !ok {
			continue
		}

		if row.translationId <= translationId {
			if !row.isDelete {
				ans = append(ans, row.data)
			}

			continue
		}

		tmp := row.undoLog
		for tmp != nil && tmp.translationId > translationId {
			tmp = tmp.nextUndoLog
		}

		if tmp == nil {
			continue
		}

		ans = append(ans, tmp.data)
	}

	return ans
}

func (r *Record) Delete(translationId int, keys ...string) {
	if len(keys) == 0 {
		return
	}

	for _, key := range keys {
		r.dataMpLock.RLock()
		row, ok := r.dataMp[key]
		r.dataMpLock.RUnlock()

		if !ok {
			continue
		}

		if row.translationId <= translationId && !row.isDelete {
			row.undoLog = &UndoLog{
				Row: &Row{
					mutex:         sync.RWMutex{},
					id:            row.id,
					translationId: row.translationId,
					isDelete:      false,
					undoLog:       nil,
					key:           "",
					data:          nil,
				},
				nextUndoLog: nil,
			}
			row.isDelete = true
			row.translationId = translationId
		}
	}
}

type Row struct {
	mutex sync.RWMutex

	id            int
	translationId int

	isDelete bool

	undoLog *UndoLog

	key  string
	data interface{}
}

type UndoLog struct {
	*Row
	nextUndoLog *UndoLog
}
