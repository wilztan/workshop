package resource

import(
	"fmt"
	"github.com/wilztan/workshop/models"
	"github.com/wilztan/workshop/db"
)

func GetAllProduct(limit, offset int64,search string) (products []models.Product, err error){
	aditionalQuery:=""
	if len(search)>0{
		aditionalQuery = `WHERE LOWER(product_name) LIKE '%`+search+`%'`
	}
	query:= fmt.Sprintf(`
		SELECT * FROM ws_product %s LIMIT %d OFFSET %d 
	`,aditionalQuery, limit,offset)
	
	rows, err :=db.Client.Query(query)
	if err!=nil{
		fmt.Println(err)
	}

	defer rows.Close()

	for rows.Next(){
		p:=models.Product{}
		errScan :=rows.Scan(
			&p.ID,
			&p.ProductName,
			&p.ProductDesc,
			&p.CreateTime,
			&p.UpdateTime,
		)
		if errScan!=nil{
			continue
		}
		products= append(products,p)
	}



	return
}

func CreateProduct(name,desc string) (p models.Product, err error){

	query := fmt.Sprintf(`INSERT INTO ws_product (product_name,product_description,create_time, update_time) VALUES('%s','%s',NOW(),NOW()) RETURNING id,product_name,product_description,create_time, update_time`, name, desc)

	err= db.Client.QueryRow(query).Scan(&p.ID,&p.ProductName, &p.ProductDesc, &p.CreateTime, &p.UpdateTime)
		
	return
}