CREATE TABLE IF NOT EXISTS "users" (
    "id" INTEGER NOT NULL,
    "username" TEXT NOT NULL,
    "email" TEXT NOT NULL,
    "firstname" TEXT NOT NULL,
    "lastname" TEXT NOT NULL,
    "password" TEXT NOT NULL,
    "failed_attempts" INTEGER NOT NULL DEFAULT 0,
    "security_stamp" TEXT NOT NULL,
    "last_login" TEXT NOT NULL,
    "permission_admin" INTEGER NOT NULL DEFAULT 0,
    PRIMARY KEY("id")
);

CREATE TABLE IF NOT EXISTS "orders" (
    "id" INTEGER NOT NULL,
    "product_id" INTEGER NOT NULL,
    "price" INTEGER NOT NULL,
    "purchaser_name" TEXT NOT NULL,
    "purchaser_email" TEXT NOT NULL,
    PRIMARY KEY("id")
);