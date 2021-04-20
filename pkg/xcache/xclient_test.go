package xcache

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCacheClient(t *testing.T) {
	fn := func() (map[string]string, error) {
		return map[string]string{
			"key_a": "a",
			"key_b": "b",
		}, nil
	}
	Convey("Xache.Xclient", t, func() {
		xClient := NewCacheClient()
		So(xClient, ShouldNotBeNil)

		v, _ := xClient.Get("key_a", fn)
		So(v, ShouldEqual, "a")

		v, _ = xClient.Get("key_b", fn)
		So(v, ShouldEqual, "b")
	})


}
