package utils

import "strconv"

type TreeMenuStruct struct {
	Menu_icon      string
	Menu_id        string
	Menu_name      string
	Menu_route     string
	Menu_up_id     string
	Menu_Dept      string
	Menu_leaf_flag string
	Menu_img       string
	Menu_color     string
}

func GetJSONMenuTree(node []TreeMenuStruct, id string, d int, result *[]TreeMenuStruct) {

	for _, val := range node {
		if val.Menu_up_id == id {

			var oneline TreeMenuStruct
			oneline.Menu_icon = val.Menu_icon
			oneline.Menu_id = val.Menu_id
			oneline.Menu_name = val.Menu_name
			oneline.Menu_up_id = val.Menu_up_id
			oneline.Menu_route = val.Menu_route
			oneline.Menu_leaf_flag = val.Menu_leaf_flag
			oneline.Menu_Dept = strconv.Itoa(d)
			oneline.Menu_img = val.Menu_img
			oneline.Menu_color = val.Menu_color
			*result = append(*result, oneline)
			GetJSONMenuTree(node, val.Menu_id, d+1, result)
		}
	}
}
