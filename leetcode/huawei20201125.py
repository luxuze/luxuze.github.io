def removeLess(s):
    mp = {}
    arr = []

    for i in s:
        if not i in mp:
            mp[i] = 0
        mp[i]+=1

    little = min(mp.values())

    for k,v in mp.items():
        if v == little:
            arr.append(k)
    for a in arr:
        s = s.replace(a, "")

    return s if s else "empty"

def disk():
    def check_disk(storage_str):
        tmp = ""
        val = 0
        for s in storage_str:
            if s == "T":
                val += int(tmp)*1024*1024
                tmp = ""
                continue
            if s == "G":
                val += int(tmp)*1024
                tmp = ""
                continue
            if s == "M":
                val += int(tmp)
                tmp = ""
                continue
            tmp+=s
        return val

    disk_list = []

    num = input("请输入个数:")
    for _ in range(int(num)):
        disk_list.append(input("请输入磁盘大小:"))

    for d in sorted(disk_list, key=lambda x:check_disk(x)):
        print(d)

def solution():
    def intervalIntersection(A, B):
        if A[0] > B[0]:
            A, B = B, A
        if A[1] < B[0]:
            return None
        return [B[0], min(A[1], B[1])]

    def merge(intervals):
        if not intervals:
            return []
        stack=[]
        n=len(intervals)
        intervals.sort()
        for i,interval in enumerate(intervals):
            left,right=interval
            if stack and stack[-1][1]>=left:
                stack[-1][1]=max(stack[-1][1],right)
            else:
                stack.append(interval)
        return stack


    l = [[0,3], [1,3], [3,5], [3,6]] #这里是输入的值
    ans = []

    for i in range(len(l)):
        for j in range(len(l)):
            if j <= i:
                continue
            ans.append(intervalIntersection(l[i], l[j]))
    print(merge(ans))

if __name__ == "__main__":
    # print(removeLess("aabcc"))
    # print(disk())

    # print(
    #     intervalIntersection([1,2], [2,9])
    #     )
    solution()
