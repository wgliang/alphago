package alphago

import (
	"reflect"
	"testing"
	"time"
)

func Test_TimeNow(t *testing.T) {
	_ = TimeNow()
}

func Test_TimeFormat_1(t *testing.T) {
	te := Str2Time("2017-07-18 16:16:36")
	if !reflect.DeepEqual(TimeFormat("Y-m-d H:i:s", te), "2017-07-18 16:16:36") {
		t.Error(`TimeFormat("Y-m-d H:i:s", te), "2017-07-18 16:16:36"`)
	}
}

func Test_TimeFormat_2(t *testing.T) {
	te := Str2Time("2017-07-18 16:16:36")
	if !reflect.DeepEqual(TimeFormat("Y-m-d", te), "2017-07-18") {
		t.Error(`TimeFormat("Y-m-d", te), "2017-07-18"`)
	}
}

func Test_TimeFormat_3(t *testing.T) {
	te := Str2Time("2017-07-18 16:16:36")
	if !reflect.DeepEqual(TimeFormat("Y-m-d H:i", te), "2017-07-18 16:16") {
		t.Error(`TimeFormat("Y-m-d H:i", te), "2017-07-18 16:16"`)
	}
}

func Test_TimeFormat_4(t *testing.T) {
	te := Str2Time("2017-07-18 16:16:36")
	if !reflect.DeepEqual(TimeFormat("Y-m-d H", te), "2017-07-18 16") {
		t.Error(`TimeFormat("Y-m-d H", te), "2017-07-18 16"`)
	}
}

func Test_TimeFormat_5(t *testing.T) {
	te := Str2Time("2017-07-18 16:16:36")
	if !reflect.DeepEqual(TimeFormat("y-m-d H:i:s", te), "17-07-18 16:16:36") {
		t.Error(`TimeFormat("y-m-d H:i:s", te), "17-07-18 16:16:36"`)
	}
}

func Test_TimeFormat_6(t *testing.T) {
	te := Str2Time("2017-07-18 16:16:36")
	if !reflect.DeepEqual(TimeFormat("Y-m-d g:i:s", te), "2017-07-18 4:16:36") {
		t.Error(`TimeFormat("Y-m-d g:i:s", te), "2017-07-18 4:16:36"`)
	}
}

func Test_TimeFormat_7(t *testing.T) {
	te := Str2Time("2017-07-18 16:16:36")
	if !reflect.DeepEqual(TimeFormat("Y-m-d H:i:s A", te), "2017-07-18 16:16:36 PM") {
		t.Error(`TimeFormat("Y-m-d H:i:s A", te), "2017-07-18 16:16:36"`)
	}
}

func Test_Str2LocalTime_1(t *testing.T) {
	_ = Str2LocalTime("2017-07-18 16:16:36 -0700 MST")
}

func Test_Str2LocalTime_2(t *testing.T) {
	if !reflect.DeepEqual(Str2LocalTime(""), time.Time{}) {
		t.Error(`TimeFormat("Y-m-d H:i:s", te), "2017-07-18 16:16:36"`)
	}
}

func Test_Str2Time_1(t *testing.T) {
	if !reflect.DeepEqual(Str2Time(""), time.Time{}) {
		t.Error(`TimeFormat("Y-m-d H:i:s", te), "2017-07-18 16:16:36"`)
	}
}

func Test_Str2Time_2(t *testing.T) {
	_ = Str2Time("2017-07-18 16:16:36")
}

func Test_IsWeekend_1(t *testing.T) {
	if !reflect.DeepEqual(IsWeekend(Str2Time("2017-07-18 16:16:36")), false) {
		t.Error(`IsWeekend(Str2Time("2017-07-18 16:16:36")), false`)
	}
}

func Test_IsWeekend_2(t *testing.T) {
	if !reflect.DeepEqual(IsWeekend(Str2Time("2017-07-22 16:16:36")), true) {
		t.Error(`IsWeekend(Str2Time("2017-07-18 16:16:36")), false`)
	}
}
