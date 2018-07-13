package example

import "fmt"

/* Repository will store any Database handler.
Querying, or Creating/ Inserting into any database will stored here.
This layer will act for CRUD to database only.
No business process happen here. Only plain function to Database.
*/

func updateSomething() {

	fmt.Println("UPDATED")
}
