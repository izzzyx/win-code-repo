package main

import (
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"
)

const n = 255

type User struct {
	Email  string `mcl:"email"`
	Name   string `mcl:"name"`
	Age    int    `mcl:"age"`
	Github string `mcl:"github" default:"a8m"`
}

type UserKv struct {
	Email   string `kv:"email,omitempty"`
	Name    string `kv:"name,omitempty"`
	Github  string `kv:"github,omitempty"`
	private string
}

func main() {
	// var (
	// 	a int8
	// 	b int16
	// 	c uint
	// 	d float32
	// )

	// SetValueOfNumberType(&a)
	// SetValueOfNumberType(&b)
	// SetValueOfNumberType(&c)
	// SetValueOfNumberType(&d)

	// // todo a 的值为什么是0？
	// fmt.Printf("%+v\n", a)
	// fmt.Printf("%+v\n", b)
	// fmt.Printf("%+v\n", c)
	// fmt.Printf("%+v\n", d)

	// m0 := make(map[int]string)
	// m1 := make(map[interface{}]string)
	// m2 := make(map[interface{}]interface{})

	// s := "1=foo,2=bar,3=baz"
	// fmt.Println(decodeMap(s, &m0), m0)
	// fmt.Println(decodeMap(s, &m1), m1)
	// fmt.Println(decodeMap(s, &m2), m2)

	// var (
	// 	v0 *User
	// 	v2 = new(User)
	// 	v3 struct{Name string}
	// 	s = "Name=Ariel,Github=a8m"
	// )

	// fmt.Println(decodeStruct(s, v0), v0)
	// fmt.Println(decodeStruct(s, v2), v2)
	// fmt.Println(decodeStruct(s, &v3), v3)

	// var (
	// 	u = UserKv{Name: "Ariel", Github: "a8m"}
	// 	v = struct{
	// 		A, B, C string
	// 	}{
	// 		"foo",
	// 		"bar",
	// 		"baz",
	// 	}
	// 	w = &UserKv{}
	// )

	// us, _ := encode(u)
	// vs, _ := encode(v)
	// ws, err := encode(w)
	
	// fmt.Println(us)
	// fmt.Println(vs)
	// fmt.Println(ws)
	// fmt.Println(err)
}

// 读取结构体的tags
func ReadTagsOfStruct() {
	var u interface{} = User{}

	// get the static type of u, which represents by reflect.Type
	t := reflect.TypeOf(u)
	// use t.Kind() get the Kind obj, the reflect.Kind is the type of reflect.Type, which means the underlying type of u
	// 也就是说，会返回go的内置类型
	if t.Kind() != reflect.Struct {
		return
	}
	for i := 0; i < t.NumField(); i++ {
		// func (t *Type) Field(index int) StructField return a StructField obj
		// it panics if t /is not a struct
		// it panics if i is out of range
		f := t.Field(i)
		fmt.Println(f.Tag.Get("mcl"), f.Tag.Get("default"))
	}
}

// get和set结构体的fields
func GetAndSetFieldsOfStruct() {
	u := &User{Name: "Ariel"}
	// Elem returns the value that the pointer u points to
	v := reflect.ValueOf(u).Elem()
	f := v.FieldByName("Github")
	// if this field is defined and can be set
	if !f.IsValid() || !f.CanSet() {
		return
	}
	// use f.String to get value of field
	if f.Kind() != reflect.String || f.String() != "" {
		return
	}
	f.SetString("a8m")
	fmt.Printf("Github username change to %q\n", u.Github)
}

// 在不知道slice的类型的情况下，给slice填充string值
func FullFillSliceWithString() {
	var a []string
	var b []interface{}
	var c []io.Writer
	fill(&a)
	fill(&b)
	fill(&c)
	fmt.Printf("%+v", a)
	fmt.Printf("%+v", b)
	fmt.Printf("%+v", c)
}

func fill(i interface{}) error {
	v := reflect.ValueOf(i)
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("non-pointer %v", v.Type())
	}
	// get the value that the pointer points to
	v = v.Elem()
	// check type of param
	if v.Kind() != reflect.Slice {
		return fmt.Errorf("non-slice param %v", i)
	}
	// make a slice use v.Type and set to v
	// func MakeSlice(typ reflect.Type, len, cap int)
	v.Set(reflect.MakeSlice(v.Type(), 3, 3))
	if !canAssign(v.Index(0)) {
		return fmt.Errorf("can't assign string to params")
	}
	for i, w := range []string{"foo", "bar", "baz"} {
		v.Index(i).Set(reflect.ValueOf(w))
	}
	return nil
}

func canAssign(v reflect.Value) bool {
	// v是string类型或者空结构体
	return v.Kind() == reflect.String || (v.Kind() == reflect.Interface && v.NumMethod() == 0)
}

// 给一个number类型设置值
func SetValueOfNumberType(i interface{}) error {
	v := reflect.ValueOf(i)
	// v必须是可以设置值的（也就是指针）
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("can not set value")
	}
	// get the value that v points to
	v = v.Elem()
	switch v.Kind() {
	// 带符号和无符号的整型需要不同的set方法
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if v.OverflowInt(n) {
			return fmt.Errorf("the n cannot be represent by v's type")
		}
		v.SetInt(n) // reflect包里面的int底层一律用int64
	case reflect.Uint, reflect.Uintptr, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if v.OverflowUint(n) {
			return fmt.Errorf("the n cannot be represent by v's type")
		}
		v.SetUint(n)
	case reflect.Float32, reflect.Float64:
		if v.OverflowFloat(n) {
			return fmt.Errorf("the n cannot be represent by v's type")
		}
		v.SetFloat(n)
	default:
		return fmt.Errorf("can't assign value to a non-number type")
	}
	return nil
}

// 解析内容是kv pairs的字符串并写入map结构
func decodeMap(s string, i interface{}) error {
	v := reflect.ValueOf(i)
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("not-pointer %v", v.Type())
	}
	v = v.Elem()
	t := v.Type()
	if v.IsNil() {
		fmt.Printf("got a nil i, create a new map")
		v.Set(reflect.MakeMap(t))
	}
	// assume that the input is valid, so that ignore check
	for _, kv := range strings.Split(s, ",") {
		s := strings.Split(kv, "=")
		n, err := strconv.Atoi(s[0])
		if err != nil {
			return fmt.Errorf("input kv pair has non integer key")
		}
		k, e := reflect.ValueOf(n), reflect.ValueOf(s[1])
		kType := t.Key() // 可以调用Type的key()方法获取map的key的类型，如果t不是map会引起panic？
		// check if input key type can converted to map key type
		if !k.Type().ConvertibleTo(kType) {
			return fmt.Errorf("input key type cannot convert to map key type")
		}
		// 为啥要convert啊？
		k = k.Convert(kType)
		// 获取map的value的type
		et := t.Elem()
		// 如果map的value不是map && map的value不能转成
		if et.Kind() != v.Kind() && !e.Type().ConvertibleTo(et) {
			return fmt.Errorf("input value type cannot conver to map value type")
		}
		v.SetMapIndex(k, e.Convert(et))
	}
	return nil
}

// 解析内容是kv paris的字符串并写入struct
func decodeStruct(s string, i interface{}) error {
	v := reflect.ValueOf(i)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return fmt.Errorf("require non-nil input")
	}
	v = v.Elem()
	// assume that the input is valid
	for _, kv := range strings.Split(s, ",") {
		s := strings.Split(kv, "=")
		f := v.FieldByName(s[0])
		// make sure field is defined and can be set value
		if !f.IsValid() || !f.CanSet() {
			continue
		}
		f.SetString(s[1])
	}
	return nil
}
// 将struct转为kv pair字符串
func encode(i interface{}) (string, error) {
	v := reflect.ValueOf(i)
	t := v.Type()
	// v.Kind和t.Kind有什么区别？
	if t.Kind() != reflect.Struct {
		return "", fmt.Errorf("type %s is not struct", t.Kind())
	}
	var s []string
	// t.NumField和v.NumField有什么区别？
	for i:=0 ; i< t.NumField(); i++ {
		f := t.Field(i)
		// 获取unexport的field name
		// 在go version >= 1.17时，用StructField.IsExported代替
		if f.PkgPath != "" {
			continue
		}
		fv := v.Field(i)
		key, omit := readTag(f)
		if omit && fv.String() == "" {
			continue
		}
		s = append(s, fmt.Sprintf("%s=%s", key, fv.String()))
	}
	return strings.Join(s, ","), nil
}

func readTag(f reflect.StructField) (string, bool) {
	val, ok := f.Tag.Lookup("kv")
	if !ok {
		return f.Name, false
	}
	opts := strings.Split(val, ",")
	omit := false
	if len(opts) == 2 {
		omit = opts[1] == "omitempty"
	}
	return opts[0], omit
}
// check对象的底层类型是不是接口
// warp a reflect.Value with pointer(T => *T)
// 函数调用
// 1. 调用一个无参数且没有返回值的方法
// 2. 调用一个有多个参数且有返回值的方法
// 3. 动态调用一个方法，类似于template/text pacakge
// 4. 调用一个接收可变参数的方法
// 5. 在运行时创建一个方法
