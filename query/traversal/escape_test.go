package traversal

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEscape(t *testing.T) {
	Convey("Given a string value for a graph property", t, func() {
		Convey("When 'escape' is called", func() {
			result := escape("text\\ text' \"text\nte😊xt\t")
			Convey("Then result should be escaped", func() {
				So(result, ShouldEqual, `'text\\ text\' "text\nte😊xt\t'`)
			})
		})
	})
}
