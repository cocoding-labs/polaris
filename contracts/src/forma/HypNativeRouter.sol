// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.4;

import {TokenRouter} from "@hyperlane-xyz/core/contracts/token/libs/TokenRouter.sol";
import {INativeMinter} from "./precompile/INativeMinter.sol";

contract HypNativeRouter is TokenRouter {

    INativeMinter public immutable nativeMinter;

    constructor(address precompile, address _mailbox) TokenRouter(_mailbox) {
        nativeMinter = INativeMinter(precompile);
    }

    function transferRemote(
        uint32 _destination,
        bytes32 _recipient,
        uint256 _amount
    ) public payable virtual override returns (bytes32 messageId) {
        require(msg.value >= _amount, "Native: amount exceeds msg.value");
        uint256 gasPayment = msg.value - _amount;
        return _transferRemote(_destination, _recipient, _amount, gasPayment);
    }

    function balanceOf(
        address _account
    ) external view override returns (uint256) {
        return _account.balance;
    }

    function _transferFromSender(
        uint256 _amount
    ) internal override returns (bytes memory) {
        nativeMinter.burn(address(this), _amount);
        return bytes(""); // no metadata
    }

    function _transferTo(
        address _recipient,
        uint256 _amount,
        bytes calldata // no metadata
    ) internal virtual override {
        nativeMinter.mint(_recipient, _amount);
    }

}