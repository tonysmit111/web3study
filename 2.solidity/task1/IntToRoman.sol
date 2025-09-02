// SPDX-License-Identifier: MIT
pragma solidity ^0.8;


// 实现整数转罗马数字
contract IntToRoman {
    mapping(uint16=>string) public m;
    uint16[] public ks;

    constructor() {
        m[1000] = "M";
        m[900] = "CM";
        m[500] = "D";
        m[400] = "CD";
        m[100] = "C";
        m[90] = "XC";
        m[50] = "L";
        m[40] = "XL";
        m[10] = "X";
        m[9] = "IX";
        m[5] = "V";
        m[4] = "IV";
        m[1] = "I";
        ks = [1000,900,500,400,100,90,50,40,10,9,5,4,1];
    }

    function intToRoman(uint16 num) public view returns (string memory) {
        string memory s;
        for (uint8 i = 0;i<ks.length; i++)  {
            while (num > ks[i] ) {
                num -= ks[i];
                string memory ss = m[ks[i]];
                s = strConcat(s, ss);
            }
        }
        return s;
    }

    function strConcat(string memory _a, string memory _b) public pure returns (string memory){
        bytes memory _ba = bytes(_a);
        bytes memory _bb = bytes(_b);
        string memory ret = new string(_ba.length + _bb.length);
        bytes memory bret = bytes(ret);
        uint k = 0;
        for (uint i = 0; i < _ba.length; i++) bret[k++] = _ba[i];
        for (uint i = 0; i < _bb.length; i++) bret[k++] = _bb[i];
        return string(ret);
    }

}