package response

import "backend-pack/internal/app/model"

type PackResponse struct {
	Tutorial  model.Tutorial `json:"tutorial"`
	PackLog   model.PackLog  `json:"pack_log"`
	StartPage string         `json:"start_page"`
	FileName  string         `json:"file_name"`
}
