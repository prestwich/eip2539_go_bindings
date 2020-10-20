package main

import (
	"C"
	"errors"
	"unsafe"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

const (
	MAX_OUTPUT_LEN                 = 256
	MAX_ERR_LEN                    = 256
	G1ADD_OPERATION_RAW_VALUE      = 1
	G1MUL_OPERATION_RAW_VALUE      = 2
	G1MULTIEXP_OPERATION_RAW_VALUE = 3
	G2ADD_OPERATION_RAW_VALUE      = 4
	G2MUL_OPERATION_RAW_VALUE      = 5
	G2MULTIEXP_OPERATION_RAW_VALUE = 6
	PAIR_OPERATION_RAW_VALUE       = 7
)

//export c_eip2539_perform_operation
func c_eip2539_perform_operation(op C.char, i *C.char, iLen uint32, o *C.char, oLen *uint32, e *C.char, eLen *uint32) C.int {
	iBuff := C.GoBytes(unsafe.Pointer(i), C.int(iLen))
	oBuff := (*[MAX_OUTPUT_LEN]byte)(unsafe.Pointer(o))
	eBuff := (*[MAX_ERR_LEN]byte)(unsafe.Pointer(e))

	opType := int(op)
	var res []byte
	var err error

	precompilesMap := vm.PrecompiledContractsEspresso

	switch opType {
	case G1ADD_OPERATION_RAW_VALUE:
		res, err = precompilesMap[common.BytesToAddress([]byte{0x13})].Run(iBuff) // TODO
	case G1MUL_OPERATION_RAW_VALUE:
		res, err = precompilesMap[common.BytesToAddress([]byte{0x14})].Run(iBuff)
	case G1MULTIEXP_OPERATION_RAW_VALUE:
		res, err = precompilesMap[common.BytesToAddress([]byte{0x15})].Run(iBuff)
	case G2ADD_OPERATION_RAW_VALUE:
		res, err = precompilesMap[common.BytesToAddress([]byte{0x16})].Run(iBuff)
	case G2MUL_OPERATION_RAW_VALUE:
		res, err = precompilesMap[common.BytesToAddress([]byte{0x17})].Run(iBuff)
	case G2MULTIEXP_OPERATION_RAW_VALUE:
		res, err = precompilesMap[common.BytesToAddress([]byte{0x18})].Run(iBuff)
	case PAIR_OPERATION_RAW_VALUE:
		res, err = precompilesMap[common.BytesToAddress([]byte{0x19})].Run(iBuff)
	default:
		err = errors.New("invalid operation type")
	}

	if err != nil {
		errDescr := err.Error()
		if len(errDescr) == 0 {
			*eLen = uint32(0)
			return 1
		}
		errDescrBytes := []byte(errDescr)
		errDescrByteLen := len(errDescrBytes)
		*eLen = uint32(errDescrByteLen)
		copied := copy(eBuff[0:], errDescrBytes)
		if copied != errDescrByteLen {
			println("Invalid number of bytes copied for an error")
		}
		return 1
	}
	oBytes := res
	resLen := len(oBytes)
	*oLen = uint32(len(oBytes))
	copied := copy(oBuff[0:], oBytes)
	if copied != resLen {
		println("Invalid number of bytes copied for result")
	}
	return 0
}

func main() {}
