package helpers

import "go.uber.org/dig"

func MultiProvideDI(container *dig.Container, constructorList []any) (err error) {
	for _, constructor := range constructorList {
		err = container.Provide(constructor)

		if err != nil {
			return
		}
	}

	return
}
