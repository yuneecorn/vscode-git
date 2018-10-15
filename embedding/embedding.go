package main

import (
	"fmt"
)

//구조체 임베딩 사용하기
//구조체 임베딩을 사용하면 다른 언어의 상속과 동일한 형태가 됩니다.
// 물론 구조체 안에 여러 개의 구조체를 임베딩하면 다중 상속도 구현할 수 있습니다.
// 하지만 Go 언어에서는 복잡한 다중 상속 대신 인터페이스를 주로 활용합니다.
type Person struct { // 사람 구조체 정의
	name string
	age  int
}

func (p *Person) greeting() { // 인사(greeting) 함수 작성
	fmt.Println("Hello~")
}

type Student1 struct {
	p      Person // 학생 구조체 안에 사람 구조체를 필드로 가지고 있음. Has-a 관계
	school string
	grade  int
}

type Student2 struct {
	Person // 필드명 없이 타입만 선언하면 포함(Is-a) 관계가 됨
	//구조체 임베딩을 할 때 필드명을 사용하지 않는 필드를 익명 필드(anonymous field)라고 합니다.
	school string
	grade  int
}

type Student3 struct {
	Person //Inherit by embedding the struct by value
	school string
	grade  int
}

func (p *Student3) greeting() { // 이름이 같은 메서드를 정의하면 오버라이딩됨
	fmt.Println("Hello Students~")
}

type Student4 struct {
	*Person //Inherit by embedding a pointer to a struct
	school  string
	grade   int
}

func main() {

	var s1 Student1

	s1.p.greeting() // Hello~

	var s2 Student2

	s2.Person.greeting() // Hello~
	s2.greeting()        // Hello~

	var s3 Student3
	var s4 Student4

	s3.Person.greeting() // Hello~
	s3.greeting()        // Hello Students~
	fmt.Println(s3.Person)

	s4.Person.greeting()
	fmt.Println(s4.Person)

}
