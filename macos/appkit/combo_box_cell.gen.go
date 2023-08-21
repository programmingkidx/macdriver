// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ComboBoxCell] class.
var ComboBoxCellClass = _ComboBoxCellClass{objc.GetClass("NSComboBoxCell")}

type _ComboBoxCellClass struct {
	objc.Class
}

// An interface definition for the [ComboBoxCell] class.
type IComboBoxCell interface {
	ITextFieldCell
	RemoveItemWithObjectValue(object objc.IObject)
	NoteNumberOfItemsChanged()
	IndexOfItemWithObjectValue(object objc.IObject) int
	InsertItemWithObjectValueAtIndex(object objc.IObject, index int)
	ScrollItemAtIndexToVisible(index int)
	RemoveAllItems()
	RemoveItemAtIndex(index int)
	AddItemWithObjectValue(object objc.IObject)
	AddItemsWithObjectValues(objects []objc.IObject)
	DeselectItemAtIndex(index int)
	ScrollItemAtIndexToTop(index int)
	SelectItemWithObjectValue(object objc.IObject)
	SelectItemAtIndex(index int)
	ItemObjectValueAtIndex(index int) objc.Object
	CompletedString(string_ string) string
	ReloadData()
	DataSource() ComboBoxCellDataSourceWrapper
	SetDataSource(value PComboBoxCellDataSource)
	SetDataSourceObject(valueObject objc.IObject)
	IndexOfSelectedItem() int
	ItemHeight() float64
	SetItemHeight(value float64)
	ObjectValueOfSelectedItem() objc.Object
	HasVerticalScroller() bool
	SetHasVerticalScroller(value bool)
	Completes() bool
	SetCompletes(value bool)
	UsesDataSource() bool
	SetUsesDataSource(value bool)
	ObjectValues() []objc.Object
	NumberOfItems() int
	NumberOfVisibleItems() int
	SetNumberOfVisibleItems(value int)
	IsButtonBordered() bool
	SetButtonBordered(value bool)
	IntercellSpacing() foundation.Size
	SetIntercellSpacing(value foundation.Size)
}

// The user interface of a combo box. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell?language=objc
type ComboBoxCell struct {
	TextFieldCell
}

func ComboBoxCellFrom(ptr unsafe.Pointer) ComboBoxCell {
	return ComboBoxCell{
		TextFieldCell: TextFieldCellFrom(ptr),
	}
}

func (cc _ComboBoxCellClass) Alloc() ComboBoxCell {
	rv := objc.Call[ComboBoxCell](cc, objc.Sel("alloc"))
	return rv
}

func ComboBoxCell_Alloc() ComboBoxCell {
	return ComboBoxCellClass.Alloc()
}

func (cc _ComboBoxCellClass) New() ComboBoxCell {
	rv := objc.Call[ComboBoxCell](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewComboBoxCell() ComboBoxCell {
	return ComboBoxCellClass.New()
}

func (c_ ComboBoxCell) Init() ComboBoxCell {
	rv := objc.Call[ComboBoxCell](c_, objc.Sel("init"))
	return rv
}

func (c_ ComboBoxCell) InitTextCell(string_ string) ComboBoxCell {
	rv := objc.Call[ComboBoxCell](c_, objc.Sel("initTextCell:"), string_)
	return rv
}

// Initializes a text field cell that displays the specified string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfieldcell/1642278-inittextcell?language=objc
func NewComboBoxCellTextCell(string_ string) ComboBoxCell {
	instance := ComboBoxCellClass.Alloc().InitTextCell(string_)
	instance.Autorelease()
	return instance
}

func (c_ ComboBoxCell) InitImageCell(image IImage) ComboBoxCell {
	rv := objc.Call[ComboBoxCell](c_, objc.Sel("initImageCell:"), objc.Ptr(image))
	return rv
}

// Returns an NSCell object initialized with the specified image and set to have the cell’s default menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscell/1533898-initimagecell?language=objc
func NewComboBoxCellImageCell(image IImage) ComboBoxCell {
	instance := ComboBoxCellClass.Alloc().InitImageCell(image)
	instance.Autorelease()
	return instance
}

// Removes all occurrences of the specified object from the combo box’s internal item list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/1410295-removeitemwithobjectvalue?language=objc
func (c_ ComboBoxCell) RemoveItemWithObjectValue(object objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("removeItemWithObjectValue:"), object)
}

// Informs the combo box that the number of items in its data source has changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/1410268-notenumberofitemschanged?language=objc
func (c_ ComboBoxCell) NoteNumberOfItemsChanged() {
	objc.Call[objc.Void](c_, objc.Sel("noteNumberOfItemsChanged"))
}

// Searches the combo box’s internal item list for the given object and returns the matching index number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/1410283-indexofitemwithobjectvalue?language=objc
func (c_ ComboBoxCell) IndexOfItemWithObjectValue(object objc.IObject) int {
	rv := objc.Call[int](c_, objc.Sel("indexOfItemWithObjectValue:"), object)
	return rv
}

// Inserts an object at the specified location in the internal item list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/1410297-insertitemwithobjectvalue?language=objc
func (c_ ComboBoxCell) InsertItemWithObjectValueAtIndex(object objc.IObject, index int) {
	objc.Call[objc.Void](c_, objc.Sel("insertItemWithObjectValue:atIndex:"), object, index)
}

// Scrolls the combo box’s pop-up list vertically so that the item at the given index is visible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/1410272-scrollitematindextovisible?language=objc
func (c_ ComboBoxCell) ScrollItemAtIndexToVisible(index int) {
	objc.Call[objc.Void](c_, objc.Sel("scrollItemAtIndexToVisible:"), index)
}

// Removes all items from the combo box’s internal item list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/1410290-removeallitems?language=objc
func (c_ ComboBoxCell) RemoveAllItems() {
	objc.Call[objc.Void](c_, objc.Sel("removeAllItems"))
}

// Removes the object at the specified location from the combo box’s internal item list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/1410276-removeitematindex?language=objc
func (c_ ComboBoxCell) RemoveItemAtIndex(index int) {
	objc.Call[objc.Void](c_, objc.Sel("removeItemAtIndex:"), index)
}

// Adds the specified object to the internal item list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/1410286-additemwithobjectvalue?language=objc
func (c_ ComboBoxCell) AddItemWithObjectValue(object objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("addItemWithObjectValue:"), object)
}

// Adds multiple objects to the internal item list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/1410300-additemswithobjectvalues?language=objc
func (c_ ComboBoxCell) AddItemsWithObjectValues(objects []objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("addItemsWithObjectValues:"), objects)
}

// Deselects the pop-up list item at the given index if it’s selected. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/1410282-deselectitematindex?language=objc
func (c_ ComboBoxCell) DeselectItemAtIndex(index int) {
	objc.Call[objc.Void](c_, objc.Sel("deselectItemAtIndex:"), index)
}

// Scrolls the combo box’s pop-up list vertically so that the item at the given index is as close to the top as possible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/1410288-scrollitematindextotop?language=objc
func (c_ ComboBoxCell) ScrollItemAtIndexToTop(index int) {
	objc.Call[objc.Void](c_, objc.Sel("scrollItemAtIndexToTop:"), index)
}

// Selects the first pop-up list item that corresponds to the specified object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/1410280-selectitemwithobjectvalue?language=objc
func (c_ ComboBoxCell) SelectItemWithObjectValue(object objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("selectItemWithObjectValue:"), object)
}

// Selects the pop-up list row at the given index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/1410292-selectitematindex?language=objc
func (c_ ComboBoxCell) SelectItemAtIndex(index int) {
	objc.Call[objc.Void](c_, objc.Sel("selectItemAtIndex:"), index)
}

// Returns the object located at the specified location in the internal item list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/1410255-itemobjectvalueatindex?language=objc
func (c_ ComboBoxCell) ItemObjectValueAtIndex(index int) objc.Object {
	rv := objc.Call[objc.Object](c_, objc.Sel("itemObjectValueAtIndex:"), index)
	return rv
}

// Returns a string from the combo box’s pop-up list that starts with the given substring. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/1410264-completedstring?language=objc
func (c_ ComboBoxCell) CompletedString(string_ string) string {
	rv := objc.Call[string](c_, objc.Sel("completedString:"), string_)
	return rv
}

// Marks the combo box as needing redisplay, so that it will reload the data for visible pop-up items and draw the new values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/1410248-reloaddata?language=objc
func (c_ ComboBoxCell) ReloadData() {
	objc.Call[objc.Void](c_, objc.Sel("reloadData"))
}

// The object that provides the data displayed in the combo box’s pop-up list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/1410266-datasource?language=objc
func (c_ ComboBoxCell) DataSource() ComboBoxCellDataSourceWrapper {
	rv := objc.Call[ComboBoxCellDataSourceWrapper](c_, objc.Sel("dataSource"))
	return rv
}

// The object that provides the data displayed in the combo box’s pop-up list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/1410266-datasource?language=objc
func (c_ ComboBoxCell) SetDataSource(value PComboBoxCellDataSource) {
	po0 := objc.WrapAsProtocol("NSComboBoxCellDataSource", value)
	objc.Call[objc.Void](c_, objc.Sel("setDataSource:"), po0)
}

// The object that provides the data displayed in the combo box’s pop-up list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/1410266-datasource?language=objc
func (c_ ComboBoxCell) SetDataSourceObject(valueObject objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("setDataSource:"), objc.Ptr(valueObject))
}

// The index of the last item selected from the pop-up list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/1410274-indexofselecteditem?language=objc
func (c_ ComboBoxCell) IndexOfSelectedItem() int {
	rv := objc.Call[int](c_, objc.Sel("indexOfSelectedItem"))
	return rv
}

// The height of each item in the combo box’s pop-up list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/1410306-itemheight?language=objc
func (c_ ComboBoxCell) ItemHeight() float64 {
	rv := objc.Call[float64](c_, objc.Sel("itemHeight"))
	return rv
}

// The height of each item in the combo box’s pop-up list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/1410306-itemheight?language=objc
func (c_ ComboBoxCell) SetItemHeight(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setItemHeight:"), value)
}

// The object corresponding to the last item selected from the pop-up list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/1410278-objectvalueofselecteditem?language=objc
func (c_ ComboBoxCell) ObjectValueOfSelectedItem() objc.Object {
	rv := objc.Call[objc.Object](c_, objc.Sel("objectValueOfSelectedItem"))
	return rv
}

// A Boolean value that indicates if the combo box displays a vertical scroller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/1410252-hasverticalscroller?language=objc
func (c_ ComboBoxCell) HasVerticalScroller() bool {
	rv := objc.Call[bool](c_, objc.Sel("hasVerticalScroller"))
	return rv
}

// A Boolean value that indicates if the combo box displays a vertical scroller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/1410252-hasverticalscroller?language=objc
func (c_ ComboBoxCell) SetHasVerticalScroller(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setHasVerticalScroller:"), value)
}

// A Boolean value that indicates if the combo box tries to complete text entered by the user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/1410262-completes?language=objc
func (c_ ComboBoxCell) Completes() bool {
	rv := objc.Call[bool](c_, objc.Sel("completes"))
	return rv
}

// A Boolean value that indicates if the combo box tries to complete text entered by the user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/1410262-completes?language=objc
func (c_ ComboBoxCell) SetCompletes(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setCompletes:"), value)
}

// A Boolean value that indicates if the combo box uses an external data source to populate its pop-up list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/1410254-usesdatasource?language=objc
func (c_ ComboBoxCell) UsesDataSource() bool {
	rv := objc.Call[bool](c_, objc.Sel("usesDataSource"))
	return rv
}

// A Boolean value that indicates if the combo box uses an external data source to populate its pop-up list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/1410254-usesdatasource?language=objc
func (c_ ComboBoxCell) SetUsesDataSource(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setUsesDataSource:"), value)
}

// The combo box’s internal item list in an array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/1410304-objectvalues?language=objc
func (c_ ComboBoxCell) ObjectValues() []objc.Object {
	rv := objc.Call[[]objc.Object](c_, objc.Sel("objectValues"))
	return rv
}

// The total number of items in the pop-up list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/1410260-numberofitems?language=objc
func (c_ ComboBoxCell) NumberOfItems() int {
	rv := objc.Call[int](c_, objc.Sel("numberOfItems"))
	return rv
}

// The maximum number of items visible in the pop-up list at any one time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/1410298-numberofvisibleitems?language=objc
func (c_ ComboBoxCell) NumberOfVisibleItems() int {
	rv := objc.Call[int](c_, objc.Sel("numberOfVisibleItems"))
	return rv
}

// The maximum number of items visible in the pop-up list at any one time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/1410298-numberofvisibleitems?language=objc
func (c_ ComboBoxCell) SetNumberOfVisibleItems(value int) {
	objc.Call[objc.Void](c_, objc.Sel("setNumberOfVisibleItems:"), value)
}

// A Boolean value that indicates whether the combo box button displays a border. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/1410246-buttonbordered?language=objc
func (c_ ComboBoxCell) IsButtonBordered() bool {
	rv := objc.Call[bool](c_, objc.Sel("isButtonBordered"))
	return rv
}

// A Boolean value that indicates whether the combo box button displays a border. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/1410246-buttonbordered?language=objc
func (c_ ComboBoxCell) SetButtonBordered(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setButtonBordered:"), value)
}

// The spacing between cells in the combo box’s pop-up list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/1410270-intercellspacing?language=objc
func (c_ ComboBoxCell) IntercellSpacing() foundation.Size {
	rv := objc.Call[foundation.Size](c_, objc.Sel("intercellSpacing"))
	return rv
}

// The spacing between cells in the combo box’s pop-up list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscomboboxcell/1410270-intercellspacing?language=objc
func (c_ ComboBoxCell) SetIntercellSpacing(value foundation.Size) {
	objc.Call[objc.Void](c_, objc.Sel("setIntercellSpacing:"), value)
}