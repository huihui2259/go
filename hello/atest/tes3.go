package main

type User struct {
	ID   int
	Age  int
	Name string
}

func test1() map[int]User {
	m := map[int]User{}
	for i := 0; i < 10; i++ {
		user := User{ID: i, Name: "nihao"}
		m[i] = user
	}
	return m
}

func test(id int) *User {
	m := test1()
	for k, v := range m {
		if k == id {
			return &v
		}
	}
	return nil
}

func test2(id int) *User {
	user1 := User{ID: 1, Name: "nihao"}
	return &user1
}
