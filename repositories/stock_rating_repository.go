package repositories

import (
	"database/sql"
	"example/hello/models"
	"log"
)

type StockRatingRepository struct {
	DB *sql.DB
}

func NewStockRatingRepository(db *sql.DB) *StockRatingRepository {
	return &StockRatingRepository{DB: db}
}

func(r *StockRatingRepository) GetAll()([]models.StockRating, error){
	rows, err := r.DB.Query("SELECT * FROM stock_ratings")
	
	if err !=nil{
		return nil, err
	}

	defer rows.Close()
	 var stockRatings []models.StockRating
	 for rows.Next(){
		var stockRating models.StockRating
		if err:=rows.Scan(&stockRating.ID, &stockRating.Ticker, &stockRating.Target_from, &stockRating.Target_to, &stockRating.Rating_from, & stockRating.Rating_to, &stockRating.Action, &stockRating.Brokerage, &stockRating.Company); err !=nil{
			log.Fatal("Erro al escanear fila:", err)
			continue
		}
		stockRatings=append(stockRatings, stockRating)
	 }
	return stockRatings, nil
}