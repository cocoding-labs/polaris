// SPDX-License-Identifier: MIT

pragma solidity ^0.8.4;

interface ICompress {

    function inflate(bytes[] memory data) view external returns (bytes[] memory);

    function deflate(bytes[] memory data) view external returns (bytes[] memory);

}
