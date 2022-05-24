package configs

type commonConfigurations struct {
	Service  serviceConfigurations
	Database databaseConfigurations
}

type serviceConfigurations struct {
	Name   string
	Detail string
	Port   string
}

type databaseConfigurations struct {
	Name string
}

func loadCommonConfigurations() commonConfigurations {
	var cc commonConfigurations

	cc.Service.Name = "matar"
	cc.Service.Detail = "user related"
	cc.Service.Port = "9000"

	cc.Database.Name = "user"
	return cc
}
