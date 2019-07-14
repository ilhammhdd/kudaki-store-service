package externals

type RetrieveItems struct{}

func (ri RetrieveItems) Work() interface{} {
	// usecase := usecases.RetrieveItems{DBO: kudakimysql.NewDBOperation()}

	// ede := EventDrivenExternal{
	// 	eventDrivenAdapter: new(adapters.RetrieveItems),
	// 	eventDrivenUsecase: usecase,
	// 	eventName:          events.StoreTopic_RETRIEVE_ITEMS_REQUESTED.String(),
	// 	inTopics:           []string{events.StoreTopic_RETRIEVE_ITEMS_REQUESTED.String()},
	// 	outTopic:           events.StoreTopic_ITEMS_RETRIEVED.String()}

	// ede.handle()
	return nil
}
