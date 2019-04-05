package jobcan

import (
	"log"
	"os"
	"testing"
)

func TestTouch(t *testing.T) {
	err := Touch(os.Getenv("EMAIL"), os.Getenv("PASSWORD"))
	log.Print(err)
}
