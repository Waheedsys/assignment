package stores

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Waheedsys/entities/entities"
	"testing"
)

func Test_GetUsers(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock db: %v", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"name", "age", "phone_number", "Email"}).
		AddRow("waheed", 23, "+123456789", "waheed@123.com")
	mock.ExpectQuery("SELECT UserName,UserAge,Phone_number,Email FROM User").
		WillReturnRows(rows)

	store := NewDetails(db)
	result, err := store.GetUsers()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(result) != 1 {
		t.Fatalf("expected 2 users, got %d", len(result))
	}
	if len(result) < 1 {
		t.Errorf("expected length of result 1 but got %v", len(result))
	}

	if result[0].UserName != "waheed" {
		t.Errorf("expected name waheed but got %v", result[0].UserName)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}

// test update
func Test_GetUsersByName(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("failed to create mock db: %v", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT UserName, UserAge, Phone_number, Email FROM User WHERE Username = ?").
		WithArgs("waheed1").
		WillReturnRows(sqlmock.NewRows([]string{"UserName", "UserAge", "Phone_number", "Email"}).
			AddRow("waheed1", 33, "+123456089", "waheed@123.com"))

	store := NewDetails(db)
	result, err := store.GetUsersByName("waheed1")
	if err != nil {
		t.Errorf("error while getting user")
	}
	if result.UserName != "waheed1" {
		t.Errorf("expected name waheed but got %v", result.UserName)
	}
	if result.Email != "waheed@123.com" {
		t.Errorf("expected email waheed@123.com but got %v", result.Email)
	}
	if result.Phone_number != "+123456089" {
		t.Errorf("expected number +1230056789 but got %v", result.Phone_number)
	}
	if result.UserAge != 33 {
		t.Errorf("expected age 23 but got %v", result.UserAge)
	}
	// non existing user
	mock.ExpectQuery("SELECT UserName, UserAge, Phone_number, Email FROM User WHERE Username = ?").
		WithArgs("non_existing_user").
		WillReturnError(sql.ErrNoRows)

	_, err = store.GetUsersByName("non_existing_user")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	expectedErr := "user with username 'non_existing_user' not found"
	if err.Error() != expectedErr {
		t.Errorf("expected error '%s', got '%s'", expectedErr, err.Error())
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unmet expectations: %v", err)
	}
}

func Test_AddUsers(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("failed to create mock db: %v", err)
	}
	defer db.Close()
	testbook := &entities.Users{
		"waheed1",
		33,
		"+123456089",
		"waheed@123.com",
	}

	mock.ExpectExec("INSERT INTO User (UserName, UserAge, Phone_number, Email) VALUES (?, ?, ?, ?)").
		WithArgs(testbook.UserName, testbook.UserAge, testbook.Phone_number, testbook.Email).
		WillReturnResult(sqlmock.NewResult(1, 1))

	store := NewDetails(db)
	err = store.AddUsers(testbook)
	if err != nil {
		t.Errorf("error while adding the user")
	}
	//mock.ExpectExec("INSERT INTO User").
	//	WithArgs(testbook.UserName, testbook.UserAge, testbook.Phone_number, testbook.Email).
	//	WillReturnError(fmt.Errorf("'ExecQuery: actual sql: INSERT INTO User (UserName, UserAge, Phone_number, Email) VALUES (?, ?, ?, ?) does not equal to expected INSERT INTO User'"))
	//userStore := NewDetails(db)
	//
	//err = userStore.AddUsers(testbook)
	//if err == nil {
	//	t.Fatalf("expected error, got nil")
	//}
	//
	//expectedErr := "database error"
	//if err.Error() != expectedErr {
	//	t.Errorf("expected error '%s', got '%s'", expectedErr, err.Error())
	//}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}

func Test_DeleteUsers(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock db: %v", err)
	}
	defer db.Close()

	mock.ExpectExec("DELETE FROM User WHERE UserName = ?").
		WithArgs("waheed").
		WillReturnResult(sqlmock.NewResult(0, 1))
	store := NewDetails(db)
	err = store.DeleteUsers("waheed")
	if err != nil {
		t.Errorf("error while deleting")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}

// test update
func Test_Updatebook(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("failed to create mock db: %v", err)
	}
	defer db.Close()
	testbook := &entities.Users{
		"waheed",
		23,
		"+1230056789",
		"waheed@123.com",
	}

	mock.ExpectExec("UPDATE User SET Email = ? WHERE UserName = ?").
		WithArgs(testbook.Email, testbook.UserName).
		WillReturnResult(sqlmock.NewResult(0, 1))
	store := NewDetails(db)

	err = store.UpdateUsers(testbook.UserName, testbook)
	if err != nil {
		t.Errorf("error while updating user")
	}
}
