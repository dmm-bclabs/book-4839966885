// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// DestructibleABI is the input ABI used to generate the binding from.
const DestructibleABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"destroyAndSend\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// Destructible is an auto generated Go binding around an Ethereum contract.
type Destructible struct {
	DestructibleCaller     // Read-only binding to the contract
	DestructibleTransactor // Write-only binding to the contract
	DestructibleFilterer   // Log filterer for contract events
}

// DestructibleCaller is an auto generated read-only Go binding around an Ethereum contract.
type DestructibleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DestructibleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DestructibleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DestructibleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DestructibleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DestructibleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DestructibleSession struct {
	Contract     *Destructible     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DestructibleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DestructibleCallerSession struct {
	Contract *DestructibleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// DestructibleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DestructibleTransactorSession struct {
	Contract     *DestructibleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// DestructibleRaw is an auto generated low-level Go binding around an Ethereum contract.
type DestructibleRaw struct {
	Contract *Destructible // Generic contract binding to access the raw methods on
}

// DestructibleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DestructibleCallerRaw struct {
	Contract *DestructibleCaller // Generic read-only contract binding to access the raw methods on
}

// DestructibleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DestructibleTransactorRaw struct {
	Contract *DestructibleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDestructible creates a new instance of Destructible, bound to a specific deployed contract.
func NewDestructible(address common.Address, backend bind.ContractBackend) (*Destructible, error) {
	contract, err := bindDestructible(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Destructible{DestructibleCaller: DestructibleCaller{contract: contract}, DestructibleTransactor: DestructibleTransactor{contract: contract}, DestructibleFilterer: DestructibleFilterer{contract: contract}}, nil
}

// NewDestructibleCaller creates a new read-only instance of Destructible, bound to a specific deployed contract.
func NewDestructibleCaller(address common.Address, caller bind.ContractCaller) (*DestructibleCaller, error) {
	contract, err := bindDestructible(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DestructibleCaller{contract: contract}, nil
}

// NewDestructibleTransactor creates a new write-only instance of Destructible, bound to a specific deployed contract.
func NewDestructibleTransactor(address common.Address, transactor bind.ContractTransactor) (*DestructibleTransactor, error) {
	contract, err := bindDestructible(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DestructibleTransactor{contract: contract}, nil
}

// NewDestructibleFilterer creates a new log filterer instance of Destructible, bound to a specific deployed contract.
func NewDestructibleFilterer(address common.Address, filterer bind.ContractFilterer) (*DestructibleFilterer, error) {
	contract, err := bindDestructible(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DestructibleFilterer{contract: contract}, nil
}

// bindDestructible binds a generic wrapper to an already deployed contract.
func bindDestructible(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DestructibleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Destructible *DestructibleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Destructible.Contract.DestructibleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Destructible *DestructibleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Destructible.Contract.DestructibleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Destructible *DestructibleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Destructible.Contract.DestructibleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Destructible *DestructibleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Destructible.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Destructible *DestructibleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Destructible.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Destructible *DestructibleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Destructible.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Destructible *DestructibleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Destructible.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Destructible *DestructibleSession) Owner() (common.Address, error) {
	return _Destructible.Contract.Owner(&_Destructible.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Destructible *DestructibleCallerSession) Owner() (common.Address, error) {
	return _Destructible.Contract.Owner(&_Destructible.CallOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Destructible *DestructibleTransactor) Destroy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Destructible.contract.Transact(opts, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Destructible *DestructibleSession) Destroy() (*types.Transaction, error) {
	return _Destructible.Contract.Destroy(&_Destructible.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Destructible *DestructibleTransactorSession) Destroy() (*types.Transaction, error) {
	return _Destructible.Contract.Destroy(&_Destructible.TransactOpts)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0xf5074f41.
//
// Solidity: function destroyAndSend(_recipient address) returns()
func (_Destructible *DestructibleTransactor) DestroyAndSend(opts *bind.TransactOpts, _recipient common.Address) (*types.Transaction, error) {
	return _Destructible.contract.Transact(opts, "destroyAndSend", _recipient)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0xf5074f41.
//
// Solidity: function destroyAndSend(_recipient address) returns()
func (_Destructible *DestructibleSession) DestroyAndSend(_recipient common.Address) (*types.Transaction, error) {
	return _Destructible.Contract.DestroyAndSend(&_Destructible.TransactOpts, _recipient)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0xf5074f41.
//
// Solidity: function destroyAndSend(_recipient address) returns()
func (_Destructible *DestructibleTransactorSession) DestroyAndSend(_recipient common.Address) (*types.Transaction, error) {
	return _Destructible.Contract.DestroyAndSend(&_Destructible.TransactOpts, _recipient)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Destructible *DestructibleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Destructible.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Destructible *DestructibleSession) RenounceOwnership() (*types.Transaction, error) {
	return _Destructible.Contract.RenounceOwnership(&_Destructible.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Destructible *DestructibleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Destructible.Contract.RenounceOwnership(&_Destructible.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Destructible *DestructibleTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Destructible.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Destructible *DestructibleSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Destructible.Contract.TransferOwnership(&_Destructible.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Destructible *DestructibleTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Destructible.Contract.TransferOwnership(&_Destructible.TransactOpts, _newOwner)
}

// DestructibleOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Destructible contract.
type DestructibleOwnershipRenouncedIterator struct {
	Event *DestructibleOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *DestructibleOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestructibleOwnershipRenounced)
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
		it.Event = new(DestructibleOwnershipRenounced)
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
func (it *DestructibleOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestructibleOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestructibleOwnershipRenounced represents a OwnershipRenounced event raised by the Destructible contract.
type DestructibleOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Destructible *DestructibleFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*DestructibleOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Destructible.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DestructibleOwnershipRenouncedIterator{contract: _Destructible.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Destructible *DestructibleFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *DestructibleOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Destructible.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestructibleOwnershipRenounced)
				if err := _Destructible.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// DestructibleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Destructible contract.
type DestructibleOwnershipTransferredIterator struct {
	Event *DestructibleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DestructibleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestructibleOwnershipTransferred)
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
		it.Event = new(DestructibleOwnershipTransferred)
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
func (it *DestructibleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestructibleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestructibleOwnershipTransferred represents a OwnershipTransferred event raised by the Destructible contract.
type DestructibleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Destructible *DestructibleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DestructibleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Destructible.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DestructibleOwnershipTransferredIterator{contract: _Destructible.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Destructible *DestructibleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DestructibleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Destructible.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestructibleOwnershipTransferred)
				if err := _Destructible.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
