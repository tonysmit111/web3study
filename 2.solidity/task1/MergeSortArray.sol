// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

// 合并两个有序数组 (Merge Sorted Array)
// 题目描述：将两个有序数组合并为一个有序数组。
contract MergeSortArray {
    function merge(int256[] memory arr1, int256[] memory arr2) public pure returns (int256[] memory) {
        uint256 size = arr1.length + arr2.length;
        int256[] memory rarr = new int256[](size);

        uint256 i = 0;
        uint256 j = 0;
        uint256 k = 0;

        // 合并两个数组，直到其中一个耗尽
        while (i < arr1.length && j < arr2.length) {
            if (arr1[i] < arr2[j]) {
                rarr[k] = arr1[i];
                i++;
            } else {
                rarr[k] = arr2[j];
                j++;
            }
            k++;
        }

        // 复制arr1的剩余元素（如果有）
        while (i < arr1.length) {
            rarr[k] = arr1[i];
            i++;
            k++;
        }

        // 复制arr2的剩余元素（如果有）
        while (j < arr2.length) {
            rarr[k] = arr2[j];
            j++;
            k++;
        }

        return rarr;
    }

    function merge2(int256[] memory arr1, int256[] memory arr2) public pure returns (int256[] memory) {
        uint256 size = arr1.length + arr2.length;
        int256[] memory rarr = new int256[](size);

        uint256 i = 0;
        uint256 j = 0;

        for (uint256 n = 0; n < rarr.length; n++) {
            if (i>=arr1.length && j < arr2.length) {
                rarr[n]=arr2[j];
                j++;
                continue;
            }
            if (j>=arr2.length && i < arr1.length) {
                rarr[n]=arr1[i];
                i++;
                continue;
            }
            if (arr1[i] < arr2[j]) {
                rarr[n] = arr1[i];
                i++;
            } else {
                rarr[n] = arr2[j];
                j++;
            }
        }
        return rarr;
    }
}
