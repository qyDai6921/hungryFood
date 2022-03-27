package define

type OrderReq struct {
	Zipcode     string     `json:"zip_code"`
	DeliverTime int        `json:"deliver_time"`
	Address     string     `json:"address"`
	ChefId      string     `json:"chef_id"`
	Price       int        `json:"price"`
	Tip         int        `json:"tip"`
	Tax         int        `json:"tax"`
	Total       int        `json:"total"`
	MenuItems   []MenuItem `json:"menu_items"`
	Username    string     `json:"username"`
	Phone       string     `json:"phone"`
	Remark      string     `json:"remark"`
}
