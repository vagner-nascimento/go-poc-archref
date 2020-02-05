package dataamqp

type customerQueueInfo struct {
	Queue   queueInfo
	Message messageInfo
}

func getCustomerInfo() customerQueueInfo {
	q := queueInfo{
		QName:         "q-customer",
		QDurable:      false,
		QDeleteUnused: false,
		QExclusive:    false,
		QNoWait:       false,
	}

	m := messageInfo{
		MConsumer:  "go-poc-archref",
		MAutoAct:   true,
		MExclusive: false,
		MNoLocal:   false,
		MNoWait:    false,
	}

	info := customerQueueInfo{
		Queue:   q,
		Message: m,
	}

	return info
}
