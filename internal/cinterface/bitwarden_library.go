package cinterface

import (
	"fmt"
	"unsafe"
)

/*
#cgo linux,amd64 LDFLAGS:   -L ./lib/linux-x64 -lbitwarden_c -lm
#cgo linux,arm64 LDFLAGS:   -L ./lib/linux-arm64 -lbitwarden_c -lm
#cgo darwin,amd64 LDFLAGS:  -L ./lib/darwin-x64 -lbitwarden_c -framework Security -framework SystemConfiguration
#cgo darwin,arm64 LDFLAGS:  -L ./lib/darwin-arm64 -lbitwarden_c -framework Security -framework SystemConfiguration
#cgo windows,amd64 LDFLAGS: -L ./lib/windows-x64 -l:libbitwarden_c.a -ladvapi32 -lbcrypt -lcrypt32 -lcryptnet -lkernel32 -lncrypt -lntdll -luserenv -lws2_32 -lmsvcrt -loleaut32 -lruntimeobject
#include <stdlib.h>
typedef void* ClientPtr;
extern char* run_command(const char *command, ClientPtr client);
extern ClientPtr init(const char *clientSettings);
extern void free_mem(ClientPtr client);
*/
import "C"

type ClientPointer struct {
	Pointer C.ClientPtr
}

type BitwardenLibrary interface {
	Init(clientSettings string) (ClientPointer, error)
	FreeMem(client ClientPointer)
	RunCommand(command string, client ClientPointer) (string, error)
}

type BitwardenLibraryImpl struct{}

func NewBitwardenLibrary() BitwardenLibrary {
	return &BitwardenLibraryImpl{}
}

func (b *BitwardenLibraryImpl) Init(clientSettings string) (ClientPointer, error) {
	css := C.CString(clientSettings)
	defer C.free(unsafe.Pointer(css))
	ptr := C.init(css)
	if ptr == nil {
		return ClientPointer{}, fmt.Errorf("initialization failed")
	}
	return ClientPointer{Pointer: ptr}, nil
}

func (b *BitwardenLibraryImpl) FreeMem(client ClientPointer) {
	C.free_mem(client.Pointer)
}

func (b *BitwardenLibraryImpl) RunCommand(command string, client ClientPointer) (string, error) {
	cmds := C.CString(command)
	defer C.free(unsafe.Pointer(cmds))
	cstr := C.run_command(cmds, client.Pointer)
	if cstr == nil {
		return "", fmt.Errorf("run command failed")
	}
	defer C.free(unsafe.Pointer(cstr))
	return C.GoString(cstr), nil
}
