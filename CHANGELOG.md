## Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]
## [0.11.0] - 2022-09-24
### Added
- `AddItemTags` method to add tags to an item.
- `RemoveItemTags` method to remove tags from an item.

## [0.10.0] - 2022-09-24
### Updated
- `Transaction` to have new fields
  - `Debit` increases the account asset.
  - `Crebit` decreases the account asset.
  - `Amount` the value and validation value with which the account asset
  increases or decreases.
- `Item` to have new fields.
  - `SKU` is the stock keeping units or unique reference such as a barcode.
  - `Unit` the smallest measurement of the item.
  - `Quantity` the number of units the item consists of.
- Update the `Service` methods related to these types.

## [0.9.1] - 2022-09-23
### Updated
- Updated the `Transaction` type to include a `BusinessName` field.
- Updated the `GetAccountTransactions`, `GetUserTransactions`, 
`CreateTransaction` and `UpdateTransaction` methods to use the `BusinessName`
field.

## [0.9.0] - 2022-09-22
### Updated
- Migrate all bank accounts to only account.

## [0.8.0] - 2022-09-19
### Added
- The Get User Transactions method. To enable getting all a user's transactions
based on the user's UUID.

## [0.7.0] - 2022-06-24
### Added
- The Get Tag methods
  - `GetTags` to get all system default tags.
  - `GetUserTags` to get all user tags based on the user UUID.
  - `GetOrganisationTags` to get all organisation tags based on the organisation UUID.

## [0.6.0] - 2022-06-24
### Added
- The CRUD methods.
  - `CreateTag` to create new tag.
  - `UpdateTag` to update a tag.
  - `DeleteTag` to delete a tag.

## [0.5.0] - 2022-06-20
### Added
- The CRUD methods.
  - `CreateItem` to create a new item.
  - `UpdateItem` to update an item.
  - `DeleteItem` to delete an item.

## [Released]
## [0.4.0] - 2022-06-17
### Added
- Comparison functions to compare types.
  - `EqualTags`.
  - `EqualItem` and `EqualItems`.
  - `EqualTransaction` and `EqualTransactions`.
- The CRUD Transaction methods.
  - `CreateTransaction` to create a new transaction.
  - `GetAccountTransactions` to get all the transactions for a bank account.
  - `UpdateTransaction` to update a transaction.
  - `DeleteTransaction` to delete a transaction.

### Updated
- Transaction to have an `AccountUUID` instead of `AccountUUID` for a more
generic description.

## [0.3.0] - 2022-06-17
### Added
- The CRUD Bank Account methods.
  - `CreateAccount` to create a new bank account.
  - `UpdateAccount` to update a bank account.
  - `DeleteAccount` to delete a bank account.
  - `GetUserAccounts` to get all the users bank accounts.
  - `GetOrganisationAccounts` to get all the organisations bank accounts.

## [0.2.0] - 2022-06-14
### Added
- The `Service` type
  - `SetURL` method to make the service a `microtest.Mock` interface.
- The microservice-package `NewService` function to be able to create a new 
msp instance.

## [0.1.0] - 2022-06-14
### Added
- Some basic types
  - Bank
  - Account
  - Transaction
  - Item
  - Tag
## [0.0.0] - 2022-05-02
### Added
- Initial commit.


