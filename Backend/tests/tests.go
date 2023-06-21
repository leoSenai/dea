package tests

import (
	repositoryUtils "api/repository/utils"
	"fmt"
)

func Run() {

	//Tests is a module where you can run commands to test specific functions
	//or procedures, use and abuse of tests to make sure that the system wil work fine.

	exist1 := repositoryUtils.VerifyUserExistanceByDocument(0)
	exist2 := repositoryUtils.VerifyUserExistanceByDocument(1)

	fmt.Println(exist1, exist2)

}
