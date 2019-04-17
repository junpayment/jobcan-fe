package jobcan

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTouch(t *testing.T) {
	t.Run("touch", func(t *testing.T) {
		err := Touch(os.Getenv("EMAIL"), os.Getenv("PASSWORD"))
		assert.Nil(t, err)
	})
	t.Run("check in", func(t *testing.T) {
		err := Touch(os.Getenv("EMAIL"), os.Getenv("PASSWORD"), "check_in", "junpayment")
		assert.Nil(t, err)
	})
	t.Run("check out", func(t *testing.T) {
		err := Touch(os.Getenv("EMAIL"), os.Getenv("PASSWORD"), "check_out", "junpayment")
		assert.Nil(t, err)
	})
}

func TestSlack(t *testing.T) {
	t.Run("slack", func(t *testing.T) {
		err := slack("junpayment", "テスト")
		assert.Nil(t, err)
	})
}
