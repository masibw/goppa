package loader

import "github.com/masibw/goppa/domain/valueobject"

type Loader interface {
	Load(fileName string) (testData []valueobject.TestEvent, err error)
}
