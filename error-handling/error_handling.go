package erratum

func Use(opener ResourceOpener, input string) (err error) {
	resource, err := opener()
	for err != nil {
		resource, err = opener()

		if _, ok := err.(TransientError); ok {
			continue
		}

		if err != nil {
			return
		}
	}

	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
			if f, ok := r.(FrobError); ok {
				resource.Defrob(f.defrobTag)
			}
		}
		resource.Close()
	}()

	resource.Frob(input)
	return
}
