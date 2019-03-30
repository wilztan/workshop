package models

import "time"

type Product struct {
    ID          int64     `json:"id"`
    ProductName string    `json:"product_name"`
	ProductDesc string    `json:"product_description"`
	CreateTime  time.Time `json:"create_time"`
	UpdateTime time.Time  `json:"update_time"`
    // Stats        Statistic `json:"statistic"`
}

// type Statistic struct {
//     View int64 `json:"view"`
// }

type News struct {
    Status       string    `json:"status"`
    TotalResults int    `json:"totalResults"`
    Articles     []Article `json:"articles"`
}

type Article struct {
    Author string `json:"author"`
    Title  string `json:"title"`
}