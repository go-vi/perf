package slice

import (
	"encoding/json"
)

func Find(slice []byte, key interface{}) (flag bool, err error) {
	var listHolder []interface{}
	var keyHolder interface{}

	err = json.Unmarshal(slice, &listHolder)

	if err != nil {
		return false, err
	}

	b, err := json.Marshal(key)
	if err != nil {
		return false, err
	}

	err = json.Unmarshal(b, &keyHolder)
	if err != nil {
		return false, err
	}

	for _, v := range listHolder {
		if v == keyHolder {
			return true, nil
		}
	}
	return false, nil
}
