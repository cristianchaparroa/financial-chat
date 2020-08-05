package fixtures

func GetLoginRequest() map[string]interface{} {

	return map[string]interface{}{
		"email":    "test@gmail.com",
		"password": "12345",
	}
}

func GetEmptyLoginRequest() map[string]interface{} {

	return map[string]interface{}{
		"email":    "",
		"password": "",
	}
}
