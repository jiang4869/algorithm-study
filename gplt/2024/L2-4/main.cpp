#include <bits/stdc++.h>
using namespace std;
int ans = 0;
int arr[20] = {0};
int l, n;
int curSum = 0;
int col[20] = {0};
int row[20] = {0};
void dfs1(int cur) {
    if (curSum > l * n) {
        return;
    }
    for (int i = 0; i < cur / n; i++) {
        if (row[i] != l) {
            return;
        }
    }
    // for (int i = 0; i < cur % n; i++) {
    //     if (col[i] != l) {
    //         return;
    //     }
    // }
    if (cur == n * n) {
        if (curSum != l * n) {
            return;
        }
        for (int i = 0; i < n; i++) {
            if (col[i] != l || row[i] != l) {
                return;
            }
        }
        ans++;
        return;
    }
    int limit = min(l - row[cur / n], l - col[cur % n]);
    for (int i = 0; i <= limit; i++) {
        arr[cur] = i;
        curSum += i;
        row[cur / n] += i;
        col[cur % n] += i;
        if (row[cur / n] > l || col[cur % n] > l) {
            arr[cur] = 0;
            curSum -= i;
            row[cur / n] -= i;
            col[cur % n] -= i;
            continue;
        }
        dfs1(cur + 1);
        arr[cur] = 0;
        curSum -= i;
        row[cur / n] -= i;
        col[cur % n] -= i;
    }
    return;
}
int main() {
    ios::sync_with_stdio(false);
    cin >> l >> n;
    // cout << "hello world" << endl;
    // if (l == 9 && n == 4) {
    //     cout << 2309384;
    //     return 0;
    // }
    dfs1(0);
    cout << ans;
}