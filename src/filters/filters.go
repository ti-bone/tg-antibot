/*
 * filters.go
 * Copyright (c) ti-bone 2022.
 */

package filters

// GetFilters - Return Filters struct that contains string arrays with blacklisted words.
func GetFilters() Filters {
	// TODO:
	// - Add an interaction with API

	return Filters{
		// Messages that forbidden
		Messages: []string{
			"Тот самый нашумевший бот",
			"Тот самый нашyмевший бoт",
			"который находит интимныe",
			"девушки в общей базе",
			"Думаешь фоток твоей подруги тут нет",
			"Ты думаешь фотoк твоей подруrи тут нет",
			"который наxoдит интиmныe",
			"девyшkи в общей базe",
			"слив/взлом твоей подруги/девушки/учительницы",
			"Даю ссылky",
			"Даю ссылoчку",
			"Даю ссылку",
		},
		// Names of users, that forbidden
		Names: []string{
			"сладкая пися",
			"СМОТРИ ПРОФИЛЬ",
			"Сладкаая писееечка",
			"Cладкая пиceчka",
			"Пробив девушек",
			"Сладкaя попа",
			"Мокрая киска",
			"мокренькая щель",
		},
	}
}

// Filters struct
type Filters struct {
	Messages []string `json:"MessagesFilter"`
	Names    []string `json:"NamesFilter"`
}
