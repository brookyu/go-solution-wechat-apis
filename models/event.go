package models

import (
	orm "github.com/brookyu/go-solution-wechat-apis/global"
)

type Event struct {
	Id                     string `json:"id" gorm:"type:char(36);primary_key"`     //
	EventTitle             string `json:"eventTitle" gorm:"type:longtext;"`        //
	EventStartDate         string `json:"eventStartDate" gorm:"type:datetime(3);"` //
	EventEndDate           string `json:"eventEndDate" gorm:"type:datetime(3);"`   //
	IsCurrent              string `json:"isCurrent" gorm:"type:tinyint;"`          //
	LuckRatio              string `json:"luckRatio" gorm:"type:int;"`              //
	NumberOfGiftsForChat   string `json:"numberOfGiftsForChat" gorm:"type:int;"`   //
	NumberOfGiftsForSurvey string `json:"numberOfGiftsForSurvey" gorm:"type:int;"` //
	TagName                string `json:"tagName" gorm:"type:longtext;"`           //
	UserTagId              string `json:"userTagId" gorm:"type:int;"`              //
	AgendaId               string `json:"agendaId" gorm:"type:char(36);"`          //
	BackgroundId           string `json:"backgroundId" gorm:"type:char(36);"`      //
	AboutEventId           string `json:"aboutEventId" gorm:"type:char(36);"`      //
	InstructionsId         string `json:"instructionsId" gorm:"type:char(36);"`    //
	SurveyId               string `json:"surveyId" gorm:"type:char(36);"`          //
	SpeakersId             string `json:"speakersId" gorm:"type:longtext;"`        //
	RegisterFormId         string `json:"registerFormId" gorm:"type:char(36);"`    //
	InteractionCode        string `json:"interactionCode" gorm:"type:longtext;"`   //
	ScanMessage            string `json:"scanMessage" gorm:"type:longtext;"`       //
	ScanNewsId             string `json:"scanNewsId" gorm:"type:char(36);"`        //
	CloudVideoId           string `json:"cloudVideoId" gorm:"type:char(36);"`      //
	Tags                   string `json:"tags" gorm:"type:longtext;"`              //
	DataScope              string `json:"dataScope" gorm:"-"`
	Params                 string `json:"params"  gorm:"-"`
	Model
}

func (Event) TableName() string {
	return "event"
}

// 创建Event
func (e *Event) Create() (Event, error) {
	var doc Event
	result := orm.Eloquent.Table(e.TableName()).Create(&e)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取Event
func (e *Event) Get() (Event, error) {
	var doc Event
	table := orm.Eloquent.Table(e.TableName())

	if e.Id != "" {
		table = table.Where("id = ?", e.Id)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

// 更新Event
func (e *Event) Update(id int) (update Event, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id = ?", id).First(&update).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = orm.Eloquent.Table(e.TableName()).Model(&update).Updates(&e).Error; err != nil {
		return
	}
	return
}

// 删除Event
func (e *Event) Delete(id int) (success bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id = ?", id).Delete(&Event{}).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

//批量删除
func (e *Event) BatchDelete(id []int) (Result bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("id in (?)", id).Delete(&Event{}).Error; err != nil {
		return
	}
	Result = true
	return
}
