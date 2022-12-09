/**
 * @Time    :2022/12/9 10:03
 * @Author  :Xiaoyu.Zhang
 */

package commons

import "strings"

// ArrayToStr
/**
 *  @Description:
 *  @param a
 *  @return str
 */
func ArrayToStr(a []string) (str string) {
	if len(a) == 0 {
		return
	}
	var build strings.Builder
	for i, s := range a {
		build.WriteString(s)
		if i != len(a)-1 {
			build.WriteString(",")
		}
	}
	return build.String()
}
