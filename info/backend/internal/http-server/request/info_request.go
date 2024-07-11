package request

type GetInfoRequest struct {
	PassportSerie  int `form:"passportSerie" binding:"required"`
	PassportNumber int `form:"passportNumber" binding:"required"`
}
