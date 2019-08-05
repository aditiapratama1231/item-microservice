package service

import (
	"github.com/jinzhu/gorm"

	models "github.com/aditiapratama1231/item-microservice/database/models"
	payload "github.com/aditiapratama1231/item-microservice/pkg/request/payload"
)

type ItemService interface {
	GetItems() payload.GetItemResponse
	CreateItem(payload.CreateItemRequest) payload.ItemResponse
	UpdateItem(payload.UpdateItemRequest) payload.ItemResponse
	ShowItem(string) payload.ItemResponse
	DeleteItem(string) payload.ItemResponse
}

type itemService struct{}

var query *gorm.DB

func NewIttemService(db *gorm.DB) ItemService {
	query = db
	return itemService{}
}

func (itemService) GetItems() payload.GetItemResponse {
	var (
		items []models.Item
	)

	query.Find(&items)

	return payload.GetItemResponse{
		Data: items,
	}
}

func (itemService) CreateItem(data payload.CreateItemRequest) payload.ItemResponse {
	item := models.Item{
		Title:       data.Data.Title,
		Description: data.Data.Description,
		IsFinish:    false,
	}

	err := query.Create(&item).Error
	if err != nil {
		return payload.ItemResponse{
			Message:    "Item Failed To Create : " + err.Error(),
			StatusCode: 500,
			Err:        true,
		}
	}

	return payload.ItemResponse{
		Message:    "Item Created Successfully",
		Data:       item,
		StatusCode: 200,
	}
}

func (itemService) UpdateItem(data payload.UpdateItemRequest) payload.ItemResponse {
	var item models.Item

	if query.Find(&item, data.ID).RecordNotFound() {
		return payload.ItemResponse{
			Message:    "Item not found",
			StatusCode: 404,
			Err:        true,
		}
	}

	item.Title = data.Data.Title
	item.Description = data.Data.Description
	item.IsFinish = data.Data.IsFinish

	err := query.Save(&item).Error
	if err != nil {
		return payload.ItemResponse{
			Message:    "Failed To Update Item : " + err.Error(),
			StatusCode: 500,
			Err:        true,
		}
	}

	return payload.ItemResponse{
		Message:    "Update Item Success",
		StatusCode: 200,
		Data:       item,
	}
}

func (itemService) ShowItem(id string) payload.ItemResponse {
	var item models.Item

	if query.Find(&item, id).RecordNotFound() {
		return payload.ItemResponse{
			Message:    "Item not found",
			StatusCode: 404,
			Err:        true,
		}
	}

	return payload.ItemResponse{
		Message:    "Item Retrieved succesfully",
		Data:       item,
		StatusCode: 200,
	}
}

func (itemService) DeleteItem(id string) payload.ItemResponse {
	var item models.Item

	if query.Find(&item, id).RecordNotFound() {
		return payload.ItemResponse{
			Message:    "Item not found",
			StatusCode: 404,
			Err:        true,
		}
	}

	err := query.Delete(&item, id).Error
	if err != nil {
		return payload.ItemResponse{
			Message:    "Failed To Delete item : " + err.Error(),
			StatusCode: 500,
			Err:        true,
		}
	}

	return payload.ItemResponse{
		Message:    "Item Successfully deleted",
		StatusCode: 204,
		Err:        true,
	}
}
