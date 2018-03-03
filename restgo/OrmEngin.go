package restgo

import (
	"github.com/go-xorm/xorm"
)
var mapxormengin map[string]*xorm.Engine =make( map[string]*xorm.Engine)

func SetEngin(key string,e *xorm.Engine){
	mapxormengin[key] = e
}
func OrmEngin(keys ...string)(e *xorm.Engine){
	if len(keys)==0{
		return mapxormengin["default"]
	}else{
		return mapxormengin[keys[0]]
	}

}