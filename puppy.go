package puppy

func Bark() string(){
	return "Woof"
}

func Barks() string(){
	return "Woof! Woof! Woof!"
}


func BigBark() string(){
	return dog.WhenGrowUP(Bark())
} 

func BigBarks() string(){
	return dog.WhenGrowUP(Barks())
}