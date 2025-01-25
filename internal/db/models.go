// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"database/sql"
	"time"
)

type Ammunition struct {
	ID        int64         `json:"id"`
	Name      string        `json:"name"`
	CostGp    float64       `json:"cost_gp"`
	Weight    sql.NullInt64 `json:"weight"`
	Quantity  sql.NullInt64 `json:"quantity"`
	CreatedAt sql.NullTime  `json:"created_at"`
	UpdatedAt sql.NullTime  `json:"updated_at"`
}

type Armor struct {
	ID              int64        `json:"id"`
	Name            string       `json:"name"`
	ArmorClass      int64        `json:"armor_class"`
	CostGp          int64        `json:"cost_gp"`
	DamageReduction int64        `json:"damage_reduction"`
	Weight          int64        `json:"weight"`
	ArmorType       string       `json:"armor_type"`
	MovementRate    int64        `json:"movement_rate"`
	CreatedAt       sql.NullTime `json:"created_at"`
	UpdatedAt       sql.NullTime `json:"updated_at"`
}

type Character struct {
	ID               int64     `json:"id"`
	UserID           int64     `json:"user_id"`
	Name             string    `json:"name"`
	Class            string    `json:"class"`
	Level            int64     `json:"level"`
	MaxHp            int64     `json:"max_hp"`
	CurrentHp        int64     `json:"current_hp"`
	Strength         int64     `json:"strength"`
	Dexterity        int64     `json:"dexterity"`
	Constitution     int64     `json:"constitution"`
	Intelligence     int64     `json:"intelligence"`
	Wisdom           int64     `json:"wisdom"`
	Charisma         int64     `json:"charisma"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	ExperiencePoints int64     `json:"experience_points"`
}

type CharacterInventory struct {
	ID                   int64          `json:"id"`
	CharacterID          int64          `json:"character_id"`
	ItemType             string         `json:"item_type"`
	ItemID               int64          `json:"item_id"`
	Quantity             int64          `json:"quantity"`
	ContainerInventoryID sql.NullInt64  `json:"container_inventory_id"`
	EquipmentSlotID      sql.NullInt64  `json:"equipment_slot_id"`
	Notes                sql.NullString `json:"notes"`
	CreatedAt            time.Time      `json:"created_at"`
	UpdatedAt            time.Time      `json:"updated_at"`
}

type Container struct {
	ID             int64          `json:"id"`
	Name           string         `json:"name"`
	CostGp         float64        `json:"cost_gp"`
	Weight         sql.NullInt64  `json:"weight"`
	CapacityWeight int64          `json:"capacity_weight"`
	CapacityItems  sql.NullInt64  `json:"capacity_items"`
	Description    sql.NullString `json:"description"`
	CreatedAt      sql.NullTime   `json:"created_at"`
	UpdatedAt      sql.NullTime   `json:"updated_at"`
}

type Equipment struct {
	ID          int64          `json:"id"`
	Name        string         `json:"name"`
	CostGp      float64        `json:"cost_gp"`
	Weight      sql.NullInt64  `json:"weight"`
	Description sql.NullString `json:"description"`
	CreatedAt   sql.NullTime   `json:"created_at"`
	UpdatedAt   sql.NullTime   `json:"updated_at"`
}

type EquipmentSlot struct {
	ID          int64          `json:"id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
}

type RangedWeapon struct {
	ID          int64          `json:"id"`
	Name        string         `json:"name"`
	WeaponType  string         `json:"weapon_type"`
	CostGp      sql.NullInt64  `json:"cost_gp"`
	Weight      int64          `json:"weight"`
	RateOfFire  string         `json:"rate_of_fire"`
	RangeShort  sql.NullInt64  `json:"range_short"`
	RangeMedium sql.NullInt64  `json:"range_medium"`
	RangeLong   sql.NullInt64  `json:"range_long"`
	Damage      sql.NullString `json:"damage"`
	CreatedAt   sql.NullTime   `json:"created_at"`
	UpdatedAt   sql.NullTime   `json:"updated_at"`
}

type RangedWeaponProperty struct {
	ID          int64  `json:"id"`
	Symbol      string `json:"symbol"`
	Description string `json:"description"`
}

type RangedWeaponPropertyLink struct {
	WeaponID   int64 `json:"weapon_id"`
	PropertyID int64 `json:"property_id"`
}

type Session struct {
	Token     string    `json:"token"`
	UserID    int64     `json:"user_id"`
	ExpiresAt time.Time `json:"expires_at"`
}

type Shield struct {
	ID           int64        `json:"id"`
	Name         string       `json:"name"`
	CostGp       int64        `json:"cost_gp"`
	Weight       int64        `json:"weight"`
	DefenseBonus int64        `json:"defense_bonus"`
	CreatedAt    sql.NullTime `json:"created_at"`
	UpdatedAt    sql.NullTime `json:"updated_at"`
}

type User struct {
	ID           int64     `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	CreatedAt    time.Time `json:"created_at"`
}

type Weapon struct {
	ID              int64          `json:"id"`
	Name            string         `json:"name"`
	Reach           int64          `json:"reach"`
	CostGp          int64          `json:"cost_gp"`
	Weight          int64          `json:"weight"`
	RangeShort      sql.NullInt64  `json:"range_short"`
	RangeMedium     sql.NullInt64  `json:"range_medium"`
	RangeLong       sql.NullInt64  `json:"range_long"`
	AttacksPerRound sql.NullString `json:"attacks_per_round"`
	Damage          string         `json:"damage"`
	CreatedAt       sql.NullTime   `json:"created_at"`
	UpdatedAt       sql.NullTime   `json:"updated_at"`
}

type WeaponProperty struct {
	ID          int64  `json:"id"`
	Symbol      string `json:"symbol"`
	Description string `json:"description"`
}

type WeaponPropertyLink struct {
	WeaponID   int64 `json:"weapon_id"`
	PropertyID int64 `json:"property_id"`
}
