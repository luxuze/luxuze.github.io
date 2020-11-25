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

    return sorted(disk_list, key=lambda x:check_disk(x))

if __name__ == "__main__":
    print(removeLess("aabcc"))
    print(disk())

