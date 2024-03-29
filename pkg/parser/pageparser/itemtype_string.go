// Code generated by "stringer -type ItemType item.go"; DO NOT EDIT.

package pageparser

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[tError-0]
	_ = x[tEOF-1]
	_ = x[TypeHTMLStart-2]
	_ = x[TypeLeadSummaryDivider-3]
	_ = x[TypeFrontMatterYAML-4]
	_ = x[TypeFrontMatterTOML-5]
	_ = x[TypeFrontMatterJSON-6]
	_ = x[TypeFrontMatterORG-7]
	_ = x[TypeEmoji-8]
	_ = x[TypeIgnore-9]
	_ = x[tLeftDelimScNoMarkup-10]
	_ = x[tRightDelimScNoMarkup-11]
	_ = x[tLeftDelimScWithMarkup-12]
	_ = x[tRightDelimScWithMarkup-13]
	_ = x[tScClose-14]
	_ = x[tScName-15]
	_ = x[tScNameInline-16]
	_ = x[tScParam-17]
	_ = x[tScParamVal-18]
	_ = x[tText-19]
	_ = x[tKeywordMarker-20]
}

const _ItemType_name = "tErrortEOFTypeHTMLStartTypeLeadSummaryDividerTypeFrontMatterYAMLTypeFrontMatterTOMLTypeFrontMatterJSONTypeFrontMatterORGTypeEmojiTypeIgnoretLeftDelimScNoMarkuptRightDelimScNoMarkuptLeftDelimScWithMarkuptRightDelimScWithMarkuptScClosetScNametScNameInlinetScParamtScParamValtTexttKeywordMarker"

var _ItemType_index = [...]uint16{0, 6, 10, 23, 45, 64, 83, 102, 120, 129, 139, 159, 180, 202, 225, 233, 240, 253, 261, 272, 277, 291}

func (i ItemType) String() string {
	if i < 0 || i >= ItemType(len(_ItemType_index)-1) {
		return "ItemType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ItemType_name[_ItemType_index[i]:_ItemType_index[i+1]]
}
