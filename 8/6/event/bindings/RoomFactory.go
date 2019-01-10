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

// RoomFactoryABI is the input ABI used to generate the binding from.
const RoomFactoryABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"createRoom\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"destroyAndSend\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_creator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_room\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_depositedValue\",\"type\":\"uint256\"}],\"name\":\"RoomCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// RoomFactory is an auto generated Go binding around an Ethereum contract.
type RoomFactory struct {
	RoomFactoryCaller     // Read-only binding to the contract
	RoomFactoryTransactor // Write-only binding to the contract
	RoomFactoryFilterer   // Log filterer for contract events
}

// RoomFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RoomFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoomFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RoomFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoomFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RoomFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoomFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RoomFactorySession struct {
	Contract     *RoomFactory      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RoomFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RoomFactoryCallerSession struct {
	Contract *RoomFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// RoomFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RoomFactoryTransactorSession struct {
	Contract     *RoomFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// RoomFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RoomFactoryRaw struct {
	Contract *RoomFactory // Generic contract binding to access the raw methods on
}

// RoomFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RoomFactoryCallerRaw struct {
	Contract *RoomFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// RoomFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RoomFactoryTransactorRaw struct {
	Contract *RoomFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRoomFactory creates a new instance of RoomFactory, bound to a specific deployed contract.
func NewRoomFactory(address common.Address, backend bind.ContractBackend) (*RoomFactory, error) {
	contract, err := bindRoomFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RoomFactory{RoomFactoryCaller: RoomFactoryCaller{contract: contract}, RoomFactoryTransactor: RoomFactoryTransactor{contract: contract}, RoomFactoryFilterer: RoomFactoryFilterer{contract: contract}}, nil
}

// NewRoomFactoryCaller creates a new read-only instance of RoomFactory, bound to a specific deployed contract.
func NewRoomFactoryCaller(address common.Address, caller bind.ContractCaller) (*RoomFactoryCaller, error) {
	contract, err := bindRoomFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RoomFactoryCaller{contract: contract}, nil
}

// NewRoomFactoryTransactor creates a new write-only instance of RoomFactory, bound to a specific deployed contract.
func NewRoomFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*RoomFactoryTransactor, error) {
	contract, err := bindRoomFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RoomFactoryTransactor{contract: contract}, nil
}

// NewRoomFactoryFilterer creates a new log filterer instance of RoomFactory, bound to a specific deployed contract.
func NewRoomFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*RoomFactoryFilterer, error) {
	contract, err := bindRoomFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RoomFactoryFilterer{contract: contract}, nil
}

// bindRoomFactory binds a generic wrapper to an already deployed contract.
func bindRoomFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RoomFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RoomFactory *RoomFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RoomFactory.Contract.RoomFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RoomFactory *RoomFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoomFactory.Contract.RoomFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RoomFactory *RoomFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RoomFactory.Contract.RoomFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RoomFactory *RoomFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RoomFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RoomFactory *RoomFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoomFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RoomFactory *RoomFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RoomFactory.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RoomFactory *RoomFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RoomFactory.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RoomFactory *RoomFactorySession) Owner() (common.Address, error) {
	return _RoomFactory.Contract.Owner(&_RoomFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RoomFactory *RoomFactoryCallerSession) Owner() (common.Address, error) {
	return _RoomFactory.Contract.Owner(&_RoomFactory.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_RoomFactory *RoomFactoryCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RoomFactory.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_RoomFactory *RoomFactorySession) Paused() (bool, error) {
	return _RoomFactory.Contract.Paused(&_RoomFactory.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_RoomFactory *RoomFactoryCallerSession) Paused() (bool, error) {
	return _RoomFactory.Contract.Paused(&_RoomFactory.CallOpts)
}

// CreateRoom is a paid mutator transaction binding the contract method 0x3be272aa.
//
// Solidity: function createRoom() returns()
func (_RoomFactory *RoomFactoryTransactor) CreateRoom(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoomFactory.contract.Transact(opts, "createRoom")
}

// CreateRoom is a paid mutator transaction binding the contract method 0x3be272aa.
//
// Solidity: function createRoom() returns()
func (_RoomFactory *RoomFactorySession) CreateRoom() (*types.Transaction, error) {
	return _RoomFactory.Contract.CreateRoom(&_RoomFactory.TransactOpts)
}

// CreateRoom is a paid mutator transaction binding the contract method 0x3be272aa.
//
// Solidity: function createRoom() returns()
func (_RoomFactory *RoomFactoryTransactorSession) CreateRoom() (*types.Transaction, error) {
	return _RoomFactory.Contract.CreateRoom(&_RoomFactory.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_RoomFactory *RoomFactoryTransactor) Destroy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoomFactory.contract.Transact(opts, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_RoomFactory *RoomFactorySession) Destroy() (*types.Transaction, error) {
	return _RoomFactory.Contract.Destroy(&_RoomFactory.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_RoomFactory *RoomFactoryTransactorSession) Destroy() (*types.Transaction, error) {
	return _RoomFactory.Contract.Destroy(&_RoomFactory.TransactOpts)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0xf5074f41.
//
// Solidity: function destroyAndSend(_recipient address) returns()
func (_RoomFactory *RoomFactoryTransactor) DestroyAndSend(opts *bind.TransactOpts, _recipient common.Address) (*types.Transaction, error) {
	return _RoomFactory.contract.Transact(opts, "destroyAndSend", _recipient)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0xf5074f41.
//
// Solidity: function destroyAndSend(_recipient address) returns()
func (_RoomFactory *RoomFactorySession) DestroyAndSend(_recipient common.Address) (*types.Transaction, error) {
	return _RoomFactory.Contract.DestroyAndSend(&_RoomFactory.TransactOpts, _recipient)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0xf5074f41.
//
// Solidity: function destroyAndSend(_recipient address) returns()
func (_RoomFactory *RoomFactoryTransactorSession) DestroyAndSend(_recipient common.Address) (*types.Transaction, error) {
	return _RoomFactory.Contract.DestroyAndSend(&_RoomFactory.TransactOpts, _recipient)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_RoomFactory *RoomFactoryTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoomFactory.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_RoomFactory *RoomFactorySession) Pause() (*types.Transaction, error) {
	return _RoomFactory.Contract.Pause(&_RoomFactory.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_RoomFactory *RoomFactoryTransactorSession) Pause() (*types.Transaction, error) {
	return _RoomFactory.Contract.Pause(&_RoomFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RoomFactory *RoomFactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoomFactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RoomFactory *RoomFactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _RoomFactory.Contract.RenounceOwnership(&_RoomFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RoomFactory *RoomFactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RoomFactory.Contract.RenounceOwnership(&_RoomFactory.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_RoomFactory *RoomFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _RoomFactory.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_RoomFactory *RoomFactorySession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _RoomFactory.Contract.TransferOwnership(&_RoomFactory.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_RoomFactory *RoomFactoryTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _RoomFactory.Contract.TransferOwnership(&_RoomFactory.TransactOpts, _newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_RoomFactory *RoomFactoryTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoomFactory.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_RoomFactory *RoomFactorySession) Unpause() (*types.Transaction, error) {
	return _RoomFactory.Contract.Unpause(&_RoomFactory.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_RoomFactory *RoomFactoryTransactorSession) Unpause() (*types.Transaction, error) {
	return _RoomFactory.Contract.Unpause(&_RoomFactory.TransactOpts)
}

// RoomFactoryOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the RoomFactory contract.
type RoomFactoryOwnershipRenouncedIterator struct {
	Event *RoomFactoryOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *RoomFactoryOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoomFactoryOwnershipRenounced)
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
		it.Event = new(RoomFactoryOwnershipRenounced)
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
func (it *RoomFactoryOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoomFactoryOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoomFactoryOwnershipRenounced represents a OwnershipRenounced event raised by the RoomFactory contract.
type RoomFactoryOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_RoomFactory *RoomFactoryFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*RoomFactoryOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _RoomFactory.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RoomFactoryOwnershipRenouncedIterator{contract: _RoomFactory.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_RoomFactory *RoomFactoryFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *RoomFactoryOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _RoomFactory.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoomFactoryOwnershipRenounced)
				if err := _RoomFactory.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// RoomFactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RoomFactory contract.
type RoomFactoryOwnershipTransferredIterator struct {
	Event *RoomFactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RoomFactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoomFactoryOwnershipTransferred)
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
		it.Event = new(RoomFactoryOwnershipTransferred)
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
func (it *RoomFactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoomFactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoomFactoryOwnershipTransferred represents a OwnershipTransferred event raised by the RoomFactory contract.
type RoomFactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_RoomFactory *RoomFactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RoomFactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RoomFactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RoomFactoryOwnershipTransferredIterator{contract: _RoomFactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_RoomFactory *RoomFactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RoomFactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RoomFactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoomFactoryOwnershipTransferred)
				if err := _RoomFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// RoomFactoryPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the RoomFactory contract.
type RoomFactoryPauseIterator struct {
	Event *RoomFactoryPause // Event containing the contract specifics and raw log

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
func (it *RoomFactoryPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoomFactoryPause)
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
		it.Event = new(RoomFactoryPause)
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
func (it *RoomFactoryPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoomFactoryPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoomFactoryPause represents a Pause event raised by the RoomFactory contract.
type RoomFactoryPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_RoomFactory *RoomFactoryFilterer) FilterPause(opts *bind.FilterOpts) (*RoomFactoryPauseIterator, error) {

	logs, sub, err := _RoomFactory.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &RoomFactoryPauseIterator{contract: _RoomFactory.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_RoomFactory *RoomFactoryFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *RoomFactoryPause) (event.Subscription, error) {

	logs, sub, err := _RoomFactory.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoomFactoryPause)
				if err := _RoomFactory.contract.UnpackLog(event, "Pause", log); err != nil {
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

// RoomFactoryRoomCreatedIterator is returned from FilterRoomCreated and is used to iterate over the raw logs and unpacked data for RoomCreated events raised by the RoomFactory contract.
type RoomFactoryRoomCreatedIterator struct {
	Event *RoomFactoryRoomCreated // Event containing the contract specifics and raw log

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
func (it *RoomFactoryRoomCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoomFactoryRoomCreated)
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
		it.Event = new(RoomFactoryRoomCreated)
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
func (it *RoomFactoryRoomCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoomFactoryRoomCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoomFactoryRoomCreated represents a RoomCreated event raised by the RoomFactory contract.
type RoomFactoryRoomCreated struct {
	Creator        common.Address
	Room           common.Address
	DepositedValue *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRoomCreated is a free log retrieval operation binding the contract event 0x6849f7a409ad97d39c5ffa074bf765330bf1a574da99d4c4831196ecd77ea8da.
//
// Solidity: e RoomCreated(_creator indexed address, _room address, _depositedValue uint256)
func (_RoomFactory *RoomFactoryFilterer) FilterRoomCreated(opts *bind.FilterOpts, _creator []common.Address) (*RoomFactoryRoomCreatedIterator, error) {

	var _creatorRule []interface{}
	for _, _creatorItem := range _creator {
		_creatorRule = append(_creatorRule, _creatorItem)
	}

	logs, sub, err := _RoomFactory.contract.FilterLogs(opts, "RoomCreated", _creatorRule)
	if err != nil {
		return nil, err
	}
	return &RoomFactoryRoomCreatedIterator{contract: _RoomFactory.contract, event: "RoomCreated", logs: logs, sub: sub}, nil
}

// WatchRoomCreated is a free log subscription operation binding the contract event 0x6849f7a409ad97d39c5ffa074bf765330bf1a574da99d4c4831196ecd77ea8da.
//
// Solidity: e RoomCreated(_creator indexed address, _room address, _depositedValue uint256)
func (_RoomFactory *RoomFactoryFilterer) WatchRoomCreated(opts *bind.WatchOpts, sink chan<- *RoomFactoryRoomCreated, _creator []common.Address) (event.Subscription, error) {

	var _creatorRule []interface{}
	for _, _creatorItem := range _creator {
		_creatorRule = append(_creatorRule, _creatorItem)
	}

	logs, sub, err := _RoomFactory.contract.WatchLogs(opts, "RoomCreated", _creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoomFactoryRoomCreated)
				if err := _RoomFactory.contract.UnpackLog(event, "RoomCreated", log); err != nil {
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

// RoomFactoryUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the RoomFactory contract.
type RoomFactoryUnpauseIterator struct {
	Event *RoomFactoryUnpause // Event containing the contract specifics and raw log

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
func (it *RoomFactoryUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoomFactoryUnpause)
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
		it.Event = new(RoomFactoryUnpause)
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
func (it *RoomFactoryUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoomFactoryUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoomFactoryUnpause represents a Unpause event raised by the RoomFactory contract.
type RoomFactoryUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_RoomFactory *RoomFactoryFilterer) FilterUnpause(opts *bind.FilterOpts) (*RoomFactoryUnpauseIterator, error) {

	logs, sub, err := _RoomFactory.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &RoomFactoryUnpauseIterator{contract: _RoomFactory.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_RoomFactory *RoomFactoryFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *RoomFactoryUnpause) (event.Subscription, error) {

	logs, sub, err := _RoomFactory.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoomFactoryUnpause)
				if err := _RoomFactory.contract.UnpackLog(event, "Unpause", log); err != nil {
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
