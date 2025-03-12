package repositories

import (
	"database/sql"
	"example/hello/models"
	"fmt"
	"strings"
	"time"
)

type StockRatingRepository struct {
	DB *sql.DB
}

func NewStockRatingRepository(db *sql.DB) *StockRatingRepository {
	return &StockRatingRepository{DB: db}
}

func(r *StockRatingRepository) GetAll(filters map[string]string, sortBy []string, order []string, limit int, offset int)([]models.StockRatingGet, int, error){
	query := `SELECT *,
    ((target_to - target_from) / target_from) * 100 AS growth FROM stock_ratings`
	args := []interface{}{}
	whereClauses := []string{}

	i := 1
	for key, value := range filters {
		whereClauses = append(whereClauses, fmt.Sprintf("%s ILIKE $%d", key, i))
		args = append(args, "%"+value+"%")
		i++
	}
	countQuery := "SELECT COUNT(*) FROM stock_ratings"

	if len(whereClauses) > 0 {
		query += " WHERE " + strings.Join(whereClauses, " OR ")
		countQuery += " WHERE " + strings.Join(whereClauses, " OR ")

	}

	fmt.Println(countQuery)

	var totalRecords int
	err := r.DB.QueryRow(countQuery, args...).Scan(&totalRecords)
	if err != nil {
		return nil,0, err
	}

	validSortFields := map[string]bool{
		"ticker": true, "company": true, "action": true,
		"brokerage": true, "rating_from": true, "rating_to": true,
		 "time": true, "target_from": true, "target_to": true, "growth":true,
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
		return nil,0, err
	}

	defer rows.Close()

	var stocks []models.StockRatingGet
	for rows.Next() {
		var stock models.StockRatingGet
		if err := rows.Scan(
			&stock.ID, &stock.Ticker, &stock.Target_from, &stock.Target_to,
			&stock.Company, &stock.Action, &stock.Brokerage, &stock.Rating_from,
			&stock.Rating_to, &stock.Time,&stock.Growth,
		); err != nil {
			return nil,0, err
		}
		stocks = append(stocks, stock)
	}

	return stocks, totalRecords, nil

}

func(r *StockRatingRepository) GetOne(id string)(*models.StockRatingGet, error){
	row := r.DB.QueryRow("SELECT * FROM stock_ratings WHERE id= $1 ::UUID",id)

		var stockRating models.StockRatingGet
		if err:=row.Scan(&stockRating.ID,&stockRating.Ticker, &stockRating.Target_from, &stockRating.Target_to, &stockRating.Company, &stockRating.Action, &stockRating.Brokerage, &stockRating.Rating_from, & stockRating.Rating_to, &stockRating.Time); err !=nil{
			return nil, err
		}
	 
	return &stockRating, nil
}

func (r *StockRatingRepository) Create(stockRating *models.StockRatingCreate, time *time.Time)error{
	_, err:=r.DB.Exec("INSERT INTO stock_ratings"+
	"(ticker, target_from, target_to, company, action, brokerage, rating_from, rating_to, time)"+ 
	"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)", 
	&stockRating.Ticker, &stockRating.Target_from, &stockRating.Target_to, &stockRating.Company, &stockRating.Action, &stockRating.Brokerage, &stockRating.Rating_from, &stockRating.Rating_to, &time)

	if err != nil{
		return err
	}
	return nil
}

func (r *StockRatingRepository) Update(stock *models.StockRatingGet) error {
	query := `
		UPDATE stock_ratings 
		SET ticker = $1, target_from = $2, target_to = $3, company = $4, action = $5, 
		    brokerage = $6, rating_from = $7, rating_to = $8, time = $9 
		WHERE id = $10`
	
	_, err := r.DB.Exec(query, 
		stock.Ticker, stock.Target_from, stock.Target_to, 
		stock.Company, stock.Action, stock.Brokerage, 
		stock.Rating_from, stock.Rating_to, stock.Time, 
		stock.ID,
	)
	if err != nil {
		return err
	}
	return nil
}


func (r *StockRatingRepository) Delete(id string)error{
	_, err:=r.DB.Exec("DELETE FROM stock_ratings WHERE id= $1 ::UUID",id)
	if err != nil{
		return err
	}
	return nil
}