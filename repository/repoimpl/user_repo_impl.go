package repoimpl

import (
	"database/sql"
	"fmt"
	models "go-postgres/model"
	repo "go-postgres/repository"
	m "net/mail"
	"strconv"
)

type UserRepoImpl struct {
	Db *sql.DB
}

func NewUserRepo(db *sql.DB) repo.UserRepo {
	return &UserRepoImpl{
		Db: db,
	}
}
func (u *UserRepoImpl) Create(Userst *[]models.User) {
	var name, gender, email string
	var UserId int
	fmt.Println("Enter The User-Id")
	fmt.Scan(&UserId)
	fmt.Println("Enter The Name")
	fmt.Scan(&name)
	fmt.Println("Enter The Gender")
	fmt.Scan(&gender)
	fmt.Println("Enter The email 000000000")
	fmt.Scan(&email)
	User := models.User{ID: UserId, Name: name, Gender: gender, Email: email}
	err := Insert(u, User)
	if err != nil {
		fmt.Println(err)
		return
	}
	*Userst = append(*Userst, User)
}

func (u *UserRepoImpl) Select() {
	users := make([]models.User, 0)
	rows, err := u.Db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		user := models.User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Gender, &user.Email)
		if err != nil {
			break
		}
		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}
	if len(users) == 0 {
		fmt.Println("The Table is Empty")
		return
	}
	fmt.Println(users)
}

func Insert(u *UserRepoImpl, user models.User) error {
	insertStatement := `
	INSERT INTO users (id, name, gender, email)
	VALUES ($1, $2, $3, $4)`
	_, err := u.Db.Exec(insertStatement, user.ID, user.Name, user.Gender, user.Email)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Record added: ", user)
	return nil
}
func (u *UserRepoImpl) Delete() {
	delId := 0
	fmt.Println("Enter the User-ID for Deletion")
	fmt.Scan(&delId)
	deleteStmt := `Delete from users where id=` + strconv.Itoa(delId)
	sqlresult, err := u.Db.Exec(deleteStmt)
	if err != nil {
		fmt.Println(err)
		return
	}
	affectRows, rr := sqlresult.RowsAffected()
	if rr != nil {
		fmt.Print(rr)
		return
	}
	if affectRows == 0 {
		fmt.Println("User-Id does not exist")
		return
	}
	fmt.Println("Record Deleted Success fullly: ")
}
func (u *UserRepoImpl) Update() {
	emailId := ""
	userId := 0
	user := models.User{}
	fmt.Println("Enter Your User-ID ")
	fmt.Scan(&userId)
	rows, err := u.Db.Exec("SELECT * FROM users where id=" + strconv.Itoa(userId))
	if err != nil {
		fmt.Println(err)
		return
	}
	n, _ := rows.RowsAffected()
	if n == 0 {
		fmt.Println("Invalid User Id ")
		return
	}

	fmt.Println("Enter your Updated E-mail")
	fmt.Scan(&emailId)
	mail, Err := u.Db.Query("SELECT * FROM users where id=$1", userId)
	if Err != nil {
		fmt.Println(Err)
	}
	for mail.Next() {
		err := mail.Scan(&user.ID, &user.Name, &user.Gender, &user.Email)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	if emailId == user.Email {
		fmt.Println(" Previous mail & updated mail are same")
		return
	}
	_, err = m.ParseAddress(emailId)
	if err != nil {
		fmt.Println("Invalid Email")
		return
	}

	updateStmt := `update "users" set "email"=$1 where "id"=$2`
	_, e := u.Db.Exec(updateStmt, emailId, userId)
	if e != nil {
		fmt.Println(userId)
		fmt.Println(e)
		return
	}
	fmt.Println("Successfully Updated")
}
