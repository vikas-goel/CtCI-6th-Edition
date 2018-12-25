package main

import (
	"container/list"
	"fmt"
	"math/rand"
	"time"
)

type AnimalType int
const (
	Cat AnimalType = 0
	Dog AnimalType = 1
	End AnimalType = 2
	Any AnimalType = 3
)

var names = [...]string{"cat-", "dog-"}

type Animal struct {
	Name string
	When time.Time
}

type Queue struct {
	cats, dogs *list.List
}

func (this Queue) Init() *Queue {
	this.cats, this.dogs = list.New(), list.New()
	return &this
}

func (this *Queue) Enqueue(name string, ofType AnimalType) {
	if this == nil {
		return
	}

	animal := Animal{Name: name, When: time.Now()}
	if ofType == Cat {
		this.cats.PushBack(&animal)
	} else {
		this.dogs.PushBack(&animal)
	}
}

func (this *Queue) DequeueAny() (animal *Animal) {
	if this == nil || (this.cats.Len() == 0 && this.dogs.Len() == 0) {
		return
	}

	cat := this.cats.Front()
	dog := this.dogs.Front()
	if cat == nil {
		animal = dog.Value.(*Animal)
	} else if dog == nil {
		animal = cat.Value.(*Animal)
	} else if cat.Value.(*Animal).When.Nanosecond() <= dog.Value.(*Animal).When.Nanosecond() {
		animal = cat.Value.(*Animal)
	} else {
		animal = dog.Value.(*Animal)
	}

	if  cat != nil && animal == cat.Value {
		this.cats.Remove(cat)
	} else {
		this.dogs.Remove(dog)
	}

	return
}

func (this *Queue) DequeueCat() *Animal {
	if this == nil || this.cats.Len() == 0 {
		return nil
	}

	cat := this.cats.Front()
	this.cats.Remove(cat)
	return cat.Value.(*Animal)
}

func (this *Queue) DequeueDog() *Animal {
	if this == nil || this.dogs.Len() == 0 {
		return nil
	}

	dog := this.dogs.Front()
	this.dogs.Remove(dog)
	return dog.Value.(*Animal)
}

func main() {
	q := Queue{}.Init()

	fmt.Print("Enqueuing animals:")
	for i := 0; i < 10; i++ {
		which := rand.Intn(int(End))
		name := fmt.Sprintf("%s%d", names[which], i)
		q.Enqueue(name, AnimalType(which))
		fmt.Print(" ", name)
		time.Sleep(1)
	}
	fmt.Println()

	fmt.Print("Dequeuing animals:")
	for i := 0; i < 50; i++ {
		var animal *Animal
		which := rand.Intn(int(Any))
		if AnimalType(which) == Cat {
			animal = q.DequeueCat()
			if animal != nil {
				fmt.Print(" Cat::", animal.Name)
			}
		} else if AnimalType(which) == Dog {
			animal = q.DequeueDog()
			if animal != nil {
				fmt.Print(" Dog::", animal.Name)
			}
		} else {
			animal = q.DequeueAny()
			if animal != nil {
				fmt.Print(" Any::", animal.Name)
			}
		}
	}
	fmt.Println()
}
