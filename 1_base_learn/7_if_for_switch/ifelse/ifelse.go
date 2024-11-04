package main

func main() {

}

func IfUsingNewVariable(start int, end int) {
	if distance := end - start; distance > 100 {
		println("distance:", distance)
	} else if distance < 0 {
		println("distance:", distance)
	} else {
		println("distance:", distance)
	}

	//println("distance:", distance)

}
