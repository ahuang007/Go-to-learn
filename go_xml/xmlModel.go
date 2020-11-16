package main

type Reusult struct {
	Person []Person
}

type Person struct {
	Name      string
	Age       string
	Career    string
	Interests Interests
}

type Interests struct {
	Interest []string
}
