package model

type Product struct {
	ProductID            string  `json: "productId" gorm:"column:product_id"`
	ProductName          string  `json: "productName" gorm:"column:product_name"`
	ProductIntro         string  `json: "productIntro" gorm:"column:product_intro"`
	CategoryID           string  `json: "categoryId" gorm:"column:category_id"`
	ProductCoverImg      string  `json: "productCoverImg" gorm:"column:product_cover_img"`
	ProductBanner        string  `json: "productId" gorm:"column:product_banner"`
	OriginalPrice        float32 `json: "productPrice" gorm:"column:original_price"`
	SellingPrice         string  `json: "productBanner" gorm:"column:selling_price"`
	StockNum             string  `json: "stockNum" gorm:"column:stock_num"`
	Tag                  string  `json: "tag" gorm:"column:tag"`
	SellStatus           string  `json: "sellStatus" gorm:"column:sell_status"`
	CreateUser           string  `json: "createUser" gorm:"column:create_user"`
	UpdateUser           string  `json: "updateUser" gorm:"column:update_user"`
	ProductDetailContent string  `json: "ProductDetailContent" gorm:"column:product_detail_content"`
	isDeleted            string  `json: "isDeleted" gorm:"column:is_deleted"`
}
