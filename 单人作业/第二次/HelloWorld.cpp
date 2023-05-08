// maxSubarraySum.cpp : 此文件包含 "main" 函数。程序执行将在此处开始并结束。
//

#include <iostream>
#include <vector>
using namespace std;

int maxSubArray(vector<int>& a) {
    int n = a.size();
    vector<int> f(n, a[0]);
    int ans = a[0];
    for (int i = 1; i < n; i++) {
        f[i] = a[i] + f[i - 1] > a[i] ? a[i] + f[i - 1] : a[i];
        ans = ans > f[i] ? ans : f[i];
    }
    return ans;
}


int main() {

    vector<int> arr = { 1, 2, 3, 4, 5 };
    cout << maxSubArray(arr) << "\n";

    return 0;
}