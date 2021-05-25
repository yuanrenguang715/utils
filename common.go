package utils

import (
	"bytes"
	"crypto/md5"
	cr "crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
	"math/rand"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

//随机数种子
var Rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

//序列化
func ToString(v interface{}) string {
	data, _ := json.Marshal(v)
	return string(data)
}

//md5加密
func MD5(data string) string {
	m := md5.Sum([]byte(data))
	return hex.EncodeToString(m[:])
}

func MD5Encrypt16(data string) string {
	m := md5.Sum([]byte(data))
	return hex.EncodeToString(m[:])[8:24]
}

// 获取随机数
func GetRandNumber(n int) int {
	return Rnd.Intn(n)
}

func UrlEncode(s string) string {
	return url.QueryEscape(s)
}

func VersionToInt(version string) int {
	version = strings.Replace(version, ".", "", -1)
	n, _ := strconv.Atoi(version)
	return n
}

func Version2Flt(appversion string) (version float64) {
	arr := strings.SplitN(appversion, ".", 2)
	if len(arr) > 1 {
		arr[1] = strings.Replace(arr[1], ".", "", -1)
		appversion = arr[0] + "." + arr[1]
	}
	version, _ = strconv.ParseFloat(appversion, 64)
	return
}

//sort map by key
func SortMapByKey2Str(m map[string]interface{}) string {
	// To store the keys in slice in sorted order
	var keys []string
	var s string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	// To perform the opertion you want
	for _, k := range keys {
		if m[k] != nil {
			s += k + "=" + fmt.Sprint(m[k]) + "&"
		}
	}
	return strings.TrimSuffix(s, "&")
}

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func Json2Map(jsonstr []byte) (s map[string]string, err error) {
	var result map[string]string
	if err := json.Unmarshal(jsonstr, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// 将字符串数组转化为逗号分割的字符串形式  ["str1","str2","str3"] >>> "str1,str2,str3"
func StrListToString(strList []string) (str string) {
	if len(strList) > 0 {
		for k, v := range strList {
			if k == 0 {
				str = v
			} else {
				str = str + "," + v
			}
		}
		return
	}
	return ""
}

func RandIntNum(min, max int64) int64 {
	maxBigInt := big.NewInt(max)
	i, _ := cr.Int(cr.Reader, maxBigInt)
	iInt64 := i.Int64()
	if iInt64 < min {
		iInt64 = RandIntNum(min, max) //应该用参数接一下
	}
	return iInt64
}

func PageCount(count, pagesize int) int {
	if count%pagesize > 0 {
		return count/pagesize + 1
	} else {
		return count / pagesize
	}
}

func StartIndex(page, pagesize int) int {
	if page > 1 {
		return (page - 1) * pagesize
	}
	return 0
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func SetField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		return errors.New("Provided value type didn't match obj field type")
	}

	structFieldValue.Set(val)
	return nil
}

func ReadFile(path string) ([]byte, error) {
	inputFile, inputError := os.Open(path)
	if inputError != nil {
		fmt.Println("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		// exit the function on error
		return []byte(""), inputError
	}
	defer inputFile.Close()
	fb, err := ioutil.ReadAll(inputFile)
	return fb, err
}

// 新建上传请求
func NewUploadRequest(linkUrl string, params map[string]string, name, path string) (*http.Request, error) {
	fp, err := os.Open(path) // 打开文件句柄
	if err != nil {
		return nil, err
	}
	defer fp.Close()
	body := &bytes.Buffer{}                                       // 初始化body参数
	writer := multipart.NewWriter(body)                           // 实例化multipart
	part, err := writer.CreateFormFile(name, filepath.Base(path)) // 创建multipart 文件字段
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, fp) // 写入文件数据到multipart
	for key, val := range params {
		_ = writer.WriteField(key, val) // 写入body中额外参数
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", linkUrl, body) // 新建请求
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "multipart/form-data") // 设置请求头,!!!非常重要，否则远端无法识别请求
	return req, nil
}

func Float64Decimal(value float64,bint int) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%."+strconv.Itoa(bint)+"f", value), 64)
	return value
}