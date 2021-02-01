package model

type Product struct {
	ProductId            string  `jsson: "productId" gorm:"column:product_id"`
	ProductName          string  `jsson: "productName" gorm:"column:product_name"`
	ProductIntro         string  `jsson: "productIntro" gorm:"column:product_intro"`
	CategoryId           string  `jsson: "categoryId" gorm:"column:category_id"`
	ProductCoverImg      string  `jsson: "productCoverImg" gorm:"column:product_cover_img"`
	ProductBanner        string  `jsson: "productId" gorm:"column:product_banner"`
	OriginalPrice        float32 `jsson: "productPrice" gorm:"column:original_price"`
	SellingPrice         string  `jsson: "productBanner" gorm:"column:selling_price"`
	StockNum             string  `jsson: "stockNum" gorm:"column:stock_num"`
	Tag                  string  `jsson: "tag" gorm:"column:tag"`
	SellStatus           string  `jsson: "sellStatus" gorm:"column:sell_status"`
	CreateUser           string  `jsson: "createUser" gorm:"column:create_user"`
	UpdateUser           string  `jsson: "updateUser" gorm:"column:update_user"`
	ProductDetailContent string  `jsson: "ProductDetailContent" gorm:"column:product_detail_content"`
	isDeleted            string  `jsson: "isDeleted" gorm:"column:is_deleted"`
}
