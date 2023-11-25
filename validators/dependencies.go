package validators

import (
	"os/exec"
)

func ValidatePythonInstalled() (bool, bool) {
	checkPython := exec.Command("which", "python")
	checkPythonThree := exec.Command("which", "python3")
	checkPip := exec.Command("which", "pip")

	boolPython := true
	boolPythonThree := true
	boolPip := true

	_, errPython := checkPython.Output()
	_, errPythonThree := checkPythonThree.Output()
	_, errPythonPip := checkPip.Output()

	if errPython != nil {
		boolPython = false
	}

	if errPythonThree != nil {
		boolPythonThree = false
	}

	if errPythonPip != nil {
		boolPip = false
	}

	dependPython := pythonBoleanHelper(boolPython, boolPythonThree)

	return dependPython, boolPip

}

func pythonBoleanHelper(python bool, three bool) bool {
	var check bool = true

	if python && !three {
		check = true
	}

	if !python && three {
		check = true
	}

	if !python && !three {
		check = false
	}

	return check
}
