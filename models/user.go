package models

type User struct {
	UserName     *string `db:"user_name"`
	Role         int     `db:"role"`
	Mail         *string `db:"mail"`
	Salt         []byte  `db:"salt"`
	PasswordHash []byte  `db:"passwd_hash"`
}

func NewUser(login *string, mail *string, salt []byte, hash []byte) *User {
	return &User{
		UserName:     login,
		Role:         0,
		Salt:         salt,
		Mail:         mail,
		PasswordHash: hash,
	}
}

type RegisterRequest struct {
	Login        string `json:"login"`
	Password     string `json:"password"`
	PasswordConf string `json:"password_conf"`
	Email        string `json:"email"`
}

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserResponse struct {
	Id       int    `json:"id"`
	UserName string `json:"user_name"`
}

//create table if not exists  Users(
// 	user_name varchar(255) not null,
// 	passwd_hash BYTEA NOT NULL,
// 	role integer,
// 	mail varchar(255),
//  salt BYTEA NOT NULL,
//	UNIQUE(user_name)
// );
