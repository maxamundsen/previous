INSERT INTO
    "user" (
        "username",
        "email",
        "firstname",
        "lastname",
        "password",
        "failed_attempts",
        "last_login",
        "security_stamp",
        "permission_admin"
    )
VALUES
    (
        "admin",
        "admin@adminland.org",
        "Example",
        "User",
        "$2a$10$bgBsgWhmTS85fV9nok4a..XcN.Xy4zabvjlUY5dwwLZwdcZ8SSRv.",
        0,
        "",
        "ss",
        1
    ),
    (
        "user",
        "user@example.org",
        "Normal",
        "User",
        "$2a$10$HPIIg27Xqa4zUH8YbqAktu30xn.7Sep0W/AO9cMjgfsoyD1tqbDwO",
        0,
        "",
        "ss",
        0
    );

INSERT INTO
    "order" (
        "product_id",
        "price",
        "purchaser_name",
        "purchaser_email"
    )
VALUES
    (1, 1000, "Max Amundsen", "max@example.com"),
    (1, 1000, "Max Amundsen", "max@example.com"),
    (2, 25000, "John Doe", "john@example.com"),
    (3, 51200, "Jane Doe", "jane@example.com"),
    (3, 51200, "Jane Doe", "jane@example.com"),
    (4, 51200, "Jane Doe", "jane@example.com"),
    (4, 102400, "Max Amundsen", "max@example.com"),
    (
        4,
        102400,
        "Little Bobby Tables",
        "bobby@example.com"
    ),
    (
        4,
        102400,
        "Alice Enwonderland",
        "alice@example.com"
    ),
    (4, 102400, "John Smith", "johnsmith@example.com"),
    (5, 204800, "Jane Doe", "jane@example.com"),
    (5, 204800, "Jane Doe", "jane@example.com"),
    (5, 204800, "Jane Doe", "jane@example.com"),
    (5, 204800, "Jane Doe", "jane@example.com"),
    (5, 204800, "Jane Doe", "jane@example.com"),
    (5, 204800, "Jane Doe", "jane@example.com"),
    (3, 51200, "Jane Doe", "jane@example.com"),
    (3, 51200, "Jane Doe", "jane@example.com"),
    (3, 51200, "Jane Doe", "jane@example.com"),
    (3, 51200, "Jane Doe", "jane@example.com"),
    (3, 51200, "Jane Doe", "jane@example.com"),
    (3, 51200, "Jane Doe", "jane@example.com"),
    (3, 51200, "Jane Doe", "jane@example.com");