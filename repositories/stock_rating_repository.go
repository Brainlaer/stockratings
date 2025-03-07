package repositories

import (
	"database/sql"
	"example/hello/models"
	"time"
)

type StockRatingRepository struct {
	DB *sql.DB
}

func NewStockRatingRepository(db *sql.DB) *StockRatingRepository {
	return &StockRatingRepository{DB: db}
}

func(r *StockRatingRepository) GetAll()([]models.StockRatingGet, error){
	rows, err := r.DB.Query("SELECT * FROM stock_ratings")
	
	if err !=nil{
		return nil, err
	}

	defer rows.Close()
	 var stockRatings []models.StockRatingGet
	 for rows.Next(){
		var stockRating models.StockRatingGet
		if err:=rows.Scan(&stockRating.ID, &stockRating.Ticker, &stockRating.Target_from, &stockRating.Target_to, &stockRating.Rating_from, & stockRating.Rating_to, &stockRating.Action, &stockRating.Brokerage, &stockRating.Company, &stockRating.Time); err !=nil{
			return nil, err
		}
		stockRatings=append(stockRatings, stockRating)
	 }
	return stockRatings, nil
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

func (r *StockRatingRepository) Delete(id string)error{
	_, err:=r.DB.Exec("DELETE FROM stock_ratings WHERE id= $1 ::UUID",id)
	if err != nil{
		return err
	}
	return nil
}