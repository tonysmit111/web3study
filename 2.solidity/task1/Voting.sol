// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

// 创建一个名为Voting的合约，包含以下功能：

// 一个mapping来存储候选人的得票数
// 一个vote函数，允许用户投票给某个候选人
// 一个getVotes函数，返回某个候选人的得票数
// 一个resetVotes函数，重置所有候选人的得票数
contract Voting {
    mapping (string name=> uint) public votes;
    string[] public candidates;

    function vote(string calldata name) public {
        if (votes[name] == 0 ){
            candidates.push(name);
        }
        votes[name] += 1;
    }

    function getVotes(string calldata name) public view returns (uint) {
        return votes[name];
    }

    function resetVotes() public {
        for (uint i = 0;i < candidates.length; i++) {
            votes[candidates[i]] = 0;
        }
        delete candidates;
    }
}