package utils_test

import (
	"AG-Installer/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDownloads(t *testing.T) {
	actual, _ := utils.Downloads("https://asdfsfd.com", "aasdf")
	expected := true
	assert.Equal(t, expected, actual, "기대값과 결과값이 다릅니다.")
}
