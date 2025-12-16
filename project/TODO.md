# TODO: Implement User Requests

## 1. Improve Password Prompt

- [x] Update go.mod to add golang.org/x/term
- [x] Modify PromptPassword in helpers.go to use masked input (\*)
- [x] Update main.go to prompt for password interactively if not provided

## 2. Enhance Password Validation

- [x] Update ValidatePassword in helpers.go to require:
  - Uppercase letter
  - Lowercase letter
  - Number
  - Symbol
  - Length 8-12 characters

## 3. Update User Information

- [x] Ensure update action validates user data before updating
- [x] Add validation in main.go for update action

## 4. Improve Location Collection

- [x] Add simple location prompt with predefined options (e.g., cities)
- [x] Modify create/update to use location prompt

## 5. View Database

- [x] Add new "rawlist" action to show encrypted data and user count
- [x] Keep existing "list" for plaintext

## 6. Comprehensive Validation

- [x] Update user.go Validate method:
  - Email: Must be @gmail.com
  - Phone: 10-15 digits
  - Address: Not empty
  - Name: 2-50 characters
- [x] Apply validations in create and update actions

## Followup

- [x] Run go mod tidy
- [x] Test changes with existing tests
- [x] Run CLI to verify functionality