package request

type CreateBandRecruitmentRequest struct {
	Title           string   `json:"title"`
	Description     string   `json:"description"`
	Genre           string   `json:"genre"`
	Location        string   `json:"location"`
	RecruitingParts []string `json:"recruitingParts"`
	SkillLevel      string   `json:"skillLevel"`
	Contact         string   `json:"contact"`
	Deadline        *string  `json:"deadline"`
	Status          string   `json:"status"`
}

type UpdateBandRecruitmentRequest struct {
	Title           string   `json:"title"`
	Description     string   `json:"description"`
	Genre           string   `json:"genre"`
	Location        string   `json:"location"`
	RecruitingParts []string `json:"recruitingParts"`
	SkillLevel      string   `json:"skillLevel"`
	Contact         string   `json:"contact"`
	Deadline        *string  `json:"deadline"`
	Status          string   `json:"status"`
}

type ApplyBandRecruitmentRequest struct {
	Message string `json:"message"`
}
