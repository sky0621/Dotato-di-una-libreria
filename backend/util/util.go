package util

import (
	"strings"

	"github.com/google/uuid"
	"google.golang.org/appengine"
)

// IsLocal ...
func IsLocal() bool {
	// 現状は appengine パッケージを使っているのでこの判定。
	// 変えたくなった時に判定ロジックが多方面に散らばってると修正面倒なのでユーティリティ化。
	return appengine.IsDevAppServer()
}

// CreateUniqueID ...
func CreateUniqueID() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}
