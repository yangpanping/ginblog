package databases

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/olivere/elastic/v7"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//索引结构
const maping = `
"mappings": {
    "properties": {
      "id":{
        "type": "keyword"
      },
      "name":{
        "type": "text",
        "analyzer": "ik_max_word"
      },
      "type":{
        "type": "keyword"
      },
      "attachment": {
        "properties": {
          "content":{
            "type": "text",
            "analyzer": "ik_smart"
          }
        }
      }
    }
  }`

//创建文件结构
type FileType struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Content string `json:"content"`
}
type abc struct {
	Took    int      `json:"took"`
	TimeOut string   `json:"time_out"`
	Shards  string   `_shards`
	Hits    []string `json:"hits"`
}

func GetESClient() (*elastic.Client, error) {

	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))

	fmt.Println("ES initialized...")

	return client, err

}

//建立索引
func CrateIndex(indexName string) {
	isexist := IndexIs(indexName)
	//存在就不进行创建，返回错误
	if isexist {
		log.Fatal("index is existed,can`t created")
		return
	}
	client, err := GetESClient()
	if err != nil {
		log.Fatal(err)
	}
	//创建索引
	indexCreate, err := client.CreateIndex(indexName).BodyString(maping).Do(context.Background())
	if err != nil {
		log.Fatal("创建index失败", err)
	}
	if !indexCreate.Acknowledged {
		log.Fatal("创建不成功")
	} else {
		fmt.Println("创建index成功")
	}

}

//测试索引是否存在
func IndexIs(indexName string) bool {
	client, err := GetESClient()
	if err != nil {
		log.Fatal(err)
	}
	exists, err := client.IndexExists(indexName).Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return exists

}

//写入数据
func CrateEssay(fileName string) {
	content := FileToBase64(fileName)
	fileData := &FileType{
		Name:    "简历",
		Type:    "docx",
		Content: content,
	}
	client, err := GetESClient()
	if err != nil {
		log.Fatal(err)
	}
	res, err := client.Index().Index("docwrite").Pipeline("attachment").BodyJson(fileData).Do(context.Background())
	if err != nil {
		log.Fatal("写入数据失败", err)
	}
	fmt.Println("写入数据成功，", res.Id, res.Index, res.Type)
}

//读取文件
func FileToBase64(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	//把数据转换为base64
	sEnc := base64.StdEncoding.EncodeToString(data)
	return sEnc
}

//查询文件
func SearchFile(data string) {
	client, err := GetESClient()
	if err != nil {
		log.Fatal(err)
	}
	phraseQuery := elastic.NewMatchPhraseQuery("query.match.attachment.content.query", "美").Analyzer("ik_smart")

	res, _ := client.Search().Index("docwrite").Query(phraseQuery).Do(context.Background())
	fmt.Println(res)
}

//采用http请求查询文件
func SerachFileByhttp() {
	url := "http://localhost:9200/docwrite/_search"
	conttentType := "application/json"
	data := `{
	"_source": [
		"name"
	],
	"query": {
		"match": {
			"attachment.content": {
				"query": "美",
				"analyzer": "ik_smart"
			}
		}
	},
	"highlight": {
		"boundary_scanner_locale": "zh_CN",
		"fields": {
			"attachment.content": {
				"pre_tags": [
					"<h3>"
				],
				"post_tags": [
					"</h3>"
				]
			}
		}
	}
}`
	resp, err := http.Post(url, conttentType, strings.NewReader(data))
	if err != nil {
		fmt.Printf("post failed,err :%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed ,err :%v\n", err)
		return
	}
	fmt.Println(string(b))

}
