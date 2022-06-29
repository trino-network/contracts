/*
1. 使用USDT按设定比率兑换TNO
2. 使用ether兑换TNO，价格按uniswap来
usdt: 0xdAC17F958D2ee523a2206206994597C13D831ec7

 */


// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";

interface ITrinoToken is IERC20 {
    function mint(address to, uint256 amount) external;
}

interface AggregatorV3Interface {
    function latestRoundData()
        external
        view
        returns (
            uint80 roundId,
            int answer,
            uint startedAt,
            uint updatedAt,
            uint80 answeredInRound
        );
}


contract TrinoSwap is Initializable , PausableUpgradeable, AccessControlUpgradeable, ReentrancyGuardUpgradeable {
    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");

    ITrinoToken token;
    IERC20 usdt;
    AggregatorV3Interface internal priceFeed;

    enum SwapMode {
        Mint,
        Tranfer
    }

    SwapMode public mode;
    uint256 public rate;

    address payable public beneficiary;
    
    event SwapWithUSDT(uint256 amount, uint256 spend, address recipient);
    event SwapWithETH(uint256 amount, uint256 spend, address recipient);
 
    // oracle: 0x5f4eC3Df9cbd43714FE2740f5E3616155c5b8419
    function initialize(address _token, address _usdt, address _oracle, SwapMode _mode,  uint256 _rate) public initializer {
        require(_token != address(0), "zero address");
        require(_rate != 0, "zero");
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(PAUSER_ROLE, msg.sender);
        token = ITrinoToken(_token);
        mode = _mode;
        usdt = IERC20(_usdt);
        rate = _rate;
        beneficiary = payable(msg.sender);
        priceFeed = AggregatorV3Interface(_oracle);
    }

    function setMode(SwapMode _mode) public onlyRole(DEFAULT_ADMIN_ROLE) {
        mode = _mode;
    }

    function setRate(uint256 _rate) public onlyRole(DEFAULT_ADMIN_ROLE) {
        require(_rate != 0, "zero");
        rate = _rate;
    }

    function setBeneficiary(address payable _beneficiary) public onlyRole(DEFAULT_ADMIN_ROLE) {
        require(_beneficiary != address(0), "zero address");
        beneficiary = _beneficiary;
    }

    function swapWithUSDT(uint256 amount, address recipient) public  whenNotPaused {
        require(amount > 0, "zero amount");

        uint256 spend = amount  / rate / 1e12;
        usdt.transferFrom(msg.sender, beneficiary,  spend);

        _send(amount, recipient);

        emit SwapWithUSDT(amount, spend, recipient);
    }

    function swapWithETH(uint256 amount, address recipient) public payable whenNotPaused nonReentrant {
        // 判断eth是否足够
        uint256 spend = amount / rate / uint256(getLatestPrice() ) * 1e8;
 
        require(spend <= msg.value, "insufficient eth");
        uint256 delta = msg.value - spend;

        // 转代币
        _send(amount, recipient);
        // 转出花费的
        _sendValue(beneficiary, spend);
        // 返还多余的eth
        if (delta != 0) {
            _sendValue(payable(msg.sender), delta);
        }

        emit SwapWithETH(amount, spend, recipient);
    }
  
    function _send(uint256 amount, address recipient) internal {
        if (recipient == address(0)) {
            recipient = msg.sender;
        }

        if (mode == SwapMode.Mint) {
            token.mint(recipient, amount);
        } else {
            token.transfer(recipient, amount);
        }
    }

    function getLatestPrice() public view returns (int) {
        (
            ,
            int price,
            ,
            ,
        ) = priceFeed.latestRoundData();
        // for ETH / USD price is scaled up by 10 ** 8
        return price;
    }

       /*
     * @description: 暂停
     * @param {public} onlyRole
     * @return {*}
     */
    function pause() public onlyRole(PAUSER_ROLE) {
        _pause();
    }

    /*
     * @description: 取消暂停
     * @param {public} onlyRole
     * @return {*}
     */
    function unpause() public onlyRole(PAUSER_ROLE) {
        _unpause();
    }

    function _sendValue(address payable recipient, uint256 amount) internal {
        require(address(this).balance >= amount, "Address: insufficient balance");

        (bool success, ) = recipient.call{value: amount}("");
        require(success, "Address: unable to send value, recipient may have reverted");
    }
}