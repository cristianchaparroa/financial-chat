package fixtures

func GetInvalidRegisterRequest() map[string]interface{} {
	return map[string]interface{}{
		"email":    "test@gmail.com",
		"password": "",
		"name":     "John Doe",
	}
}

func GetRegisterRequest() map[string]interface{} {
	return map[string]interface{}{
		"email":    "test@gmail.com",
		"password": "1234567",
		"name":     "John Doe",
	}
}
