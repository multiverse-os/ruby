package rbmarshal

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"strings"

	//"math/big"
	"reflect"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

const (
	SUPPORTED_MAJOR_VERSION = 4
	SUPPORTED_MINOR_VERSION = 8

	NIL_SIGN         = '0'
	TRUE_SIGN        = 'T'
	FALSE_SIGN       = 'F'
	FIXNUM_SIGN      = 'i'
	RAWSTRING_SIGN   = '"'
	SYMBOL_SIGN      = ':'
	SYMBOL_LINK_SIGN = ';'
	OBJECT_SIGN      = 'o'
	OBJECT_LINK_SIGN = '@'
	ARRAY_SIGN       = '['
	IVAR_SIGN        = 'I'
	HASH_SIGN        = '{'
	BIGNUM_SIGN      = 'l'
	REGEXP_SIGN      = '/'
	CLASS_SIGN       = 'c'
	MODULE_SIGN      = 'm'
)

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r: bufio.NewReader(r)}
}

type Decoder struct {
	r       *bufio.Reader
	objects []interface{}
	symbols []string
}

func (d *Decoder) unmarshal() interface{} {
	typ, _ := d.r.ReadByte()

	switch typ {
	case NIL_SIGN: // 0 - nil
		return nil
	case TRUE_SIGN: // T - true
		return true
	case FALSE_SIGN: // F - false
		return false
	case FIXNUM_SIGN: // i - integer
		return d.parseInt()
	case RAWSTRING_SIGN: // " - string
		return d.parseString()
	case SYMBOL_SIGN: // : - symbol
		return d.parseSymbol()
	case SYMBOL_LINK_SIGN: // ; - symbol symlink
		return d.parseSymLink()
	case OBJECT_LINK_SIGN: // @ - object link
		return d.parseObjectLink()
	case IVAR_SIGN: // I - IVAR (encoded string or regexp)
		return d.parseIvar()
	case ARRAY_SIGN: // [ - array
		return d.parseArray()
	case OBJECT_SIGN: // o - object
		panic("not supported.")
	case HASH_SIGN: // { - hash
		return d.parseHash()
	case BIGNUM_SIGN: // l - bignum
		return d.parseBignum()
		// panic("not supported.")
	case REGEXP_SIGN: // / - regexp
		panic("not supported.")
	case CLASS_SIGN: // c - class
		panic("not supported.")
	case MODULE_SIGN: // m -module
		panic("not supported.")
	default:
		return nil
	}
}

func (d *Decoder) parseInt() int {
	var result int
	b, _ := d.r.ReadByte()
	c := int(int8(b))
	if c == 0 {
		return 0
	} else if 5 < c && c < 128 {
		return c - 5
	} else if -129 < c && c < -5 {
		return c + 5
	}
	cInt8 := int8(b)
	if cInt8 > 0 {
		result = 0
		for i := int8(0); i < cInt8; i++ {
			n, _ := d.r.ReadByte()
			result |= int(uint(n) << (8 * uint(i)))
		}
	} else {
		result = -1
		c = -c
		for i := 0; i < c; i++ {
			n, _ := d.r.ReadByte()
			result &= ^(0xff << uint(8*i))
			result |= int(n) << uint(8*i)
		}
	}
	return result
}

func (d *Decoder) parseBignum() int64 {
	b, _ := d.r.ReadByte()
	isNegative := b == '-'

	length := d.parseInt()
	buff := make([]byte, length*2)
	_, _ = d.r.Read(buff)

	result := int64(0)
	for i, d := range buff {
		result |= int64(d) << (uint(i) * 8)
	}

	if isNegative {
		return -1 * result
	} else {
		return result
	}
}

func (d *Decoder) parseSymbol() string {
	symbol := d.parseString()
	d.symbols = append(d.symbols, symbol)
	return symbol
}

func (d *Decoder) parseSymLink() string {
	index := d.parseInt()
	return d.symbols[index]
}

func (d *Decoder) parseObjectLink() interface{} {
	index := d.parseInt()
	if index >= len(d.objects) {
		return nil
	}
	return d.objects[index]
}

func (d *Decoder) parseString() string {
	length := d.parseInt()

	var result []byte

	readed := 0
	for readed<length {
		str := make([]byte, length-readed)
		n, _ := d.r.Read(str[:])
		result = append(result, str[:n]...)
		readed += n
	}

	return string(result)
}

type iVar struct {
	str string
	key string
	value interface{}
}

func gbk2utf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	return ioutil.ReadAll(reader)
}

func utf82gbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	return ioutil.ReadAll(reader)
}

func (d *Decoder) parseIvar() string {
	str := d.unmarshal()

	ivar := iVar{}

	symbolCharLen := d.parseInt()
	if symbolCharLen == 1 {
		// 第一个值key，第二个value，目前未用上
		key := d.unmarshal().(string) // :E
		value := d.unmarshal()                // T
		d.objects = append(d.objects, value)

		ivar.key = key
		ivar.value = value
	}

	strString := str.(string)

	// 需要准确编码处理
 	if ivar.key == "encoding" && ivar.value != nil {
 		encoding := ivar.value.(string)
		if (strings.ToLower(encoding) == "gbk") {
			decodeStr, err := gbk2utf8([]byte(strString))
			if err != nil {
				return ""
			}
			strString  = string(decodeStr)
		}
	}
	ivar.str = strString

	d.objects = append(d.objects, ivar)
	return strString
}

func (d *Decoder) parseArray() interface{} {
	size := d.parseInt()
	arr := make([]interface{}, size)

	for i := 0; i < int(size); i++ {
		arr[i] = d.unmarshal()
	}
	return arr
}

func (d *Decoder) parseHash() interface{} {
	size := d.parseInt()
	hash := make(map[string]interface{}, size)

	for i := 0; i < int(size); i++ {
		key := d.unmarshal()
		value := d.unmarshal()
		if key == nil {
			hash["nil"] = value
		} else {
			hash[key.(string)] = value
		}

	}

	return hash
}

func (d *Decoder) Decode(v interface{}) error {
	major, err := d.r.ReadByte()
	minor, err := d.r.ReadByte()

	if err != nil {
		return errors.New("cant decode MAJOR, MINOR version")
	}
	if major != SUPPORTED_MAJOR_VERSION || minor > SUPPORTED_MINOR_VERSION {
		return errors.New("unsupported marshal version")
	}

	val := reflect.ValueOf(v)

	if val.Kind() != reflect.Ptr {
		return errors.New("pointer need.")
	}

	r := d.unmarshal()
	if r == nil {
		v = nil
		return nil
	}

	decode(r, reflect.ValueOf(v))

	return nil
}


type Encoder struct {
	w            *bufio.Writer
	symbols      map[string]int
	symbolsIndex int
	objects      map[*interface{}]int
	objectsIndex int
	stringObj    iVar
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{
		w:            bufio.NewWriter(w),
		symbols:      map[string]int{},
		symbolsIndex: 0,
		objects:      map[*interface{}]int{},
		objectsIndex: 0,
	}
}

func (e *Encoder) Encode(v interface{}) error {
	if _, err := e.w.Write([]byte{SUPPORTED_MAJOR_VERSION, SUPPORTED_MINOR_VERSION}); err != nil {
		return err
	}

	if err := e.marshal(v); err != nil {
		return err
	}

	_ = e.w.Flush()
	return nil
}

func (e *Encoder) marshal(v interface{}) error {
	vKind := reflect.TypeOf(v).Kind()
	val := reflect.ValueOf(v)
	typ := reflect.TypeOf(v)

	if vKind == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}

	return e._marshalVal(val, typ)
}

func (e *Encoder) _marshalVal(val reflect.Value, typ reflect.Type) error {
	switch typ.Kind() {
	case reflect.Bool:
		return e.encBool(val.Bool())
	case reflect.Int:
		_ = e.w.WriteByte(FIXNUM_SIGN)
		return e.encInt(int(val.Int()))
	case reflect.Int64:
		_ = e.w.WriteByte(BIGNUM_SIGN)
		return e.encBigInt(val.Int())
	case reflect.String:
		return e.encString(val.String())
	case reflect.Struct:
		return e.encStruct(val, typ)
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		return e.encArray(val)
	}
	return nil
}

type TagInfo struct {
	name    string
	keyType string
}

func parseTag(tag string) (TagInfo, error) {
	tagInfo := TagInfo{}
	tags := strings.Split(tag, ";")
	tagInfo.name = tags[0]

	for _, t := range tags[1:] {
		split := strings.Split(t, ":")
		if len(split) != 2 {
			return tagInfo, errors.New("not valid struct tag")
		}

		if split[0] == "key" {
			tagInfo.keyType = split[1]
		}
	}

	return tagInfo, nil
}

func (e *Encoder) encStruct(value reflect.Value, typ reflect.Type) error {
	// 计算总共有几个是标记了ruby的
	rubyTag := 0
	for i := 0; i < value.NumField(); i++ {
		field := typ.Field(i)
		tagName := field.Tag.Get("ruby")
		if tagName != "" {
			rubyTag += 1
		}
	}
	err := e.w.WriteByte(HASH_SIGN)
	if err != nil {
		return err
	}

	err = e.encInt(int(rubyTag))
	if err != nil {
		return err
	}

	for i := 0; i < value.NumField(); i++ {
		field := typ.Field(i)
		rubyTag := field.Tag.Get("ruby")
		if rubyTag == "" {
			continue
		}

		info, err := parseTag(rubyTag)
		if err != nil {
			return err
		}

		if info.keyType == "symbol" || len(info.keyType) == 0 {
			err = e._encSymbol(info.name)
			if err != nil {
				return err
			}
		} else if info.keyType == "string"{
			err = e.encString(info.name)
			if err != nil {
				return err
			}
		} else {
			return errors.New(fmt.Sprintf("unknown tag key %s", info.keyType))
		}

		// 获取field的值，然后将值编码进去
		val := value.Field(i)
		fieldType := field.Type

		if fieldType.Kind() == reflect.Ptr {
			val = val.Elem()
			fieldType = fieldType.Elem()
		}

		err = e._marshalVal(val, fieldType)
		if err != nil {
			return err
		}
	}

	return nil
}

func (e *Encoder) encArray(val reflect.Value) error {
	err := e.w.WriteByte(ARRAY_SIGN)
	if err != nil {
		return err
	}

	// 获取数组或者slice的长度
	length := val.Len()

	err = e.encInt(length)
	if err != nil {
		return err
	}

	for i := 0; i < length; i++ {
		elem := val.Index(i)
		elemTyp := elem.Type()

		vKind := elemTyp.Kind()

		if vKind == reflect.Ptr {
			elem = elem.Elem()
			elemTyp = elemTyp.Elem()
		}

		err := e._marshalVal(elem, elemTyp)
		if err != nil {
			return err
		}
	}

	return nil
}

func (e *Encoder) encBool(val bool) error {
	if val {
		return e.w.WriteByte(TRUE_SIGN)
	}
	return e.w.WriteByte(FALSE_SIGN)
}

func (e *Encoder) encInt(i int) error {
	var len int

	if i == 0 {
		return e.w.WriteByte(0)
	} else if 0 < i && i < 123 {
		return e.w.WriteByte(byte(i + 5))
	} else if -124 < i && i <= -1 {
		return e.w.WriteByte(byte(i - 5))
	} else if 122 < i && i <= 0xff {
		len = 1
	} else if 0xff < i && i <= 0xffff {
		len = 2
	} else if 0xffff < i && i <= 0xffffff {
		len = 3
	} else if 0xffffff < i && i <= 0x3fffffff {
		//for compatibility with 32bit Ruby, Fixnum should be less than 1073741824
		len = 4
	} else if -0x100 <= i && i < -123 {
		len = -1
	} else if -0x10000 <= i && i < -0x100 {
		len = -2
	} else if -0x1000000 <= i && i < -0x100000 {
		len = -3
	} else if -0x40000000 <= i && i < -0x1000000 {
		//for compatibility with 32bit Ruby, Fixnum should be greater than -1073741825
		len = -4
	}

	if err := e.w.WriteByte(byte(len)); err != nil {
		return err
	}
	if len < 0 {
		len = -len
	}

	for c := 0; c < len; c++ {
		if err := e.w.WriteByte(byte(i >> uint(8*c) & 0xff)); err != nil {
			return err
		}
	}

	return nil
}

/**
这是一个半成品，只能针对int64做解码（可以满足一部分需求）)
*/
func (e *Encoder) encBigInt(v int64) error {
	buff := bytes.NewBuffer(nil)
	// 写入符号
	if v >= 0 {
		_ = e.w.WriteByte('+')
	} else {
		_ = e.w.WriteByte('-')
	}

	// 计算字节数
	sz := 0

	d := v
	if v < 0 {
		d = -d
	}

	for {
		sz++
		d >>= 8
		if d == 0 {
			break;
		}
	}

	length := int(math.Ceil(float64(sz) / 2))
	// 因为是int64，所以一半的长度是4个字节
	if err := e.encInt(length); err != nil {
		return err
	}

	d = v
	if v < 0 {
		d = -d
	}

	w := 0
	for i := 0; i < 64; i += 8 {
		buff.WriteByte(byte(d & 0xff))
		w++
		d >>= 8
		if d == 0 {
			break;
		}
	}
	for w < sz {
		buff.WriteByte(0)
		w++
	}

	_, err := e.w.Write(buff.Bytes())
	return err
}

func (e *Encoder) _encRawString(str string) error {
	// | len (Fixnum) | stirng |
	if err := e.encInt(len(str)); err != nil {
		return err
	}

	_, err := e.w.WriteString(str)
	return err
}

func (e *Encoder) encString(str string) error {
	// | I | " | RawString( string ) | FixNum( 1 ) | Symbol( E ) | True |
	if _, err := e.w.Write([]byte{IVAR_SIGN, RAWSTRING_SIGN}); err != nil {
		return err
	}
	if err := e._encRawString(str); err != nil {
		return err
	}
	if err := e.encInt(1); err != nil {
		return err
	}
	if err := e._encSymbol("E"); err != nil {
		return err
	}
	return e.encBool(true)
}

func (e *Encoder) _encSymbol(str string) error {
	if index, ok := e.symbols[str]; ok {
		if err := e.w.WriteByte(SYMBOL_LINK_SIGN); err != nil {
			return err
		}
		return e.encInt(index)
	}

	e.symbols[str] = e.symbolsIndex
	e.symbolsIndex++

	if err := e.w.WriteByte(SYMBOL_SIGN); err != nil {
		return err
	}
	if err := e.encInt(len(str)); err != nil {
		return err
	}
	_, err := e.w.WriteString(str)
	return err
}

func (e *Encoder) _encObject() error {
	return nil
}
