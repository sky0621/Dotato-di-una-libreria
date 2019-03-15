package valueobject

import (
	"strings"

	"github.com/google/uuid"
)

// UniqueID ... 各エンティティをユニークに識別するIDとして利用
type UniqueID string

// CreateUniqueID ...
func CreateUniqueID() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}
