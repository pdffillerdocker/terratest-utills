// Copyright (c) 2021, airSlate, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package terratest_utills

import (
    "strings"
    "fmt"
    "io/ioutil"
    "gopkg.in/yaml.v2"
)

func GetVarFromYamlFile (filename string, conf interface{}) (error) {

    buf, err := ioutil.ReadFile(filename)
    if err != nil {
        return err
    }

    return yaml.Unmarshal(buf, conf)
}

func MergeVarFromFiles (globalconffile string, conffile string, conf interface{}) (error) {

    err := GetVarFromYamlFile(globalconffile, conf)
    if err != nil {
        return err
    }

    return GetVarFromYamlFile(conffile, conf)
}

func MapStringToMapList(mapString map[string]string) (map[string][]string) {

    mapList := make(map[string][]string)

    for key, val := range mapString{

        // remove [ on begin and ] on end of string
        val = val[1:len(val)-1]

        // split from string to list
        list := strings.Split(val, " ")

        // append items to mapList
        for _, item := range list {
            mapList[key] = append(mapList[key], fmt.Sprintf("%v", item))
        }

    }
    return mapList
}
