#include <iostream>
#include <vector>
#include <map>
#include <functional>
using namespace std;

const int inf = 1E9;
signed main() {

    freopen("Source.txt", "r", stdin);

    int n;
    cin >> n;

    vector<vector<int>> dis(n, vector<int>(n, inf)); // 二维数组表示[i, j]之间的距离
    map<string, int> mp; // 记录景点的（名称->编号）的映射
    vector<string> name(n); // 记录景点名称

    for (int i = 0; i < n; i++) {
        cin >> name[i];
        mp[name[i]] = i;
        name[i] = name[i];
        dis[i][i] = 0;
    }

    int m;
    cin >> m;
    for (int i = 0; i < m; i++) {
        string a, b;
        int w;
        cin >> a >> b >> w;
        int u = mp[a];
        int v = mp[b];

        dis[u][v] = min(dis[u][v], w);
        dis[u][v] = min(dis[u][v], w);
    }

    vector<vector<int>> Path(n, vector<int>(n, -1)); // Path[i, j]记录从 i 到 j 的最短路径上经过的第一个点
    for (int k = 0; k < n; k++) {
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) {
                if (dis[i][j] > dis[i][k] + dis[k][j]) {
                    dis[i][j] = dis[i][k] + dis[k][j];
                    Path[i][j] = k;
                }
            }
        }
    }

    int x; // 将要访问的景点数量
    vector<int> place(n); // 记录将要访问的景点
    vector<int> path; // 最优访问路径搜索辅助数组
    vector<int> route; // 最优访问路径
    vector<bool> vis(n); // 景点是否被搜索
    int ans = inf;
    function<void(int, int)> dfs = [&](int u, int res) {
        if (u == x) {
            if (res < ans) {
                ans = res;
                route = path;
            }
            return;
        }

        for (int i = 0; i < x; i++) {
            if (vis[place[i]]) {
                continue;
            }
            path[u] = place[i];
            vis[place[i]] = true;
            dfs(u + 1, res + dis[path[u - 1]][place[i]]);
            vis[place[i]] = false;
        }
    };

    while (cin >> x) {
        path.resize(x);
        fill(vis.begin(), vis.end(), false);

        for (int i = 0; i < x; i++) {
            string a;
            cin >> a;
            place[i] = mp[a];
        }

        for (int i = 0; i < x; i++) {
            path[0] = place[i];
            vis[place[i]] = true;
            dfs(1, 0);
            vis[place[i]] = false;
        }

        vector<int> finalRoute = { route[0] }; // 加上中间转移节点的最终访问路径

        for (int i = 1; i < route.size(); i++) {
            int u = route[i - 1], v = route[i];
            for (int j = Path[u][v]; j != v && j != -1; j = Path[j][v]) {
                finalRoute.push_back(j);
            }
            finalRoute.push_back(v);
        }

        for (auto p : finalRoute) {
            cout << name[p] << " ";
        } cout << "\n";

    }

    return 0;
}
