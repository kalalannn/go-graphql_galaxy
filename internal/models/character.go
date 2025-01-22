package models

import "time"

/* Postgres character schema:
=# \d character
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

type Character struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	Name            string    `json:"name" gorm:"not null"`
	Gender          string    `json:"gender"`
	Ability         string    `json:"ability" gorm:"not null"`
	MinimalDistance float64   `json:"minimal_distance" gorm:"not null"`
	Weight          float64   `json:"weight"`
	Born            time.Time `json:"born" gorm:"not null"`
	InSpaceSince    time.Time `json:"in_space_since"`
	BeerConsumption int       `json:"beer_consumption" gorm:"not null"`
	KnowsTheAnswer  bool      `json:"knows_the_answer" gorm:"not null"`

	// has many Nemeses
	Nemeses []Nemesis `json:"nemeses" gorm:"foreignKey:CharacterID;references:ID"`
}

func (Character) TableName() string {
	return "character"
}
