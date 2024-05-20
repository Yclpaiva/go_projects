package tools

import "time"

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"test": {
		AuthToken: "12345",
		Username:  "test",
	},
	"yuri": {
		AuthToken: "1234AB",
		Username:  "yuri",
	},
}

var mockCoinsDetails = map[string]CoinsDetails{
	"test": {
		Coins:    100,
		Username: "test",
	},
	"yuri": {
		Coins:    1000,
		Username: "yuri",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// Simulation of DB Calling
	time.Sleep(1 * time.Second)

	clientData := LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinsDetails {
	// Simulation of DB Calling
	time.Sleep(1 * time.Second)

	clientData := CoinsDetails{}
	clientData, ok := mockCoinsDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
