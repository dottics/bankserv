## Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]
## [1.3.0] - 2023-06-06
### Added
- The `CreateAccountBalance` method to create a new account balance.

## [1.2.0] - 2023-06-06
### Added
- The `GetAccountBalance` method to get the balance of a bank account.

## [1.1.0] - 2023-06-06
### Added
- The `CreateTransactionBatch` method to create a batch of transactions.
- This method also bring in the initial support for inter-service communication
  using a http header to pass an authentication token, as an API key.

## [1.0.0] - 2023-03-21
### Updated
**Breaking Change**
- Fetching of accounts and transactions are done with an entity UUID instead.
  User, Organisations and more is rather modelled as a generic entity
  irrespective of what the entity really is. As this does not influence the
  bank-service.
- New methods `GetEntityTransactions` and `GetEntityAccounts`.
- Removed methods `GetUserAccounts`, `GetUserTransactions` and
  `GetOrganisationTransactions`.

## [0.14.0] - 2023-03-05
### Added
- `Query` method to allow for rapid development of bank summary queries to find
  effective use and views of transactional data for users.

## [0.13.0] - 2023-02-1
### Added
- Method `UpdateItemTags` to update and set all the tags on a item.

## [0.12.0] - 2022-10-17
### Added
- Method `GetTransaction` that accepts a UUID parameter and returns the
respective transaction from the bank service.

## [0.11.2] - 2022-09-26
### Updated
- Added the type `Transaction.AccountUUID` field, as this is  required to be 
able to create a transaction.

## [0.11.1] - 2022-09-26
### Updated
- Updated the type `Account`'s `AccountNumber string `json:"account_number"` to
`Number string `json:"number"`. Making more logical sense as `Account.Number`.
- Removed the type `Transaction.AccountUUID` field, as this is a feature in to
come in the future.

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


