package globalvar

var ENV string = "dev"

func SetEnv(env string) {
	if env == "" {
		ENV = "dev"
	} else {
		ENV = env
	}
}

func GetEnv() string {
	return ENV
}

func IsProd() bool {
	return ENV == "prod"
}

func GetHost() string {
	if IsProd() {
		return "https://three-cats.top"
	}
	return "http://localhost:8080"
}
