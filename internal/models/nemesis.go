package models

/* Postgres nemesis schema:
=# \d nemesis
	is_alive     | boolean |                    | not null          |
	years        | integer |                    |                   |
	id           | integer |                    | not null          | nextval('nemesis_id_seq'::regclass)
	character_id | integer |                    |                   |

Indexes:
	"nemesis_pkey" PRIMARY KEY, btree (id)
    "fki_Character Id" btree (id)
References:
	"character" FOREIGN KEY (character_id) REFERENCES "character"(id) NOT VALID

Referenced by:
	TABLE "secret" CONSTRAINT "nemesis" FOREIGN KEY (nemesis_id) REFERENCES nemesis(id)
*/

type Nemesis struct {
	ID          uint `json:"id" gorm:"primaryKey"`
	IsAlive     bool `json:"is_alive" gorm:"not null"`
	Years       int  `json:"years"`
	CharacterID uint `json:"character_id"`

	// has many Secrets
	Secrets []Secret `json:"secrets" gorm:"foreignKey:NemesisID;references:ID"`

	// belongs to Character
	Character Character `gorm:"foreignKey:CharacterID;references:ID"`
}

func (Nemesis) TableName() string {
	return "nemesis"
}
