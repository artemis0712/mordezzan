package server

import (
	"database/sql"
	"time"

	"github.com/marbh56/mordezzan/internal/db"
	"github.com/marbh56/mordezzan/internal/rules"
	"github.com/marbh56/mordezzan/internal/rules/currency"
)

func interfaceToNullString(v interface{}) sql.NullString {
	if v == nil {
		return sql.NullString{}
	}
	if s, ok := v.(string); ok {
		return sql.NullString{
			String: s,
			Valid:  true,
		}
	}
	return sql.NullString{}
}

func interfaceToNullInt64(v interface{}) sql.NullInt64 {
	if v == nil {
		return sql.NullInt64{}
	}
	switch i := v.(type) {
	case int64:
		return sql.NullInt64{Int64: i, Valid: true}
	case int:
		return sql.NullInt64{Int64: int64(i), Valid: true}
	case sql.NullInt64:
		return i
	}
	return sql.NullInt64{}
}

// Represents a single item in a character's inventory
type InventoryItem struct {
	ID                   int64          `json:"id"`
	CharacterID          int64          `json:"character_id"`
	ItemType             string         `json:"item_type"`
	ItemID               int64          `json:"item_id"`
	ItemName             string         `json:"item_name"`
	ItemWeight           int            `json:"item_weight"`
	Quantity             int64          `json:"quantity"`
	ContainerInventoryID sql.NullInt64  `json:"container_inventory_id"`
	EquipmentSlotID      sql.NullInt64  `json:"equipment_slot_id"`
	SlotName             sql.NullString `json:"slot_name"`
	Damage               sql.NullString `json:"damage"`
	AttacksPerRound      sql.NullString `json:"attacks_per_round"`
	MovementRate         sql.NullInt64  `json:"movement_rate"`
	Notes                sql.NullString `json:"notes"`
	CreatedAt            time.Time      `json:"created_at"`
	UpdatedAt            time.Time      `json:"updated_at"`
}

// Contains inventory statistics and calculated values
type InventoryStats struct {
	TotalWeight         int    `json:"total_weight"`
	EquippedWeight      int    `json:"equipped_weight"`
	CarriedWeight       int    `json:"carried_weight"`
	ContainersWeight    int    `json:"containers_weight"`
	CoinWeight          int    `json:"coin_weight"`
	BaseEncumbered      int    `json:"base_encumbered"`
	BaseHeavyEncumbered int    `json:"base_heavy_encumbered"`
	MaximumCapacity     int    `json:"maximum_capacity"`
	EncumbranceLevel    string `json:"encumbrance_level"`
}

func classGetsFighterBonus(class string) bool {
	fighterClasses := map[string]bool{
		"Fighter":   true,
		"Barbarian": true,
		"Beserker":  true,
		"Huntsman":  true,
		"Paladin":   true,
		"Ranger":    true,
		"Warlock":   true,
	}
	return fighterClasses[class]
}

// Complete character view model including inventory
type CharacterViewModel struct {
	ID         int64  `json:"id"`
	UserID     int64  `json:"user_id"`
	Name       string `json:"name"`
	Class      string `json:"class"`
	Level      int64  `json:"level"`
	MaxHp      int64  `json:"max_hp"`
	CurrentHp  int64  `json:"current_hp"`
	ArmorClass int    `json:"armor_class"`

	// Ability scores with modifiers
	Strength          int64                   `json:"strength"`
	StrengthModifiers rules.StrengthModifiers `json:"strength_modifiers"`

	Dexterity          int64                    `json:"dexterity"`
	DexterityModifiers rules.DexterityModifiers `json:"dexterity_modifiers"`

	Constitution          int64                       `json:"constitution"`
	ConstitutionModifiers rules.ConstitutionModifiers `json:"constitution_modifiers"`

	Intelligence          int64                       `json:"intelligence"`
	IntelligenceModifiers rules.IntelligenceModifiers `json:"intelligence_modifiers"`

	Wisdom          int64                 `json:"wisdom"`
	WisdomModifiers rules.WisdomModifiers `json:"wisdom_modifiers"`

	Charisma          int64                   `json:"charisma"`
	CharismaModifiers rules.CharismaModifiers `json:"charisma_modifiers"`

	// Combat information
	CombatMatrix []int64 `json:"combat_matrix"`
	SavingThrow  int64   `json:"saving_throw"`

	// Inventory organization
	EquippedItems  []InventoryItem           `json:"equipped_items"`
	CarriedItems   []InventoryItem           `json:"carried_items"`
	ContainerItems map[int64][]InventoryItem `json:"container_items"`

	// Calculated inventory statistics
	InventoryStats InventoryStats `json:"inventory_stats"`
	PlatinumPieces int64
	GoldPieces     int64
	ElectrumPieces int64
	SilverPieces   int64
	CopperPieces   int64

	// Experience points and level progression
	ExperiencePoints int64 `json:"experience_points"`
	NextLevelXP      int64 `json:"next_level_xp"`
	XPNeeded         int64 `json:"xp_needed"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Helper function to safely convert interface{} to string
func interfaceToString(v interface{}) string {
	if v == nil {
		return ""
	}
	if s, ok := v.(string); ok {
		return s
	}
	return ""
}

// Helper function to safely convert interface{} to int
func interfaceToInt(v interface{}) int {
	if v == nil {
		return 0
	}
	if i, ok := v.(int64); ok {
		return int(i)
	}
	if i, ok := v.(int); ok {
		return i
	}
	return 0
}

// Helper functions for templates
func add(a, b int) int {
	return a + b
}

// Creates a new character view model with inventory data
func NewCharacterViewModel(c db.Character, inventory []db.GetCharacterInventoryRow) CharacterViewModel {
	vm := CharacterViewModel{
		ID:               c.ID,
		UserID:           c.UserID,
		Name:             c.Name,
		Class:            c.Class,
		Level:            c.Level,
		MaxHp:            c.MaxHp,
		CurrentHp:        c.CurrentHp,
		Strength:         c.Strength,
		Dexterity:        c.Dexterity,
		Constitution:     c.Constitution,
		Intelligence:     c.Intelligence,
		Wisdom:           c.Wisdom,
		Charisma:         c.Charisma,
		CreatedAt:        c.CreatedAt,
		UpdatedAt:        c.UpdatedAt,
		ExperiencePoints: c.ExperiencePoints,

		// Initialize modifiers
		StrengthModifiers:     rules.CalculateStrengthModifiers(c.Strength),
		DexterityModifiers:    rules.CalculateDexterityModifiers(c.Dexterity),
		ConstitutionModifiers: rules.CalculateConstitutionModifiers(c.Constitution),
		IntelligenceModifiers: rules.CalculateIntelligenceModifiers(c.Intelligence),
		WisdomModifiers:       rules.CalculateWisdomModifiers(c.Wisdom),
		CharismaModifiers:     rules.CalculateCharismaModifiers(c.Charisma),

		// Initialize inventory containers
		ContainerItems: make(map[int64][]InventoryItem),
	}

	if classGetsFighterBonus(c.Class) {
		vm.StrengthModifiers.ExtraordinaryFeat += 8
	}

	// Get class progression
	progression := rules.GetClassProgression(vm.Class)

	// Calculate XP needed for next level
	vm.XPNeeded = progression.GetXPForNextLevel(c.ExperiencePoints)

	// Get XP required for current level
	for _, level := range progression.Levels {
		if level.XPRequired > c.ExperiencePoints {
			vm.NextLevelXP = level.XPRequired
			break
		}
	}

	// Calculate base AC
	baseAC := 9
	var armorAC int64
	var shieldBonus int64

	// Check equipped items for armor and shield
	for _, item := range inventory {
		if item.EquipmentSlotID.Valid {
			switch item.ItemType {
			case "armor":
				// Get AC from armor
				var ac sql.NullInt64
				switch v := item.ArmorClass.(type) {
				case int64:
					ac = sql.NullInt64{Int64: v, Valid: true}
				case sql.NullInt64:
					ac = v
				}
				if ac.Valid {
					armorAC = ac.Int64
				}
			case "shield":
				// Get defense bonus from shield
				var bonus sql.NullInt64
				switch v := item.DefenseBonus.(type) {
				case int64:
					bonus = sql.NullInt64{Int64: v, Valid: true}
				case sql.NullInt64:
					bonus = v
				}
				if bonus.Valid {
					shieldBonus = bonus.Int64
				}
			}
		}
	}

	// If armor is equipped, use its AC instead of base AC
	if armorAC > 0 {
		baseAC = int(armorAC)
	}

	// Apply shield bonus if any
	totalAC := baseAC - int(shieldBonus)

	// Apply Dexterity modifier
	totalAC -= vm.DexterityModifiers.DefenseAdj

	vm.ArmorClass = totalAC

	// Initialize inventory stats with encumbrance thresholds
	encumbranceThresholds := rules.CalculateEncumbranceThresholds(c.Strength, c.Constitution)
	vm.InventoryStats = InventoryStats{
		BaseEncumbered:      encumbranceThresholds.BaseEncumbered,
		BaseHeavyEncumbered: encumbranceThresholds.BaseHeavyEncumbered,
		MaximumCapacity:     encumbranceThresholds.MaximumCapacity,
	}

	// Add coin weight to total weight
	coinage := currency.Purse{
		PlatinumPieces: c.PlatinumPieces,
		GoldPieces:     c.GoldPieces,
		ElectrumPieces: c.ElectrumPieces,
		SilverPieces:   c.SilverPieces,
		CopperPieces:   c.CopperPieces,
	}
	vm.PlatinumPieces = c.PlatinumPieces
	vm.GoldPieces = c.GoldPieces
	vm.ElectrumPieces = c.ElectrumPieces
	vm.SilverPieces = c.SilverPieces
	vm.CopperPieces = c.CopperPieces
	vm.InventoryStats.CoinWeight = int(currency.GetTotalWeight(&coinage))

	for _, item := range inventory {
		damage := sql.NullString{}
		switch d := item.Damage.(type) {
		case string:
			damage = sql.NullString{String: d, Valid: true}
		case sql.NullString:
			damage = d
		}

		invItem := InventoryItem{
			ID:                   item.ID,
			CharacterID:          item.CharacterID,
			ItemType:             item.ItemType,
			ItemID:               item.ItemID,
			ItemName:             interfaceToString(item.ItemName),
			ItemWeight:           interfaceToInt(item.ItemWeight),
			Quantity:             item.Quantity,
			ContainerInventoryID: item.ContainerInventoryID,
			EquipmentSlotID:      item.EquipmentSlotID,
			SlotName:             item.SlotName,
			Damage:               damage,
			AttacksPerRound:      interfaceToNullString(item.AttacksPerRound),
			MovementRate:         interfaceToNullInt64(item.MovementRate),
			Notes:                item.Notes,
			CreatedAt:            item.CreatedAt,
			UpdatedAt:            item.UpdatedAt,
		}

		itemTotalWeight := invItem.ItemWeight * int(invItem.Quantity)

		if invItem.EquipmentSlotID.Valid {
			vm.EquippedItems = append(vm.EquippedItems, invItem)
			vm.InventoryStats.EquippedWeight += itemTotalWeight
		} else if invItem.ContainerInventoryID.Valid {
			containerID := invItem.ContainerInventoryID.Int64
			vm.ContainerItems[containerID] = append(vm.ContainerItems[containerID], invItem)
			vm.InventoryStats.ContainersWeight += itemTotalWeight
		} else {
			vm.CarriedItems = append(vm.CarriedItems, invItem)
			vm.InventoryStats.CarriedWeight += itemTotalWeight
		}
	}

	// Calculate total weight and encumbrance level
	vm.InventoryStats.TotalWeight = vm.InventoryStats.EquippedWeight +
		vm.InventoryStats.CarriedWeight +
		vm.InventoryStats.ContainersWeight

	// Determine encumbrance level
	switch {
	case vm.InventoryStats.TotalWeight > vm.InventoryStats.MaximumCapacity:
		vm.InventoryStats.EncumbranceLevel = "Over"
	case vm.InventoryStats.TotalWeight > vm.InventoryStats.BaseHeavyEncumbered:
		vm.InventoryStats.EncumbranceLevel = "Heavy"
	case vm.InventoryStats.TotalWeight > vm.InventoryStats.BaseEncumbered:
		vm.InventoryStats.EncumbranceLevel = "Encumbered"
	default:
		vm.InventoryStats.EncumbranceLevel = "None"
	}

	// Calculate FA and generate combat matrix row
	fa := rules.CalculateFightingAbility(c.Class, c.Level)
	vm.CombatMatrix = make([]int64, 19) // -9 to 9 AC
	for ac := -9; ac <= 9; ac++ {
		vm.CombatMatrix[ac+9] = rules.GetTargetNumber(fa, int64(ac))
	}

	// Get saving throw value
	progression = rules.GetClassProgression(vm.Class)
	vm.SavingThrow = progression.GetSavingThrow(vm.Level)

	return vm
}
