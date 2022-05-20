//SPDX-License-Identifier: UNLICENSED
pragma solidity >=0.8.1;

contract owned {
    address public owner;

    constructor(){
        owner = msg.sender;
    }

    modifier onlyOwner {
        require(msg.sender == owner);
        _;
    }

    function transferOwnership(address newOwner) onlyOwner public {
        owner = newOwner;
    }
}

contract Advertise is owned{

    struct AdItem{
        string name;
        string configInJson;
    }

    mapping(address=>bool)public Administrators;
    AdItem[] public advertisements;
    mapping(string=>uint) public sIdx;

    constructor() {
        Administrators[msg.sender]=true;
        advertisements.push(AdItem("", ""));
    }
    modifier onlyAdmin{
        require(Administrators[msg.sender],"invalid operator");
        _;
    }
    function addAdmin(address admin) public onlyOwner{
        Administrators[admin] = true;
    }
    function removeAdmin(address admin) public onlyOwner{
        delete  Administrators[admin];
    }


    function addItem(string memory name, string memory config)public onlyAdmin{
        require(sIdx[name] == 0, "dulicated name");

        sIdx[name]=advertisements.length;
        advertisements.push(AdItem(name, config));
    }

    function removeItem(string memory name) public onlyAdmin{
        uint idx = sIdx[name];
        require(idx > 0, "no such item");

        if (idx == advertisements.length - 1){
            advertisements.pop();
            sIdx[name]=0;
            return ;
        }

        AdItem memory lastItem = advertisements[advertisements.length - 1];
        advertisements[idx] = lastItem;
        advertisements.pop();
        sIdx[name]=0;
        sIdx[lastItem.name] = idx;
    }

    function changeAd(string memory name, string memory config)public onlyAdmin{
        uint idx = sIdx[name];
        require(idx > 0, "no such item");

        advertisements[idx].configInJson = config;
    }

    function AdList() public view returns(AdItem[] memory){
        return advertisements;
    }

    function QueryByOne(string memory name) public view returns (string memory){
        uint idx = sIdx[name];
        if (idx == 0){
            return "";
        }

        return advertisements[idx].configInJson;
    }
}