package models

import "time"

type RefreshToken struct {
	JTI       *int64     `db:"jti"`
	UserID    int        `db:"user_id"`
	IAT       time.Time  `db:"iat"`
	EXP       time.Time  `db:"exp"`
	RevokedAt *time.Time `db:"revoked_at"`
	CreatedAt time.Time  `db:"created_at"`
}

// CREATE TABLE RefreshTokens (
// 	   jti BIGSERIAL PRIMARY KEY,
//     user_id INTEGER NOT NULL,
//     iat TIMESTAMP WITH TIME ZONE NOT NULL,
//     exp TIMESTAMP WITH TIME ZONE NOT NULL,
//     revoked_at TIMESTAMP WITH TIME ZONE NULL,
//     created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
//     FOREIGN KEY (user_id)
//         REFERENCES users(id)
//         ON DELETE CASCADE
// );

// ALTER TABLE RefreshTokens ADD COLUMN jti BIGSERIAL PRIMARY KEY;
// ALTER TABLE RefreshTokens DROP COLUMN jti;
