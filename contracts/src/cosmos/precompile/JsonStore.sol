// SPDX-License-Identifier: MIT

pragma solidity ^0.8.4;

interface IJsonStore {

  /////////////////////////////////////// READ METHODS (all) //////////////////////////////////////////

  function get(address key, uint256 slot) view external returns(string memory);

  function get(address key, uint256 slot, string memory path) view external returns(string memory);

  function dataURI(address key, uint256 slot) view external returns(string memory);

  ////////////////////////////////////// WRITE METHODS (only sender) //////////////////////////////////////////

  function set(uint256 slot, string memory jsonBlob) external returns(bool);

  function set(uint256 slot, string memory path, string memory value) external returns(bool);

  function set(uint256 slot, string[] memory path, string[] memory value) external returns(bool);

  function subReplace(uint256 slot, string memory searchPath, string memory replacePath, string memory value) external returns(bool);

  function subReplace(uint256 slot, string memory searchPath, string[] memory replacePaths, string[] memory values) external returns(bool);

  function remove(uint256 slot, string memory path) external returns(bool);

  /////////////////////////////////////// UTIL METHODS //////////////////////////////////////////

  function get(string memory jsonBlob, string memory path) view external returns(string memory);

  function dataURI(string memory jsonBlob) view external returns(string memory);

  function set(string memory jsonBlob, string memory path, string memory value) view external returns(string memory);

  function set(string memory jsonBlob, string[] memory path, string[] memory value) view external returns(string memory);

  function subReplace(string memory jsonBlob, string memory searchPath, string memory replacePath, string memory value) view external returns(string memory);

  function subReplace(string memory jsonBlob, string memory searchPath, string[] memory replacePaths, string[] memory values) view external returns(string memory);

  function remove(string memory jsonBlob, string memory path) view external returns(string memory);

}
