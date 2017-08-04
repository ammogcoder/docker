package sysx

// These functions will be generated by generate.sh
//    $ GOOS=darwin GOARCH=386 ./generate.sh xattr
//    $ GOOS=darwin GOARCH=amd64 ./generate.sh xattr

//sys  getxattr(path string, attr string, dest []byte, pos int, options int) (sz int, err error)
//sys  setxattr(path string, attr string, data []byte, flags int) (err error)
//sys  removexattr(path string, attr string, options int) (err error)
//sys  listxattr(path string, dest []byte, options int) (sz int, err error)
//sys  Fchmodat(dirfd int, path string, mode uint32, flags int) (err error)

const (
	xattrNoFollow = 0x01
)

func listxattrFollow(path string, dest []byte) (sz int, err error) {
	return listxattr(path, dest, 0)
}

// Listxattr calls syscall getxattr
func Listxattr(path string) ([]string, error) {
	return listxattrAll(path, listxattrFollow)
}

// Removexattr calls syscall getxattr
func Removexattr(path string, attr string) (err error) {
	return removexattr(path, attr, 0)
}

// Setxattr calls syscall setxattr
func Setxattr(path string, attr string, data []byte, flags int) (err error) {
	return setxattr(path, attr, data, flags)
}

func getxattrFollow(path, attr string, dest []byte) (sz int, err error) {
	return getxattr(path, attr, dest, 0, 0)
}

// Getxattr calls syscall getxattr
func Getxattr(path, attr string) ([]byte, error) {
	return getxattrAll(path, attr, getxattrFollow)
}

func listxattrNoFollow(path string, dest []byte) (sz int, err error) {
	return listxattr(path, dest, xattrNoFollow)
}

// LListxattr calls syscall listxattr with XATTR_NOFOLLOW
func LListxattr(path string) ([]string, error) {
	return listxattrAll(path, listxattrNoFollow)
}

// LRemovexattr calls syscall removexattr with XATTR_NOFOLLOW
func LRemovexattr(path string, attr string) (err error) {
	return removexattr(path, attr, xattrNoFollow)
}

// Setxattr calls syscall setxattr with XATTR_NOFOLLOW
func LSetxattr(path string, attr string, data []byte, flags int) (err error) {
	return setxattr(path, attr, data, flags|xattrNoFollow)
}

func getxattrNoFollow(path, attr string, dest []byte) (sz int, err error) {
	return getxattr(path, attr, dest, 0, xattrNoFollow)
}

// LGetxattr calls syscall getxattr with XATTR_NOFOLLOW
func LGetxattr(path, attr string) ([]byte, error) {
	return getxattrAll(path, attr, getxattrNoFollow)
}
