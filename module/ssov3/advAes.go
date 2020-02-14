package ssov3

import (
	"rest-shell/pkg/utils/syslog"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type ecb struct {
	b cipher.Block
	blockSize int
}

func Base64URLDecode(data string) ([]byte, error) {
	var missing = (4 - len(data)%4) % 4
	data += strings.Repeat("=", missing)
	base64.URLEncoding.DecodeString(data)
	return base64.URLEncoding.DecodeString(data)
}

func Base64UrlSafeEncode(source []byte) string {
	// Base64 Url Safe is the same as Base64 but does not contain '/' and '+' (replaced by '_' and '-') and trailing '=' are removed.
	bytearr := base64.StdEncoding.EncodeToString(source)
	safeurl := strings.Replace(string(bytearr), "/", "_", -1)
	safeurl = strings.Replace(safeurl, "+", "-", -1)
	safeurl = strings.Replace(safeurl, "=", "", -1)
	return safeurl
}

func AesDecrypt(crypted, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("err is:", err)
		return nil
	}
	blockMode := NewECBDecrypter(block)
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	fmt.Println("source is :", origData, string(origData))
	return origData
}

func AesEncrypt(src, key string) []byte {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println("key error1", err)
	}
	if src == "" {
		fmt.Println("plain content empty")
	}
	ecb := NewECBEncrypter(block)
	content := []byte(src)
	content = PKCS5Padding(content, block.BlockSize())
	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)
	return crypted
}


func SrpNameDecode(encodeStr string,appid string,srpToken string, srpkeyStr string,defaultAdminKey string )(string ,error){
	decodeByte ,err := Base64URLDecode(srpToken)
	encodeByte  ,err := Base64URLDecode(encodeStr)
	if err != nil {
		fmt.Println(fmt.Sprintf("[ normal ] - [ debug ] generate client name decode error :%v",err))
		fmt.Println(fmt.Sprintf("[ normal ] - [ error ] generate client name decode error"))
		return "",err
	}
	srpkey := []byte(srpkeyStr)
	srpdecode := AesDecrypt(decodeByte,srpkey)
	str := string(srpdecode)
	if len(str)<=12{
		return "",errors.New("srpToken invalidate")
	}
	re :=regexp.MustCompile("[0-9]*")
	if !re.MatchString(str[0:10]){
		return "",errors.New("srpToken invalidate")
	}
	TokenTimeStr:= str[0:10]
	src := TokenTimeStr +"-"+defaultAdminKey+ "-" + appid
	has := md5.Sum([]byte(src))
	md5str1 := fmt.Sprintf("%x", has)
	key := md5str1[0:24]
	srpNameByte :=AesDecrypt(encodeByte,[]byte(key))
	/*if err!=nil {
		fmt.Println(fmt.Sprintf("[ normal ] - [ debug ] generate srp name decode error :%v",err))
		fmt.Println(fmt.Sprintf("[ normal ] - [ error ] generate srp name decode error"))
		return "",err
	}*/
	return string(srpNameByte),nil
}

func GetSrpToken(timenow int64) string {

	curTime := strconv.FormatInt(timenow, 10)
	src := curTime + "-" + os.Getenv("SrpName")
	fmt.Printf("----------\r\nsrp token source data:%s\r\n", src)
	key := "ssoisno12345678987654321"
	crypted := AesEncrypt(src, key)
	result := Base64UrlSafeEncode(crypted)
	fmt.Printf("----------\r\nsrp token base64UrlSafe result:%s\r\n", result)
	return result
}


func  SrpNameEncode (appid string,serviceName string,timeNow int64,defaultAdminKey string)(string,error){
	curTime := strconv.FormatInt(timeNow, 10)
	src := curTime +"-"+defaultAdminKey+"-" + appid
	has := md5.Sum([]byte(src))
	md5str1 := fmt.Sprintf("%x", has)
	fmt.Println(md5str1)
	fmt.Println(len(md5str1))
	key := md5str1[0:24]
	fmt.Println(key)
	fmt.Println(len(key))

	crypted := AesEncrypt(serviceName, key)
/*	if err!=nil{
		fmt.Println(fmt.Sprintf("[ normal ] - [ debug ] generate srp name encode error :%v",err))
		fmt.Println(fmt.Sprintf("[ normal ] - [ error ] generate srp name encode error"))
		return "",err
	}*/
	result := Base64UrlSafeEncode(crypted)
	return result,nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	if unpadding >=length || unpadding <0 {
		return nil
	}
	return origData[:(length - unpadding)]
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

type ecbEncrypter ecb

// NewECBEncrypter returns a BlockMode which encrypts in electronic code book
// mode, using the given Block.
func NewECBEncrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbEncrypter)(newECB(b))
}
func (x *ecbEncrypter) BlockSize() int { return x.blockSize }
func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

type ecbDecrypter ecb

// NewECBDecrypter returns a BlockMode which decrypts in electronic code book
// mode, using the given Block.
func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbDecrypter)(newECB(b))
}
func (x *ecbDecrypter) BlockSize() int { return x.blockSize }
func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		LOG.Error("crypto/cipher: input not full blocks")
		return
		//panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		LOG.Error("crypto/cipher: output smaller than input")
		return
		//panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}
//longzhang.li add 2018.11.14