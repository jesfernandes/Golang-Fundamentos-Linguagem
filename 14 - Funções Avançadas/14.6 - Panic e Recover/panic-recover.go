package main

import "fmt"

func recoverExecution() {
	if r := recover(); r != nil {
		fmt.Println("Successfully recovered execution!")
	}
}

func mediaStudent(n1, n2 float32) bool {
	defer recoverExecution()
	media := (n1 + n2) / 2
	if media < 2 {
		panic("very low grade, reinforcement of disciplines is necessary")
	} else if media > 6 {
		return true
	}
	return false
}

func verifyStudentApproval(n1, n2 float32) bool {
	defer fmt.Println("Complete average calculation.")
	fmt.Println("Entering the function to check if the student is approved:")
	return mediaStudent(n1, n2)
}

func main() {

	fmt.Println("Student approval Result : ", verifyStudentApproval(1, 0))
	//o programa é finalizado após o resultado ser do tipo panic
	// pode ser usado a funcao recover, mesmo sem usar a funcao panic
	//Panic nao é a melhor maneira de se tratar erros
	fmt.Println("Program finished")

}
