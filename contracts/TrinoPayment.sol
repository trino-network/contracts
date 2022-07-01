// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";


contract TrinoPayment is Initializable , PausableUpgradeable, AccessControlUpgradeable {
    
    IERC20 token;
    mapping(string => uint256) public orders;
    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");

    event Pay(string invoice_id, address recipient, uint256 amount);
    event Withdraw(address recipient, uint256 amount);

    function initialize(address _token) public initializer {
        token = IERC20(_token);
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(PAUSER_ROLE, msg.sender);
    }


    function pay(string calldata invoice_id, address recipient, uint256 amount) public whenNotPaused {
        require(recipient != address(0), "recipient can't be zero address");
        require(token.transferFrom(msg.sender, recipient, amount), "transfer failed");
        orders[invoice_id] = amount;
        emit Pay(invoice_id, recipient, amount);
    }

    function withdraw(address recipient) public onlyRole(DEFAULT_ADMIN_ROLE)  {
        uint256 amount = token.balanceOf(address(this));
        token.transfer(recipient, amount);
        emit Withdraw(recipient, amount);
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
}