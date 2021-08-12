package user

type User struct {
	UID             string `json:"uid"`
	Email           string `json:"email"`
	Name            string `json:"name"`
	Username        string `json:"username"`
	Avatar          string `json:"avatar"`
	PlatformVersion int    `json:"platformVersion"`
	Billing         struct {
		Plan        string      `json:"plan"`
		Period      interface{} `json:"period"`
		Trial       interface{} `json:"trial"`
		Cancelation interface{} `json:"cancelation"`
		Addons      interface{} `json:"addons"`
	} `json:"billing"`
	Bio      string `json:"bio"`
	Website  string `json:"website"`
	Profiles []struct {
		Service string `json:"service"`
		Link    string `json:"link"`
	} `json:"profiles"`
}
