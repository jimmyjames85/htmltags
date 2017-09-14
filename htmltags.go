package htmltags

import (
	"bytes"
	"fmt"
)

type Document struct {
	bytes.Buffer
}

func (d *Document) HTML(attr ...string)                { fmt.Fprintf(d, "%s", HTML(attr...)) }
func (d *Document) HTMLClose()                         { fmt.Fprintf(d, "%s", HTMLClose()) }
func (d *Document) H1(attr ...string)                  { fmt.Fprintf(d, "%s", H1(attr...)) }
func (d *Document) H1Close()                           { fmt.Fprintf(d, "%s", H1Close()) }
func (d *Document) H2(attr ...string)                  { fmt.Fprintf(d, "%s", H2(attr...)) }
func (d *Document) H2Close()                           { fmt.Fprintf(d, "%s", H2Close()) }
func (d *Document) H3(attr ...string)                  { fmt.Fprintf(d, "%s", H3(attr...)) }
func (d *Document) H3Close()                           { fmt.Fprintf(d, "%s", H3Close()) }
func (d *Document) H4(attr ...string)                  { fmt.Fprintf(d, "%s", H4(attr...)) }
func (d *Document) H4Close()                           { fmt.Fprintf(d, "%s", H4Close()) }
func (d *Document) Header1(header string)              { fmt.Fprintf(d, "%s", Header1(header)) }
func (d *Document) Header2(header string)              { fmt.Fprintf(d, "%s", Header2(header)) }
func (d *Document) Header3(header string)              { fmt.Fprintf(d, "%s", Header3(header)) }
func (d *Document) Header4(header string)              { fmt.Fprintf(d, "%s", Header4(header)) }
func (d *Document) BR()                                { fmt.Fprintf(d, "%s", BR()) }
func (d *Document) HR()                                { fmt.Fprintf(d, "%s", HR()) }
func (d *Document) P(attr ...string)                   { fmt.Fprintf(d, "%s", P(attr...)) }
func (d *Document) PClose()                            { fmt.Fprintf(d, "%s", PClose()) }
func (d *Document) Table(attr ...string)               { fmt.Fprintf(d, "%s", Table(attr...)) }
func (d *Document) TableClose()                        { fmt.Fprintf(d, "%s", TableClose()) }
func (d *Document) TD(attr ...string)                  { fmt.Fprintf(d, "%s", TD(attr...)) }
func (d *Document) TDClose()                           { fmt.Fprintf(d, "%s", TDClose()) }
func (d *Document) TR(attr ...string)                  { fmt.Fprintf(d, "%s", TR(attr...)) }
func (d *Document) TRClose()                           { fmt.Fprintf(d, "%s", TRClose()) }
func (d *Document) TH(attr ...string)                  { fmt.Fprintf(d, "%s", TH(attr...)) }
func (d *Document) THClose(attr ...string)             { fmt.Fprintf(d, "%s", THClose(attr...)) }
func (d *Document) TableRow(cells ...interface{})      { fmt.Fprintf(d, "%s", TableRow(cells...)) }
func (d *Document) TableHeader(headers ...interface{}) { fmt.Fprintf(d, "%s", TableHeader(headers...)) }

func (d *Document) Paragraph(data string, attr ...string) {
	fmt.Fprintf(d, "%s", Paragraph(data, attr...))
}

func (d *Document) Print(format string) (int, error) {
	return d.Printf(format, []interface{}{})
}

func (d *Document) Println(format string, a ...interface{}) (int, error) {
	return d.Printf(format+"\n", a...)
}

func (d *Document) Printf(format string, a ...interface{}) (int, error) {
	return fmt.Fprintf(d, format, a...)
}

func HTML(attr ...string) string    { return tagWithAttrNL("html", attr...) }
func HTMLClose() string             { return "</html>\n" }
func H1(attr ...string) string      { return tagWithAttr("h1", attr...) }
func H1Close() string               { return closeTag("h1") }
func H2(attr ...string) string      { return tagWithAttr("h2", attr...) }
func H2Close() string               { return closeTag("h2") }
func H3(attr ...string) string      { return tagWithAttr("h3", attr...) }
func H3Close() string               { return closeTag("h3") }
func H4(attr ...string) string      { return tagWithAttr("h4", attr...) }
func H4Close() string               { return closeTag("h4") }
func Header1(header string) string  { return fmt.Sprintf("%s%s%s", H1(), header, H1Close()) }
func Header2(header string) string  { return fmt.Sprintf("%s%s%s", H2(), header, H2Close()) }
func Header3(header string) string  { return fmt.Sprintf("%s%s%s", H3(), header, H3Close()) }
func Header4(header string) string  { return fmt.Sprintf("%s%s%s", H4(), header, H4Close()) }
func BR() string                    { return "<br>\n" }
func HR() string                    { return "<hr>\n" }
func P(attr ...string) string       { return tagWithAttr("p", attr...) }
func PClose() string                { return closeTag("p") }
func Table(attr ...string) string   { return tagWithAttrNL("table", attr...) }
func TableClose() string            { return closeTag("table") }
func TD(attr ...string) string      { return tagWithAttr("td", attr...) }
func TDClose() string               { return closeTag("td") }
func TR(attr ...string) string      { return tagWithAttrNL("tr", attr...) }
func TRClose() string               { return closeTag("tr") }
func TH(attr ...string) string      { return tagWithAttr("th", attr...) }
func THClose(attr ...string) string { return closeTag("th") }

func TableRow(cells ...interface{}) string {
	ret := new(bytes.Buffer)
	fmt.Fprintf(ret, TR())
	for _, c := range cells {
		fmt.Fprintf(ret, "%s%v%s", TD(), c, TDClose())
	}
	fmt.Fprintf(ret, TRClose())
	return ret.String()
}

func TableHeader(headers ...interface{}) string {
	ret := new(bytes.Buffer)
	fmt.Fprintf(ret, TR())
	for _, h := range headers {
		fmt.Fprintf(ret, "%s%v%s", TH(), h, THClose())
	}
	fmt.Fprintf(ret, TRClose())
	return ret.String()
}

func Paragraph(data string, attr ...string) string {
	return fmt.Sprintf("%s%s%s", P(attr...), data, PClose())
}

func tagWithAttrNL(tag string, attr ...string) string {
	return fmt.Sprintf("%s\n", tagWithAttr(tag, attr...))
}

func tagWithAttr(tag string, attr ...string) string {
	ret := new(bytes.Buffer)
	fmt.Fprintf(ret, "<%s", tag)
	for _, a := range attr {
		fmt.Fprintf(ret, " %s", a)
	}
	fmt.Fprintf(ret, ">")
	return ret.String()

}

func closeTag(tag string) string {
	var d Document
	d.BR()
	return fmt.Sprintf("</%s>\n", tag)
}
