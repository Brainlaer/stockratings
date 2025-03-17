package repositories

import (
	"database/sql"
	"example/hello/models"
	"fmt"
	"strings"
)

type StockRepository struct {
	DB *sql.DB
}

func NewStockRepository(db *sql.DB) *StockRepository {
	return &StockRepository{DB: db}
}

func (r *StockRepository) GetAll(filters map[string]interface{}, sortBy []string, order []string, limit int, offset int, queryStrong bool) ([]models.StockResponseGet, int, error) {
	query := `SELECT *,
    ((target_to - target_from) / target_from) * 100 AS growth FROM stock_ratings`
	args := []interface{}{}
	whereClauses := []string{}
	countQuery := "SELECT COUNT(*) FROM stock_ratings"
	var totalRecords int

	i := 1
	for key, value := range filters {
		if key == "rating_to" || key == "rating_from" || key == "action" {
			vals := value.([]string)
			placeholders := []string{}
			for _, v := range vals {
				placeholders = append(placeholders, fmt.Sprintf("LOWER(%s) ILIKE $%d", key, i))
				args = append(args, "%"+strings.ToLower(v)+"%")
				i++
			}
			whereClauses = append(whereClauses, fmt.Sprintf("(%s)", strings.Join(placeholders, " OR ")))
		} else {
			strValue, _ := value.(string)
			whereClauses = append(whereClauses, fmt.Sprintf("%s ILIKE $%d", key, i))
			args = append(args, "%"+strValue+"%")
			i++
		}
	}

	if len(whereClauses) > 0 {
		if !queryStrong{
			query += " WHERE " + strings.Join(whereClauses, " OR ")
			countQuery += " WHERE " + strings.Join(whereClauses, " OR ")
		}else{
			query += " WHERE " + strings.Join(whereClauses, " AND ")
			countQuery += " WHERE " + strings.Join(whereClauses, " AND ")
		}
	}

	err := r.DB.QueryRow(countQuery, args...).Scan(&totalRecords)
	if err != nil {
		return nil, 0, err
	}

	validSortFields := map[string]bool{
		"ticker": true, "company": true, "action": true,
		"brokerage": true, "rating_from": true, "rating_to": true,
		"time": true, "target_from": true, "target_to": true, "growth": true,
	}
	var orderClauses []string

	for idx, column := range sortBy {
		if validSortFields[column] {
			dir := "ASC"
			if idx < len(order) && (order[idx] == "asc" || order[idx] == "desc") {
				dir = strings.ToUpper(order[idx])
			}
			orderClauses = append(orderClauses, fmt.Sprintf("%s %s", column, dir))
		}
	}

	if len(orderClauses) > 0 {
		query += " ORDER BY " + strings.Join(orderClauses, ", ")
	} else {
		query += " ORDER BY time DESC"
	}

	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", i, i+1)
	args = append(args, limit, offset)
	
	rows, err := r.DB.Query(query, args...)

	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()

	var stocks []models.StockResponseGet
	for rows.Next() {
		var stock models.StockResponseGet
		if err := rows.Scan(
			&stock.ID, &stock.Ticker, &stock.Target_from, &stock.Target_to,
			&stock.Company, &stock.Action, &stock.Brokerage, &stock.Rating_from,
			&stock.Rating_to, &stock.Time, &stock.Growth,
		); err != nil {
			return nil, 0, err
		}
		stocks = append(stocks, stock)
	}

	return stocks, totalRecords, nil

}

func (r *StockRepository) GetOne(id string) (*models.StockResponseGet, error) {
	row := r.DB.QueryRow(`SELECT *,((target_to - target_from) / target_from) * 100 AS growth
	FROM stock_ratings WHERE id= $1 ::UUID`, id)

	var stock models.StockResponseGet
	if err := row.Scan(&stock.ID, &stock.Ticker, &stock.Target_from, &stock.Target_to, &stock.Company, &stock.Action, &stock.Brokerage, &stock.Rating_from, &stock.Rating_to, &stock.Time, &stock.Growth); err != nil {
		return nil, err
	}

	return &stock, nil
}

func (r *StockRepository) Create(stock *models.Stock) error {
	_, err := r.DB.Exec("INSERT INTO stock_ratings"+
		"(ticker, target_from, target_to, company, action, brokerage, rating_from, rating_to, time)"+
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		&stock.Ticker, &stock.Target_from, &stock.Target_to, &stock.Company, &stock.Action, &stock.Brokerage, &stock.Rating_from, &stock.Rating_to, &stock.Time)

	if err != nil {
		return err
	}
	return nil
}

func (r *StockRepository) Update(id string, stock *models.Stock) error {
	query := `
		UPDATE stock_ratings 
		SET ticker = $1, target_from = $2, target_to = $3, company = $4, action = $5, 
		    brokerage = $6, rating_from = $7, rating_to = $8, time = $9 
		WHERE id = $10 ::UUID`

	_, err := r.DB.Exec(query,
		stock.Ticker, stock.Target_from, stock.Target_to,
		stock.Company, stock.Action, stock.Brokerage,
		stock.Rating_from, stock.Rating_to, stock.Time,
		id,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *StockRepository) Delete(id string) error {
	_, err := r.DB.Exec("DELETE FROM stock_ratings WHERE id= $1 ::UUID", id)
	if err != nil {
		return err
	}
	return nil
}
