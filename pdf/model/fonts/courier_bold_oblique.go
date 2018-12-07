/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE.md', which is part of this source code package.
 */
/*
 * The embedded character metrics specified in this file are distributed under the terms listed in
 * ./afms/MustRead.html.
 */

package fonts

import (
	"github.com/unidoc/unidoc/pdf/core"
	"github.com/unidoc/unidoc/pdf/internal/textencoding"
)

// FontCourierBoldOblique represents the Courier-BoldOblique font.
// This is a built-in font and it is assumed that every reader has access to it.
type FontCourierBoldOblique struct {
	encoder textencoding.TextEncoder
}

// NewFontCourierBoldOblique returns a new instance of the font with a default encoder set
// (WinAnsiEncoding).
func NewFontCourierBoldOblique() FontCourierBoldOblique {
	font := FontCourierBoldOblique{}
	font.encoder = textencoding.NewWinAnsiTextEncoder()
	return font
}

// Encoder returns the font's text encoder.
func (font FontCourierBoldOblique) Encoder() textencoding.TextEncoder {
	return font.encoder
}

// GetGlyphCharMetrics returns character metrics for a given glyph.
func (font FontCourierBoldOblique) GetGlyphCharMetrics(glyph GlyphName) (CharMetrics, bool) {
	metrics, has := CourierBoldObliqueCharMetrics[glyph]
	if !has {
		return metrics, false
	}

	return metrics, true
}

// ToPdfObject returns a primitive PDF object representation of the font.
func (font FontCourierBoldOblique) ToPdfObject() core.PdfObject {
	fontDict := core.MakeDict()
	fontDict.Set("Type", core.MakeName("Font"))
	fontDict.Set("Subtype", core.MakeName("Type1"))
	fontDict.Set("BaseFont", core.MakeName("Courier-BoldOblique"))
	fontDict.Set("Encoding", font.encoder.ToPdfObject())

	return core.MakeIndirectObject(fontDict)
}

// CourierBoldObliqueCharMetrics are the font metrics loaded from afms/Courier-BoldOblique.afm.
// See afms/MustRead.html for license information.
var CourierBoldObliqueCharMetrics = map[GlyphName]CharMetrics{
	"A":              {GlyphName: "A", Wx: 600.000000, Wy: 0.000000},
	"AE":             {GlyphName: "AE", Wx: 600.000000, Wy: 0.000000},
	"Aacute":         {GlyphName: "Aacute", Wx: 600.000000, Wy: 0.000000},
	"Abreve":         {GlyphName: "Abreve", Wx: 600.000000, Wy: 0.000000},
	"Acircumflex":    {GlyphName: "Acircumflex", Wx: 600.000000, Wy: 0.000000},
	"Adieresis":      {GlyphName: "Adieresis", Wx: 600.000000, Wy: 0.000000},
	"Agrave":         {GlyphName: "Agrave", Wx: 600.000000, Wy: 0.000000},
	"Amacron":        {GlyphName: "Amacron", Wx: 600.000000, Wy: 0.000000},
	"Aogonek":        {GlyphName: "Aogonek", Wx: 600.000000, Wy: 0.000000},
	"Aring":          {GlyphName: "Aring", Wx: 600.000000, Wy: 0.000000},
	"Atilde":         {GlyphName: "Atilde", Wx: 600.000000, Wy: 0.000000},
	"B":              {GlyphName: "B", Wx: 600.000000, Wy: 0.000000},
	"C":              {GlyphName: "C", Wx: 600.000000, Wy: 0.000000},
	"Cacute":         {GlyphName: "Cacute", Wx: 600.000000, Wy: 0.000000},
	"Ccaron":         {GlyphName: "Ccaron", Wx: 600.000000, Wy: 0.000000},
	"Ccedilla":       {GlyphName: "Ccedilla", Wx: 600.000000, Wy: 0.000000},
	"D":              {GlyphName: "D", Wx: 600.000000, Wy: 0.000000},
	"Dcaron":         {GlyphName: "Dcaron", Wx: 600.000000, Wy: 0.000000},
	"Dcroat":         {GlyphName: "Dcroat", Wx: 600.000000, Wy: 0.000000},
	"Delta":          {GlyphName: "Delta", Wx: 600.000000, Wy: 0.000000},
	"E":              {GlyphName: "E", Wx: 600.000000, Wy: 0.000000},
	"Eacute":         {GlyphName: "Eacute", Wx: 600.000000, Wy: 0.000000},
	"Ecaron":         {GlyphName: "Ecaron", Wx: 600.000000, Wy: 0.000000},
	"Ecircumflex":    {GlyphName: "Ecircumflex", Wx: 600.000000, Wy: 0.000000},
	"Edieresis":      {GlyphName: "Edieresis", Wx: 600.000000, Wy: 0.000000},
	"Edotaccent":     {GlyphName: "Edotaccent", Wx: 600.000000, Wy: 0.000000},
	"Egrave":         {GlyphName: "Egrave", Wx: 600.000000, Wy: 0.000000},
	"Emacron":        {GlyphName: "Emacron", Wx: 600.000000, Wy: 0.000000},
	"Eogonek":        {GlyphName: "Eogonek", Wx: 600.000000, Wy: 0.000000},
	"Eth":            {GlyphName: "Eth", Wx: 600.000000, Wy: 0.000000},
	"Euro":           {GlyphName: "Euro", Wx: 600.000000, Wy: 0.000000},
	"F":              {GlyphName: "F", Wx: 600.000000, Wy: 0.000000},
	"G":              {GlyphName: "G", Wx: 600.000000, Wy: 0.000000},
	"Gbreve":         {GlyphName: "Gbreve", Wx: 600.000000, Wy: 0.000000},
	"Gcommaaccent":   {GlyphName: "Gcommaaccent", Wx: 600.000000, Wy: 0.000000},
	"H":              {GlyphName: "H", Wx: 600.000000, Wy: 0.000000},
	"I":              {GlyphName: "I", Wx: 600.000000, Wy: 0.000000},
	"Iacute":         {GlyphName: "Iacute", Wx: 600.000000, Wy: 0.000000},
	"Icircumflex":    {GlyphName: "Icircumflex", Wx: 600.000000, Wy: 0.000000},
	"Idieresis":      {GlyphName: "Idieresis", Wx: 600.000000, Wy: 0.000000},
	"Idotaccent":     {GlyphName: "Idotaccent", Wx: 600.000000, Wy: 0.000000},
	"Igrave":         {GlyphName: "Igrave", Wx: 600.000000, Wy: 0.000000},
	"Imacron":        {GlyphName: "Imacron", Wx: 600.000000, Wy: 0.000000},
	"Iogonek":        {GlyphName: "Iogonek", Wx: 600.000000, Wy: 0.000000},
	"J":              {GlyphName: "J", Wx: 600.000000, Wy: 0.000000},
	"K":              {GlyphName: "K", Wx: 600.000000, Wy: 0.000000},
	"Kcommaaccent":   {GlyphName: "Kcommaaccent", Wx: 600.000000, Wy: 0.000000},
	"L":              {GlyphName: "L", Wx: 600.000000, Wy: 0.000000},
	"Lacute":         {GlyphName: "Lacute", Wx: 600.000000, Wy: 0.000000},
	"Lcaron":         {GlyphName: "Lcaron", Wx: 600.000000, Wy: 0.000000},
	"Lcommaaccent":   {GlyphName: "Lcommaaccent", Wx: 600.000000, Wy: 0.000000},
	"Lslash":         {GlyphName: "Lslash", Wx: 600.000000, Wy: 0.000000},
	"M":              {GlyphName: "M", Wx: 600.000000, Wy: 0.000000},
	"N":              {GlyphName: "N", Wx: 600.000000, Wy: 0.000000},
	"Nacute":         {GlyphName: "Nacute", Wx: 600.000000, Wy: 0.000000},
	"Ncaron":         {GlyphName: "Ncaron", Wx: 600.000000, Wy: 0.000000},
	"Ncommaaccent":   {GlyphName: "Ncommaaccent", Wx: 600.000000, Wy: 0.000000},
	"Ntilde":         {GlyphName: "Ntilde", Wx: 600.000000, Wy: 0.000000},
	"O":              {GlyphName: "O", Wx: 600.000000, Wy: 0.000000},
	"OE":             {GlyphName: "OE", Wx: 600.000000, Wy: 0.000000},
	"Oacute":         {GlyphName: "Oacute", Wx: 600.000000, Wy: 0.000000},
	"Ocircumflex":    {GlyphName: "Ocircumflex", Wx: 600.000000, Wy: 0.000000},
	"Odieresis":      {GlyphName: "Odieresis", Wx: 600.000000, Wy: 0.000000},
	"Ograve":         {GlyphName: "Ograve", Wx: 600.000000, Wy: 0.000000},
	"Ohungarumlaut":  {GlyphName: "Ohungarumlaut", Wx: 600.000000, Wy: 0.000000},
	"Omacron":        {GlyphName: "Omacron", Wx: 600.000000, Wy: 0.000000},
	"Oslash":         {GlyphName: "Oslash", Wx: 600.000000, Wy: 0.000000},
	"Otilde":         {GlyphName: "Otilde", Wx: 600.000000, Wy: 0.000000},
	"P":              {GlyphName: "P", Wx: 600.000000, Wy: 0.000000},
	"Q":              {GlyphName: "Q", Wx: 600.000000, Wy: 0.000000},
	"R":              {GlyphName: "R", Wx: 600.000000, Wy: 0.000000},
	"Racute":         {GlyphName: "Racute", Wx: 600.000000, Wy: 0.000000},
	"Rcaron":         {GlyphName: "Rcaron", Wx: 600.000000, Wy: 0.000000},
	"Rcommaaccent":   {GlyphName: "Rcommaaccent", Wx: 600.000000, Wy: 0.000000},
	"S":              {GlyphName: "S", Wx: 600.000000, Wy: 0.000000},
	"Sacute":         {GlyphName: "Sacute", Wx: 600.000000, Wy: 0.000000},
	"Scaron":         {GlyphName: "Scaron", Wx: 600.000000, Wy: 0.000000},
	"Scedilla":       {GlyphName: "Scedilla", Wx: 600.000000, Wy: 0.000000},
	"Scommaaccent":   {GlyphName: "Scommaaccent", Wx: 600.000000, Wy: 0.000000},
	"T":              {GlyphName: "T", Wx: 600.000000, Wy: 0.000000},
	"Tcaron":         {GlyphName: "Tcaron", Wx: 600.000000, Wy: 0.000000},
	"Tcommaaccent":   {GlyphName: "Tcommaaccent", Wx: 600.000000, Wy: 0.000000},
	"Thorn":          {GlyphName: "Thorn", Wx: 600.000000, Wy: 0.000000},
	"U":              {GlyphName: "U", Wx: 600.000000, Wy: 0.000000},
	"Uacute":         {GlyphName: "Uacute", Wx: 600.000000, Wy: 0.000000},
	"Ucircumflex":    {GlyphName: "Ucircumflex", Wx: 600.000000, Wy: 0.000000},
	"Udieresis":      {GlyphName: "Udieresis", Wx: 600.000000, Wy: 0.000000},
	"Ugrave":         {GlyphName: "Ugrave", Wx: 600.000000, Wy: 0.000000},
	"Uhungarumlaut":  {GlyphName: "Uhungarumlaut", Wx: 600.000000, Wy: 0.000000},
	"Umacron":        {GlyphName: "Umacron", Wx: 600.000000, Wy: 0.000000},
	"Uogonek":        {GlyphName: "Uogonek", Wx: 600.000000, Wy: 0.000000},
	"Uring":          {GlyphName: "Uring", Wx: 600.000000, Wy: 0.000000},
	"V":              {GlyphName: "V", Wx: 600.000000, Wy: 0.000000},
	"W":              {GlyphName: "W", Wx: 600.000000, Wy: 0.000000},
	"X":              {GlyphName: "X", Wx: 600.000000, Wy: 0.000000},
	"Y":              {GlyphName: "Y", Wx: 600.000000, Wy: 0.000000},
	"Yacute":         {GlyphName: "Yacute", Wx: 600.000000, Wy: 0.000000},
	"Ydieresis":      {GlyphName: "Ydieresis", Wx: 600.000000, Wy: 0.000000},
	"Z":              {GlyphName: "Z", Wx: 600.000000, Wy: 0.000000},
	"Zacute":         {GlyphName: "Zacute", Wx: 600.000000, Wy: 0.000000},
	"Zcaron":         {GlyphName: "Zcaron", Wx: 600.000000, Wy: 0.000000},
	"Zdotaccent":     {GlyphName: "Zdotaccent", Wx: 600.000000, Wy: 0.000000},
	"a":              {GlyphName: "a", Wx: 600.000000, Wy: 0.000000},
	"aacute":         {GlyphName: "aacute", Wx: 600.000000, Wy: 0.000000},
	"abreve":         {GlyphName: "abreve", Wx: 600.000000, Wy: 0.000000},
	"acircumflex":    {GlyphName: "acircumflex", Wx: 600.000000, Wy: 0.000000},
	"acute":          {GlyphName: "acute", Wx: 600.000000, Wy: 0.000000},
	"adieresis":      {GlyphName: "adieresis", Wx: 600.000000, Wy: 0.000000},
	"ae":             {GlyphName: "ae", Wx: 600.000000, Wy: 0.000000},
	"agrave":         {GlyphName: "agrave", Wx: 600.000000, Wy: 0.000000},
	"amacron":        {GlyphName: "amacron", Wx: 600.000000, Wy: 0.000000},
	"ampersand":      {GlyphName: "ampersand", Wx: 600.000000, Wy: 0.000000},
	"aogonek":        {GlyphName: "aogonek", Wx: 600.000000, Wy: 0.000000},
	"aring":          {GlyphName: "aring", Wx: 600.000000, Wy: 0.000000},
	"asciicircum":    {GlyphName: "asciicircum", Wx: 600.000000, Wy: 0.000000},
	"asciitilde":     {GlyphName: "asciitilde", Wx: 600.000000, Wy: 0.000000},
	"asterisk":       {GlyphName: "asterisk", Wx: 600.000000, Wy: 0.000000},
	"at":             {GlyphName: "at", Wx: 600.000000, Wy: 0.000000},
	"atilde":         {GlyphName: "atilde", Wx: 600.000000, Wy: 0.000000},
	"b":              {GlyphName: "b", Wx: 600.000000, Wy: 0.000000},
	"backslash":      {GlyphName: "backslash", Wx: 600.000000, Wy: 0.000000},
	"bar":            {GlyphName: "bar", Wx: 600.000000, Wy: 0.000000},
	"braceleft":      {GlyphName: "braceleft", Wx: 600.000000, Wy: 0.000000},
	"braceright":     {GlyphName: "braceright", Wx: 600.000000, Wy: 0.000000},
	"bracketleft":    {GlyphName: "bracketleft", Wx: 600.000000, Wy: 0.000000},
	"bracketright":   {GlyphName: "bracketright", Wx: 600.000000, Wy: 0.000000},
	"breve":          {GlyphName: "breve", Wx: 600.000000, Wy: 0.000000},
	"brokenbar":      {GlyphName: "brokenbar", Wx: 600.000000, Wy: 0.000000},
	"bullet":         {GlyphName: "bullet", Wx: 600.000000, Wy: 0.000000},
	"c":              {GlyphName: "c", Wx: 600.000000, Wy: 0.000000},
	"cacute":         {GlyphName: "cacute", Wx: 600.000000, Wy: 0.000000},
	"caron":          {GlyphName: "caron", Wx: 600.000000, Wy: 0.000000},
	"ccaron":         {GlyphName: "ccaron", Wx: 600.000000, Wy: 0.000000},
	"ccedilla":       {GlyphName: "ccedilla", Wx: 600.000000, Wy: 0.000000},
	"cedilla":        {GlyphName: "cedilla", Wx: 600.000000, Wy: 0.000000},
	"cent":           {GlyphName: "cent", Wx: 600.000000, Wy: 0.000000},
	"circumflex":     {GlyphName: "circumflex", Wx: 600.000000, Wy: 0.000000},
	"colon":          {GlyphName: "colon", Wx: 600.000000, Wy: 0.000000},
	"comma":          {GlyphName: "comma", Wx: 600.000000, Wy: 0.000000},
	"commaaccent":    {GlyphName: "commaaccent", Wx: 600.000000, Wy: 0.000000},
	"copyright":      {GlyphName: "copyright", Wx: 600.000000, Wy: 0.000000},
	"currency":       {GlyphName: "currency", Wx: 600.000000, Wy: 0.000000},
	"d":              {GlyphName: "d", Wx: 600.000000, Wy: 0.000000},
	"dagger":         {GlyphName: "dagger", Wx: 600.000000, Wy: 0.000000},
	"daggerdbl":      {GlyphName: "daggerdbl", Wx: 600.000000, Wy: 0.000000},
	"dcaron":         {GlyphName: "dcaron", Wx: 600.000000, Wy: 0.000000},
	"dcroat":         {GlyphName: "dcroat", Wx: 600.000000, Wy: 0.000000},
	"degree":         {GlyphName: "degree", Wx: 600.000000, Wy: 0.000000},
	"dieresis":       {GlyphName: "dieresis", Wx: 600.000000, Wy: 0.000000},
	"divide":         {GlyphName: "divide", Wx: 600.000000, Wy: 0.000000},
	"dollar":         {GlyphName: "dollar", Wx: 600.000000, Wy: 0.000000},
	"dotaccent":      {GlyphName: "dotaccent", Wx: 600.000000, Wy: 0.000000},
	"dotlessi":       {GlyphName: "dotlessi", Wx: 600.000000, Wy: 0.000000},
	"e":              {GlyphName: "e", Wx: 600.000000, Wy: 0.000000},
	"eacute":         {GlyphName: "eacute", Wx: 600.000000, Wy: 0.000000},
	"ecaron":         {GlyphName: "ecaron", Wx: 600.000000, Wy: 0.000000},
	"ecircumflex":    {GlyphName: "ecircumflex", Wx: 600.000000, Wy: 0.000000},
	"edieresis":      {GlyphName: "edieresis", Wx: 600.000000, Wy: 0.000000},
	"edotaccent":     {GlyphName: "edotaccent", Wx: 600.000000, Wy: 0.000000},
	"egrave":         {GlyphName: "egrave", Wx: 600.000000, Wy: 0.000000},
	"eight":          {GlyphName: "eight", Wx: 600.000000, Wy: 0.000000},
	"ellipsis":       {GlyphName: "ellipsis", Wx: 600.000000, Wy: 0.000000},
	"emacron":        {GlyphName: "emacron", Wx: 600.000000, Wy: 0.000000},
	"emdash":         {GlyphName: "emdash", Wx: 600.000000, Wy: 0.000000},
	"endash":         {GlyphName: "endash", Wx: 600.000000, Wy: 0.000000},
	"eogonek":        {GlyphName: "eogonek", Wx: 600.000000, Wy: 0.000000},
	"equal":          {GlyphName: "equal", Wx: 600.000000, Wy: 0.000000},
	"eth":            {GlyphName: "eth", Wx: 600.000000, Wy: 0.000000},
	"exclam":         {GlyphName: "exclam", Wx: 600.000000, Wy: 0.000000},
	"exclamdown":     {GlyphName: "exclamdown", Wx: 600.000000, Wy: 0.000000},
	"f":              {GlyphName: "f", Wx: 600.000000, Wy: 0.000000},
	"fi":             {GlyphName: "fi", Wx: 600.000000, Wy: 0.000000},
	"five":           {GlyphName: "five", Wx: 600.000000, Wy: 0.000000},
	"fl":             {GlyphName: "fl", Wx: 600.000000, Wy: 0.000000},
	"florin":         {GlyphName: "florin", Wx: 600.000000, Wy: 0.000000},
	"four":           {GlyphName: "four", Wx: 600.000000, Wy: 0.000000},
	"fraction":       {GlyphName: "fraction", Wx: 600.000000, Wy: 0.000000},
	"g":              {GlyphName: "g", Wx: 600.000000, Wy: 0.000000},
	"gbreve":         {GlyphName: "gbreve", Wx: 600.000000, Wy: 0.000000},
	"gcommaaccent":   {GlyphName: "gcommaaccent", Wx: 600.000000, Wy: 0.000000},
	"germandbls":     {GlyphName: "germandbls", Wx: 600.000000, Wy: 0.000000},
	"grave":          {GlyphName: "grave", Wx: 600.000000, Wy: 0.000000},
	"greater":        {GlyphName: "greater", Wx: 600.000000, Wy: 0.000000},
	"greaterequal":   {GlyphName: "greaterequal", Wx: 600.000000, Wy: 0.000000},
	"guillemotleft":  {GlyphName: "guillemotleft", Wx: 600.000000, Wy: 0.000000},
	"guillemotright": {GlyphName: "guillemotright", Wx: 600.000000, Wy: 0.000000},
	"guilsinglleft":  {GlyphName: "guilsinglleft", Wx: 600.000000, Wy: 0.000000},
	"guilsinglright": {GlyphName: "guilsinglright", Wx: 600.000000, Wy: 0.000000},
	"h":              {GlyphName: "h", Wx: 600.000000, Wy: 0.000000},
	"hungarumlaut":   {GlyphName: "hungarumlaut", Wx: 600.000000, Wy: 0.000000},
	"hyphen":         {GlyphName: "hyphen", Wx: 600.000000, Wy: 0.000000},
	"i":              {GlyphName: "i", Wx: 600.000000, Wy: 0.000000},
	"iacute":         {GlyphName: "iacute", Wx: 600.000000, Wy: 0.000000},
	"icircumflex":    {GlyphName: "icircumflex", Wx: 600.000000, Wy: 0.000000},
	"idieresis":      {GlyphName: "idieresis", Wx: 600.000000, Wy: 0.000000},
	"igrave":         {GlyphName: "igrave", Wx: 600.000000, Wy: 0.000000},
	"imacron":        {GlyphName: "imacron", Wx: 600.000000, Wy: 0.000000},
	"iogonek":        {GlyphName: "iogonek", Wx: 600.000000, Wy: 0.000000},
	"j":              {GlyphName: "j", Wx: 600.000000, Wy: 0.000000},
	"k":              {GlyphName: "k", Wx: 600.000000, Wy: 0.000000},
	"kcommaaccent":   {GlyphName: "kcommaaccent", Wx: 600.000000, Wy: 0.000000},
	"l":              {GlyphName: "l", Wx: 600.000000, Wy: 0.000000},
	"lacute":         {GlyphName: "lacute", Wx: 600.000000, Wy: 0.000000},
	"lcaron":         {GlyphName: "lcaron", Wx: 600.000000, Wy: 0.000000},
	"lcommaaccent":   {GlyphName: "lcommaaccent", Wx: 600.000000, Wy: 0.000000},
	"less":           {GlyphName: "less", Wx: 600.000000, Wy: 0.000000},
	"lessequal":      {GlyphName: "lessequal", Wx: 600.000000, Wy: 0.000000},
	"logicalnot":     {GlyphName: "logicalnot", Wx: 600.000000, Wy: 0.000000},
	"lozenge":        {GlyphName: "lozenge", Wx: 600.000000, Wy: 0.000000},
	"lslash":         {GlyphName: "lslash", Wx: 600.000000, Wy: 0.000000},
	"m":              {GlyphName: "m", Wx: 600.000000, Wy: 0.000000},
	"macron":         {GlyphName: "macron", Wx: 600.000000, Wy: 0.000000},
	"minus":          {GlyphName: "minus", Wx: 600.000000, Wy: 0.000000},
	"mu":             {GlyphName: "mu", Wx: 600.000000, Wy: 0.000000},
	"multiply":       {GlyphName: "multiply", Wx: 600.000000, Wy: 0.000000},
	"n":              {GlyphName: "n", Wx: 600.000000, Wy: 0.000000},
	"nacute":         {GlyphName: "nacute", Wx: 600.000000, Wy: 0.000000},
	"ncaron":         {GlyphName: "ncaron", Wx: 600.000000, Wy: 0.000000},
	"ncommaaccent":   {GlyphName: "ncommaaccent", Wx: 600.000000, Wy: 0.000000},
	"nine":           {GlyphName: "nine", Wx: 600.000000, Wy: 0.000000},
	"notequal":       {GlyphName: "notequal", Wx: 600.000000, Wy: 0.000000},
	"ntilde":         {GlyphName: "ntilde", Wx: 600.000000, Wy: 0.000000},
	"numbersign":     {GlyphName: "numbersign", Wx: 600.000000, Wy: 0.000000},
	"o":              {GlyphName: "o", Wx: 600.000000, Wy: 0.000000},
	"oacute":         {GlyphName: "oacute", Wx: 600.000000, Wy: 0.000000},
	"ocircumflex":    {GlyphName: "ocircumflex", Wx: 600.000000, Wy: 0.000000},
	"odieresis":      {GlyphName: "odieresis", Wx: 600.000000, Wy: 0.000000},
	"oe":             {GlyphName: "oe", Wx: 600.000000, Wy: 0.000000},
	"ogonek":         {GlyphName: "ogonek", Wx: 600.000000, Wy: 0.000000},
	"ograve":         {GlyphName: "ograve", Wx: 600.000000, Wy: 0.000000},
	"ohungarumlaut":  {GlyphName: "ohungarumlaut", Wx: 600.000000, Wy: 0.000000},
	"omacron":        {GlyphName: "omacron", Wx: 600.000000, Wy: 0.000000},
	"one":            {GlyphName: "one", Wx: 600.000000, Wy: 0.000000},
	"onehalf":        {GlyphName: "onehalf", Wx: 600.000000, Wy: 0.000000},
	"onequarter":     {GlyphName: "onequarter", Wx: 600.000000, Wy: 0.000000},
	"onesuperior":    {GlyphName: "onesuperior", Wx: 600.000000, Wy: 0.000000},
	"ordfeminine":    {GlyphName: "ordfeminine", Wx: 600.000000, Wy: 0.000000},
	"ordmasculine":   {GlyphName: "ordmasculine", Wx: 600.000000, Wy: 0.000000},
	"oslash":         {GlyphName: "oslash", Wx: 600.000000, Wy: 0.000000},
	"otilde":         {GlyphName: "otilde", Wx: 600.000000, Wy: 0.000000},
	"p":              {GlyphName: "p", Wx: 600.000000, Wy: 0.000000},
	"paragraph":      {GlyphName: "paragraph", Wx: 600.000000, Wy: 0.000000},
	"parenleft":      {GlyphName: "parenleft", Wx: 600.000000, Wy: 0.000000},
	"parenright":     {GlyphName: "parenright", Wx: 600.000000, Wy: 0.000000},
	"partialdiff":    {GlyphName: "partialdiff", Wx: 600.000000, Wy: 0.000000},
	"percent":        {GlyphName: "percent", Wx: 600.000000, Wy: 0.000000},
	"period":         {GlyphName: "period", Wx: 600.000000, Wy: 0.000000},
	"periodcentered": {GlyphName: "periodcentered", Wx: 600.000000, Wy: 0.000000},
	"perthousand":    {GlyphName: "perthousand", Wx: 600.000000, Wy: 0.000000},
	"plus":           {GlyphName: "plus", Wx: 600.000000, Wy: 0.000000},
	"plusminus":      {GlyphName: "plusminus", Wx: 600.000000, Wy: 0.000000},
	"q":              {GlyphName: "q", Wx: 600.000000, Wy: 0.000000},
	"question":       {GlyphName: "question", Wx: 600.000000, Wy: 0.000000},
	"questiondown":   {GlyphName: "questiondown", Wx: 600.000000, Wy: 0.000000},
	"quotedbl":       {GlyphName: "quotedbl", Wx: 600.000000, Wy: 0.000000},
	"quotedblbase":   {GlyphName: "quotedblbase", Wx: 600.000000, Wy: 0.000000},
	"quotedblleft":   {GlyphName: "quotedblleft", Wx: 600.000000, Wy: 0.000000},
	"quotedblright":  {GlyphName: "quotedblright", Wx: 600.000000, Wy: 0.000000},
	"quoteleft":      {GlyphName: "quoteleft", Wx: 600.000000, Wy: 0.000000},
	"quoteright":     {GlyphName: "quoteright", Wx: 600.000000, Wy: 0.000000},
	"quotesinglbase": {GlyphName: "quotesinglbase", Wx: 600.000000, Wy: 0.000000},
	"quotesingle":    {GlyphName: "quotesingle", Wx: 600.000000, Wy: 0.000000},
	"r":              {GlyphName: "r", Wx: 600.000000, Wy: 0.000000},
	"racute":         {GlyphName: "racute", Wx: 600.000000, Wy: 0.000000},
	"radical":        {GlyphName: "radical", Wx: 600.000000, Wy: 0.000000},
	"rcaron":         {GlyphName: "rcaron", Wx: 600.000000, Wy: 0.000000},
	"rcommaaccent":   {GlyphName: "rcommaaccent", Wx: 600.000000, Wy: 0.000000},
	"registered":     {GlyphName: "registered", Wx: 600.000000, Wy: 0.000000},
	"ring":           {GlyphName: "ring", Wx: 600.000000, Wy: 0.000000},
	"s":              {GlyphName: "s", Wx: 600.000000, Wy: 0.000000},
	"sacute":         {GlyphName: "sacute", Wx: 600.000000, Wy: 0.000000},
	"scaron":         {GlyphName: "scaron", Wx: 600.000000, Wy: 0.000000},
	"scedilla":       {GlyphName: "scedilla", Wx: 600.000000, Wy: 0.000000},
	"scommaaccent":   {GlyphName: "scommaaccent", Wx: 600.000000, Wy: 0.000000},
	"section":        {GlyphName: "section", Wx: 600.000000, Wy: 0.000000},
	"semicolon":      {GlyphName: "semicolon", Wx: 600.000000, Wy: 0.000000},
	"seven":          {GlyphName: "seven", Wx: 600.000000, Wy: 0.000000},
	"six":            {GlyphName: "six", Wx: 600.000000, Wy: 0.000000},
	"slash":          {GlyphName: "slash", Wx: 600.000000, Wy: 0.000000},
	"space":          {GlyphName: "space", Wx: 600.000000, Wy: 0.000000},
	"sterling":       {GlyphName: "sterling", Wx: 600.000000, Wy: 0.000000},
	"summation":      {GlyphName: "summation", Wx: 600.000000, Wy: 0.000000},
	"t":              {GlyphName: "t", Wx: 600.000000, Wy: 0.000000},
	"tcaron":         {GlyphName: "tcaron", Wx: 600.000000, Wy: 0.000000},
	"tcommaaccent":   {GlyphName: "tcommaaccent", Wx: 600.000000, Wy: 0.000000},
	"thorn":          {GlyphName: "thorn", Wx: 600.000000, Wy: 0.000000},
	"three":          {GlyphName: "three", Wx: 600.000000, Wy: 0.000000},
	"threequarters":  {GlyphName: "threequarters", Wx: 600.000000, Wy: 0.000000},
	"threesuperior":  {GlyphName: "threesuperior", Wx: 600.000000, Wy: 0.000000},
	"tilde":          {GlyphName: "tilde", Wx: 600.000000, Wy: 0.000000},
	"trademark":      {GlyphName: "trademark", Wx: 600.000000, Wy: 0.000000},
	"two":            {GlyphName: "two", Wx: 600.000000, Wy: 0.000000},
	"twosuperior":    {GlyphName: "twosuperior", Wx: 600.000000, Wy: 0.000000},
	"u":              {GlyphName: "u", Wx: 600.000000, Wy: 0.000000},
	"uacute":         {GlyphName: "uacute", Wx: 600.000000, Wy: 0.000000},
	"ucircumflex":    {GlyphName: "ucircumflex", Wx: 600.000000, Wy: 0.000000},
	"udieresis":      {GlyphName: "udieresis", Wx: 600.000000, Wy: 0.000000},
	"ugrave":         {GlyphName: "ugrave", Wx: 600.000000, Wy: 0.000000},
	"uhungarumlaut":  {GlyphName: "uhungarumlaut", Wx: 600.000000, Wy: 0.000000},
	"umacron":        {GlyphName: "umacron", Wx: 600.000000, Wy: 0.000000},
	"underscore":     {GlyphName: "underscore", Wx: 600.000000, Wy: 0.000000},
	"uogonek":        {GlyphName: "uogonek", Wx: 600.000000, Wy: 0.000000},
	"uring":          {GlyphName: "uring", Wx: 600.000000, Wy: 0.000000},
	"v":              {GlyphName: "v", Wx: 600.000000, Wy: 0.000000},
	"w":              {GlyphName: "w", Wx: 600.000000, Wy: 0.000000},
	"x":              {GlyphName: "x", Wx: 600.000000, Wy: 0.000000},
	"y":              {GlyphName: "y", Wx: 600.000000, Wy: 0.000000},
	"yacute":         {GlyphName: "yacute", Wx: 600.000000, Wy: 0.000000},
	"ydieresis":      {GlyphName: "ydieresis", Wx: 600.000000, Wy: 0.000000},
	"yen":            {GlyphName: "yen", Wx: 600.000000, Wy: 0.000000},
	"z":              {GlyphName: "z", Wx: 600.000000, Wy: 0.000000},
	"zacute":         {GlyphName: "zacute", Wx: 600.000000, Wy: 0.000000},
	"zcaron":         {GlyphName: "zcaron", Wx: 600.000000, Wy: 0.000000},
	"zdotaccent":     {GlyphName: "zdotaccent", Wx: 600.000000, Wy: 0.000000},
	"zero":           {GlyphName: "zero", Wx: 600.000000, Wy: 0.000000},
}
