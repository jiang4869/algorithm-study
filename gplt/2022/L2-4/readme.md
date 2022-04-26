#题解
根据题意，我们可以通过每个人的关系构造出一个图，途中的每个边都是有向的。

一个人对其他人的距离相当于图中某个点到其他点的最短距离。

可以通过Dijkstra算法求出每个人对其他人的距离。即求出距离感`D`（Dij表示一个人 i 在一个异性 j 眼中的距离感）。

求出D后根据题意就能找出大众情人了。

不过用go实现的时候，实在是跑不过去，最后一个点一直TLE，这里给出C++的写法

```cpp
#include<bits/stdc++.h>
using namespace std;
const int maxn = 550;
struct node {
	int to, dis;
	node(int a, int b) :to(a), dis(b) {}
};
vector<node> G[maxn];
int D[maxn][maxn];
const int INF = 9e8;
typedef pair<int, int> pii;
int n;
void dijkstra(int s) {
	vector<int> dis(n + 10, INF);
	vector<bool>sign(n + 10, false);
	priority_queue<pii, vector<pii>, greater<pii>> que;
	que.push(pii(0, s));
	dis[s] = 0;
	while (que.size()) {
		auto tmp = que.top();
		que.pop();
		if (sign[tmp.second])continue;
		sign[tmp.second] = true;
		for (auto& val : G[tmp.second]) {
			if (dis[val.to] > dis[tmp.second] + val.dis) {
				dis[val.to] = dis[tmp.second] + val.dis;
				que.push(pii(dis[val.to], val.to));
			}
		}
	}
	for (int i = 1; i <= n; i++) {
		D[i][s] = dis[i];
	}
}

int main() {
	ios::sync_with_stdio(false);
	cin >> n;
	vector<char> sex(n + 10);
	for (int i = 1; i <= n; i++) {
		char ch;
		int k;
		cin >> ch >> k;
		sex[i] = ch;
		for (int j = 0; j < k; j++) {
			int dis, num;
			cin >> num >> ch >> dis;
			G[i].push_back(node(num, dis));
		}
	}
	for (int i = 1; i <= n; i++) {
		dijkstra(i);
	}
	vector<int> maxs(n + 10);
	for (int i = 1; i <= n; i++) {
		for (int j = 1; j <= n; j++) {
			if (sex[i] != sex[j]) {
				maxs[i] = max(maxs[i], D[i][j]);
			}
		}
	}
	int max1, max2;
	max1 = max2 = INT_MAX;
	for (int i = 1; i <= n; i++) {
		if (sex[i] == 'F') {
			max1 = min(max1, maxs[i]);
		}
		else {
			max2 = min(max2, maxs[i]);
		}
	}
	bool flag = false;
	for (int i = 1; i <= n; i++) {
		if (sex[i] == 'F' && maxs[i]==max1) {
			if (flag) cout << " ";
			cout << i;
			flag = true;
		}
	}
	cout << endl;
	flag = false;
	for (int i = 1; i <= n; i++) {
		if (sex[i] == 'M' && maxs[i] == max2) {
			if (flag) cout << " ";
			cout << i;
			flag = true;
		}
	}
	cout << endl;
	return 0;
}
```
