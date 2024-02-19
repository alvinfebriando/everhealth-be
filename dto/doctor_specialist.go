package dto

import "github.com/alvinfebriando/project-batman-be/entity"

type DoctorSpecialistRes struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

func NewDoctorSpecialistres(d *entity.DoctorSpecialist) DoctorSpecialistRes {
	return DoctorSpecialistRes{Id: d.Id, Name: d.Name, Image: d.Image}
}
