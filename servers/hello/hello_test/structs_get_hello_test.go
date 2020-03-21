package hello_test

import "github.com/bmsandoval/ekube/db/models"

type GetHelloTestData struct {
	BaseTestData
	MockGetRelease *MockGetRelease
}

type MockGetRelease struct {
	OutGreetings []models.Greetings
	OutError   error
}
