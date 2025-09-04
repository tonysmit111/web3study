// SPDX-License-Identifier: MIT

pragma solidity ^0.8;

contract RomanToInt {
    mapping(string => uint256) public m;

    string[] public ks;

    constructor() {
        m["M"] = 1000;
        m["CM"] = 900;
        m["D"] = 500;
        m["CD"] = 400;
        m["C"] = 100;
        m["XC"] = 90;
        m["L"] = 50;
        m["XL"] = 40;
        m["X"] = 10;
        m["IX"] = 9;
        m["V"] = 5;
        m["IV"] = 4;
        m["I"] = 1;
        ks.push("CM");
        ks.push("CD");
        ks.push("XC");
        ks.push("XL");
        ks.push("IX");
        ks.push("IV");
        ks.push("M");
        ks.push("D");
        ks.push("C");
        ks.push("L");
        ks.push("X");
        ks.push("V");
        ks.push("I");
    }

    function romanToInt(string memory roman) public view returns (uint) {
        bytes memory sb = bytes(roman);
        uint num=0;
        while (sb.length > 0) {
            for (uint8 i = 0;i<ks.length; i++) {
                bytes memory flag = bytes(ks[i]);
                if (flag.length == 2 && sb.length >= 2 && flag[0]==sb[0] && flag[1] == sb[1] ) {
                    sb = trim(2, sb);
                    num += m[ks[i]];
                    break;
                } else if (flag.length == 1 && flag[0] == sb[0]) {
                    sb = trim(1, sb);
                    num += m[ks[i]];
                    break;
                }
            }
        }
        return num;
    }

    function trim(uint8 fromIndex, bytes memory arr) public pure returns(bytes memory) {
        bytes memory _arr = new bytes(arr.length - fromIndex);
        for (uint8 i =0 ; i < _arr.length ; i++ ) {
            _arr[i] = arr[fromIndex + i];
        }
        return _arr;
    }

    function toBytes(string memory s) public pure returns(bytes memory) {
        return bytes(s);
    }

    function toString(bytes memory b) public pure returns(string memory) {
        return string(b);
    }
}
