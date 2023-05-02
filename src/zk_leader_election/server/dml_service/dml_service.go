package dml_service

import "time"

const (
	StandbyService = 0
	InorderService = 1
)

func newDMLService(timeout time.Duration) DMLService {
	dao := newDMLServiceDAO()
	log := newDMLServiceLog(dao)
	context := newDMLServiceContext(log, timeout)
	return context
}

func newStandByService() DMLService {
	return newDMLServiceStandBy()
}

func NewReplaceableDMLService() ReplaceableDMLService {
	// todo: config file as param
	standby := newDMLServiceStandBy()
	inorder := newDMLService(time.Second * 3)

	srvMap := make(map[interface{}]DMLService)
	srvMap[InorderService] = inorder
	srvMap[StandbyService] = standby

	replaceable := newDMLServiceReplaceable()
	replaceable.RegisterServices(srvMap)
	if err := replaceable.ActivateService(StandbyService); err != nil {
		panic(err) // never reach
	}
	return replaceable
}
