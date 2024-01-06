import functools

MOD = 10 ** 9 + 7


@functools.lru_cache(None)
def f(s: str) -> int:
    @functools.lru_cache(None)
    def f(i: int, pre: int, isLimit: bool, isNum: bool) -> int:
        if i == len(s):
            return int(isNum)
        res = 0
        if not isNum:
            res += f(i + 1, pre, False, False)

        up = int(s[i]) if isLimit else 9
        for d in range(1 - int(isNum), up + 1):
            if pre == -1 or abs(d - pre) == 1:
                res = (res + f(i + 1, d, isLimit and up == d, True)) % MOD
        return res % MOD

    return f(0, -1, True, False)


class Solution:
    def countSteppingNumbers(self, low: str, high: str) -> int:
        return (f(high) - f(str(int(low) - 1))) % MOD


if __name__ == '__main__':
    s = Solution()
    print(s.countSteppingNumbers("1", "11"))
    print(s.countSteppingNumbers("90", "101"))
