package main
 
import (
	"fmt"
	"os"
	"path/filepath"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
	"sync"
	"time"
)



type Compliancer struct {
	regexpStr string
	fullpath string
}

func (obj Compliancer) IsDetected() bool {

	content := getFileContent(obj.fullpath)
	fragments := getContentFragments(content)
	IsDetected := compiliceAllFragment(fragments, obj.regexpStr)

    return IsDetected
}

func readExcel(filepath string )string {
	return "";
}

func readDoc(filepath string )string {
	return ""
}

func readImage(filepath string)string{
	return ""
}

func readPdf(filepath string)string{
	return ""
}

func readText(filepath string) string{
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}   
   return string(content)
}

func getFileContent(fullpath string) string {
	content := ""

	switch filepath.Ext(fullpath){
	case ".pdf":
		content = readPdf(fullpath)
	case ".xls":
		content = readExcel(fullpath)
	case ".doc":
		content = readDoc(fullpath)
	default:
		content = readText(fullpath)
	}
	return content
}

func getContentFragments(content string) []string{
	fragments := strings.Split(content," ")
	return fragments
}

func compiliceAllFragment(fragments []string, regexpStr string)bool{
	for _, fragment := range fragments {
		has_wrong_data := compilice(fragment,regexpStr)
		if true == has_wrong_data{
			return true
		}
	}
	return false
}

func compilice(content string,regexpStr string)bool{
	matchbool,err := regexp.MatchString(regexpStr,content)
	if err != nil{
		panic("[400] regexp match error")
	}
	return matchbool
}

func worker(id int, wg *sync.WaitGroup) {

}

func main() {
	
	var detectedFiles []string
	var wg sync.WaitGroup

	root := "D:\\"
	reStr := "^[A-Z]{1}[1-2]{1}[0-9]{8}$"

	s := time.Now()
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

			if false == info.IsDir(){
				c := Compliancer{regexpStr:reStr}
				c.fullpath = path

				if true == c.IsDetected() {
				// 	// fmt.Println("[+]",path)
					detectedFiles = append(detectedFiles,path)
					fmt.Println("[+] ",path)
				}else{
					fmt.Println("[-]",path)
				}
			}

        return nil
    })
    
	wg.Wait()

	if err != nil {
        panic(err)
    }

	fmt.Println("done.", time.Since(s))
	for _, element := range detectedFiles {
		// element is the element from someSlice for where we are
		fmt.Println("found ->",element)
	}


}