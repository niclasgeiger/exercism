package erratum

import "errors"

const testVersion = 2

func Use(opener func() (Resource, error), input string) (err error) {
	var resource Resource
	for res, resErr := opener(); ; res, resErr = opener() {
		if resErr == nil {
			resource = res
			break
		} else {
			switch e := resErr.(type) {
			case (TransientError):
				{
					continue
				}
			default:
				{
					return e
				}
			}
		}
	}
	defer func() {
		defer resource.Close()
		if r := recover(); r != nil {
			switch v := r.(type) {
			case (FrobError):
				{
					resource.Defrob(v.defrobTag)
					err = v.inner
				}
			default:
				{
					err = errors.New("meh")
				}
			}
		}
	}()
	resource.Frob(input)
	return err
}
