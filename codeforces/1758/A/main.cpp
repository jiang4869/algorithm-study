#include <bits/stdc++.h>
using namespace std;
int main() {
    ios::sync_with_stdio(false);
    int t;
    cin >> t;
    while (t--) {
        string str;
        cin >> str;
        int size = str.length();
        vector<char> v(2 * size);
        int i;
        for (i = 0; i < str.size(); i++) {
            v[i] = v[2 * size - i - 1] = str[i];
        }

        for (i = 0; i < 2 * size; i++) {
            cout << v[i];
        }
        cout << endl;
    }
    return 0;
}