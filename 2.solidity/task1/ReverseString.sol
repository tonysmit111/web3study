// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

// 反转字符串 (Reverse String)

// 题目描述：反转一个字符串。输入 "abcde"，输出 "edcba"

contract ReverseString {

    function reverse(string calldata str) public pure returns(string memory) {
        bytes memory bstr = bytes(str);
        uint len = bstr.length;
        for (uint i = 0; i < len / 2; i++) {
            (bstr[i], bstr[len - 1 - i]) = (bstr[len - 1 - i], bstr[i]);
        }
        return string(bstr);
    }
}

        // Student memory stu = Student({name:"tom", age:18, courses:cs });