package models

import "time"

type Character struct {
	ID              uint       `json:"id" gorm:"primaryKey"`
	Name            string     `json:"name" gorm:"not null"`
	Gender          *string    `json:"gender"`
	Ability         string     `json:"ability" gorm:"not null"`
	MinimalDistance float64    `json:"minimal_distance" gorm:"not null"`
	Weight          *float64   `json:"weight"`
	Born            time.Time  `json:"born" gorm:"not null"`
	InSpaceSince    *time.Time `json:"in_space_since"`
	BeerConsumption int        `json:"beer_consumption" gorm:"not null"`
	KnowsTheAnswer  bool       `json:"knows_the_answer" gorm:"not null"`

	// has many Nemeses
	Nemeses []*Nemesis `json:"nemeses" gorm:"foreignKey:CharacterID;references:ID"`
}

type Nemesis struct {
	ID          uint `json:"id" gorm:"primaryKey"`
	IsAlive     bool `json:"is_alive" gorm:"not null"`
	Years       *int `json:"years"`
	CharacterID uint `json:"character_id"`

	// has many Secrets
	Secrets []*Secret `json:"secrets" gorm:"foreignKey:NemesisID;references:ID"`

	// belongs to Character
	Character Character `gorm:"foreignKey:CharacterID;references:ID"`
}

type Secret struct {
	ID         uint  `json:"id" gorm:"primaryKey"`
	SecretCode int64 `json:"secret_code" gorm:"not null"`
	NemesisID  uint  `json:"nemesis_id" gorm:"not null"`

	// belongs to Nemesis
	Nemesis *Nemesis `gorm:"foreignKey:NemesisID;references:ID"`
}

func (Character) TableName() string {
	return "character"
}

func (Nemesis) TableName() string {
	return "nemesis"
}

func (Secret) TableName() string {
	return "secret"
}

/* Postgres schemas:
=#> character
	id               | integer                     | not null          | nextval('character_id_seq'::regclass)
	name             | text                        | not null          |
	gender           | text                        |                   |
	ability          | text                        | not null          |
	minimal_distance | numeric                     | not null          |
	weight           | numeric                     |                   |
	born             | timestamp without time zone | not null          |
	in_space_since   | timestamp without time zone |                   |
	beer_consumption | integer                     | not null          |
	knows_the_answer | boolean                     | not null          |

Indexes:
	"character_pkey" PRIMARY KEY, btree (id)
Referenced by:
	TABLE "nemesis" CONSTRAINT "character" FOREIGN KEY (character_id) REFERENCES "character"(id) NOT VALID

=#> nemesis
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

=#> secret
	id          | integer |                    | not null          | nextval('secrete_id_seq'::regclass)
	secret_code | bigint  |                    | not null          |
	nemesis_id  | integer |                    | not null          |

Indexes:
	"secrete_pkey" PRIMARY KEY, btree (id)
References:
	"nemesis" FOREIGN KEY (nemesis_id) REFERENCES nemesis(id)
*/
