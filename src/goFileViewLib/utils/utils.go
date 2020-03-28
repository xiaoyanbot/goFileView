package utils

import(
	"os"
	"io"
	"log"
	"fmt"
	"path"
	"time"
	"strings"
	"os/exec"
	"math/rand"
	"crypto/md5"
)

/**
 * 比较字符串
 * 获得的字符串，和 flag 之间比较
 * onlinePerview?url=http://localhost:9090/123.docx
 * onlinePreview
 */
func ComparePath(filePathGive string, flag string) bool {

	//log.Println("filePathGive:" + filePathGive)
	//log.Println("flag:" + flag)
	//log.Println("截取:" + filePathGive[0:len(flag)])

	if len(filePathGive) >= len(flag){
		if strings.Compare(filePathGive[0:len(flag)], flag) == 0 {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func ConvertToPDF(filePath string) string {
	commandName := "libreoffice"
	params := []string{"--invisible","--convert-to","pdf","--outdir","cache/pdf/",filePath}
	if _,ok := interactiveToexec(commandName,params);ok{
		resultPath := "cache/pdf/"+ strings.Split(path.Base(filePath),".")[0] + ".pdf"
		if ok,_ := PathExists(resultPath);ok{
			return resultPath
		} else {
			return ""
		}
	} else {
		return ""
	}
}

func ConvertToImg(filePath string) string{
	fileName := strings.Split(path.Base(filePath),".")[0]
	//fileDir := path.Dir(filePath)
	fileExt := path.Ext(filePath)
	if fileExt != ".pdf" {
		return ""
	}
	os.Mkdir("cache/convert/"+fileName, os.ModePerm)
	commandName := "convert"
	params := []string{"-density","130",filePath,"cache/convert/"+fileName+"/%d.jpg"}
	if _,ok := interactiveToexec(commandName,params);ok{
		resultPath := "cache/convert/" + strings.Split(path.Base(filePath),".")[0]
		if ok,_ := PathExists(resultPath);ok{
			return resultPath
		} else {
			return ""
		}
	} else {
		return ""
	}
}

func interactiveToexec(commandName string,params []string) (string,bool){
	cmd := exec.Command(commandName,params...)
	buf, err := cmd.Output()
	if err != nil{
		log.Println("Error: <",err,"> when exec command read out buffer")
		return "",false
	} else {
		return string(buf),true
	}
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func GetFileMD5(filePath string) string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Println("Error: <",err,"> when open file to get md5")
		return ""
	}
	defer f.Close()
	md5hash := md5.New()
	if _, err := io.Copy(md5hash, f); err != nil {
		log.Println("Error: <",err,"> when get md5")
		return ""
	}
	f.Close()
	return fmt.Sprintf("%x",md5hash.Sum(nil))
}

func randString(len int) string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
    bytes := make([]byte, len)
    for i := 0; i < len; i++ {
        b := r.Intn(26) + 65
        bytes[i] = byte(b)
    }
    return string(bytes)
}

func IsInArr(key string,arr []string) bool {
	for i:= 0;i<len(arr);i++{
		if key == arr[i]{
			return true
		}
	}
	return false
}