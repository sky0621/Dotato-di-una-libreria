package model

import "Dotato-di-una-libreria/backend2/domain/valueobject"

// Notice ... お知らせ
type Notice struct {
	ID       string
	Title    string
	Sentence string
	Severity valueobject.NoticeSeverity
}
