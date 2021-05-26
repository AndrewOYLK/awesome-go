package patchyaml

import "reflect"

func PatchData(source map[interface{}]interface{}, dst interface{}) map[interface{}]interface{} {
	tdst := dst.(map[interface{}]interface{})
	if dst != nil {
		for k, v := range source {
			if _, ok := tdst[k]; !ok {
				tdst[k] = v
				continue
			}

			// <TODO> for the type of slice and more
			switch reflect.TypeOf(v).Kind() {
			case reflect.Map:
				tdst[k] = PatchData(v.(map[interface{}]interface{}), tdst[k])
			default:
				tdst[k] = v
			}
		}
	}
	return tdst
}
