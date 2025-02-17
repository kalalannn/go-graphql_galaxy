package entities

const SecretTableName = "secret"

func (SecretEntity) TableName() string {
	return SecretTableName
}

type SecretEntity struct {
	ID         uint  `gorm:"primaryKey"`
	SecretCode int64 `gorm:"not null"`
	NemesisID  uint  `gorm:"not null"`

	// belongs to Nemesis
	Nemesis NemesisEntity `gorm:"foreignKey:NemesisID;references:ID"`
}

/* Postgres schemas:
=#> secret
	id          | integer | not null
	secret_code | bigint  | not null
	nemesis_id  | integer | not null

Indexes:
	"secrete_pkey" PRIMARY KEY, btree (id)
References:
	"nemesis" FOREIGN KEY (nemesis_id) REFERENCES nemesis(id)
*/
