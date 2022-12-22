package comman

type Config struct {
	App Application `json:"app"`
	DB  DataBase    `json:"db"`
}

type DataBase struct {
	IP       string `json:"ip"`
	PORT     string `json:"port"`
	UserName string `json:"username"`
	PASSWORD string `json:"password"`
	DATABASE string `json:"database"`
}

type Application struct {
	MYIP   string `json:"myip"`
	MYPORT int    `json:"myport"`
}
