package main

import (
	"fmt"
	"github.com/skorolevskiy/fakeproducts-api-client/api-client"
	"log"
	"time"
)

func main() {
	storeClient, err := api_client.NewClient(time.Second * 10)

	if err != nil {
		log.Fatal(err)
	}

	/* Make request for get all Products */
	//assets, err := storeClient.GetProducts()
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//for _, asset := range assets {
	//	fmt.Println(asset.Info())
	//}

	/* Make request for get one Product for ID */
	//asset, err := storeClient.GetProduct(20)
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(asset.Info())

	/* Make request for get all Categories */
	//categories, err := storeClient.GetCategories()
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//for _, category := range categories {
	//	fmt.Println(category)
	//}

	/* Make request for Products from Category */
	//assets, err := storeClient.GetProductsFromCategory("jewelery")
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//for _, asset := range assets {
	//	fmt.Println(asset.Info())
	//}

	/* Make request for Carts */
	assets, err := storeClient.GetCarts()

	if err != nil {
		log.Fatal(err)
	}

	for _, asset := range assets {
		fmt.Println(asset)
	}
}
