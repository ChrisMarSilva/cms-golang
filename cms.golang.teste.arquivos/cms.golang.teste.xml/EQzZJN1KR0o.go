package main

import (
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/antchfx/xmlquery"
	"github.com/antchfx/xpath"
)

const data1 = `
<plan>
<unit_amount_in_cents>
 <USD type="integer">4000</USD>
 <GBP type="integer">4000</GBP>
</unit_amount_in_cents>
<setup_fee_in_cents>
 <USD type="integer">4000</USD>
 <GBP type="integer">4000</GBP>
</setup_fee_in_cents>
</plan>
`

const data2 = `
<ValueList>
  <ArraySize>2</ArraySize>
  <v89BNZMpdlWXkuv>value1</v89BNZMpdlWXkuv>
  <v89N83oCrGhI7jh>value2</v89N83oCrGhI7jh>
</ValueList>
`

func main() {
	//fmt.Println(data1)
	//fmt.Println(data2)

	xmlstr := `
	<foo>
		<tag1>toto</tag1> 
		<tag2>tutu</tag2> 
		<foo2>
			<tag3>titi</tag3>
		</foo2> 
		<tag4>parsed</tag4> 
		<tag5>blabla</tag5>
	</foo>
	`

	var xmlstruct xmlTest

	err := xml.Unmarshal([]byte(xmlstr), &xmlstruct)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(xmlstr)
	//fmt.Printf("resp: %+v\n", xmlstruct)

	fetched := xmlstruct.Unparsed.GetContentByName("tag5")
	if len(fetched) > 0 {
		fmt.Printf("Getted by name: %+v\n", fetched)
	}

	s := `<?xml version="1.0" encoding="UTF-8" ?>
<breakfast_menu>
	<food>
		<name price="10">Berry-Berry Belgian Waffles</name>
		<description>Light Belgian waffles</description>
		<calories>900</calories>
	</food>
	<food>
		<name price="20">French Toast</name>
		<description>Thick slices</description>
		<calories>600</calories>
	</food>
	<food>
		<name price="30">Homestyle Breakfast</name>
		<description>Two eggs, bacon or sausage</description>
		<calories>950</calories>
	</food>	
</breakfast_menu>`

	doc, err := xmlquery.Parse(strings.NewReader(s))
	if err != nil {
		panic(err)
	}

	root := xmlquery.FindOne(doc, "//breakfast_menu")
	if n := root.SelectElement("//food/name"); n != nil {
		fmt.Printf("Name #%s\n", n.InnerText())
	}

	if n := root.SelectElement("//food[2]/name"); n != nil {
		fmt.Printf("Name #%s\n", n.InnerText())
	}

	for i, n := range xmlquery.Find(doc, "//food/name/@price") {
		fmt.Printf("Price #%d %s\n", i, n.InnerText())
	}

	for i, n := range xmlquery.Find(doc, "//food/calories") {
		fmt.Printf("Calories #%d %s\n", i, n.InnerText())
	}

	if n := root.SelectElement("//food[2]/name"); n != nil {
		fmt.Printf("Attr #%s\n", n.Attr)
	}

	if n := root.SelectElement("//food[2]/name"); n != nil {
		fmt.Printf("Data #%s\n", n.Data)
	}

	node := xmlquery.FindOne(doc, "//breakfast_menu/food[2]")
	if n := node.SelectElement("//description"); n != nil {
		fmt.Printf("Description #%s\n", n.InnerText())
	}

	expr, err := xpath.Compile("sum(//breakfast_menu/food/name/@price)")
	price := expr.Evaluate(xmlquery.CreateXPathNavigator(doc)).(float64)

	fmt.Printf("Total price: %f\n", price)

	countexpr, err := xpath.Compile("count(//breakfast_menu/food)")
	count := countexpr.Evaluate(xmlquery.CreateXPathNavigator(doc))

	fmt.Printf("Food Node Counts: %f\n", count)
}

type Value struct {
	ID    string
	Value string
}

type ValueList struct {
	Values []Value
}

type xmlTest struct {
	Tag4     string       `xml:"tag4"`
	Unparsed UnparsedTags `xml:",any"`
}

type UnparsedTag struct {
	XMLName xml.Name
	Content string `xml:",chardata"`
	//FullContent   string `xml:",innerxml"` // for debug purpose, allow to see what's inside some tags
}

// UnparsedTags store tags not handled by Unmarshal in a map, it should be labelled with `xml",any"`
type UnparsedTags map[string]string

func (m *UnparsedTags) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if *m == nil {
		*m = UnparsedTags{}
	}

	e := UnparsedTag{}
	err := d.DecodeElement(&e, &start)
	if err != nil {
		return err
	}

	//if _, ok := (*m)[e.XMLName.Local]; ok {
	//	return fmt.Errorf("UnparsedTags: UnmarshalXML: Tag %s:  multiple entries with the same name", e.XMLName.Local)
	//}
	(*m)[e.XMLName.Local] = e.Content

	return nil
}

func (u *UnparsedTags) GetContentByName(name string) string {
	return ((map[string]string)(*u))[name]
}

// UnparsedTags contains a list of tags not handled by unmarshall, it should be marked with `xml",any"`
type UnparsedTags struct {
	Tags []UnparsedTag
}

// UnparsedTag contains the tag informations
type UnparsedTag struct {
	XMLName xml.Name
	Content string `xml:",chardata"`
	//FullContent   string `xml:",innerxml"` // for debug purpose, allow to what's inside some tags
}

func (u *UnparsedTags) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var e UnparsedTag
	err := d.DecodeElement(&e, &start)
	if err != nil {
		return err
	}

	u.Tags = append(u.Tags, e)
	return nil
}

func (u *UnparsedTags) GetContentByName(name string) []string {
	content := make([]string, 0, 1)
	for _, t := range u.Tags {
		if t.XMLName.Local != name {
			continue
		}
		content = append(content, t.Content)
	}
	return content
}
