package main

type numIdentErr int

func (numIdentErr) Error() string {
	return "Variable identifiers may not start with a number."
}

type invalRuneInIdentErr int

func (invalRuneInIdentErr) Error() string {
	return "Variable identifiers must only contain printable runes."
}

type unknownVarErr int

func (unknownVarErr) Error() string {
	return "The variable identifier does not exist."
}

type varExistsErr int

func (varExistsErr) Error() string {
	return "The variable identifier conflicts with a preexisting variable."
}

type flagValErr int

func (flagValErr) Error() string {
	return "The flag was not defined with a value."
}
