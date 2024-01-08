package rpn

type Input struct {
	characters []rune
	nextChar   rune
}

func NewInput(input string) *Input {
	characters := []rune(input)
	if len(characters) == 0 {
		return nil
	}

	i := &Input{
		characters: characters,
		nextChar:   characters[0],
	}
	return i
}

// Eat consumes a character from the input and returns the next one.
func (i *Input) Eat() rune {
	if len(i.characters) == 0 {
		return 0
	}

	i.characters = i.characters[1:]
	if len(i.characters) == 0 {
		i.nextChar = 0
	} else {
		i.nextChar = i.characters[0]
	}

	return i.nextChar
}

func (i *Input) NextChar() rune {
	return i.nextChar
}
