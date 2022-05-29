package system

type Service interface {
	GetName() string
	Start(string) error
}

var handlers = make(map[string]Service)

func SetService(service Service) {
	handlers[service.GetName()] = service
}

func GetService(name string) Service {
	return handlers[name]
}

func AllService() (services []string) {
	for srv := range handlers {
		services = append(services, srv)
	}
	return services
}
