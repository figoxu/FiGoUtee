package Figo

import (
	"time"
	"net/http"
	"github.com/go-martini/martini"
	"github.com/quexer/utee"
	"strings"
	"strconv"
)

type ParamHelper struct {
	context map[string]string
	param   martini.Params
}

func (p *ParamHelper) Float64(name string) float64 {
	r, pure := wp_func_float64(p.param, name)
	p.context[name] = pure
	return r
}
func (p *ParamHelper) Bool(name string) bool {
	r, pure := wp_func_Bool(p.param, name)
	p.context[name] = pure
	return r
}
func (p *ParamHelper) Int(name string, defaultVs ...int) int {
	r, pure := wp_func_Int(p.param, name)
	p.context[name] = pure
	return r
}
func (p *ParamHelper) Int64(name string) int64 {
	r, pure := wp_func_Int64(p.param, name)
	p.context[name] = pure
	return r
}
func (p *ParamHelper) Uint64(name string) uint64 {
	r, pure := wp_func_Uint64(p.param, name)
	p.context[name] = pure
	return r
}
func (p *ParamHelper) Time(name, format string) time.Time {
	r, pure := wp_func_time(p.param, name, format)
	p.context[name] = pure
	return r
}
func (p *ParamHelper) TimeLoc(name, format string, loc *time.Location) time.Time {
	r, pure := wp_func_time_loc(p.param, name, format, loc)
	p.context[name] = pure
	return r
}
func (p *ParamHelper) String(name string) string {
	r, pure := wp_func_string(p.param, name)
	p.context[name] = pure
	return r
}
func (p *ParamHelper) IntArr(name, separate string) []int {
	r, pure := wp_func_IntArray(p.param, name, separate)
	p.context[name] = pure
	return r
}

func (p *ParamHelper) Params()map[string]string{
	return p.context
}

func Mid_helper_param(c martini.Context, param martini.Params, w http.ResponseWriter) {
	c.Map(&ParamHelper{
		param:   param,
		context: make(map[string]string),
	})
}


func wp_func_string(param martini.Params, name string) (result string, pure string) {
	pure = param[name]
	return strings.TrimSpace(pure), pure
}

func wp_func_float64(param martini.Params, name string) (result float64, pure string) {
	pure = param[name]
	v, err := strconv.ParseFloat(pure, 64)
	utee.Chk(err)
	return v, pure
}

func wp_func_Bool(param martini.Params, name string) (result bool, pure string) {
	pure = param[name]
	v, err := strconv.ParseBool(pure)
	utee.Chk(err)
	return v, pure
}

func wp_func_Int(param martini.Params, name string, defaultVs ...int) (result int, pure string) {
	pure = param[name]
	v, err := strconv.ParseInt(pure, 10, 32)
	if err != nil && len(defaultVs) > 0 {
		return defaultVs[0], ""
	}
	utee.Chk(err)
	return int(v), pure
}

func wp_func_Int64(param martini.Params, name string) (result int64, pure string) {
	pure = param[name]
	v, err := strconv.ParseInt(pure, 10, 64)
	utee.Chk(err)
	return v, pure
}

func wp_func_Uint64(param martini.Params, name string) (result uint64, pure string) {
	pure = param[name]
	v, err := TpUint64(pure)
	utee.Chk(err)
	return v, pure
}

func wp_func_time(param martini.Params, name, format string) (result time.Time, pure string) {
	pure = param[name]
	t, err := time.Parse(format, pure)
	utee.Chk(err)
	return t, pure
}

func wp_func_time_loc(param martini.Params, name, format string, loc *time.Location) (result time.Time, pure string) {
	pure = param[name]
	t, err := time.ParseInLocation(format, pure, loc)
	utee.Chk(err)
	return t, pure
}

func wp_func_IntArray(param martini.Params, name, separate string) (result []int, pure string) {
	pure = param[name]
	sv := strings.TrimSpace(pure)
	svs := strings.Split(sv, separate)
	ivs := make([]int, 0)
	for _, v := range svs {
		if v == "" {
			continue
		}
		if iv, err := strconv.ParseInt(v, 10, 32); err != nil {
			ivs = append(ivs, int(iv))
		}
	}
	return ivs, pure

}