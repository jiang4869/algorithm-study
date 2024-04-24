#include <bits/stdc++.h>
using namespace std;
int main() {
    ios::sync_with_stdio(false);
    int n, m;
    cin >> n >> m;
    unordered_map<int, unordered_set<int>> mmid;
    for (int i = 1; i <= n; i++) {
        int k, num;
        cin >> k;
        for (int j = 1; j <= k; j++) {
            cin >> num;
            mmid[num].insert(i);
        }
    }
    int q;
    cin >> q;
    int a, b;
    for (int i = 0; i < q; i++) {
        cin >> a >> b;
        auto s1 = mmid[a];
        auto s2 = mmid[b];
        int count = 0;
        for (auto it = s1.begin(); it != s1.end(); it++) {
            if (s2.find(*it) != s2.end()) {
                count++;
            }
        }
        cout << count << endl;
    }

    return 0;
}