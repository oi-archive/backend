package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strconv"
)

func APIMetadata(c *gin.Context) {
	c.JSON(200, gin.H{
		"static": false,
	})
}

func APIProblemSetList(c *gin.Context) {
	c.JSON(200, &ProblemSetFile)
}

func APIProblemSetMetadata(c *gin.Context) {
	problemset := c.Param("problemset")
	for _, i := range ProblemSets {
		if i.Id == problemset {
			c.JSON(200, gin.H{
				"name":    i.Name,
				"problem": len(i.ProblemArray),
				"page":    i.MaxPage,
			})
			return
		}
	}
	c.JSON(404, gin.H{
		"errorCode":    404,
		"errorMessage": "Problem set not found.",
	})
}

func APIProblemSetPage(c *gin.Context) {
	problemset := c.Param("problemset")
	page, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		c.JSON(404, gin.H{
			"errorCode":    404,
			"errorMessage": "page invalid",
		})
		return
	}

	for _, i := range ProblemSets {
		if i.Id == problemset {
			if page < 0 || page > i.MaxPage {
				c.JSON(404, gin.H{
					"errorCode":    404,
					"errorMessage": "page invalid",
				})
				return
			}
			c.JSON(200, i.ProblemArray[50*page-49:min(50*page, len(i.ProblemArray))])
			return
		}
	}
	c.JSON(404, gin.H{
		"errorCode":    404,
		"errorMessage": "Problem set not found.",
	})
}

func APIProblem(c *gin.Context) {
	problemset := c.Param("problemset")
	problem := c.Param("problem")
	for _, i := range ProblemSets {
		if i.Id == problemset {
			if _, ok := i.ProblemMap[problem]; !ok {
				c.JSON(404, gin.H{
					"errorCode":    404,
					"errorMessage": "Problem not found",
				})
				return
			}
			b, err := ioutil.ReadFile(BasePath + "/" + i.Id + "/" + problem + "/main.json")
			if err != nil {
				ThrowUnknownError(c, fmt.Errorf("error can't read file %s/%s/main.json : [%w] ", i.Id, problem, err))
				return
			}
			var res interface{}
			err = json.Unmarshal(b, &res)
			if err != nil {
				ThrowUnknownError(c, fmt.Errorf("error can't parse file %s/%s/main.json : [%w] ", i.Id, problem, err))
				return
			}
			des, err := ioutil.ReadFile(BasePath + "/" + i.Id + "/" + problem + "/description.md")
			if err != nil {
				ThrowUnknownError(c, fmt.Errorf("error can't read file %s/%s/description.md : [%w] ", i.Id, problem, err))
				return
			}
			if res, ok := res.(map[string]interface{}); ok {
				res["description"] = string(des)
				c.JSON(200, res)
			} else {
				ThrowUnknownError(c, fmt.Errorf("error can't parse file %s/%s/main.json", i.Id, problem))
				return
			}
			return
		}
	}
	c.JSON(404, gin.H{
		"errorCode":    404,
		"errorMessage": "Problem set not found.",
	})
}
