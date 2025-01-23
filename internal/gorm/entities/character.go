package entities

import (
	"time"
)

func (CharacterEntity) TableName() string {
	return "character"
}

type CharacterEntity struct {
	ID              uint       `gorm:"primaryKey"`
	Name            string     `gorm:"not null"`
	Gender          *string    ``
	Ability         string     `gorm:"not null"`
	MinimalDistance float64    `gorm:"not null"`
	Weight          *float64   ``
	Born            time.Time  `gorm:"not null"`
	InSpaceSince    *time.Time ``
	BeerConsumption int32      `gorm:"not null"`
	KnowsTheAnswer  bool       `gorm:"not null"`

	// has many Nemeses
	Nemeses []*NemesisEntity `gorm:"foreignKey:CharacterID;references:ID"`
}

/*
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
*/
