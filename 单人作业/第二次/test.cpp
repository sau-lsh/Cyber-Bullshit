#include "pch.h"
#include "../HelloWorld/HelloWord.h"
#include <time.h>

//TEST(TestCaseName, TestName) {
//  EXPECT_EQ(1, 1);
//  EXPECT_TRUE(true);
//}
//
//TEST(MaxSubArrayTest1, SimpleTestCase) {
//    vector<int> arr = { 1, 2, 3, 4, 5 };
//    EXPECT_EQ(maxSubArray(arr), 15);
//}
//
//TEST(MaxSubArrayTest2, ZeroTestCase) {
//    vector<int> arr = { 0, 0, 0, 0, 0 };
//    EXPECT_EQ(maxSubArray(arr), 0);
//}

#define __TESTTIME__ 100
#define __TESTSIZE__ 1000
TEST(MaxSubArrayTest, TestCase) {
    srand(time(0));
    vector<int> arr;
    for (int t = 0; t < __TESTTIME__; t++) {
        int n = rand() % __TESTSIZE__ + 1;
        arr.resize(n);
        for (int i = 0; i < n; i++) {
            arr[i] = rand() % __TESTSIZE__ - 100;
        }
        random_shuffle(arr.begin(), arr.end());
        int res = 0;
        for (int i = 0; i < n; i++) {
            for (int j = i; j < n; j++) {
                int sum = 0;
                for (int k = i; k <= j; k++) {
                    sum += arr[k];
                }
                res = max(res, sum);
            }
        }
        EXPECT_EQ(maxSubArray(arr), res);
    }
}
