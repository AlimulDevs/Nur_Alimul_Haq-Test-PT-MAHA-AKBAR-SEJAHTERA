package userService

func (ui *UserImplementation) Delete(id string) error {
	err := ui.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
