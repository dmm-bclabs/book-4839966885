// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// RoomABI is the input ABI used to generate the binding from.
const RoomABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"active\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"activate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_reward\",\"type\":\"uint256\"},{\"name\":\"_dest\",\"type\":\"address\"},{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"sendReward\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"refundToOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deactivate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rewardSent\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"destroyAndSend\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_creator\",\"type\":\"address\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_depositor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_depositedValue\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_dest\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_reward\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"RewardSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_dest\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_refundedBalance\",\"type\":\"uint256\"}],\"name\":\"RefundedToOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"Deactivate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"Activate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// Room is an auto generated Go binding around an Ethereum contract.
type Room struct {
	RoomCaller     // Read-only binding to the contract
	RoomTransactor // Write-only binding to the contract
	RoomFilterer   // Log filterer for contract events
}

// RoomCaller is an auto generated read-only Go binding around an Ethereum contract.
type RoomCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoomTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RoomTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoomFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RoomFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoomSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RoomSession struct {
	Contract     *Room             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RoomCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RoomCallerSession struct {
	Contract *RoomCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RoomTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RoomTransactorSession struct {
	Contract     *RoomTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RoomRaw is an auto generated low-level Go binding around an Ethereum contract.
type RoomRaw struct {
	Contract *Room // Generic contract binding to access the raw methods on
}

// RoomCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RoomCallerRaw struct {
	Contract *RoomCaller // Generic read-only contract binding to access the raw methods on
}

// RoomTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RoomTransactorRaw struct {
	Contract *RoomTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRoom creates a new instance of Room, bound to a specific deployed contract.
func NewRoom(address common.Address, backend bind.ContractBackend) (*Room, error) {
	contract, err := bindRoom(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Room{RoomCaller: RoomCaller{contract: contract}, RoomTransactor: RoomTransactor{contract: contract}, RoomFilterer: RoomFilterer{contract: contract}}, nil
}

// NewRoomCaller creates a new read-only instance of Room, bound to a specific deployed contract.
func NewRoomCaller(address common.Address, caller bind.ContractCaller) (*RoomCaller, error) {
	contract, err := bindRoom(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RoomCaller{contract: contract}, nil
}

// NewRoomTransactor creates a new write-only instance of Room, bound to a specific deployed contract.
func NewRoomTransactor(address common.Address, transactor bind.ContractTransactor) (*RoomTransactor, error) {
	contract, err := bindRoom(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RoomTransactor{contract: contract}, nil
}

// NewRoomFilterer creates a new log filterer instance of Room, bound to a specific deployed contract.
func NewRoomFilterer(address common.Address, filterer bind.ContractFilterer) (*RoomFilterer, error) {
	contract, err := bindRoom(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RoomFilterer{contract: contract}, nil
}

// bindRoom binds a generic wrapper to an already deployed contract.
func bindRoom(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RoomABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Room *RoomRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Room.Contract.RoomCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Room *RoomRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Room.Contract.RoomTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Room *RoomRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Room.Contract.RoomTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Room *RoomCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Room.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Room *RoomTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Room.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Room *RoomTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Room.Contract.contract.Transact(opts, method, params...)
}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() constant returns(bool)
func (_Room *RoomCaller) Active(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Room.contract.Call(opts, out, "active")
	return *ret0, err
}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() constant returns(bool)
func (_Room *RoomSession) Active() (bool, error) {
	return _Room.Contract.Active(&_Room.CallOpts)
}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() constant returns(bool)
func (_Room *RoomCallerSession) Active() (bool, error) {
	return _Room.Contract.Active(&_Room.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Room *RoomCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Room.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Room *RoomSession) Owner() (common.Address, error) {
	return _Room.Contract.Owner(&_Room.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Room *RoomCallerSession) Owner() (common.Address, error) {
	return _Room.Contract.Owner(&_Room.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Room *RoomCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Room.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Room *RoomSession) Paused() (bool, error) {
	return _Room.Contract.Paused(&_Room.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Room *RoomCallerSession) Paused() (bool, error) {
	return _Room.Contract.Paused(&_Room.CallOpts)
}

// RewardSent is a free data retrieval call binding the contract method 0x8f134025.
//
// Solidity: function rewardSent( uint256) constant returns(bool)
func (_Room *RoomCaller) RewardSent(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Room.contract.Call(opts, out, "rewardSent", arg0)
	return *ret0, err
}

// RewardSent is a free data retrieval call binding the contract method 0x8f134025.
//
// Solidity: function rewardSent( uint256) constant returns(bool)
func (_Room *RoomSession) RewardSent(arg0 *big.Int) (bool, error) {
	return _Room.Contract.RewardSent(&_Room.CallOpts, arg0)
}

// RewardSent is a free data retrieval call binding the contract method 0x8f134025.
//
// Solidity: function rewardSent( uint256) constant returns(bool)
func (_Room *RoomCallerSession) RewardSent(arg0 *big.Int) (bool, error) {
	return _Room.Contract.RewardSent(&_Room.CallOpts, arg0)
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_Room *RoomTransactor) Activate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Room.contract.Transact(opts, "activate")
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_Room *RoomSession) Activate() (*types.Transaction, error) {
	return _Room.Contract.Activate(&_Room.TransactOpts)
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_Room *RoomTransactorSession) Activate() (*types.Transaction, error) {
	return _Room.Contract.Activate(&_Room.TransactOpts)
}

// Deactivate is a paid mutator transaction binding the contract method 0x51b42b00.
//
// Solidity: function deactivate() returns()
func (_Room *RoomTransactor) Deactivate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Room.contract.Transact(opts, "deactivate")
}

// Deactivate is a paid mutator transaction binding the contract method 0x51b42b00.
//
// Solidity: function deactivate() returns()
func (_Room *RoomSession) Deactivate() (*types.Transaction, error) {
	return _Room.Contract.Deactivate(&_Room.TransactOpts)
}

// Deactivate is a paid mutator transaction binding the contract method 0x51b42b00.
//
// Solidity: function deactivate() returns()
func (_Room *RoomTransactorSession) Deactivate() (*types.Transaction, error) {
	return _Room.Contract.Deactivate(&_Room.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_Room *RoomTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Room.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_Room *RoomSession) Deposit() (*types.Transaction, error) {
	return _Room.Contract.Deposit(&_Room.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_Room *RoomTransactorSession) Deposit() (*types.Transaction, error) {
	return _Room.Contract.Deposit(&_Room.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Room *RoomTransactor) Destroy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Room.contract.Transact(opts, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Room *RoomSession) Destroy() (*types.Transaction, error) {
	return _Room.Contract.Destroy(&_Room.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Room *RoomTransactorSession) Destroy() (*types.Transaction, error) {
	return _Room.Contract.Destroy(&_Room.TransactOpts)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0xf5074f41.
//
// Solidity: function destroyAndSend(_recipient address) returns()
func (_Room *RoomTransactor) DestroyAndSend(opts *bind.TransactOpts, _recipient common.Address) (*types.Transaction, error) {
	return _Room.contract.Transact(opts, "destroyAndSend", _recipient)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0xf5074f41.
//
// Solidity: function destroyAndSend(_recipient address) returns()
func (_Room *RoomSession) DestroyAndSend(_recipient common.Address) (*types.Transaction, error) {
	return _Room.Contract.DestroyAndSend(&_Room.TransactOpts, _recipient)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0xf5074f41.
//
// Solidity: function destroyAndSend(_recipient address) returns()
func (_Room *RoomTransactorSession) DestroyAndSend(_recipient common.Address) (*types.Transaction, error) {
	return _Room.Contract.DestroyAndSend(&_Room.TransactOpts, _recipient)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Room *RoomTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Room.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Room *RoomSession) Pause() (*types.Transaction, error) {
	return _Room.Contract.Pause(&_Room.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Room *RoomTransactorSession) Pause() (*types.Transaction, error) {
	return _Room.Contract.Pause(&_Room.TransactOpts)
}

// RefundToOwner is a paid mutator transaction binding the contract method 0x3e4e10d6.
//
// Solidity: function refundToOwner() returns()
func (_Room *RoomTransactor) RefundToOwner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Room.contract.Transact(opts, "refundToOwner")
}

// RefundToOwner is a paid mutator transaction binding the contract method 0x3e4e10d6.
//
// Solidity: function refundToOwner() returns()
func (_Room *RoomSession) RefundToOwner() (*types.Transaction, error) {
	return _Room.Contract.RefundToOwner(&_Room.TransactOpts)
}

// RefundToOwner is a paid mutator transaction binding the contract method 0x3e4e10d6.
//
// Solidity: function refundToOwner() returns()
func (_Room *RoomTransactorSession) RefundToOwner() (*types.Transaction, error) {
	return _Room.Contract.RefundToOwner(&_Room.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Room *RoomTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Room.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Room *RoomSession) RenounceOwnership() (*types.Transaction, error) {
	return _Room.Contract.RenounceOwnership(&_Room.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Room *RoomTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Room.Contract.RenounceOwnership(&_Room.TransactOpts)
}

// SendReward is a paid mutator transaction binding the contract method 0x2673bd74.
//
// Solidity: function sendReward(_reward uint256, _dest address, _id uint256) returns()
func (_Room *RoomTransactor) SendReward(opts *bind.TransactOpts, _reward *big.Int, _dest common.Address, _id *big.Int) (*types.Transaction, error) {
	return _Room.contract.Transact(opts, "sendReward", _reward, _dest, _id)
}

// SendReward is a paid mutator transaction binding the contract method 0x2673bd74.
//
// Solidity: function sendReward(_reward uint256, _dest address, _id uint256) returns()
func (_Room *RoomSession) SendReward(_reward *big.Int, _dest common.Address, _id *big.Int) (*types.Transaction, error) {
	return _Room.Contract.SendReward(&_Room.TransactOpts, _reward, _dest, _id)
}

// SendReward is a paid mutator transaction binding the contract method 0x2673bd74.
//
// Solidity: function sendReward(_reward uint256, _dest address, _id uint256) returns()
func (_Room *RoomTransactorSession) SendReward(_reward *big.Int, _dest common.Address, _id *big.Int) (*types.Transaction, error) {
	return _Room.Contract.SendReward(&_Room.TransactOpts, _reward, _dest, _id)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Room *RoomTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Room.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Room *RoomSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Room.Contract.TransferOwnership(&_Room.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Room *RoomTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Room.Contract.TransferOwnership(&_Room.TransactOpts, _newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Room *RoomTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Room.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Room *RoomSession) Unpause() (*types.Transaction, error) {
	return _Room.Contract.Unpause(&_Room.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Room *RoomTransactorSession) Unpause() (*types.Transaction, error) {
	return _Room.Contract.Unpause(&_Room.TransactOpts)
}

// RoomActivateIterator is returned from FilterActivate and is used to iterate over the raw logs and unpacked data for Activate events raised by the Room contract.
type RoomActivateIterator struct {
	Event *RoomActivate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RoomActivateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoomActivate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RoomActivate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RoomActivateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoomActivateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoomActivate represents a Activate event raised by the Room contract.
type RoomActivate struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterActivate is a free log retrieval operation binding the contract event 0xf7e9fe69e1d05372bc855b295bc4c34a1a0a5882164dd2b26df30a26c1c8ba15.
//
// Solidity: e Activate(_sender indexed address)
func (_Room *RoomFilterer) FilterActivate(opts *bind.FilterOpts, _sender []common.Address) (*RoomActivateIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Room.contract.FilterLogs(opts, "Activate", _senderRule)
	if err != nil {
		return nil, err
	}
	return &RoomActivateIterator{contract: _Room.contract, event: "Activate", logs: logs, sub: sub}, nil
}

// WatchActivate is a free log subscription operation binding the contract event 0xf7e9fe69e1d05372bc855b295bc4c34a1a0a5882164dd2b26df30a26c1c8ba15.
//
// Solidity: e Activate(_sender indexed address)
func (_Room *RoomFilterer) WatchActivate(opts *bind.WatchOpts, sink chan<- *RoomActivate, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Room.contract.WatchLogs(opts, "Activate", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoomActivate)
				if err := _Room.contract.UnpackLog(event, "Activate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// RoomDeactivateIterator is returned from FilterDeactivate and is used to iterate over the raw logs and unpacked data for Deactivate events raised by the Room contract.
type RoomDeactivateIterator struct {
	Event *RoomDeactivate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RoomDeactivateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoomDeactivate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RoomDeactivate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RoomDeactivateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoomDeactivateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoomDeactivate represents a Deactivate event raised by the Room contract.
type RoomDeactivate struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeactivate is a free log retrieval operation binding the contract event 0x238ce44d0fada9e1348f183a8436996fb52b4c41a9cbf3af6e2dee00fcb80a9c.
//
// Solidity: e Deactivate(_sender indexed address)
func (_Room *RoomFilterer) FilterDeactivate(opts *bind.FilterOpts, _sender []common.Address) (*RoomDeactivateIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Room.contract.FilterLogs(opts, "Deactivate", _senderRule)
	if err != nil {
		return nil, err
	}
	return &RoomDeactivateIterator{contract: _Room.contract, event: "Deactivate", logs: logs, sub: sub}, nil
}

// WatchDeactivate is a free log subscription operation binding the contract event 0x238ce44d0fada9e1348f183a8436996fb52b4c41a9cbf3af6e2dee00fcb80a9c.
//
// Solidity: e Deactivate(_sender indexed address)
func (_Room *RoomFilterer) WatchDeactivate(opts *bind.WatchOpts, sink chan<- *RoomDeactivate, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Room.contract.WatchLogs(opts, "Deactivate", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoomDeactivate)
				if err := _Room.contract.UnpackLog(event, "Deactivate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// RoomDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the Room contract.
type RoomDepositedIterator struct {
	Event *RoomDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RoomDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoomDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RoomDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RoomDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoomDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoomDeposited represents a Deposited event raised by the Room contract.
type RoomDeposited struct {
	Depositor      common.Address
	DepositedValue *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: e Deposited(_depositor indexed address, _depositedValue uint256)
func (_Room *RoomFilterer) FilterDeposited(opts *bind.FilterOpts, _depositor []common.Address) (*RoomDepositedIterator, error) {

	var _depositorRule []interface{}
	for _, _depositorItem := range _depositor {
		_depositorRule = append(_depositorRule, _depositorItem)
	}

	logs, sub, err := _Room.contract.FilterLogs(opts, "Deposited", _depositorRule)
	if err != nil {
		return nil, err
	}
	return &RoomDepositedIterator{contract: _Room.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: e Deposited(_depositor indexed address, _depositedValue uint256)
func (_Room *RoomFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *RoomDeposited, _depositor []common.Address) (event.Subscription, error) {

	var _depositorRule []interface{}
	for _, _depositorItem := range _depositor {
		_depositorRule = append(_depositorRule, _depositorItem)
	}

	logs, sub, err := _Room.contract.WatchLogs(opts, "Deposited", _depositorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoomDeposited)
				if err := _Room.contract.UnpackLog(event, "Deposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// RoomOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Room contract.
type RoomOwnershipRenouncedIterator struct {
	Event *RoomOwnershipRenounced // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RoomOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoomOwnershipRenounced)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RoomOwnershipRenounced)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RoomOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoomOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoomOwnershipRenounced represents a OwnershipRenounced event raised by the Room contract.
type RoomOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Room *RoomFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*RoomOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Room.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RoomOwnershipRenouncedIterator{contract: _Room.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Room *RoomFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *RoomOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Room.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoomOwnershipRenounced)
				if err := _Room.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// RoomOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Room contract.
type RoomOwnershipTransferredIterator struct {
	Event *RoomOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RoomOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoomOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RoomOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RoomOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoomOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoomOwnershipTransferred represents a OwnershipTransferred event raised by the Room contract.
type RoomOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Room *RoomFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RoomOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Room.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RoomOwnershipTransferredIterator{contract: _Room.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Room *RoomFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RoomOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Room.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoomOwnershipTransferred)
				if err := _Room.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// RoomPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the Room contract.
type RoomPauseIterator struct {
	Event *RoomPause // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RoomPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoomPause)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RoomPause)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RoomPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoomPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoomPause represents a Pause event raised by the Room contract.
type RoomPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_Room *RoomFilterer) FilterPause(opts *bind.FilterOpts) (*RoomPauseIterator, error) {

	logs, sub, err := _Room.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &RoomPauseIterator{contract: _Room.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_Room *RoomFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *RoomPause) (event.Subscription, error) {

	logs, sub, err := _Room.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoomPause)
				if err := _Room.contract.UnpackLog(event, "Pause", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// RoomRefundedToOwnerIterator is returned from FilterRefundedToOwner and is used to iterate over the raw logs and unpacked data for RefundedToOwner events raised by the Room contract.
type RoomRefundedToOwnerIterator struct {
	Event *RoomRefundedToOwner // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RoomRefundedToOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoomRefundedToOwner)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RoomRefundedToOwner)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RoomRefundedToOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoomRefundedToOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoomRefundedToOwner represents a RefundedToOwner event raised by the Room contract.
type RoomRefundedToOwner struct {
	Dest            common.Address
	RefundedBalance *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRefundedToOwner is a free log retrieval operation binding the contract event 0xa9304f0ac97f820b11d5a66f29245fb1e76b1f1f93f46a84a71682e69d095e80.
//
// Solidity: e RefundedToOwner(_dest indexed address, _refundedBalance uint256)
func (_Room *RoomFilterer) FilterRefundedToOwner(opts *bind.FilterOpts, _dest []common.Address) (*RoomRefundedToOwnerIterator, error) {

	var _destRule []interface{}
	for _, _destItem := range _dest {
		_destRule = append(_destRule, _destItem)
	}

	logs, sub, err := _Room.contract.FilterLogs(opts, "RefundedToOwner", _destRule)
	if err != nil {
		return nil, err
	}
	return &RoomRefundedToOwnerIterator{contract: _Room.contract, event: "RefundedToOwner", logs: logs, sub: sub}, nil
}

// WatchRefundedToOwner is a free log subscription operation binding the contract event 0xa9304f0ac97f820b11d5a66f29245fb1e76b1f1f93f46a84a71682e69d095e80.
//
// Solidity: e RefundedToOwner(_dest indexed address, _refundedBalance uint256)
func (_Room *RoomFilterer) WatchRefundedToOwner(opts *bind.WatchOpts, sink chan<- *RoomRefundedToOwner, _dest []common.Address) (event.Subscription, error) {

	var _destRule []interface{}
	for _, _destItem := range _dest {
		_destRule = append(_destRule, _destItem)
	}

	logs, sub, err := _Room.contract.WatchLogs(opts, "RefundedToOwner", _destRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoomRefundedToOwner)
				if err := _Room.contract.UnpackLog(event, "RefundedToOwner", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// RoomRewardSentIterator is returned from FilterRewardSent and is used to iterate over the raw logs and unpacked data for RewardSent events raised by the Room contract.
type RoomRewardSentIterator struct {
	Event *RoomRewardSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RoomRewardSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoomRewardSent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RoomRewardSent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RoomRewardSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoomRewardSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoomRewardSent represents a RewardSent event raised by the Room contract.
type RoomRewardSent struct {
	Dest   common.Address
	Reward *big.Int
	Id     *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardSent is a free log retrieval operation binding the contract event 0x6379339f0ae63e95e65fad18ca2a7ec4e7e3f05f3cc5f7079f4d8da8cec34faa.
//
// Solidity: e RewardSent(_dest indexed address, _reward uint256, _id uint256)
func (_Room *RoomFilterer) FilterRewardSent(opts *bind.FilterOpts, _dest []common.Address) (*RoomRewardSentIterator, error) {

	var _destRule []interface{}
	for _, _destItem := range _dest {
		_destRule = append(_destRule, _destItem)
	}

	logs, sub, err := _Room.contract.FilterLogs(opts, "RewardSent", _destRule)
	if err != nil {
		return nil, err
	}
	return &RoomRewardSentIterator{contract: _Room.contract, event: "RewardSent", logs: logs, sub: sub}, nil
}

// WatchRewardSent is a free log subscription operation binding the contract event 0x6379339f0ae63e95e65fad18ca2a7ec4e7e3f05f3cc5f7079f4d8da8cec34faa.
//
// Solidity: e RewardSent(_dest indexed address, _reward uint256, _id uint256)
func (_Room *RoomFilterer) WatchRewardSent(opts *bind.WatchOpts, sink chan<- *RoomRewardSent, _dest []common.Address) (event.Subscription, error) {

	var _destRule []interface{}
	for _, _destItem := range _dest {
		_destRule = append(_destRule, _destItem)
	}

	logs, sub, err := _Room.contract.WatchLogs(opts, "RewardSent", _destRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoomRewardSent)
				if err := _Room.contract.UnpackLog(event, "RewardSent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// RoomUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the Room contract.
type RoomUnpauseIterator struct {
	Event *RoomUnpause // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RoomUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoomUnpause)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RoomUnpause)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RoomUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoomUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoomUnpause represents a Unpause event raised by the Room contract.
type RoomUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_Room *RoomFilterer) FilterUnpause(opts *bind.FilterOpts) (*RoomUnpauseIterator, error) {

	logs, sub, err := _Room.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &RoomUnpauseIterator{contract: _Room.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_Room *RoomFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *RoomUnpause) (event.Subscription, error) {

	logs, sub, err := _Room.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoomUnpause)
				if err := _Room.contract.UnpackLog(event, "Unpause", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
