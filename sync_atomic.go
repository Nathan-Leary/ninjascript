package ninjascript

import (
	sync_atomic "sync/atomic"
)

func init() {
	if _, ok := Api["sync/atomic"]; !ok {
		Api["sync/atomic"] = map[string]interface{}{}
	}
	Api["sync/atomic"]["AddInt32"] = sync_atomic.AddInt32
	Api["sync/atomic"]["AddInt64"] = sync_atomic.AddInt64
	Api["sync/atomic"]["AddUint32"] = sync_atomic.AddUint32
	Api["sync/atomic"]["AddUint64"] = sync_atomic.AddUint64
	Api["sync/atomic"]["AddUintptr"] = sync_atomic.AddUintptr
	Api["sync/atomic"]["CompareAndSwapInt32"] = sync_atomic.CompareAndSwapInt32
	Api["sync/atomic"]["CompareAndSwapInt64"] = sync_atomic.CompareAndSwapInt64
	Api["sync/atomic"]["CompareAndSwapPointer"] = sync_atomic.CompareAndSwapPointer
	Api["sync/atomic"]["CompareAndSwapUint32"] = sync_atomic.CompareAndSwapUint32
	Api["sync/atomic"]["CompareAndSwapUint64"] = sync_atomic.CompareAndSwapUint64
	Api["sync/atomic"]["CompareAndSwapUintptr"] = sync_atomic.CompareAndSwapUintptr
	Api["sync/atomic"]["LoadInt32"] = sync_atomic.LoadInt32
	Api["sync/atomic"]["LoadInt64"] = sync_atomic.LoadInt64
	Api["sync/atomic"]["LoadPointer"] = sync_atomic.LoadPointer
	Api["sync/atomic"]["LoadUint32"] = sync_atomic.LoadUint32
	Api["sync/atomic"]["LoadUint64"] = sync_atomic.LoadUint64
	Api["sync/atomic"]["LoadUintptr"] = sync_atomic.LoadUintptr
	Api["sync/atomic"]["StoreInt32"] = sync_atomic.StoreInt32
	Api["sync/atomic"]["StoreInt64"] = sync_atomic.StoreInt64
	Api["sync/atomic"]["StorePointer"] = sync_atomic.StorePointer
	Api["sync/atomic"]["StoreUint32"] = sync_atomic.StoreUint32
	Api["sync/atomic"]["StoreUint64"] = sync_atomic.StoreUint64
	Api["sync/atomic"]["StoreUintptr"] = sync_atomic.StoreUintptr
	Api["sync/atomic"]["SwapInt32"] = sync_atomic.SwapInt32
	Api["sync/atomic"]["SwapInt64"] = sync_atomic.SwapInt64
	Api["sync/atomic"]["SwapPointer"] = sync_atomic.SwapPointer
	Api["sync/atomic"]["SwapUint32"] = sync_atomic.SwapUint32
	Api["sync/atomic"]["SwapUint64"] = sync_atomic.SwapUint64
	Api["sync/atomic"]["SwapUintptr"] = sync_atomic.SwapUintptr
	Api["sync/atomic"]["Bool"] = sync_atomic.Bool{}
	Api["sync/atomic"]["Int32"] = sync_atomic.Int32{}
	Api["sync/atomic"]["Int64"] = sync_atomic.Int64{}
	// Api["sync/atomic"]["Pointer"] = sync_atomic.Pointer{}
	Api["sync/atomic"]["Uint32"] = sync_atomic.Uint32{}
	Api["sync/atomic"]["Uint64"] = sync_atomic.Uint64{}
	Api["sync/atomic"]["Uintptr"] = sync_atomic.Uintptr{}
	Api["sync/atomic"]["Value"] = sync_atomic.Value{}
	Api["sync/atomic"]["AddInt32"] = sync_atomic.AddInt32
	Api["sync/atomic"]["AddInt64"] = sync_atomic.AddInt64
	Api["sync/atomic"]["AddUint32"] = sync_atomic.AddUint32
	Api["sync/atomic"]["AddUint64"] = sync_atomic.AddUint64
	Api["sync/atomic"]["AddUintptr"] = sync_atomic.AddUintptr
	Api["sync/atomic"]["CompareAndSwapInt32"] = sync_atomic.CompareAndSwapInt32
	Api["sync/atomic"]["CompareAndSwapInt64"] = sync_atomic.CompareAndSwapInt64
	Api["sync/atomic"]["CompareAndSwapPointer"] = sync_atomic.CompareAndSwapPointer
	Api["sync/atomic"]["CompareAndSwapUint32"] = sync_atomic.CompareAndSwapUint32
	Api["sync/atomic"]["CompareAndSwapUint64"] = sync_atomic.CompareAndSwapUint64
	Api["sync/atomic"]["CompareAndSwapUintptr"] = sync_atomic.CompareAndSwapUintptr
	Api["sync/atomic"]["LoadInt32"] = sync_atomic.LoadInt32
	Api["sync/atomic"]["LoadInt64"] = sync_atomic.LoadInt64
	Api["sync/atomic"]["LoadPointer"] = sync_atomic.LoadPointer
	Api["sync/atomic"]["LoadUint32"] = sync_atomic.LoadUint32
	Api["sync/atomic"]["LoadUint64"] = sync_atomic.LoadUint64
	Api["sync/atomic"]["LoadUintptr"] = sync_atomic.LoadUintptr
	Api["sync/atomic"]["StoreInt32"] = sync_atomic.StoreInt32
	Api["sync/atomic"]["StoreInt64"] = sync_atomic.StoreInt64
	Api["sync/atomic"]["StorePointer"] = sync_atomic.StorePointer
	Api["sync/atomic"]["StoreUint32"] = sync_atomic.StoreUint32
	Api["sync/atomic"]["StoreUint64"] = sync_atomic.StoreUint64
	Api["sync/atomic"]["StoreUintptr"] = sync_atomic.StoreUintptr
	Api["sync/atomic"]["SwapInt32"] = sync_atomic.SwapInt32
	Api["sync/atomic"]["SwapInt64"] = sync_atomic.SwapInt64
	Api["sync/atomic"]["SwapPointer"] = sync_atomic.SwapPointer
	Api["sync/atomic"]["SwapUint32"] = sync_atomic.SwapUint32
	Api["sync/atomic"]["SwapUint64"] = sync_atomic.SwapUint64
	Api["sync/atomic"]["SwapUintptr"] = sync_atomic.SwapUintptr

}
