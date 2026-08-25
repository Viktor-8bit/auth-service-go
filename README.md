
### auth-service-go

Микросервис на go для создания jwt refresh & access токенов

### Таблицы

#### Table "public.users"

```
   Column    |          Type          | Collation | Nullable |              Default              
-------------+------------------------+-----------+----------+-----------------------------------
 user_name   | character varying(255) |           | not null | 
 passwd_hash | bytea                  |           | not null | 
 role        | integer                |           |          | 
 mail        | character varying(255) |           |          | 
 salt        | bytea                  |           | not null | 
 id          | bigint                 |           | not null | nextval('users_id_seq')
Indexes:
    "users_pkey" PRIMARY KEY, btree (id)
    "users_user_name_key" UNIQUE CONSTRAINT, btree (user_name)
Referenced by:
    TABLE "refreshtokens" CONSTRAINT "refreshtokens_user_id_fkey"
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
```
___________________________

#### Table "public.refreshtokens"

```
   Column   |           Type           | Collation | Nullable |                  Default         
------------+--------------------------+-----------+----------+----------------------------------
 user_id    | integer                  |           | not null | 
 iat        | timestamp with time zone |           | not null | 
 exp        | timestamp with time zone |           | not null | 
 revoked_at | timestamp with time zone |           |          | 
 created_at | timestamp with time zone |           | not null | now()
 jti        | bigint                   |           | not null | nextval('refreshtokens_jti_seq')
Indexes:
    "refreshtokens_pkey" PRIMARY KEY, btree (jti)
Foreign-key constraints:
    "refreshtokens_user_id_fkey" FOREIGN KEY (user_id)
    REFERENCES users(id) ON DELETE CASCADE
```

