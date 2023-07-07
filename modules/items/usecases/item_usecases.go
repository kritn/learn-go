package usecases

// มีหน้าที่ รับมือกับ Logic ต่างๆ ก่อนที่จะส่งข้อมูลเข้าออก Database เช่น Search, Sort, Hash

import (
	"go_cleanarc/modules/entities"
)

type itemUseCase struct {
	itemRepo entities.ItemsRepository
}

func NewItemUseCase(repo entities.ItemsRepository) entities.ItemsUseCase {
	return &itemUseCase{repo}
}

func (t *itemUseCase) GetAllItems() (res []entities.ItemRes, err error) {
	var items []entities.ItemRes
	handleErr := t.itemRepo.GetAllItems(&items)

	return items, handleErr
}

func (t *itemUseCase) CreateAItem(req *entities.ItemReq) (err error) {
	handleErr := t.itemRepo.CreateAItem(req)

	return handleErr
}

func (t *itemUseCase) GetAItem(res *entities.ItemRes, id string) (err error) {
	handleErr := t.itemRepo.GetAItem(res, id)

	return handleErr
}

func (t *itemUseCase) UpdateAItem(req *entities.ItemReq, id string) (err error) {
	handleErr := t.itemRepo.UpdateAItem(req, id)

	return handleErr
}

func (t *itemUseCase) DeleteAItem(rows *entities.RowsAffected, id string) (err error) {
	handleErr := t.itemRepo.DeleteAItem(rows, id)

	return handleErr
}
