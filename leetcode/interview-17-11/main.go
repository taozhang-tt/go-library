package main

import (
	"fmt"
)

/*
面试题 17.11. 单词距离
	https://leetcode-cn.com/problems/find-closest-lcci/
题目描述
	有个内含单词的超大文本文件，给定任意两个单词，找出在这个文件中这两个单词的最短距离(相隔单词数)。如果寻找过程在这个文件中会重复多次，而每次寻找的单词不同，你能对此优化吗?
示例：
	输入：words = ["I","am","a","student","from","a","university","in","a","city"], word1 = "a", word2 = "student"
	输出：1
	提示：

	words.length <= 100000
*/
func main() {
	words := []string{"vg","ti","te","yd","ja","vqx","a","rtg","bn","rdg","v","sxz","t","gn","bdc","bp","fk","nl","qr","qh","vo","d","ab","ui","gu","ft","kfr","a","jt","gyq","gmw","jqo","f","iy","pu","kky","cur","qug","z","gb","osn","khg","szd","zhb","g","n","wbv","p","f","h","hka","j","u","n","o","mod","hgc","bev","e","cx","jk","z","in","cua","on","nis","c","qpq","e","rs","ew","ms","d","nh","h","cvu","oup","m","a","yb","v","t","rc","s","hfa","ehn","w","dox","ptt","r","jzm","l","urm","qd","va","emi","sf","dc","h","wh","eh","z","c","mx","pm","bnn","fpk","sky","uyd","cwk","su","xqc","xv","w","ai","dd","zic","qg","vmc","yrd","os","el","ylz","jac","b","mpw","nso","nzx","d","c","nl","qo","vbe","p","lq","u","np","u","keo","gg","myh","emt","to","mfo","zlv","v","m","nm","i","lml","v","wly","rd","bu","ymn","iu","br","r","gwz","nou","qod","ei","gt","ws","j","wgo","e","sn","rem","jmy","hy","dm","qy","gmo","qz","xcm","iy","spz","yr","fc","cqo","pol","ib","u","ldd","fj","vj","v","hu","nr","e","hf","fwv","eu","dd","o","zft","j","uz","j","ce","hwn","dcs","z","xyn","o","b","kx","kg","pjv","ooh","k","r","fgj","wc","v","tjs","cq","aia","z","exs","be","py","zif","b","oy","o","aox","yf","i","tzn","ny","xst","eo","soq","waz","cr","bf","l","yf","bj","p","eez","io","er","xe","w","iqh","as","wmx","pk","ge","iet","xel","acc","yc","ek","bb","qip","igg","ejf","far","wkp","fzl","hhm","r","czh","vpw","z","jei","tmi","m","zdw","obg","cq","ot","h","qj","m","nu","cbu","ffv","u","fc","vb","n","vr","myt","so","g","pkl","gya","f","ahb","sn","ym","gk","i","nfn","e","efr","yu","xr","k","h","esf","uay","kws","czh","fwr","no","trb","wo","ukc","p","dn","me","up","jde","fko","y","ghw","v","g","lnq","ztl","ppr","s","o","wsj","z","rxc","o","b","gmv","kk","nak","i","gy","qv","yl","qn","lhs","r","n","ymy","d","ko","y","nvm","sdu","dru","k","uk","yu","yg","dds","hoh","v","iv","tu","fe","k","d","ggz","d","yw","al","vnu","b","u","f","up","krv","em","slq","cee","sbt","vl","t","tep","fz","db","q","z","m","x","l","ho","bcm","fcy","x","l","ps","e","yld","rou","q","d","ldo","rjn","nb","a","zr","kle","nyz","xu","h","q","fxe","aam","fho","kj","sw","i","jwa","joo","y","zrg","vdy","oc","o","cj","re","qw","wkk","s","of","nnj","or","p","ih","ps","o","u","xsg","hi","w","bof","mvn","lgy","p","m","muq","avm","w","f","vl","ntu","b","b","hbz","iv","qk","ttl","g","v","je","nej","e","s","hnu","pn","ygy","vos","ui","mvj","ae","sz","onr","mo","ar","ffs","oa","v","qf","zi","ou","utq","fx","yb","u","t","jy","oh","jyv","tju","w","vl","qft","s","hlb","hyb","su","ol","hzr","xd","j","kt","x","kzi","d","epx","gf","lxy","pmy","d","not","sbc","h","ial","rz","ypf","yd","p","hp","gck","d","r","hlb","r","zl","jbr","uaw","wz","pjc","v","bs","gh","rqv","msp","f","oaf","zp","mrh","az","br","ro","oj","zfa","ejr","gu","z","po","bc","p","we","yzw","qrg","j","wwv","bh","evz","m","h","m","mv","m","c","qz","b","vnx","hju","r","sdu","qum","m","eq","doo","nbw","lvm","o","kfn","t","z","fpl","yqv","nf","h","gr","i","udm","j","v","sz","w","eef","vy","gow","ro","zmx","ts","dub","g","dvy","ye","gmr","b","b","y","y","hxn","dr","a","pmy","ick","fb","ny","hzb","x","nm","sqr","n","ial","xo","aw","egi","mrm","uh","i","unn","j","kg","tk","lar","o","iy","yh","ag","ss","jx","ory","p","k","o","c","vng","d","wa","ff","fks","v","cz","c","u","rxf","ccv","yx","n","rt","pzn","pih","b","goy","znl","iw","fym","sy","hq","dc","a","bbp","ql","oyx","afi","gry","of","gfg","lys","lz","ky","ssv","k","y","dj","woq","pw","wbz","e","x","o","b","ptl","n","lte","gv","cg","sx","v","q","qm","xs","dip","r","pyr","j","i","w","qot","by","sla","fum","e","jx","s","ws","ntq","xb","sq","t","ky","zzd","ykj","no","sy","zlc","v","cj","f","mzm","n","sh","ao","t","ky","rrh","s","wb","bsx","kko","jxv","j","iwd","vzo","kh","q","ulq","tba","er","oic","ca","qee","e","phs","ew","aj","hm","a","hm","o","nn","w","g","oi","di","urj","nn","dp","di","tu","lll","u","xya","o","e","rz","y","e","ybn","nkb","o","aj","dim","my","xdf","bcb","rf","rbr","b","ip","e","lw","oez","k","zx","hvj","tmh","cb","ghy","qju","hy","d","d","j","qu","zoq","r","z","l","bmz","tb","z","hsw","ctu","yd","wx","pka","b","xrl","ob","gp","e","r","vm","gt","qxo","c","fur","djn","vd","p","zq","xm","a","g","t","c","kh","uh","nsq","vjc","mys","v","q","o","vc","wee","rms","ah","se","fg","uo","fl","n","io","z","ww","e","m","o","kif","yqu","r","fqy","e","yuc","pvf","jni","ss","fi","wn","ct","lrd","dyt","j","kxs","oon","mw","mdk","ejj","m","uv","gid","ldu","ah","d","cxh","bb","lr","c","an","ey","g","sqs","z","zx","x","ejy","qi","c","q","p","nm","fn","v","z","rfz","sus","vh","nxy","oa","xdg","vc","l","fvb","i","z","igu","ay","t","xvp","n","xrl","v","e","cq","qv","or","ff","apj","za","hlx","p","y","tz","hng","d","j","ykt","ta","lp","lc","fhs","xp","gfl","tzi","v","x","tgo","jv","xoj","gux","eyx","adg","gjc","p","gw","hnl","p","jt","mlj","p","nv","trl","tp","dap","mc","lx","g","l","kjg","cax","jct","q","t","d","ww","vwd","zcl","yhr","fft","qq","y","wwz","sek","a","mc","vkq","m","yn","ope","kna","y","ve","w","z","mwm","a","i","d","ta","s","fs","pix","n","lio","xka","wbn","kag","v"}
	word1 := "o"
	word2 := "fk"
	fmt.Println(findClosest3(words, word1, word2))
}

/*
暴力法(通过)
	先统计每个单词出现的位置，然后就是求两个数组中元素的最小差值
*/
func findClosest(words []string, word1 string, word2 string) int {
	hash, ans := make(map[string][]int, 0), len(words)
	for index, word := range words {
		hash[word] = append(hash[word], index)
	}
	for i := 0; i < len(hash[word1]); i++ {
		for j := 0; j < len(hash[word2]); j++ {
			dis := hash[word1][i] - hash[word2][j]
			if dis < 0 {
				dis = hash[word2][j] - hash[word1][i]
			}
			if dis < ans {
				ans = dis
			}
		}
	}
	return ans
}

/*
暴力法、优化了最短距离的步骤
*/
func findClosest2(words []string, word1 string, word2 string) int {
	hash := make(map[string][]int, 0)
	for index, word := range words {
		hash[word] = append(hash[word], index)
	}
	arr1, arr2 := hash[word1], hash[word2] //两个数组都是单调递增
	i, j, ans := 0, 0, abs(arr1[0]-arr2[0])
	for i < len(arr1) && j < len(arr2) {
		dis := abs(arr1[i] - arr2[j])
		if dis < ans {
			ans = dis
		}
		if arr1[i] > arr2[j] {
			j++
		} else {
			i++
		}
	}
	return ans
}

func abs(num int) int {
	if num > 0 {
		return num
	}
	return -num
}

/*
双指针
	初始化两个指针 l,r 指向 words 的第0个元素，初始化一个记录返回结果的 ans，其最大值不会超过 words 长度
	r 指针向右扫描
		r 扫到一个目标(r==word1 || r==word2)
			回头判断 words[l]、words[r]的关系
			如果 l 不是目标单词，那么将 l 指针指向 r 的位置， r 继续向右扫描
			如果 l 也是目标单词，判断 words[l] 是否 等于 words[r]
				words[l] == words[r] ? 将 l 挪动到 r
				words[l] != words[r] ？计算 r-l，并判断是否小于 ans
*/
func findClosest3(words []string, word1 string, word2 string) int {
	l, r, ans := 0, 0, len(words)
	for r < len(words) {
		if words[r] == word1 || words[r] == word2 {
			if (words[l] == word1 || words[l] == word2) && words[l] != words[r] && r-l < ans {
				ans = r - l
			}
			l = r
		}
		r++
	}
	return ans
}
