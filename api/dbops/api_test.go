package dbops

import (
    "testing"
)

func clearTables() {
    dbConn.Exec("truncate users")
    dbConn.Exec("truncate video_info")
    dbConn.Exec("truncate comments")
    dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M) {
    clearTables()
    m.Run()
    clearTables()
}

func TestUserWorkFlow(t *testing.T) {
    t.Run("Add", testAddUser)
    t.Run("GET", testGetUser)
    t.Run("DEL", testDeleteUser)
    t.Run("Reget", testRegetUser)
}

func testAddUser(t *testing.T) {
    err := AddUserCredential("lbh", "123")
    if err != nil {
        t.Errorf("Error of AddUser: %v", err)
    }
}

func testGetUser(t *testing.T) {
    pwd, err := GetUserCredential("lbh")
    if pwd != "123" || err != nil {
        t.Errorf("Error of GetUser")
    }
}

func testDeleteUser(t *testing.T) {
    err := DeleteUser("lbh", "123")
    if err != nil {
        t.Errorf("Error of DeleteUser: %v", err)
    }
}

func testRegetUser(t *testing.T) {
    pwd, err := GetUserCredential("lbh")
    if err != nil {
        t.Errorf("Error of RegetUser %v", err)
    }
    if pwd != "" {
        t.Errorf("Deleting user test failed")
    }
}
