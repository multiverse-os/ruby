package memory

import (
	"sync/atomic"
	"unsafe"

	"gvisor.dev/gvisor/pkg/atomicbitops"
	"gvisor.dev/gvisor/pkg/sentry/context"
)

// SwapUint32 implements IO.SwapUint32.
func (b *BytesIO) SwapUint32(ctx context.Context, addr Addr, new uint32, opts IOOpts) (uint32, error) {
	if _, rngErr := b.rangeCheck(addr, 4); rngErr != nil {
		return 0, rngErr
	}
	return atomic.SwapUint32((*uint32)(unsafe.Pointer(&b.Bytes[int(addr)])), new), nil
}

// CompareAndSwapUint32 implements IO.CompareAndSwapUint32.
func (b *BytesIO) CompareAndSwapUint32(ctx context.Context, addr Addr, old, new uint32, opts IOOpts) (uint32, error) {
	if _, rngErr := b.rangeCheck(addr, 4); rngErr != nil {
		return 0, rngErr
	}
	return atomicbitops.CompareAndSwapUint32((*uint32)(unsafe.Pointer(&b.Bytes[int(addr)])), old, new), nil
}

// LoadUint32 implements IO.LoadUint32.
func (b *BytesIO) LoadUint32(ctx context.Context, addr Addr, opts IOOpts) (uint32, error) {
	if _, err := b.rangeCheck(addr, 4); err != nil {
		return 0, err
	}
	return atomic.LoadUint32((*uint32)(unsafe.Pointer(&b.Bytes[int(addr)]))), nil
}
