// package main

// import "fmt"

// type user struct{}

// func (u user) userMethod(name string) {
// 	fmt.Printf("User method called with name: %s\n", name)
// }

// type userType struct {
// 	user
// }

// func (u userType) userMethod(name string) {
// 	u.user.userMethod(name)
// 	fmt.Println("UserType method called")
// }


// func main(){
// 	user:=userType{}
// 	user.userMethod("example")
// }	



package main
import "fmt"


type Database interface {
	SaveUser(name string)
}

type RealDatabase struct{}

func (r RealDatabase) SaveUser(name string) {
	fmt.Printf("User %s saved to the real database.\n", name)
}


type MockDatabase struct{}
func (m MockDatabase) SaveUser(name string) {
	fmt.Printf("Mock save for user %s.\n", name)
}
type UserService struct {
	DB Database
}

func (u UserService) SaveUser(name string) {
	fmt.Printf("Saving user %s.\n", name)
	u.DB.SaveUser(name)
}
func main() {
	realDB := RealDatabase{}
	mockDB := MockDatabase{}
	userService := UserService{DB: realDB}
	userService.SaveUser("John Doe")
	userService.DB = mockDB
	userService.SaveUser("Jane Doe")
}