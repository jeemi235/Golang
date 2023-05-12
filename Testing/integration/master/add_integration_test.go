// +build integration

package master

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestAPISum(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		t.Fatalf("could not load .env file: %v", err)
	}

	token := os.Getenv("AUTH_TOKEN")

	client := MathClient{
		Token: token,
		// Host:  "math.example.com", // This is the example used in the blog
		Host: "127.0.0.1:9000",
	}

	sum, err := client.APISum(2, 2)
	if err != nil {
		t.Errorf("No error expected, got %v", err)
	}

	if sum != 4 {
		t.Errorf("Expected %v, got %v", 4, sum)
	}
}
