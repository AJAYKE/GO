package main

import "fmt"

const IamPublicWithCapitalStart string = "first letter capital => public varibale"

func main() {
	var chicki string = "chicki chicki chicki"
	fmt.Println(chicki)
	fmt.Printf("type of chicki is: %T \n", chicki)

	var isInstaTred bool = false
	fmt.Println(isInstaTred)
	fmt.Printf("type pf isinstatrend: %T \n", isInstaTred)

	var checkint uint16 = 129
	fmt.Println(checkint)
	

	var checkfloat float32 = 124312543412451596808809790787007087087.888888841435434123354223214532
	fmt.Println(checkfloat)
	fmt.Printf("check type of float: %T \n", checkfloat)

	var emptyfloat float32
	fmt.Println(emptyfloat)
	fmt.Printf("type of empty float: %T \n", emptyfloat)

	var implicitDeclaration = "mic testing"
	fmt.Println(implicitDeclaration)

	implicitDeclaration = "changing to testing 123"
	fmt.Println(implicitDeclaration)

	nodeclaration := "yevdra nuvvu"
	fmt.Printf("check type of nodeclaration: %T \n", nodeclaration)

	fmt.Printf("type of publicvar: %T \n", IamPublicWithCapitalStart)
}