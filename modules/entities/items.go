package entities

/*interface*/
type ItemsUseCase interface {
	GetAllItems() (t []ItemRes, err error)
	CreateAItem(req *ItemReq) (err error)
	GetAItem(res *ItemRes, id string) (err error)
	UpdateAItem(req *ItemReq, id string) (err error)
	DeleteAItem(rows *RowsAffected, id string) (err error)
}

type ItemsRepository interface {
	GetAllItems(t *[]ItemRes) (err error)
	CreateAItem(req *ItemReq) (err error)
	GetAItem(res *ItemRes, id string) (err error)
	UpdateAItem(req *ItemReq, id string) (err error)
	DeleteAItem(rows *RowsAffected, id string) (err error)
}

/*model*/
/*type ItemRes struct {
	Id             int64  `json:"id"`
	Item           string `json:"item"`
	Detail         string `json:"detail"`
	Createdatetime string `json:"createdatetime"`
}*/

type ItemRes struct {
	Id             int64  `json:"id"`
	Item           string `json:"item"`
	Detail         string `json:"detail"`
	Createdatetime string `json:"createdatetime"`
}

type ItemReq struct {
	Id     int64  `json:"id"`
	Item   string `json:"item"`
	Detail string `json:"detail"`
}

type RowsAffected struct {
	NumOfRows int64 `json:"num_of_rows"`
}
