package g

var offset_gid = map[string]int64{
	"go1.4":     128,
	"go1.4.1":   128,
	"go1.4.2":   128,
	"go1.4.3":   128,
	"go1.5":     184,
	"go1.5.1":   184,
	"go1.5.2":   184,
	"go1.5.3":   184,
	"go1.5.4":   184,
	"go1.6":     192,
	"go1.6.1":   192,
	"go1.6.2":   192,
	"go1.6.3":   192,
	"go1.6.4":   192,
	"go1.7":     192,
	"go1.7.1":   192,
	"go1.7.2":   192,
	"go1.7.3":   192,
	"go1.7.4":   192,
	"go1.7.5":   192,
	"go1.7.6":   192,
	"go1.8":     192,
	"go1.8.1":   192,
	"go1.8.2":   192,
	"go1.8.3":   192,
	"go1.8.4":   192,
	"go1.8.5":   192,
	"go1.9":     152,
	"go1.9.1":   152,
	"go1.9.2":   152,
	"go1.10":    152,
	"go1.10.1":  152,
	"go1.10.2":  152,
	"go1.10.3":  152,
	"go1.10.4":  152,
	"go1.10.5":  152,
	"go1.10.6":  152,
	"go1.10.7":  152,
	"go1.10.8":  152,
	"go1.11":    152,
	"go1.11.1":  152,
	"go1.11.2":  152,
	"go1.11.3":  152,
	"go1.11.4":  152,
	"go1.11.5":  152,
	"go1.11.6":  152,
	"go1.11.7":  152,
	"go1.11.8":  152,
	"go1.11.9":  152,
	"go1.11.10": 152,
	"go1.11.11": 152,
	"go1.12":    152,
	"go1.12.1":  152,
	"go1.12.2":  152,
	"go1.12.3":  152,
	"go1.12.4":  152,
	"go1.12.5":  152,
	"go1.12.6":  152,
	"go1.12.7":  152,
	"go1.13":    152,
}

//
//import (
//	"sync/atomic"
//	"unsafe"
//)
//
//// NOTE:  Add package header and missing imports.
//
//func GoRoutineId() uint64 {
//
//	_ptr := uintptr(goRoutinegPtr())
//	_struct := (*[32]uint64) (unsafe.Pointer(_ptr))
//
//	_offset := atomic.LoadUint64(&goRoutineIdOffset)
//
//	if _offset != 0 {
//
//		return _struct[int(_offset)]
//
//	} else {
//
//		//_slow := goRoutineIdSlow ()
//		panic("func goRoutineIdSlow todo")
//		//_slow:=0
//		//_matchedCount := 0
//		//_matchedOffset := 0
//		//
//		//for _offset, _value := range _struct[:] {
//		//	if _value == _slow {
//		//		_matchedOffset = _offset
//		//		_matchedCount += 1
//		//		if _matchedCount >= 2 {
//		//			break
//		//		}
//		//	}
//		//}
//		//
//		//if _matchedCount == 1 {
//		//	atomic.StoreUint64 (&goRoutineIdOffset, uint64 (_matchedOffset))
//		//}
//		//
//		//return _slow
//	}
//}
//
//func goRoutinePtr() uint64
//
//var goRoutineIdOffset uint64 = 0
