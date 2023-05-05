package dto

import "mime/multipart"

type AvatarUpload struct {
	Id     int                   `json:"id" form:"id"`
	Table  string                `json:"table" form:"table"`
	Avatar *multipart.FileHeader `json:"avatar" form:"avatar"`
}
