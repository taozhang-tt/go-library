package main

func main() {

}

/*
先统计出每个字符的数量
按数量将字符排序
判断最多的字符是否超过限制：设所有的字符数一共有 cnt, 如果 cnt 为偶数，则任意字符数不能大于 cnt/2，如果 cnt 为奇数，则任意字符数不能大于 cnt/2 + 1
先将字符数最多的字符安排好位置：尽可能的分散
*/
func reorganizeString(s string) string {
    n := len(s)
    if n <= 1 {
        return s
    }

    cnt := [26]int{}
    maxCnt := 0
    for _, ch := range s {
        ch -= 'a'
        cnt[ch]++
        if cnt[ch] > maxCnt {
            maxCnt = cnt[ch]
        }
    }
    if maxCnt > (n+1)/2 {
        return ""
    }

    ans := make([]byte, n)
    evenIdx, oddIdx, halfLen := 0, 1, n/2
    for i, c := range cnt[:] {
        ch := byte('a' + i)
        for c > 0 && c <= halfLen && oddIdx < n {
            ans[oddIdx] = ch
            c--
            oddIdx += 2
        }
        for c > 0 {
            ans[evenIdx] = ch
            c--
            evenIdx += 2
        }
    }
    return string(ans)
}