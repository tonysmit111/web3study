// SPDX-License-Identifier: MIT
pragma solidity ^0.8;
// 合约包含以下标准 ERC20 功能：
// balanceOf：查询账户余额。
// transfer：转账。
// approve 和 transferFrom：授权和代扣转账。
// 使用 event 记录转账和授权操作。
// 提供 mint 函数，允许合约所有者增发代币。
// 提示：
// 使用 mapping 存储账户余额和授权信息。
// 使用 event 定义 Transfer 和 Approval 事件。
// 部署到sepolia 测试网，导入到自己的钱包

contract MyToken {
    address public owner;

    string private name;
    string private symbol;
    uint  private decimal = 2;


    mapping(address=>uint256) balances;
    mapping(address=>mapping(address=>uint256)) allowances;

    event Transfer(address indexed from, address indexed to, uint256 amount);
    event Approval(address indexed from, address indexed spender, uint256 amount);

    constructor(string memory _name, string memory _symbol) {
        owner = msg.sender;
        name = _name;
        symbol = _symbol;
    }

    // 余额查询
    function balanceOf(address addr) public view returns(uint256) {
        return balances[addr];
    }

    // 转账
    function transfer(address toAddr, uint256 amount) public {
        address fromAddr = msg.sender;
       _transfer(fromAddr, toAddr, amount);
    }

    function _transfer(address fromAddr, address toAddr, uint256 amount) private {
        bool appr = msg.sender == fromAddr ;
       if (!appr) {
            appr = allowances[fromAddr][msg.sender] > 0;
            require(allowances[fromAddr][msg.sender] >= amount, "no enough auth balance");
       }
       require(appr, "no auth");
       require(balances[fromAddr]>=amount, "no enough balance");
        balances[fromAddr] -= amount;
        balances[toAddr] += amount;
        emit Transfer(msg.sender, toAddr, amount);
    }

    // 授信
    function approve(address spender, uint256 amount) public {
        allowances[msg.sender][spender] += amount;
        emit Approval(msg.sender, spender, amount);
    }

    function allowance(address from, address spender) public view returns(uint256) {
        return allowances[from][spender];
    }

    // 授信转账
    function transferFrom(address fromAddr, address toAddr, uint256 amount) public {
       _transfer(fromAddr, toAddr, amount);
    }

    // 铸币
    function mint(uint256 amount) external {
        require(msg.sender == owner, "no mint auth");
        balances[msg.sender] += amount;
    }

}