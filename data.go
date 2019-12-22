package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var BasePath string

type ProblemSet struct {
	Name         string
	Id           string
	ProblemMap   map[string]string
	ProblemArray ProblemList
	Css          []byte
	MaxPage      int
}

var ProblemSets []ProblemSet

var ProblemSetFile []struct {
	Name string
	Id   string
}

type ProblemList []struct {
	Title string `json:"title"`
	Pid   string `json:"pid"`
}

func MakeReadFileError(filename string, err error) error {
	return fmt.Errorf("Read File %s error: [%w]", filename, err)
}

func UpdateData() error {
	b, err := ioutil.ReadFile(BasePath + "/problemset.json")
	if err != nil {
		return MakeReadFileError("problemset.json", err)
	}

	err = json.Unmarshal(b, &ProblemSetFile)
	if err != nil {
		return err
	}

	ProblemSets = []ProblemSet{}
	for _, i := range ProblemSetFile {
		ps := ProblemSet{Name: i.Name, Id: i.Id}
		path := BasePath + "/" + i.Id
		b, err := ioutil.ReadFile(path + "/problemlist.json")
		if err != nil {
			return MakeReadFileError(path+"/problemlist.json", err)
		}
		list := ProblemList{}
		err = json.Unmarshal(b, &list)
		if err != nil {
			return err
		}
		ps.MaxPage = (len(list)-1)/50 + 1
		ps.ProblemArray = list
		ps.ProblemMap = make(map[string]string)
		for _, j := range list {
			ps.ProblemMap[j.Pid] = j.Title
		}
		ProblemSets = append(ProblemSets, ps)
	}

	return nil

}
