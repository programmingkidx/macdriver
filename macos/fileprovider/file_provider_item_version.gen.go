// AUTO-GENERATED CODE, DO NOT MODIFY

package fileprovider

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FileProviderItemVersion] class.
var FileProviderItemVersionClass = _FileProviderItemVersionClass{objc.GetClass("NSFileProviderItemVersion")}

type _FileProviderItemVersionClass struct {
	objc.Class
}

// An interface definition for the [FileProviderItemVersion] class.
type IFileProviderItemVersion interface {
	objc.IObject
	MetadataVersion() []byte
	ContentVersion() []byte
}

// The version of the item’s content and its metadata. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovideritemversion?language=objc
type FileProviderItemVersion struct {
	objc.Object
}

func FileProviderItemVersionFrom(ptr unsafe.Pointer) FileProviderItemVersion {
	return FileProviderItemVersion{
		Object: objc.ObjectFrom(ptr),
	}
}

func (f_ FileProviderItemVersion) InitWithContentVersionMetadataVersion(contentVersion []byte, metadataVersion []byte) FileProviderItemVersion {
	rv := objc.Call[FileProviderItemVersion](f_, objc.Sel("initWithContentVersion:metadataVersion:"), contentVersion, metadataVersion)
	return rv
}

// Creates a new version object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovideritemversion/3043899-initwithcontentversion?language=objc
func NewFileProviderItemVersionWithContentVersionMetadataVersion(contentVersion []byte, metadataVersion []byte) FileProviderItemVersion {
	instance := FileProviderItemVersionClass.Alloc().InitWithContentVersionMetadataVersion(contentVersion, metadataVersion)
	instance.Autorelease()
	return instance
}

func (fc _FileProviderItemVersionClass) Alloc() FileProviderItemVersion {
	rv := objc.Call[FileProviderItemVersion](fc, objc.Sel("alloc"))
	return rv
}

func FileProviderItemVersion_Alloc() FileProviderItemVersion {
	return FileProviderItemVersionClass.Alloc()
}

func (fc _FileProviderItemVersionClass) New() FileProviderItemVersion {
	rv := objc.Call[FileProviderItemVersion](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFileProviderItemVersion() FileProviderItemVersion {
	return FileProviderItemVersionClass.New()
}

func (f_ FileProviderItemVersion) Init() FileProviderItemVersion {
	rv := objc.Call[FileProviderItemVersion](f_, objc.Sel("init"))
	return rv
}

// An opaque object used to track versions of the item’s metadata. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovideritemversion/3043900-metadataversion?language=objc
func (f_ FileProviderItemVersion) MetadataVersion() []byte {
	rv := objc.Call[[]byte](f_, objc.Sel("metadataVersion"))
	return rv
}

// A Boolean value indicating that this version predates the version returned by the file provider extension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovideritemversion/3869739-beforefirstsynccomponent?language=objc
func (fc _FileProviderItemVersionClass) BeforeFirstSyncComponent() []byte {
	rv := objc.Call[[]byte](fc, objc.Sel("beforeFirstSyncComponent"))
	return rv
}

// A Boolean value indicating that this version predates the version returned by the file provider extension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovideritemversion/3869739-beforefirstsynccomponent?language=objc
func FileProviderItemVersion_BeforeFirstSyncComponent() []byte {
	return FileProviderItemVersionClass.BeforeFirstSyncComponent()
}

// An opaque object used to track versions of the item’s content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileprovideritemversion/3043898-contentversion?language=objc
func (f_ FileProviderItemVersion) ContentVersion() []byte {
	rv := objc.Call[[]byte](f_, objc.Sel("contentVersion"))
	return rv
}