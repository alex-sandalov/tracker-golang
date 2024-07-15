package response

type GetInfoResponse struct {
	Surname    string `json:"surname" form:"surname"`
	Name       string `json:"name" form:"name"`
	Patronymic string `json:"patronymic" form:"patronymic"`
	Address    string `json:"address" form:"address"`
}
