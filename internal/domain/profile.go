package domain

import "time"

// Profile は、システムのコアとなるプロフィール情報を表すエンティティです。
// この構造体は、アプリケーションのビジネスロジック層と永続化層でのみ使用されます。
// プレゼンテーション層（例: JSONレスポンス）でこの構造体を直接使用することは避けてください。
// 代わりに、各層で必要とされるDTO（Data Transfer Object）を定義し、
// このエンティティとDTOの間でマッピング（変換）を行ってください。
// これにより、関心の分離が保たれ、各層が独立して変更可能になります。
type Profile struct {
	ID          string
	Name        string
	Affiliation string
	Bio         string
	InstagramID string
	TwitterID   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
