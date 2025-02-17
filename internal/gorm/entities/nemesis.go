package entities

const NemesisTableName = "nemesis"

func (NemesisEntity) TableName() string {
	return NemesisTableName
}

type AliveNemeses struct {
	Alive int64
	Dead  int64
}

type NemesisEntity struct {
	ID          uint   `gorm:"primaryKey"`
	IsAlive     bool   `gorm:"not null"`
	Years       *int32 ``
	CharacterID uint   ``

	// has many Secrets
	Secrets []*SecretEntity `gorm:"foreignKey:NemesisID;references:ID"`

	// belongs to Character
	Character CharacterEntity `gorm:"foreignKey:CharacterID;references:ID"`
}

/*
=#> nemesis
	is_alive     | boolean | not null
	years        | integer |
	id           | integer | not null
	character_id | integer |

Indexes:
	"nemesis_pkey" PRIMARY KEY, btree (id)
    "fki_Character Id" btree (id)
References:
	"character" FOREIGN KEY (character_id) REFERENCES "character"(id) NOT VALID

Referenced by:
	TABLE "secret" CONSTRAINT "nemesis" FOREIGN KEY (nemesis_id) REFERENCES nemesis(id)
*/
