package main

import (
	"fmt"
	"errors"
	// "reflect"
)

func main() {
	name := "Marcelo";
	age := 20;
	_, _ = name, age;

	arr2 := []int{21, 23, 90, 12093, 912930, 19239};
	arr2 = append(arr2, 9090);

	mapp := make(map[string]string);
	mapp["name"]  = "Marcelo Rafael de Brito";
	mapp["age"]   = "90";
	mapp["city"]  = "Belo Horizonte";
	mapp["state"] = "MG";

	delete(mapp, "city");

	// for k, v := range mapp {
	// 	fmt.Println(k, v);
	// }
	r, err := broadcast(100);
	if r == true {
		fmt.Println("Enviado com sucesso");
	} else {
		fmt.Println(err);
	}
	fmt.Println("___________________________________");

	j := person{height:170,weight:90,skin:"black"}
	fmt.Println(j);

	fmt.Println("___________________________________");

	status := "initial";
	name2  := "Marcelo";
	pass   := "123";
	login(&status, name2, pass);
	fmt.Println(status, name2, pass);

	for i := 0; i <= 9000; i++ {
		fmt.Println(i);
	}
}
func login(status *string, name string, pass string) {
	if name == "Marcelo" && pass == "123" {
		*status = "success"
	} else {
		*status = "error"
	}
}
type person struct {
	height int
	weight int
	skin string
}
func broadcast(value int) (bool, error) {
	fmt.Println("Voce está enviano", value);
	if value > 100 {
		return false, errors.New("Valor maior que o normal");
	}
	return true, nil;
}
