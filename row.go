package xlsx

import (
	"github.com/plandem/xlsx/format"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/options"
)

//Row is a higher level object that wraps ml.Row with functionality. Inherits functionality of Range
type Row struct {
	ml *ml.Row
	*Range
}

//Set sets options for row
func (r *Row) Set(o *options.RowOptions) {
	if o.Height > 0 {
		r.ml.Height = o.Height
		r.ml.CustomHeight = true
	}

	r.ml.OutlineLevel = o.OutlineLevel
	r.ml.Hidden = o.Hidden
	r.ml.Collapsed = o.Collapsed
	r.ml.Phonetic = o.Phonetic
}

//SetFormatting sets default style for the row. Affects cells not yet allocated in the row. In other words, this style applies to new cells.
func (r *Row) SetFormatting(styleRef format.StyleRefID) {
	r.ml.CustomFormat = true
	r.ml.Style = styleRef
}

//CopyTo copies row cells into another row with rIdx index.
//N.B.: Merged cells are not supported
func (r *Row) CopyTo(rIdx int, withOptions bool) {
	//TODO: check if sheet is opened as read stream and panic about

	if withOptions {
		//TODO: copy row options
	}

	//copy cell data
	r.Range.CopyTo(r.Range.fromCol, rIdx)
}
