package main

// ParseJSON contacts file
func ParseJSON(filename string) ([]Contact, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var contacts []Contact

	json.Unmarshal(file, &contacts)

	return contacts, nil
}
