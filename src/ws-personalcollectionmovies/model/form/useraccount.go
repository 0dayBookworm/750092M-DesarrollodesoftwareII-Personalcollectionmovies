package form

// Useraccount 
type Useraccount struct {
	Username   string `form:"Username"`
	Email      string `form:"Email"`
	Password   string `form:"Password"`
	FirstName  string `form:"FirstName"`
	SecondName string `form:"SecondName"`
	LastName   string `form:"LastName"`
	BirthDate  string  `form:"BirthDate"`
	Gender     string  `form:"Gender"`
}

func UseraccountToString(Useraccount pUseraccount) {
    useraccount := "Username:"+useraccount.Username+", Email:"+useraccount.Email+", Password:"+useraccount.Password+", FirstName:"+useraccount.FirstName+", SecondName:"+useraccount.SecondName+", LastName:"+useraccount.LastName+", BirthDate:"+useraccount.BirthDate+", Gender:"+useraccount.Gender
	return useraccount
}