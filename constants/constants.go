package constants

import "time"

const HttpPort string = "localhost:8080"
const CookieExpiryTime time.Duration = time.Duration(time.Hour * 24 * 7) // 7 days
