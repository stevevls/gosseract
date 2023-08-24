package gosseract

// #cgo CXXFLAGS: -std=c++0x
// #cgo CPPFLAGS: -I/usr/local/include
// #cgo CPPFLAGS: -Wno-unused-result
// #cgo pkg-config: lept tesseract
import "C"
