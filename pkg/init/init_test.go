package init

import "testing"

func TestInitializeConfiguration(t *testing.T) {
	host := "localhost"
	port := 8080
	err := InitializeConfiguration(host, port)
	if err != nil {
		t.Fatalf("excepted no error; but got %v", err)
	}
}
