package audit

var (
	ActionAccountLogin    = "account:login"
	ActionAccountRegister = "account:register"
	ActionAccountDelete   = "account:delete"

	ActionRequestCreate = "request:create"
	ActionRequestUpdate = "request:update"
	ActionRequestDelete = "request:delete"

	ActionJobCreate = "job:create"
	ActionJobUpdate = "job:update"
	ActionJobDelete = "job:delete"
	ActionJobCancel = "job:cancel"

	ActionCommodityCreate = "commodity:create"
	ActionCommodityUpdate = "commodity:update"
	ActionCommodityDelete = "commodity:delete"

	ActionSessionCreate = "session:create"
	ActionSessionDelete = "session:delete"

	ActionShopCreate = "shop:create"
	ActionShopUpdate = "shop:update"
	ActionShopDelete = "shop:delete"
)
