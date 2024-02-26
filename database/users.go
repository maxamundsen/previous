package database

type User struct {
	Id        int
	Email     string
	Firstname string
	Lastname  string
	Password  string
}

type Claim struct {
	Id     int
	UserId int
	Key    string
	Value  string
}

func FetchUserByEmail(email string) (User, error) {
	var user User

	sql := "SELECT * FROM users WHERE email = ?"

	stmt, err := db.Prepare(sql)

	if err != nil {
		return user, err
	}

	queryErr := stmt.QueryRow(email).Scan(&user.Id, &user.Email, &user.Firstname, &user.Lastname, &user.Password)

	if queryErr != nil {
		return user, queryErr
	}

	return user, nil
}

func InsertClaim(claim Claim) error {
	sql := `INSERT INTO claims
			(userid, claim, value)
			VALUES(?, ?, ?)
			`

	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}

	_, execErr := stmt.Exec(claim.UserId, claim.Key, claim.Value)
	if execErr != nil {
		return execErr
	}

	return nil
}

func UpdateClaim(claim Claim) error {
	sql := `UPDATE claims
			SET claim = ?,
			value = ?
			WHERE id = ?
			`

	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}

	_, execErr := stmt.Exec(claim.Key, claim.Value, claim.Id)
	if execErr != nil {
		return execErr
	}

	return nil
}

func FetchUserClaimsById(userid int) ([]Claim, error) {
	claims := make([]Claim, 0)

	sql := "SELECT id, userid, claim, value FROM claims WHERE userid = ?"

	stmt, err := db.Prepare(sql)

	if err != nil {
		return claims, err
	}

	rows, queryErr := stmt.Query(userid)

	if queryErr != nil {
		return claims, queryErr
	}

	defer rows.Close()

	for rows.Next() {
		var claim Claim
		rows.Scan(&claim.Id, &claim.UserId, &claim.Key, &claim.Value)
		claims = append(claims, claim)
	}

	return claims, nil
}

func FetchUserClaimsByIdAsMap(userid int) (map[string]string, error) {
	claims := make(map[string]string)

	sql := "SELECT claim, value FROM claims WHERE userid = ?"

	stmt, err := db.Prepare(sql)

	if err != nil {
		return claims, err
	}

	rows, queryErr := stmt.Query(userid)

	if queryErr != nil {
		return claims, queryErr
	}

	defer rows.Close()

	for rows.Next() {
		var claimKey, claimValue string
		rows.Scan(&claimKey, &claimValue)
		claims[claimKey] = claimValue
	}

	return claims, nil
}

func FetchUserClaimsByEmail(email string) (map[string]string, error) {
	claims := make(map[string]string)

	sql := "SELECT claim, value FROM claims INNER JOIN users ON claims.userid = users.id WHERE users.email = ?"

	stmt, err := db.Prepare(sql)

	if err != nil {
		return claims, err
	}

	rows, queryErr := stmt.Query(email)

	if queryErr != nil {
		return claims, queryErr
	}

	defer rows.Close()

	for rows.Next() {
		var claimKey, claimValue string
		rows.Scan(&claimKey, &claimValue)
		claims[claimKey] = claimValue
	}

	return claims, nil
}

func FetchSpecificClaim(userId int, key string) (Claim, error) {
	var claim Claim

	sql := "SELECT id, userid, claim, value FROM claims WHERE userid = ? AND claim = ?"

	stmt, err := db.Prepare(sql)

	if err != nil {
		return claim, err
	}

	queryErr := stmt.QueryRow(userId, key).Scan(&claim.Id, &claim.UserId, &claim.Key, &claim.Value)

	if queryErr != nil {
		return claim, queryErr
	}

	return claim, nil
}

func DeleteClaim(claim Claim) error {
	sql := "DELETE FROM claims WHERE userid = ? AND claim = ?"

	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}

	_, execErr := stmt.Exec(claim.UserId, claim.Key)

	if execErr != nil {
		return execErr
	}

	return nil
}

func DeleteAllClaimsByUserId(userid int) error {
	sql := "DELETE FROM claims WHERE userid = ?"

	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}

	_, execErr := stmt.Exec(userid)
	if execErr != nil {
		return execErr
	}

	return nil
}

func FetchAllUsers() ([]User, error) {
	users := make([]User, 0)

	sql := "SELECT id, email, firstname, lastname, password FROM users"

	rows, err := db.Query(sql)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user User
		rows.Scan(&user.Id, &user.Email, &user.Firstname, &user.Lastname, &user.Password)

		users = append(users, user)
	}

	return users, nil
}

func UpdateUser(user User) error {
	sql := `UPDATE users
			SET email = ?,
			firstname = ?,
			lastname = ?,
			password = ?
			WHERE id = ?`

	stmt, err := db.Prepare(sql)

	if err != nil {
		return err
	}

	_, execErr := stmt.Exec(user.Email, user.Firstname, user.Lastname, user.Password, user.Id)

	if execErr != nil {
		return err
	}

	return nil
}

func InsertUser(user User) error {
	sql := `INSERT INTO users
			(email, firstname, lastname, password)
			VALUES (?, ?, ?, ?)`

	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}

	_, execErr := stmt.Exec(user.Email, user.Firstname, user.Lastname, user.Password)
	if execErr != nil {
		return execErr
	}

	return nil
}

func DeleteUser(userid int) error {
	sql := `DELETE FROM users
			WHERE id = ?`

	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}

	_, execErr := stmt.Exec(userid)
	if execErr != nil {
		return execErr
	}

	return nil
}