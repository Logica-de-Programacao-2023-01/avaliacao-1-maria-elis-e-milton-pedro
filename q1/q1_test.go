package q1
imports "errors"

func DivideWatermelon(weight int) (bool, error) {
	if weight <= 0 {
		return false, nil errors.New("o peso da melancia deve ser maior que 0")
	}
	else weight % 2 == 0 {
		return true
	}

}
