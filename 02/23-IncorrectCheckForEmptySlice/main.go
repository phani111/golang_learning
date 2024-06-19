package main

func main() {
	handleOperations("")
	handleOperationsCorrect("")

}

func handleOperations(id string) {
	operations := getOperations(id)
	if operations != nil {
		handle(operations)
	}
}
