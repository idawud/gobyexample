package structs

type Person struct {
	name string
	age  int
}

func NewPerson(name string, age int) *Person {
	p := Person{name, age}
	return &p
}

func (p Person) GetDetails() map[string]interface{} {
	per := make(map[string]interface{}, 0)
	per["name"] = p.name
	per["age"] = p.age
	return per
}
