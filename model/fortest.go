package model

type ForTest struct {
	ID      int
	Comment string
}

func CreateTable() error {
	err := db.CreateTable(&ForTest{}).Error

	if err != nil {
		return err
	}

	return nil
}
