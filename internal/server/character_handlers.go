package server

import (
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/marbh56/mordezzan/internal/db"
	"github.com/marbh56/mordezzan/internal/rules"
)

func (s *Server) HandleCharacterDetail(w http.ResponseWriter, r *http.Request) {
	user, ok := GetUserFromContext(r.Context())
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Get character ID from URL query parameter
	characterIDStr := r.URL.Query().Get("id")
	characterID, err := strconv.ParseInt(characterIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid character ID", http.StatusBadRequest)
		return
	}

	// Get character from database
	queries := db.New(s.db)
	character, err := queries.GetCharacter(r.Context(), db.GetCharacterParams{
		ID:     characterID,
		UserID: user.UserID,
	})
	if err != nil {
		log.Printf("Error fetching character: %v", err)
		http.Error(w, "Character not found", http.StatusNotFound)
		return
	}

	// Create view model with calculated modifiers
	viewModel := NewCharacterViewModel(character)

	tmpl, err := template.ParseFiles(
		"templates/layout/base.html",
		"templates/characters/detail.html",
	)
	if err != nil {
		log.Printf("Template parsing error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := struct {
		IsAuthenticated bool
		Username        string
		Character       CharacterViewModel
		FlashMessage    string
		CurrentYear     int
	}{
		IsAuthenticated: true,
		Username:        user.Username,
		Character:       viewModel,
		FlashMessage:    r.URL.Query().Get("message"),
		CurrentYear:     time.Now().Year(),
	}

	err = tmpl.ExecuteTemplate(w, "base.html", data)
	if err != nil {
		log.Printf("Template execution error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (s *Server) HandleCharacterList(w http.ResponseWriter, r *http.Request) {
	// Get user from context (set by AuthMiddleware)
	user, ok := GetUserFromContext(r.Context())
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Get all characters for the user
	queries := db.New(s.db)
	characters, err := queries.ListCharactersByUser(r.Context(), user.UserID)
	if err != nil {
		log.Printf("Error fetching characters: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles(
		"templates/layout/base.html",
		"templates/characters/list.html",
	)
	if err != nil {
		log.Printf("Template parsing error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := struct {
		IsAuthenticated bool
		Username        string
		Characters      []db.Character
		FlashMessage    string
		CurrentYear     int
	}{
		IsAuthenticated: true,
		Username:        user.Username,
		Characters:      characters,
		FlashMessage:    r.URL.Query().Get("message"),
		CurrentYear:     time.Now().Year(),
	}

	err = tmpl.ExecuteTemplate(w, "base.html", data)
	if err != nil {
		log.Printf("Template execution error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (s *Server) HandleCharacterCreate(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.handleCharacterCreateForm(w, r)
	case http.MethodPost:
		s.handleCharacterCreateSubmission(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (s *Server) handleCharacterCreateForm(w http.ResponseWriter, r *http.Request) {
	user, ok := GetUserFromContext(r.Context())
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	tmpl, err := template.ParseFiles(
		"templates/layout/base.html",
		"templates/characters/create.html",
	)
	if err != nil {
		log.Printf("Template parsing error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := struct {
		IsAuthenticated bool
		Username        string
		FlashMessage    string
		CurrentYear     int
	}{
		IsAuthenticated: true,
		Username:        user.Username,
		FlashMessage:    r.URL.Query().Get("message"),
		CurrentYear:     time.Now().Year(),
	}

	err = tmpl.ExecuteTemplate(w, "base.html", data)
	if err != nil {
		log.Printf("Template execution error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (s *Server) handleCharacterCreateSubmission(w http.ResponseWriter, r *http.Request) {
    user, ok := GetUserFromContext(r.Context())
    if !ok {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    if err := r.ParseForm(); err != nil {
        http.Error(w, "Failed to parse form", http.StatusBadRequest)
        return
    }

    // Parse and validate form fields
    params := db.CreateCharacterParams{
        UserID: user.UserID,
        Name:   r.Form.Get("name"),
        Class:  r.Form.Get("class"),
    }

    // Validate class
    if params.Class != "Fighter" {
        http.Redirect(w, r, "/characters/create?message=Invalid character class", http.StatusSeeOther)
        return
    }

    // Parse level
    level, err := strconv.ParseInt(r.Form.Get("level"), 10, 64)
    if err != nil {
        http.Redirect(w, r, "/characters/create?message=Invalid level value", http.StatusSeeOther)
        return
    }
    params.Level = level

    // Parse base HP and constitution
    baseHP, err := strconv.ParseInt(r.Form.Get("max_hp"), 10, 64)
    if err != nil {
        http.Redirect(w, r, "/characters/create?message=Invalid HP value", http.StatusSeeOther)
        return
    }

    constitution, err := strconv.ParseInt(r.Form.Get("constitution"), 10, 64)
    if err != nil {
        http.Redirect(w, r, "/characters/create?message=Invalid constitution value", http.StatusSeeOther)
        return
    }
    params.Constitution = constitution

    // Calculate total HP using the rules package
    totalHP, err := rules.CalculateTotalHP(baseHP, level, constitution)
    if err != nil {
        http.Redirect(w, r, "/characters/create?message="+err.Error(), http.StatusSeeOther)
        return
    }
    params.MaxHp = totalHP
    params.CurrentHp = totalHP

    // Parse other ability scores
    str, err := strconv.ParseInt(r.Form.Get("strength"), 10, 64)
    if err != nil {
        http.Redirect(w, r, "/characters/create?message=Invalid strength value", http.StatusSeeOther)
        return
    }
    params.Strength = str

    dex, err := strconv.ParseInt(r.Form.Get("dexterity"), 10, 64)
    if err != nil {
        http.Redirect(w, r, "/characters/create?message=Invalid dexterity value", http.StatusSeeOther)
        return
    }
    params.Dexterity = dex

    intel, err := strconv.ParseInt(r.Form.Get("intelligence"), 10, 64)
    if err != nil {
        http.Redirect(w, r, "/characters/create?message=Invalid intelligence value", http.StatusSeeOther)
        return
    }
    params.Intelligence = intel

    wis, err := strconv.ParseInt(r.Form.Get("wisdom"), 10, 64)
    if err != nil {
        http.Redirect(w, r, "/characters/create?message=Invalid wisdom value", http.StatusSeeOther)
        return
    }
    params.Wisdom = wis

    cha, err := strconv.ParseInt(r.Form.Get("charisma"), 10, 64)
    if err != nil {
        http.Redirect(w, r, "/characters/create?message=Invalid charisma value", http.StatusSeeOther)
        return
    }
    params.Charisma = cha

    // Validate ability scores
    abilities := []int64{params.Strength, params.Dexterity, params.Constitution,
        params.Intelligence, params.Wisdom, params.Charisma}
    for _, score := range abilities {
        if score < 3 || score > 18 {
            http.Redirect(w, r, "/characters/create?message=Ability scores must be between 3 and 18",
                http.StatusSeeOther)
            return
        }
    }

    // Create character in database
    queries := db.New(s.db)
    _, err = queries.CreateCharacter(r.Context(), params)
    if err != nil {
        log.Printf("Error creating character: %v", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, "/characters?message=Character created successfully", http.StatusSeeOther)
}
