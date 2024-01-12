package constants

import "time"

const CookieExpiryTime time.Duration = time.Duration(time.Hour * 24 * 7) // 7 days
const CookieEntropy int = 33
