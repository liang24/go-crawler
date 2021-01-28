package parser

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"

	"github.com/liang24/go-crawler/engine"
	"github.com/liang24/go-crawler/model"
)

var (
	desRe    = regexp.MustCompile(`<div class="des f-cl"[^>]*>\S+ \| (?P<Age>\d+)岁 \| (?P<Education>\S+) \| (?P<Marriage>\S+) \| (?P<Height>\d+)cm \| (?P<Income>[^<]+)</div>`)
	nameRe   = regexp.MustCompile(`<span class="nickName"[^>]*>([^<]+)</span>`)
	carRe    = regexp.MustCompile(`<div class="m-btn pink"[^>]*>(\S+车)</div>`)
	houseRe  = regexp.MustCompile(`<div class="m-btn pink"[^>]*>(\S+房)</div>`)
	hokouRe  = regexp.MustCompile(`<div class="m-btn pink"[^>]*>籍贯:([^<]+)</div>`)
	xinzouRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([^<]+座)\(\d+\.\d+-\d+\.\d+\)</div>`)
	weightRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>(\d+)kg</div>`)
	idRe     = regexp.MustCompile(`https://album.zhenai.com/u/(\d+)`)
)

func ParseProfile(contents []byte, gender string, url string) engine.ParseResult {
	var match [][]byte
	profile := model.Profile{}

	profile.Gender = gender

	// Name（昵称）
	profile.Name = extractString(contents, nameRe)

	// Age、Education、Marriage、Height、Income
	match = desRe.FindSubmatch(contents)
	groupNames := desRe.SubexpNames()
	if match != nil {
		v := reflect.ValueOf(&profile).Elem()
		for i, name := range groupNames {
			if i != 0 && name != "" { // 第一个分组为空（也就是整个匹配）
				if _, ok := reflect.TypeOf(profile).FieldByName(name); !ok {
					fmt.Printf("Failed to get '%s' field.\n", name)
				} else {
					setFieldValue(v, name, string(match[i]))
				}
			}
		}
	}

	// Car
	profile.Car = extractString(contents, carRe)

	// House
	profile.House = extractString(contents, houseRe)

	// Hokou
	profile.Hokou = extractString(contents, hokouRe)

	// Xinzou
	profile.Xinzuo = extractString(contents, xinzouRe)

	// Weight
	if weight, err := strconv.Atoi(extractString(contents, weightRe)); err != nil {
		profile.Weight = weight
	}

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}

	return result
}

func setFieldValue(v reflect.Value, propertyName string, value string) {
	if len(value) == 0 {
		return
	}
	pv := v.FieldByName(propertyName)
	switch pv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if i, err := strconv.Atoi(value); err == nil {
			pv.SetInt(int64(i))
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		if i, err := strconv.ParseUint(value, 10, 64); err == nil {
			pv.SetUint(uint64(i))
		}
	case reflect.Bool:
		if i, err := strconv.ParseBool(value); err == nil {
			pv.SetBool(i)
		}
	case reflect.String:
		pv.SetString(value)
	}
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
