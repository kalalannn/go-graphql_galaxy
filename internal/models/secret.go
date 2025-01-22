package models

/* Postgres secret schema:
=# \d secret
	id          | integer |                    | not null          | nextval('secrete_id_seq'::regclass)
	secret_code | bigint  |                    | not null          |
	nemesis_id  | integer |                    | not null          |

Indexes:
	"secrete_pkey" PRIMARY KEY, btree (id)
References:
	"nemesis" FOREIGN KEY (nemesis_id) REFERENCES nemesis(id)
*/

type Secret struct {
	ID         uint  `json:"id" gorm:"primaryKey"`
	SecretCode int64 `json:"secret_code" gorm:"not null"`
	NemesisID  uint  `json:"nemesis_id" gorm:"not null"`

	// belongs to Nemesis
	Nemesis Nemesis `gorm:"foreignKey:NemesisID;references:ID"`
}

func (Secret) TableName() string {
	return "secret"
}
