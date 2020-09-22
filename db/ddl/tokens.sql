DROP TABLE if exists tokens ;
CREATE TABLE tokens (
    user_id varchar(255) primary key,
    os_type int(1) NOT NULL,
    token varchar(255),
    created_at DATETIME NOT NULL,
    updated_at DATETIME
)
