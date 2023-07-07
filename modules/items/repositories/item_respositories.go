package respositories

//มีหน้าที่ ในการรับส่ง Entities เข้าออกจาก Database หรือพูดง่ายๆ ก็คือมีหน้าที่ Query ข้อมูลจาก Database นั่นแหละ
// SQL Query

import (
	"database/sql"
	"fmt"
	"go_cleanarc/modules/entities"
	"log"
)

type itemRepository struct {
	conn *sql.DB
}

func NewItemRepository(db *sql.DB) entities.ItemsRepository {
	return &itemRepository{db}
}

func (t *itemRepository) GetAllItems(items *[]entities.ItemRes) (err error) {
	var item entities.ItemRes

	sql_query := `SELECT * FROM items`

	stmt, err := t.conn.Prepare(sql_query)
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("%s %s", err, "when Prepare"))
	}

	theRows, err := stmt.Query()
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("%s %s", err, "when Query"))
	}
	for theRows.Next() {
		err := theRows.Scan(&item.Id, &item.Item, &item.Detail, &item.Createdatetime)
		*items = append(*items, item)

		if err != nil {
			return fmt.Errorf(fmt.Sprintf("%s %s", err, "when Scan"))
		}
	}
	err = theRows.Err()
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("%s %s", err, "when theRows"))
	}

	defer stmt.Close()
	return nil
}

func (t *itemRepository) CreateAItem(req *entities.ItemReq) (err error) {

	var itemDuplicate entities.ItemRes

	theRows := t.conn.QueryRow("SELECT * FROM items WHERE item = ?;", req.Item)
	if theRows != nil {
		theRows.Scan(&itemDuplicate.Id, &itemDuplicate.Item, &itemDuplicate.Detail, &itemDuplicate.Createdatetime)
		if itemDuplicate.Id != 0 {
			return fmt.Errorf("duplicate item")
		}
	}

	query := `INSERT INTO items(item, detail) VALUES (?, ?);`

	// Query part
	stmt, err := t.conn.Prepare(query)
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("%s %s", err, "when Prepare"))
	}

	rs, err := stmt.Exec(req.Item, req.Detail)
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("%s %s", err, "when Exec"))
	}
	id, _ := rs.LastInsertId()
	req.Id = id

	return nil
}

func (t *itemRepository) GetAItem(res *entities.ItemRes, id string) (err error) {

	sql_query := `SELECT * FROM items WHERE id = ?`

	stmt, err := t.conn.Prepare(sql_query)
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("%s %s", err, "when Prepare"))
	}

	scan_err := stmt.QueryRow(id).Scan(&res.Id, &res.Item, &res.Detail, &res.Createdatetime)
	if scan_err != nil {
		return fmt.Errorf(fmt.Sprintf("%s %s", scan_err, "item not found"))
	}

	defer stmt.Close()
	return nil
}

func (t *itemRepository) UpdateAItem(req *entities.ItemReq, id string) (err error) {
	query := `UPDATE items SET item=?, detail=? WHERE id=?`

	// Query part
	stmt, err := t.conn.Prepare(query)
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("%s %s", err, "when Prepare"))
	}

	rs, err := stmt.Exec(req.Item, req.Detail, id)
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("%s %s", err, "when Exec"))
	}
	rows, _ := rs.RowsAffected()

	if rows > 0 {
		sql := `SELECT id, item, detail FROM items WHERE id = ?`

		stmt_, err := t.conn.Prepare(sql)
		if err != nil {
			return fmt.Errorf(fmt.Sprintf("%s %s", err, "when Prepare"))
		}

		scan_err := stmt_.QueryRow(id).Scan(&req.Id, &req.Item, &req.Detail)
		if scan_err != nil {
			return fmt.Errorf(fmt.Sprintf("%s %s", scan_err, "item not found"))
		}
	}
	return nil
}

func (t *itemRepository) DeleteAItem(rows *entities.RowsAffected, id string) (err error) {

	sql_query := `DELETE FROM items WHERE id = ?`

	stmt, err := t.conn.Prepare(sql_query)
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("%s %s", err, "when Prepare"))
	}
	rs, err := stmt.Exec(id)
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("%s %s", err, "when Exec"))
	}
	rowsAffected, _ := rs.RowsAffected()
	log.Println(rowsAffected)
	rows.NumOfRows = rowsAffected

	defer stmt.Close()
	return nil
}
