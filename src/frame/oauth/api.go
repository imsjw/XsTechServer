package oauth

func Init() {
	initConfig()
	if !configEnable {
		return
	}
	initRouter()
}
