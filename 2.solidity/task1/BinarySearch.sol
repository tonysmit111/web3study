// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

// 二分查找 (Binary Search)
// 题目描述：在一个有序数组中查找目标值
contract BinarySearch {
    function binarySearch(int[] memory sortedArr, int target) public pure returns(int index) {
        uint start = 0;
        uint end = sortedArr.length - 1;
        while (start <= end) {
            uint mid = (start+end)/2;
            if (target == sortedArr[mid]) {
                return int(mid);
            } else if (target > sortedArr[mid]) {
                start = mid + 1;
            } else {
                end = mid - 1;
            }
        }
        return -1;
    }
}