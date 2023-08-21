// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [UserDefaults] class.
var UserDefaultsClass = _UserDefaultsClass{objc.GetClass("NSUserDefaults")}

type _UserDefaultsClass struct {
	objc.Class
}

// An interface definition for the [UserDefaults] class.
type IUserDefaults interface {
	objc.IObject
	URLForKey(defaultName string) URL
	ObjectForKey(defaultName string) objc.Object
	SetBoolForKey(value bool, defaultName string)
	ArrayForKey(defaultName string) []objc.Object
	StringArrayForKey(defaultName string) []string
	SetFloatForKey(value float64, defaultName string)
	StringForKey(defaultName string) string
	FloatForKey(defaultName string) float64
	BoolForKey(defaultName string) bool
	RemovePersistentDomainForName(domainName string)
	SetObjectForKey(value objc.IObject, defaultName string)
	IntegerForKey(defaultName string) int
	VolatileDomainForName(domainName string) map[string]objc.Object
	ObjectIsForcedForKey(key string) bool
	SetPersistentDomainForName(domain map[string]objc.IObject, domainName string)
	SetDoubleForKey(value float64, defaultName string)
	DataForKey(defaultName string) []byte
	DoubleForKey(defaultName string) float64
	DictionaryForKey(defaultName string) map[string]objc.Object
	RegisterDefaults(registrationDictionary map[string]objc.IObject)
	SetIntegerForKey(value int, defaultName string)
	RemoveSuiteNamed(suiteName string)
	SetURLForKey(url IURL, defaultName string)
	RemoveVolatileDomainForName(domainName string)
	DictionaryRepresentation() map[string]objc.Object
	Synchronize() bool
	RemoveObjectForKey(defaultName string)
	SetVolatileDomainForName(domain map[string]objc.IObject, domainName string)
	PersistentDomainForName(domainName string) map[string]objc.Object
	AddSuiteNamed(suiteName string)
	VolatileDomainNames() []string
}

// An interface to the user’s defaults database, where you store key-value pairs persistently across launches of your app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserdefaults?language=objc
type UserDefaults struct {
	objc.Object
}

func UserDefaultsFrom(ptr unsafe.Pointer) UserDefaults {
	return UserDefaults{
		Object: objc.ObjectFrom(ptr),
	}
}

func (u_ UserDefaults) InitWithSuiteName(suitename string) UserDefaults {
	rv := objc.Call[UserDefaults](u_, objc.Sel("initWithSuiteName:"), suitename)
	return rv
}

// Creates a user defaults object initialized with the defaults for the specified database name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserdefaults/1409957-initwithsuitename?language=objc
func NewUserDefaultsWithSuiteName(suitename string) UserDefaults {
	instance := UserDefaultsClass.Alloc().InitWithSuiteName(suitename)
	instance.Autorelease()
	return instance
}

func (u_ UserDefaults) Init() UserDefaults {
	rv := objc.Call[UserDefaults](u_, objc.Sel("init"))
	return rv
}

func (uc _UserDefaultsClass) Alloc() UserDefaults {
	rv := objc.Call[UserDefaults](uc, objc.Sel("alloc"))
	return rv
}

func UserDefaults_Alloc() UserDefaults {
	return UserDefaultsClass.Alloc()
}

func (uc _UserDefaultsClass) New() UserDefaults {
	rv := objc.Call[UserDefaults](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewUserDefaults() UserDefaults {
	return UserDefaultsClass.New()
}

// Returns the URL associated with the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserdefaults/1408648-urlforkey?language=objc
func (u_ UserDefaults) URLForKey(defaultName string) URL {
	rv := objc.Call[URL](u_, objc.Sel("URLForKey:"), defaultName)
	return rv
}

// Returns the object associated with the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserdefaults/1410095-objectforkey?language=objc
func (u_ UserDefaults) ObjectForKey(defaultName string) objc.Object {
	rv := objc.Call[objc.Object](u_, objc.Sel("objectForKey:"), defaultName)
	return rv
}

// Sets the value of the specified default key to the specified Boolean value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserdefaults/1408905-setbool?language=objc
func (u_ UserDefaults) SetBoolForKey(value bool, defaultName string) {
	objc.Call[objc.Void](u_, objc.Sel("setBool:forKey:"), value, defaultName)
}

// Returns the array associated with the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserdefaults/1414792-arrayforkey?language=objc
func (u_ UserDefaults) ArrayForKey(defaultName string) []objc.Object {
	rv := objc.Call[[]objc.Object](u_, objc.Sel("arrayForKey:"), defaultName)
	return rv
}

// Returns the array of strings associated with the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserdefaults/1416414-stringarrayforkey?language=objc
func (u_ UserDefaults) StringArrayForKey(defaultName string) []string {
	rv := objc.Call[[]string](u_, objc.Sel("stringArrayForKey:"), defaultName)
	return rv
}

// Sets the value of the specified default key to the specified float value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserdefaults/1413320-setfloat?language=objc
func (u_ UserDefaults) SetFloatForKey(value float64, defaultName string) {
	objc.Call[objc.Void](u_, objc.Sel("setFloat:forKey:"), value, defaultName)
}

// Returns the string associated with the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserdefaults/1416700-stringforkey?language=objc
func (u_ UserDefaults) StringForKey(defaultName string) string {
	rv := objc.Call[string](u_, objc.Sel("stringForKey:"), defaultName)
	return rv
}

// Returns the float value associated with the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserdefaults/1414027-floatforkey?language=objc
func (u_ UserDefaults) FloatForKey(defaultName string) float64 {
	rv := objc.Call[float64](u_, objc.Sel("floatForKey:"), defaultName)
	return rv
}

// Returns the Boolean value associated with the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserdefaults/1416388-boolforkey?language=objc
func (u_ UserDefaults) BoolForKey(defaultName string) bool {
	rv := objc.Call[bool](u_, objc.Sel("boolForKey:"), defaultName)
	return rv
}

// Removes the contents of the specified persistent domain from the user’s defaults. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserdefaults/1417339-removepersistentdomainforname?language=objc
func (u_ UserDefaults) RemovePersistentDomainForName(domainName string) {
	objc.Call[objc.Void](u_, objc.Sel("removePersistentDomainForName:"), domainName)
}

// Sets the value of the specified default key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserdefaults/1414067-setobject?language=objc
func (u_ UserDefaults) SetObjectForKey(value objc.IObject, defaultName string) {
	objc.Call[objc.Void](u_, objc.Sel("setObject:forKey:"), value, defaultName)
}

// Returns the integer value associated with the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserdefaults/1407405-integerforkey?language=objc
func (u_ UserDefaults) IntegerForKey(defaultName string) int {
	rv := objc.Call[int](u_, objc.Sel("integerForKey:"), defaultName)
	return rv
}

// Returns the dictionary for the specified volatile domain. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserdefaults/1409592-volatiledomainforname?language=objc
func (u_ UserDefaults) VolatileDomainForName(domainName string) map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](u_, objc.Sel("volatileDomainForName:"), domainName)
	return rv
}

// Returns a Boolean value indicating whether the specified key is managed by an administrator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserdefaults/1408635-objectisforcedforkey?language=objc
func (u_ UserDefaults) ObjectIsForcedForKey(key string) bool {
	rv := objc.Call[bool](u_, objc.Sel("objectIsForcedForKey:"), key)
	return rv
}

// Sets a dictionary for the specified persistent domain. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserdefaults/1408187-setpersistentdomain?language=objc
func (u_ UserDefaults) SetPersistentDomainForName(domain map[string]objc.IObject, domainName string) {
	objc.Call[objc.Void](u_, objc.Sel("setPersistentDomain:forName:"), domain, domainName)
}

// Sets the value of the specified default key to the double value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserdefaults/1408646-setdouble?language=objc
func (u_ UserDefaults) SetDoubleForKey(value float64, defaultName string) {
	objc.Call[objc.Void](u_, objc.Sel("setDouble:forKey:"), value, defaultName)
}

// Returns the data object associated with the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserdefaults/1409590-dataforkey?language=objc
func (u_ UserDefaults) DataForKey(defaultName string) []byte {
	rv := objc.Call[[]byte](u_, objc.Sel("dataForKey:"), defaultName)
	return rv
}

// Returns the double value associated with the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserdefaults/1416581-doubleforkey?language=objc
func (u_ UserDefaults) DoubleForKey(defaultName string) float64 {
	rv := objc.Call[float64](u_, objc.Sel("doubleForKey:"), defaultName)
	return rv
}

// Returns the dictionary object associated with the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserdefaults/1408563-dictionaryforkey?language=objc
func (u_ UserDefaults) DictionaryForKey(defaultName string) map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](u_, objc.Sel("dictionaryForKey:"), defaultName)
	return rv
}

// Adds the contents of the specified dictionary to the registration domain. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserdefaults/1417065-registerdefaults?language=objc
func (u_ UserDefaults) RegisterDefaults(registrationDictionary map[string]objc.IObject) {
	objc.Call[objc.Void](u_, objc.Sel("registerDefaults:"), registrationDictionary)
}

// Sets the value of the specified default key to the specified integer value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserdefaults/1413614-setinteger?language=objc
func (u_ UserDefaults) SetIntegerForKey(value int, defaultName string) {
	objc.Call[objc.Void](u_, objc.Sel("setInteger:forKey:"), value, defaultName)
}

// Removes the specified domain name from the receiver’s search list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserdefaults/1408047-removesuitenamed?language=objc
func (u_ UserDefaults) RemoveSuiteNamed(suiteName string) {
	objc.Call[objc.Void](u_, objc.Sel("removeSuiteNamed:"), suiteName)
}

// Sets the value of the specified default key to the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserdefaults/1414194-seturl?language=objc
func (u_ UserDefaults) SetURLForKey(url IURL, defaultName string) {
	objc.Call[objc.Void](u_, objc.Sel("setURL:forKey:"), objc.Ptr(url), defaultName)
}

// Removes the specified volatile domain from the user’s defaults. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserdefaults/1415955-removevolatiledomainforname?language=objc
func (u_ UserDefaults) RemoveVolatileDomainForName(domainName string) {
	objc.Call[objc.Void](u_, objc.Sel("removeVolatileDomainForName:"), domainName)
}

// Returns a dictionary that contains a union of all key-value pairs in the domains in the search list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserdefaults/1415919-dictionaryrepresentation?language=objc
func (u_ UserDefaults) DictionaryRepresentation() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](u_, objc.Sel("dictionaryRepresentation"))
	return rv
}

// Waits for any pending asynchronous updates to the defaults database and returns; this method is unnecessary and shouldn't be used. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserdefaults/1414005-synchronize?language=objc
func (u_ UserDefaults) Synchronize() bool {
	rv := objc.Call[bool](u_, objc.Sel("synchronize"))
	return rv
}

// This method has no effect and shouldn't be used. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserdefaults/1407708-resetstandarduserdefaults?language=objc
func (uc _UserDefaultsClass) ResetStandardUserDefaults() {
	objc.Call[objc.Void](uc, objc.Sel("resetStandardUserDefaults"))
}

// This method has no effect and shouldn't be used. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserdefaults/1407708-resetstandarduserdefaults?language=objc
func UserDefaults_ResetStandardUserDefaults() {
	UserDefaultsClass.ResetStandardUserDefaults()
}

// Removes the value of the specified default key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserdefaults/1411182-removeobjectforkey?language=objc
func (u_ UserDefaults) RemoveObjectForKey(defaultName string) {
	objc.Call[objc.Void](u_, objc.Sel("removeObjectForKey:"), defaultName)
}

// Sets the dictionary for the specified volatile domain. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserdefaults/1413720-setvolatiledomain?language=objc
func (u_ UserDefaults) SetVolatileDomainForName(domain map[string]objc.IObject, domainName string) {
	objc.Call[objc.Void](u_, objc.Sel("setVolatileDomain:forName:"), domain, domainName)
}

// Returns a dictionary representation of the defaults for the specified domain. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserdefaults/1412197-persistentdomainforname?language=objc
func (u_ UserDefaults) PersistentDomainForName(domainName string) map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](u_, objc.Sel("persistentDomainForName:"), domainName)
	return rv
}

// Inserts the specified domain name into the receiver’s search list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserdefaults/1410294-addsuitenamed?language=objc
func (u_ UserDefaults) AddSuiteNamed(suiteName string) {
	objc.Call[objc.Void](u_, objc.Sel("addSuiteNamed:"), suiteName)
}

// Returns the shared defaults object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserdefaults/1416603-standarduserdefaults?language=objc
func (uc _UserDefaultsClass) StandardUserDefaults() UserDefaults {
	rv := objc.Call[UserDefaults](uc, objc.Sel("standardUserDefaults"))
	return rv
}

// Returns the shared defaults object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserdefaults/1416603-standarduserdefaults?language=objc
func UserDefaults_StandardUserDefaults() UserDefaults {
	return UserDefaultsClass.StandardUserDefaults()
}

// The current volatile domain names. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserdefaults/1414231-volatiledomainnames?language=objc
func (u_ UserDefaults) VolatileDomainNames() []string {
	rv := objc.Call[[]string](u_, objc.Sel("volatileDomainNames"))
	return rv
}