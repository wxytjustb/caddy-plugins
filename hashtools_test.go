package studychanger

import (
	"testing"
	"time"
)

func TestMapWithExpire_SetNx(t *testing.T) {
	var testMap = NewMapWithExpire(nil)

	val, ok := testMap.SetNx("key1", "val1", 3 * time.Second)

	t.Logf("test1 res is: %t,  val is:%s", ok, val.(string))
	if ok == false {
		t.Fail()
	}

	val, ok = testMap.SetNx("key1", "val2", 3 * time.Second)

	t.Logf("test2 res is: %t,  val is:%s", ok, val.(string))

	time.Sleep(3 * time.Second)

	val, ok = testMap.SetNx("key1", "val3", 3 * time.Second)
	t.Logf("test3 res is: %t,  val is:%s", ok, val.(string))

}
